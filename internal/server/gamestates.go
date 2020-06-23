package server

import (
	"encoding/json"
	"errors"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	cah "github.com/j4rv/cah/internal/model"
)

/*
GET GAME STATE
*/

type playerInfo struct {
	ID               int             `json:"id"`
	Name             string          `json:"name"`
	HandSize         int             `json:"handSize"`
	WhiteCardsInPlay int             `json:"whiteCardsInPlay"`
	Points           []cah.BlackCard `json:"points"`
}

type fullPlayerInfo struct {
	ID               int             `json:"id"`
	Name             string          `json:"name"`
	Hand             []cah.WhiteCard `json:"hand" db:"hand"`
	WhiteCardsInPlay []cah.WhiteCard `json:"whiteCardsInPlay"`
	Points           []cah.BlackCard `json:"points"`
}

type sinnerPlay struct {
	ID         int             `json:"id"`
	WhiteCards []cah.WhiteCard `json:"whiteCards"`
}

type gameStateResponse struct {
	ID              int            `json:"id"`
	Phase           string         `json:"phase"`
	Players         []playerInfo   `json:"players"`
	CurrCzarID      int            `json:"currentCzarID"`
	BlackCardInPlay cah.BlackCard  `json:"blackCardInPlay"`
	BlackCardsLeft  int            `json:"blackCardsLeft"`
	WhiteCardsLeft  int            `json:"whiteCardsLeft"`
	SinnerPlays     []sinnerPlay   `json:"sinnerPlays"`
	MyPlayer        fullPlayerInfo `json:"myPlayer"`
	CurrRound       int            `json:"currRound"`
	MaxRounds       int            `json:"maxRounds"`
}

var gameStateListeners = make(map[int][]chan *cah.GameState)

func startListening(gsID int, cb chan *cah.GameState) {
	gameStateListeners[gsID] = append(gameStateListeners[gsID], cb)
}

func stopListening(gsID int, cb chan *cah.GameState) {
	var cbRemoved []chan *cah.GameState
	for _, listener := range gameStateListeners[gsID] {
		if cb == listener {
			close(cb)
			continue
		}
		cbRemoved = append(cbRemoved, listener)
	}
	gameStateListeners[gsID] = cbRemoved
}

func gameStateUpdated(gs *cah.GameState) {
	for i := range gameStateListeners[gs.ID] {
		gameStateListeners[gs.ID][i] <- gs
	}
}

var upgrader = websocket.Upgrader{
	HandshakeTimeout: 5 * time.Second,
}

const (
	writeWait  = 10 * time.Second
	pingPeriod = 20 * time.Second // Must be less than pongWait.
	pongWait   = 60 * time.Second
)

func gameStateWebsocket(game cah.Game, user cah.User, w http.ResponseWriter, req *http.Request) (err error) {
	gameState, err := usecase.GameState.ByID(game.StateID)
	if err != nil {
		return
	}

	p, err := player(gameState, user)
	if err != nil {
		return
	}

	conn, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		return
	}

	gsEventChan := make(chan *cah.GameState)
	startListening(gameState.ID, gsEventChan)
	log.Printf("[Game: %s][User: %s] Started listening", game.Name, user.Username)

	defer func() {
		conn.Close()
		stopListening(gameState.ID, gsEventChan)
		log.Printf("[Game: %s][User: %s] Stopped listening", game.Name, user.Username)
	}()

	// Start by returning the current state
	conn.WriteJSON(newGameStateResponse(gameState, p))
	pingPongChan := webSocketPingPongs(conn)

	for {
		select {
		case gameState = <-gsEventChan:
			conn.SetWriteDeadline(time.Now().Add(writeWait))
			connErr := conn.WriteJSON(newGameStateResponse(gameState, p))
			if connErr != nil {
				return nil
			}
		case <-pingPongChan:
			return nil
		}
	}
}

func webSocketPingPongs(conn *websocket.Conn) chan struct{} {
	res := make(chan struct{})

	// Pong reader
	go func() {
		conn.SetReadDeadline(time.Now().Add(pongWait))
		conn.SetPongHandler(func(string) error {
			conn.SetReadDeadline(time.Now().Add(pongWait))
			return nil
		})

		for {
			_, _, err := conn.ReadMessage()
			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					log.Printf("Unexpected websocket close: %v", err)
				}
				break
			}
		}

		conn.Close()
	}()

	// Ping writer
	go func() {
		pingTicker := time.NewTicker(pingPeriod)

		for {
			<-pingTicker.C
			conn.SetWriteDeadline(time.Now().Add(writeWait))
			err := conn.WriteMessage(websocket.PingMessage, nil)
			if err != nil {
				break
			}
		}

		conn.Close()
		pingTicker.Stop()
		res <- struct{}{}
	}()

	return res
}

func newGameStateResponse(gs *cah.GameState, player *cah.Player) *gameStateResponse {
	return &gameStateResponse{
		ID:              gs.ID,
		Phase:           gs.Phase.String(),
		Players:         playersInfoFromGame(gs),
		CurrCzarID:      gs.Players[gs.CurrCzarIndex].User.ID,
		BlackCardInPlay: *gs.BlackCardInPlay,
		BlackCardsLeft:  len(gs.BlackDeck),
		WhiteCardsLeft:  len(gs.WhiteDeck),
		SinnerPlays:     sinnerPlaysFromGame(gs),
		MyPlayer:        newFullPlayerInfo(*player),
		CurrRound:       gs.CurrRound,
		MaxRounds:       gs.MaxRounds,
	}
}

func playersInfoFromGame(gs *cah.GameState) []playerInfo {
	ret := make([]playerInfo, len(gs.Players))
	for i, p := range gs.Players {
		ret[i] = newPlayerInfo(*p)
	}
	return ret
}

func newPlayerInfo(p cah.Player) playerInfo {
	return playerInfo{
		ID:               p.User.ID,
		Name:             p.User.Username,
		HandSize:         len(p.Hand),
		WhiteCardsInPlay: len(p.WhiteCardsInPlay),
		Points:           dereferenceBlackCards(p.Points),
	}
}

func newFullPlayerInfo(player cah.Player) fullPlayerInfo {
	return fullPlayerInfo{
		ID:               player.User.ID,
		Name:             player.User.Username,
		Hand:             dereferenceWhiteCards(player.Hand),
		WhiteCardsInPlay: dereferenceWhiteCards(player.WhiteCardsInPlay),
	}
}

func sinnerPlaysFromGame(gs *cah.GameState) []sinnerPlay {
	if !usecase.GameState.AllSinnersPlayedTheirCards(gs) {
		return []sinnerPlay{}
	}
	var ret []sinnerPlay
	for i, p := range gs.Players {
		if gs.IsCurrCzar(p.User) {
			i--
			continue
		}
		ret = append(ret, sinnerPlay{
			ID:         p.User.ID,
			WhiteCards: dereferenceWhiteCards(p.WhiteCardsInPlay),
		})
	}
	rand.Shuffle(len(ret), func(i, j int) {
		ret[i], ret[j] = ret[j], ret[i]
	})
	return ret
}

/*
CHOOSE WINNER
*/

type chooseWinnerPayload struct {
	Winner int `json:"winner"`
}

func chooseWinner(game cah.Game, user cah.User, w http.ResponseWriter, req *http.Request) error {
	// Decode user's payload
	var payload chooseWinnerPayload
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&payload)
	if err != nil {
		return errors.New("Misconstructed payload")
	}
	gs, err := usecase.GameState.ByID(game.StateID)
	if err != nil {
		return err
	}
	pid, err := playerIndex(gs, user)
	if err != nil {
		return err
	}
	if pid != gs.CurrCzarIndex {
		return errors.New("Only the Czar can choose the winner")
	}
	err = usecase.GameState.GiveBlackCardToWinner(payload.Winner, gs)
	if err != nil {
		return err
	}
	gameStateUpdated(gs)
	return nil
}

/*
PLAY CARDS
*/

type playCardsPayload struct {
	CardIndexes []int `json:"cardIndexes"`
}

func playCards(game cah.Game, user cah.User, w http.ResponseWriter, req *http.Request) error {
	// Decode user's payload
	var payload playCardsPayload
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&payload)
	if err != nil {
		return errors.New("Misconstructed payload")
	}
	gs, err := usecase.GameState.ByID(game.StateID)
	if err != nil {
		return err
	}
	pid, err := playerIndex(gs, user)
	if err != nil {
		return err
	}
	err = usecase.GameState.PlayWhiteCards(pid, payload.CardIndexes, gs)
	if err != nil {
		return err
	}
	gameStateUpdated(gs)
	return nil
}

// Utils

func playerIndex(g *cah.GameState, u cah.User) (int, error) {
	for i, p := range g.Players {
		if p.User.ID == u.ID {
			return i, nil
		}
	}
	return -1, errors.New("You are not playing this game")
}

func player(g *cah.GameState, u cah.User) (*cah.Player, error) {
	i, err := playerIndex(g, u)
	if err != nil {
		return &cah.Player{}, errors.New("You are not playing this game")
	}
	return g.Players[i], nil
}

package cah

type GameStateStore interface {
	Create(*GameState) (*GameState, error)
	ByID(id int) (*GameState, error)
	Update(*GameState) error
}

type GameStateUsecases interface {
	ByID(id int) (*GameState, error)
	//FetchOpen() []Game
	Create() *GameState
	GiveBlackCardToWinner(wID int, g *GameState) error
	PlayWhiteCards(p int, cs []int, g *GameState) error
	AllSinnersPlayedTheirCards(g *GameState) bool
	End(g *GameState) error
	PlayRandomWhiteCards(p int, g *GameState) error
}

type GameState struct {
	ID              int          `json:"id" db:"id"`
	Phase           Phase        `json:"phase" db:"phase"`
	Players         []*Player    `json:"players" db:"players"`
	BlackDeck       []*BlackCard `json:"blackDeck" db:"blackDeck"`
	WhiteDeck       []*WhiteCard `json:"whiteDeck" db:"whiteDeck"`
	CurrCzarIndex   int          `json:"currentCzarIndex" db:"currentCzarIndex"`
	BlackCardInPlay *BlackCard   `json:"blackCardInPlay" db:"blackCardInPlay"`
	DiscardPile     []*WhiteCard `json:"discardPile" db:"discardPile"`
	HandSize        int          `json:"handSize" db:"handSize"`
	CurrRound       int          `json:"-" db:"currRound"`
	MaxRounds       int          `json:"-" db:"maxRounds"`
}

func (s *GameState) DrawWhite() *WhiteCard {
	ret := s.WhiteDeck[0]
	s.WhiteDeck = s.WhiteDeck[1:]
	return ret
}

func (s GameState) CurrCzar() *Player {
	return s.Players[s.CurrCzarIndex]
}

func (s GameState) IsCurrCzar(u User) bool {
	return s.CurrCzar().User.ID == u.ID
}

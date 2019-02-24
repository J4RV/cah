package cah

type GameStateStore interface {
	Create(GameState) (GameState, error)
	ByID(id int) (GameState, error)
	Update(GameState) error
}

type GameStateUsecases interface {
	ByID(id int) (GameState, error)
	//FetchOpen() []Game
	Create() GameState
	GiveBlackCardToWinner(wId int, g GameState) (GameState, error)
	PlayWhiteCards(p int, cs []int, g GameState) (GameState, error)
	AllSinnersPlayedTheirCards(g GameState) bool
	End(g GameState) (GameState, error)
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

func (s GameState) Equal(other GameState) bool {
	// First we check very identifiable fields like ID or names
	if s.ID != other.ID {
		return false
	}
	// Fast comparisons before checking lists
	equal := s.Phase == other.Phase
	equal = equal && s.CurrCzarIndex == other.CurrCzarIndex
	equal = equal && s.BlackCardInPlay == other.BlackCardInPlay
	equal = equal && s.HandSize == other.HandSize
	if !equal {
		return false
	}
	// Now check lists one by one
	for i := range s.Players {
		if s.Players[i] != other.Players[i] {
			return false
		}
	}
	for i := range s.BlackDeck {
		if s.BlackDeck[i] != other.BlackDeck[i] {
			return false
		}
	}
	for i := range s.WhiteDeck {
		if s.WhiteDeck[i] != other.WhiteDeck[i] {
			return false
		}
	}
	for i := range s.DiscardPile {
		if s.DiscardPile[i] != other.DiscardPile[i] {
			return false
		}
	}
	return true
}

func (s GameState) Clone() GameState {
	res := GameState{
		ID:              s.ID,
		Phase:           s.Phase,
		CurrCzarIndex:   s.CurrCzarIndex,
		BlackCardInPlay: s.BlackCardInPlay,
		HandSize:        s.HandSize,
		Players:         make([]*Player, len(s.Players)),
		BlackDeck:       make([]*BlackCard, len(s.BlackDeck)),
		WhiteDeck:       make([]*WhiteCard, len(s.WhiteDeck)),
		DiscardPile:     make([]*WhiteCard, len(s.DiscardPile)),
	}
	copy(res.Players, s.Players)
	copy(res.BlackDeck, s.BlackDeck)
	copy(res.WhiteDeck, s.WhiteDeck)
	copy(res.DiscardPile, s.DiscardPile)
	return res
}

package cah

type GameStore interface {
	Create(Game) error
	ByID(int) (Game, error)
	ByStatePhase(...Phase) []Game
	Update(Game) error
}

type GameUsecases interface {
	Create(owner User, name, pass string) error
	ByID(int) (Game, error)
	AllOpen() []Game
	InProgressForUser(User) []Game
	UserJoins(User, Game) error
	Start(Game, GameState, ...Option) error
	Options() GameOptions
	//Start(gameID int, options ...Option) error
}

type Game struct {
	ID       int
	Owner    User
	UserID   int
	Users    []User `gorm:"many2many:game_users;"`
	Name     string
	Password string
	State    GameState
	StateID  int
}

type GameOptions interface {
	WhiteDeck([]*WhiteCard) Option
	BlackDeck([]*BlackCard) Option
	HandSize(size int) Option
	RandomStartingCzar() Option
}

type Option func(s *GameState)

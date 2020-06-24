package cah

type GameStore interface {
	Create(Game) (Game, error)
	ByID(int) (Game, error)
	ByGameStateID(int) (Game, error)
	Update(Game) error
	ByPhase(started, finished bool) ([]Game, error)
}

type GameStatsStore interface {
	Create(GameStats) (GameStats, error)
	ByGameID(int) (GameStats, error)
}

type GameUsecases interface {
	Create(owner User, name, pass string) (Game, error)
	ByID(int) (Game, error)
	ByGameStateID(int) (Game, error)
	AllOpen() ([]Game, error)
	InProgressForUser(User) ([]Game, error)
	Update(Game) error
	UserJoins(User, Game) error
	UserLeaves(User, Game) error
	Start(Game, *GameState, ...Option) error
	Options() GameOptions
}

type Game struct {
	ID       int
	Owner    User
	Users    []User `gorm:"many2many:game_users;"`
	Name     string
	Password string
	StateID  int
	Started  bool
	Finished bool
}

type GameStats struct {
	ID      int
	GameID  int
	Winners []Winner
}

type Winner struct {
	User   User
	Prizes []BlackCard
}

// GAME STATE OPTIONS

type GameOptions interface {
	WhiteDeck([]*WhiteCard) Option
	BlackDeck([]*BlackCard) Option
	HandSize(size int) Option
	RandomStartingCzar() Option
	MaxRounds(max int) Option
}

type Option func(s *GameState)

package cah

import "io"

type CardStore interface {
	CreateWhite(text, expansion string) error
	CreateBlack(text, expansion string, blanks int) error
	AllWhites() ([]*WhiteCard, error)
	AllBlacks() ([]*BlackCard, error)
	ExpansionWhites(...string) ([]*WhiteCard, error)
	ExpansionBlacks(...string) ([]*BlackCard, error)
	AvailableExpansions() ([]string, error)
}

type CardUsecases interface {
	CreateFromReaders(wdat, bdat io.Reader, expansionName string) error
	CreateFromFolder(folderPath, expansionName string) error
	AllWhites() []*WhiteCard
	AllBlacks() []*BlackCard
	ExpansionWhites(...string) []*WhiteCard
	ExpansionBlacks(...string) []*BlackCard
	AvailableExpansions() []string
}

type WhiteCard struct {
	ID        int    `json:"-" db:"white_card"`
	Text      string `json:"text" db:"text"`
	Expansion string `json:"expansion" db:"expansion"`
}

type BlackCard struct {
	ID        int    `json:"-" db:"black_card"`
	Text      string `json:"text" db:"text"`
	Expansion string `json:"expansion" db:"expansion"`
	Blanks    int    `json:"blanks" db:"blanks"`
}

package cah

import (
	"io"

	"gorm.io/gorm"
)

type CardStore interface {
	CreateWhite(text, expansion string) error
	CreateBlack(text, expansion string, blanks int) error
	AllWhites() ([]*WhiteCard, error)
	AllBlacks() ([]*BlackCard, error)
	WhitesByExpansion(...string) ([]*WhiteCard, error)
	BlacksByExpansion(...string) ([]*BlackCard, error)
	AvailableExpansions() ([]string, error)
}

type CardUsecases interface {
	CreateFromReaders(wdat, bdat io.Reader, expansionName string) error
	CreateFromFolder(folderPath, expansionName string) error
	AllWhites() []*WhiteCard
	AllBlacks() []*BlackCard
	WhitesByExpansion(...string) []*WhiteCard
	BlacksByExpansion(...string) []*BlackCard
	AvailableExpansions() []string
}

type WhiteCard struct {
	gorm.Model
	Text      string `json:"text"`
	Expansion string `json:"expansion"`
}

type BlackCard struct {
	gorm.Model
	ID        int    `json:"-"`
	Text      string `json:"text"`
	Expansion string `json:"expansion"`
	Blanks    int    `json:"blanks"`
}

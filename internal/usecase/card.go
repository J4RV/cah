package usecase

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	cah "github.com/j4rv/cah/internal/model"
)

type cardController struct {
	store cah.CardStore
}

// NewCardUsecase returns a cah.CardUsecases
func NewCardUsecase(uc *cah.Usecases, store cah.CardStore) cah.CardUsecases {
	return &cardController{store: store}
}

func (cc cardController) AllBlacks() []*cah.BlackCard {
	res, err := cc.store.AllBlacks()
	checkErr(err, "cardController.AllBlacks")
	return res
}

func (cc cardController) AllWhites() []*cah.WhiteCard {
	res, err := cc.store.AllWhites()
	checkErr(err, "cardController.AllWhites")
	return res
}

func (cc cardController) WhitesByExpansion(exps ...string) []*cah.WhiteCard {
	res, err := cc.store.WhitesByExpansion(exps...)
	checkErr(err, "cardController.WhitesByExpansion")
	return res
}

func (cc cardController) BlacksByExpansion(exps ...string) []*cah.BlackCard {
	res, err := cc.store.BlacksByExpansion(exps...)
	checkErr(err, "cardController.BlacksByExpansion")
	return res
}

func (cc cardController) AvailableExpansions() []string {
	res, err := cc.store.AvailableExpansions()
	checkErr(err, "cardController.AvailableExpansions")
	return res
}

// CreateFromReaders creates and stores cards from two readers.
// The reader should provide a card per line. A line can contain "\n"s for card line breaks.
// Lines containing only whitespace or starting with "#" are ignored
func (cc cardController) CreateFromReaders(wdat, bdat io.Reader, expansionName string) error {
	// Create cards from files
	var err error
	err = doEveryLine(wdat, func(t string) {
		text := strings.TrimSpace(t)
		if text == "" || string([]rune(text)[0]) == "#" {
			return
		}
		cc.store.CreateWhite(text, expansionName)
	})
	if err != nil {
		return err
	}
	err = doEveryLine(bdat, func(t string) {
		text := strings.TrimSpace(t)
		if text == "" || string([]rune(text)[0]) == "#" {
			return
		}
		blanks := strings.Count(t, "_")
		if blanks == 0 {
			blanks = 1
		}
		cc.store.CreateBlack(text, expansionName, blanks)
	})
	log.Println("Successfully loaded cards from expansion " + expansionName)
	return err
}

// CreateFromFolder creates and stores cards from an expansion folder
// That folder should contain two files called 'white.md' and 'black.md'
// The files content is treated as explained for the CreateCards function
func (cc cardController) CreateFromFolder(folderPath, expansionName string) error {
	wdat, err := os.Open(fmt.Sprintf("%s/white.md", folderPath))
	defer wdat.Close()
	if err != nil {
		return err
	}
	bdat, err := os.Open(fmt.Sprintf("%s/black.md", folderPath))
	defer bdat.Close()
	if err != nil {
		return err
	}
	return cc.CreateFromReaders(wdat, bdat, expansionName)
}

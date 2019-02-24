package usecase

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/j4rv/golang-stuff/cah"
)

type cardController struct {
	store cah.CardStore
}

func NewCardUsecase(store cah.CardStore) *cardController {
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

func (cc cardController) ExpansionWhites(exps ...string) []*cah.WhiteCard {
	res, err := cc.store.ExpansionWhites(exps...)
	checkErr(err, "cardController.ExpansionWhites")
	return res
}

func (cc cardController) ExpansionBlacks(exps ...string) []*cah.BlackCard {
	res, err := cc.store.ExpansionBlacks(exps...)
	checkErr(err, "cardController.ExpansionBlacks")
	return res
}

func (cc cardController) AvailableExpansions() []string {
	res, err := cc.store.AvailableExpansions()
	checkErr(err, "cardController.AvailableExpansions")
	return res
}

// CreateFromReaders creates and stores cards from two readers.
// The reader should provide a card per line. A line can contain "\n"s for card line breaks.
// Lines containing only whitespace are ignored
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

func doEveryLine(r io.Reader, fun func(string)) error {
	s := bufio.NewScanner(r)
	for s.Scan() {
		t := strings.Replace(s.Text(), "\\n", "\n", -1)
		fun(t)
	}
	return s.Err()
}

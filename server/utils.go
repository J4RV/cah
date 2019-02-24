package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/j4rv/golang-stuff/cah"
)

func writeResponse(w http.ResponseWriter, obj interface{}) {
	j, err := json.Marshal(obj)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "%s", j)
	}
}

func dereferenceWhiteCards(wcs []*cah.WhiteCard) []cah.WhiteCard {
	res := make([]cah.WhiteCard, len(wcs))
	for i, wc := range wcs {
		res[i] = *wc
	}
	return res
}

func dereferenceBlackCards(bcs []*cah.BlackCard) []cah.BlackCard {
	res := make([]cah.BlackCard, len(bcs))
	for i, bc := range bcs {
		res[i] = *bc
	}
	return res
}

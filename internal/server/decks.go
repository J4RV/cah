package server

import (
	"net/http"

	"github.com/gorilla/mux"
	cah "github.com/j4rv/cah/internal/model"
)

type tabletopDeckData struct {
	Blacks []cah.BlackCard
	Whites []cah.WhiteCard
}

func tabletopDeck(w http.ResponseWriter, req *http.Request) {
	expansion := mux.Vars(req)["expansion"]
	blacks := usecase.Card.BlacksByExpansion(expansion)
	whites := usecase.Card.WhitesByExpansion(expansion)
	execTemplate(tabletopDeckTmpl, w, http.StatusOK, tabletopDeckData{blacks, whites})
}

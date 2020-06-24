package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	cah "github.com/j4rv/cah/internal/model"
)

const missingRequiredParamsMsg = "Missing required parameters"

var errMissingRequiredParam = errors.New("missing required param")

func writeJSONResponse(w http.ResponseWriter, obj interface{}) {
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

func requiredFormParams(params ...[]string) error {
	for _, param := range params {
		if len(param) == 0 {
			return errMissingRequiredParam
		}
	}
	return nil
}

func optionalSingleFormParam(param []string) string {
	if len(param) == 0 {
		return ""
	}
	return param[0]
}

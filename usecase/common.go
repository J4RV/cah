package usecase

import "log"

func checkErr(err error, context string) {
	if err != nil {
		log.Printf("ERROR %s: %s", context, err)
	}
}

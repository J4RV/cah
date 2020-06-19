package usecase

import (
	"bufio"
	"io"
	"log"
	"strings"
)

func checkErr(err error, context string) {
	if err != nil {
		log.Printf("ERROR %s: %s", context, err)
	}
}

func doEveryLine(r io.Reader, fun func(string)) error {
	s := bufio.NewScanner(r)
	for s.Scan() {
		t := strings.Replace(s.Text(), "\\n", "\n", -1)
		fun(t)
	}
	return s.Err()
}

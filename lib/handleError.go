package lib

import (
	"log"
)

func HandleError(err error, msg string) {
    if err != nil {
        log.Fatalf("%s: %v", msg, err)
    } else {
		return;
	}
}

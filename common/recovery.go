package common

import (
	"log"
)

func Recovery() {
	if r := recover(); r != nil {
		log.Println("Recover", r)
	}
}

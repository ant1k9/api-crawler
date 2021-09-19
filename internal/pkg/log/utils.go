package log

import "log"

func FatalIfErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Info(msg string) {
	log.Println(msg)
}

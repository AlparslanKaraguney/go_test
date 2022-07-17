package logger

import (
	"log"
	"os"
	"time"
)

func Log(msg string) {
	// log to custom file
	LOG_FILE := "logger/" + time.Now().Format("2006-01-02") + ".log"

	// open log file
	logFile, err := os.OpenFile(LOG_FILE, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)

	if err != nil {
		log.Panic(err)
	}
	defer logFile.Close()

	customLog := log.New(logFile, "", log.LstdFlags|log.Lshortfile)

	customLog.Println(msg)

}

package main

import (
	"github.com/sathishkumar-manogaran/publisher-in-golang/message"
	log "github.com/sirupsen/logrus"
	"os"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.TraceLevel)
}

func main() {

	message.InitChannel()
	message.Publisher()

	//defer database.DBCon.Close()
	defer message.Channel.Close()
}

package cron

import (
	bc "application/blockchain"
	"github.com/robfig/cron/v3"
	"log"
)

const spec = "0 0 0 * * ?" // Executed daily at 0:00
//const spec = "*/10 * * * * ?" // Executed once every 10 seconds for testin

func Init() {
	c := cron.New(cron.WithSeconds()) // Support to the second level
	_, err := c.AddFunc(spec, GoRun)
	if err != nil {
		log.Printf("Scheduled task fails to open %s", err)
	}
	c.Start()
	log.Printf("Scheduled task has been opened")
	select {}
}

func GoRun() {
	log.Printf("Scheduled task has been opened")
	_, err := bc.ChannelQuery("querySequenceList", [][]byte{}) // invoke smart contract
	if err != nil {
		log.Printf("Scheduled task-querySequenceList failed%s", err.Error())
		return
	}
}

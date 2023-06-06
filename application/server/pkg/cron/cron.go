package cron

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"time"

	bc "application/blockchain"
	"application/model"

	"github.com/robfig/cron/v3"
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
	// query all authorizing records
	resp, err := bc.ChannelQuery("querySellingList", [][]byte{}) // invoke smart contract
	if err != nil {
		log.Printf("Scheduled task-queryAuthorizingList failed%s", err.Error())
		return
	}
	// deserialize json
	var data []model.Authorizing
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		log.Printf("Scheduled task-deserialize json failed%s", err.Error())
		return
	}
	for _, v := range data {
		// Filter out those with status of In-Authorizing and In-Delivery
		if v.AuthorizingStatus == model.AuthorizationStatusConstant()["publish"] ||
			v.AuthorizingStatus == model.AuthorizationStatusConstant()["delivery"] {
			// valid days
			day, _ := time.ParseDuration(fmt.Sprintf("%dh", v.AuthorizePeriod*24))
			local, _ := time.LoadLocation("Local")
			t, _ := time.ParseInLocation("2006-01-02 15:04:05", v.CreateTime, local)
			vTime := t.Add(day)
			// if time.Now() > vTime  expired
			if time.Now().Local().After(vTime) {
				// change the status to expired
				var bodyBytes [][]byte
				bodyBytes = append(bodyBytes, []byte(v.ObjectOfAuthorize))
				bodyBytes = append(bodyBytes, []byte(v.Hospital))
				bodyBytes = append(bodyBytes, []byte(v.Patient))
				bodyBytes = append(bodyBytes, []byte("expired"))
				// invoke smart contract
				resp, err := bc.ChannelExecute("updateAuthorizing", bodyBytes)
				if err != nil {
					return
				}
				var data map[string]interface{}
				if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
					return
				}
				fmt.Println(data)
			}
		}
	}
}

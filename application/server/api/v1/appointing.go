package v1

import (
	bc "application/blockchain"
	"application/pkg/app"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AppointingRequestBody struct {
	ObjectOfAppointing string `json:"objectOfAppointing"` // Object Of Appointing
	Patient            string `json:"patient"`            // patient appointed
	Hospital           string `json:"hospital"`           // hospital appointed
}

type AppointingListQueryRequestBody struct {
	Patient string `json:"patient"`
}

type AppointingListQueryByHospitalRequestBody struct {
	Hospital string `json:"hospital"`
}

type UpdateAppointingRequestBody struct {
	ObjectOfAppointing string `json:"objectOfAppointing"` // Object Of Appointing
	Patient            string `json:"patient"`            // patient appointed
	Hospital           string `json:"hospital"`           // hospital appointed
	AppointingStatus   string `json:"status"`             // the status need to be changed
}

func CreateAppointing(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(AppointingRequestBody)
	//parse Body parameter
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "failed", fmt.Sprintf("parameter error%s", err.Error()))
		return
	}
	if body.ObjectOfAppointing == "" || body.Patient == "" || body.Hospital == "" {
		appG.Response(http.StatusBadRequest, "failed", "ObjectOfAppointing, patient appointed and hospital appointed can not be empty")
		return
	}
	var bodyBytes [][]byte
	bodyBytes = append(bodyBytes, []byte(body.ObjectOfAppointing))
	bodyBytes = append(bodyBytes, []byte(body.Patient))
	bodyBytes = append(bodyBytes, []byte(body.Hospital))
	// invoke smart contract
	resp, err := bc.ChannelExecute("createAppointing", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "failed", err.Error())
		return
	}
	var data map[string]interface{}
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		appG.Response(http.StatusInternalServerError, "failed", err.Error())
		return
	}
	appG.Response(http.StatusOK, "success", data)
}

func QueryAppointingList(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(AppointingListQueryRequestBody)
	//parse Body parameter
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "failed", fmt.Sprintf("parameter error%s", err.Error()))
		return
	}
	var bodyBytes [][]byte
	if body.Patient != "" {
		bodyBytes = append(bodyBytes, []byte(body.Patient))
	}
	// invoke smart contract
	resp, err := bc.ChannelQuery("queryAppointingList", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "failed", err.Error())
		return
	}
	// deserialize json
	var data []map[string]interface{}
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		appG.Response(http.StatusInternalServerError, "failed", err.Error())
		return
	}
	appG.Response(http.StatusOK, "success", data)
}

func QueryAppointingListByHospital(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(AppointingListQueryByHospitalRequestBody)
	// parse Body parameter
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "failed", fmt.Sprintf("parameter error%s", err.Error()))
		return
	}
	if body.Hospital == "" {
		appG.Response(http.StatusBadRequest, "failed", "must use AccountId to query")
		return
	}
	var bodyBytes [][]byte
	bodyBytes = append(bodyBytes, []byte(body.Hospital))
	// invoke smart contract
	resp, err := bc.ChannelQuery("queryAppointingListByHospital", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "failed", err.Error())
		return
	}
	// deserialize json
	var data []map[string]interface{}
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		appG.Response(http.StatusInternalServerError, "failed", err.Error())
		return
	}
	appG.Response(http.StatusOK, "success", data)
}

func UpdateAppointing(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(UpdateAppointingRequestBody)
	// parse Body parameter
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "failed", fmt.Sprintf("parameter error%s", err.Error()))
		return
	}
	if body.ObjectOfAppointing == "" || body.Patient == "" || body.Hospital == "" || body.AppointingStatus == "" {
		appG.Response(http.StatusBadRequest, "failed", "parameter can not be empty")
		return
	}
	var bodyBytes [][]byte
	bodyBytes = append(bodyBytes, []byte(body.ObjectOfAppointing))
	bodyBytes = append(bodyBytes, []byte(body.Patient))
	bodyBytes = append(bodyBytes, []byte(body.Hospital))
	bodyBytes = append(bodyBytes, []byte(body.AppointingStatus))
	// invoke smart contract
	resp, err := bc.ChannelExecute("updateAppointing", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "failed", err.Error())
		return
	}
	var data map[string]interface{}
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		appG.Response(http.StatusInternalServerError, "failed", err.Error())
		return
	}
	appG.Response(http.StatusOK, "success", data)
}

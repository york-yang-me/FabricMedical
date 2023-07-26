package v1

import (
	bc "application/blockchain"
	"application/pkg/app"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AuthorizingRequestBody struct {
	ObjectOfAuthorize string  `json:"objectOfAuthorize"` // Authorize DNA test object (DNA in Authorization:RealSequenceID)
	Hospital          string  `json:"hospital"`          // Hospitals that preserve dna(Hospital AccountId)
	Price             float64 `json:"price"`             // Authorizing price
	AuthorizePeriod   int     `json:"authorizePeriod"`   // the validity of the smart contract(in days)
}

type AuthorizingByBuyRequestBody struct {
	ObjectOfAuthorize string `json:"objectOfAuthorize"` // Authorize DNA test object (DNA in Authorization:RealSequenceID)
	Hospital          string `json:"hospital"`          // Hospitals that preserve dna(Hospital AccountId)
	Patient           string `json:"patient"`           // Patients involved in testing DNA data(Patient AccountId)
}

type AuthorizingListQueryRequestBody struct {
	Hospital string `json:"hospital"` // Hospitals that preserve dna(Hospital AccountId)
}

type AuthorizingListQueryByBuyRequestBody struct {
	Patient string `json:"patient"` // Patients involved in testing DNA data(Patient AccountId)
}

type UpdateAuthorizingRequestBody struct {
	ObjectOfAuthorize string `json:"objectOfAuthorize"` // Authorize DNA test object (DNA in Authorization:RealSequenceID)
	Hospital          string `json:"hospital"`          // Hospitals that preserve dna(Hospital AccountId)
	Patient           string `json:"patient"`           // Patients involved in testing DNA data(Patient AccountId)
	AuthorizingStatus string `json:"authorizingStatus"` // authorize status
}

func CreateSelling(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(AuthorizingRequestBody)
	// parse Body parameter
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "failed", fmt.Sprintf("parameter error%s", err.Error()))
		return
	}
	if body.ObjectOfAuthorize == "" || body.Hospital == "" {
		appG.Response(http.StatusBadRequest, "failed", "ObjectOfAuthorize and Hospital can not be empty")
		return
	}
	if body.Price <= 0 || body.AuthorizePeriod <= 0 {
		appG.Response(http.StatusBadRequest, "failed", "Price and AuthorizePeriod smart contact must be valid>0")
		return
	}
	var bodyBytes [][]byte
	bodyBytes = append(bodyBytes, []byte(body.ObjectOfAuthorize))
	bodyBytes = append(bodyBytes, []byte(body.Hospital))
	bodyBytes = append(bodyBytes, []byte(strconv.FormatFloat(body.Price, 'E', -1, 64)))
	bodyBytes = append(bodyBytes, []byte(strconv.Itoa(body.AuthorizePeriod)))
	// invoke smart contract
	resp, err := bc.ChannelExecute("createAuthorizing", bodyBytes)
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

func CreateAuthorizingByBuy(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(AuthorizingByBuyRequestBody)
	// invoke smart contract
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "failed", fmt.Sprintf("parameter failed%s", err.Error()))
		return
	}
	if body.ObjectOfAuthorize == "" || body.Hospital == "" || body.Patient == "" {
		appG.Response(http.StatusBadRequest, "failed", "parameter can not be empty")
		return
	}
	var bodyBytes [][]byte
	bodyBytes = append(bodyBytes, []byte(body.ObjectOfAuthorize))
	bodyBytes = append(bodyBytes, []byte(body.Hospital))
	bodyBytes = append(bodyBytes, []byte(body.Patient))
	// invoke smart contract
	resp, err := bc.ChannelExecute("createAuthorizingByBuy", bodyBytes)
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

func QueryAuthorizingList(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(AuthorizingListQueryRequestBody)
	// invoke smart contract
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "failed", fmt.Sprintf("parameter error%s", err.Error()))
		return
	}
	var bodyBytes [][]byte
	if body.Hospital != "" {
		bodyBytes = append(bodyBytes, []byte(body.Hospital))
	}
	// invoke smart contract
	resp, err := bc.ChannelQuery("queryAuthorizingList", bodyBytes)
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

func QueryAuthorizingListByBuyer(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(AuthorizingListQueryByBuyRequestBody)
	// parse Body parameter
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "failed", fmt.Sprintf("parameter error%s", err.Error()))
		return
	}
	if body.Patient == "" {
		appG.Response(http.StatusBadRequest, "failed", "must appoint the Patients' AccountId to query")
		return
	}
	var bodyBytes [][]byte
	bodyBytes = append(bodyBytes, []byte(body.Patient))
	// invoke smart contract
	resp, err := bc.ChannelQuery("queryAuthorizingListByBuyer", bodyBytes)
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

func UpdateAuthorizing(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(UpdateAuthorizingRequestBody)
	// parse Body parameter
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "failed", fmt.Sprintf("parameter error%s", err.Error()))
		return
	}
	if body.ObjectOfAuthorize == "" || body.Hospital == "" || body.AuthorizingStatus == "" {
		appG.Response(http.StatusBadRequest, "failed", "parameter can not be empty")
		return
	}
	var bodyBytes [][]byte
	bodyBytes = append(bodyBytes, []byte(body.ObjectOfAuthorize))
	bodyBytes = append(bodyBytes, []byte(body.Hospital))
	bodyBytes = append(bodyBytes, []byte(body.Patient))
	bodyBytes = append(bodyBytes, []byte(body.AuthorizingStatus))
	// invoke smart contract
	resp, err := bc.ChannelExecute("updateAuthorizing", bodyBytes)
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

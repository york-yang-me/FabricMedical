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

type RealSequenceRequestBody struct {
	RealSequenceID string  `json:"realSequenceID"` // DNA Sequence ID
	Owner          string  `json:"owner"`          // Owner(DNA Holders)(Owner AccountId)
	TotalLength    float64 `json:"totalLength"`    // total length
	DNAContents    float64 `json:"dnaContents"`    // DNA contents
}

type RealSequenceQueryRequestBody struct {
	Owner string `json:"owner"` // Owner(DNA Holders)(Owner AccountId)
}

func CreateRealSequence(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(RealSequenceRequestBody)
	// parse Body parameter
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "failed", fmt.Sprintf("parameter error%s", err.Error()))
		return
	}
	if body.TotalLength <= 0 || body.DNAContents <= 0 || body.DNAContents > body.TotalLength {
		appG.Response(http.StatusBadRequest, "failed", "TotalLength ,DNA contents must be bigger than 0，and DNA contents < total length")
		return
	}
	var bodyBytes [][]byte
	bodyBytes = append(bodyBytes, []byte(body.RealSequenceID))
	bodyBytes = append(bodyBytes, []byte(body.Owner))
	bodyBytes = append(bodyBytes, []byte(strconv.FormatFloat(body.TotalLength, 'E', -1, 64)))
	bodyBytes = append(bodyBytes, []byte(strconv.FormatFloat(body.DNAContents, 'E', -1, 64)))
	// invoke smart contract
	resp, err := bc.ChannelExecute("createRealSequence", bodyBytes)
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

func QueryRealSequenceList(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(RealSequenceQueryRequestBody)
	// parse Body parameter
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "failed", fmt.Sprintf("parameter error%s", err.Error()))
		return
	}
	var bodyBytes [][]byte
	if body.Owner != "" {
		bodyBytes = append(bodyBytes, []byte(body.Owner))
	}
	// invoke smart contract
	resp, err := bc.ChannelQuery("queryRealSequenceList", bodyBytes)
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

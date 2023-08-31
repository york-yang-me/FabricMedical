package v1

import (
	bc "application/blockchain"
	"application/pkg/app"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type RealSequenceRequestBody struct {
	RealSequenceID string `json:"realSequenceID"` // DNA Sequence ID
	Owner          string `json:"owner"`          // Owner(DNA Holders)(Owner AccountId)
	TotalLength    int    `json:"totalLength"`    // total length
	DNAContents    string `json:"dnaContents"`    // DNA contents
	Description    string `json:"description"`    // DNA description
}

type RealSequenceQueryRequestBody struct {
	Owner string `json:"owner"` // Owner(DNA Holders)(Owner AccountId)
}

type RealSequenceUpdateRequestBody struct {
	Owner       string `json:"owner"`       // Owner(DNA Holders)(Owner AccountId)
	DNAContents string `json:"dnaContents"` // DNA contents
	Description string `json:"description"` // DNA description
	Proof       string `json:"proof"`       // proof for zk-snarks
}

func CreateRealSequence(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(RealSequenceRequestBody)
	// parse Body parameter
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "failed", fmt.Sprintf("parameter error%s", err.Error()))
		return
	}
	if body.TotalLength <= 0 || len(body.DNAContents) <= 0 || len(body.DNAContents) > body.TotalLength || len(body.Description) <= 0 {
		appG.Response(http.StatusBadRequest, "failed", "TotalLength, DNA contents, Description must be bigger than 0，and DNA contents < total length")
		return
	}

	SequenceHash := HashCalc(body.DNAContents)
	vk, proof, err := Generate(body.DNAContents, SequenceHash)

	var bodyBytes [][]byte
	bodyBytes = append(bodyBytes, []byte(body.RealSequenceID))
	bodyBytes = append(bodyBytes, []byte(body.Owner))
	bodyBytes = append(bodyBytes, []byte(strconv.Itoa(body.TotalLength)))
	bodyBytes = append(bodyBytes, []byte(SequenceHash))
	bodyBytes = append(bodyBytes, []byte(body.Description))
	bodyBytes = append(bodyBytes, []byte(vk))
	bodyBytes = append(bodyBytes, []byte(proof))

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

func ModifyRealSequenceList(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(RealSequenceUpdateRequestBody)
	// parse Body parameter
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "failed", fmt.Sprintf("parameter error%s", err.Error()))
		return
	}
	var bodyBytes [][]byte
	bodyBytes = append(bodyBytes, []byte(body.Owner))
	bodyBytes = append(bodyBytes, []byte(body.DNAContents))
	bodyBytes = append(bodyBytes, []byte(body.Description))
	bodyBytes = append(bodyBytes, []byte(body.Proof))

	// invoke smart contract
	resp, err := bc.ChannelExecute("updateRealSequence", bodyBytes)
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

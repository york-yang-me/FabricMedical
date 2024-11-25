package model

// Account virtual admin and many users
type Account struct {
	AccountId string  `json:"accountId"` // Account ID
	UserName  string  `json:"userName"`  // Account Name
	Balance   float64 `json:"balance"`   // Account Balance
}

// RealSequence Endorsement is true when dna sequence is checked, verified or valid as security, default status:false
// Only when Endorsement is false, dna sequence can be transfer
// Owner and RealSequenceID as composite key, guaranteed access to DNA sequence which belongs to themselves through 'Owner'
type RealSequence struct {
	RealSequenceID string  `json:"realSequenceId"` // DNA Sequence ID
	Owner          string  `json:"owner"`          // Owner(DNA Holders)(Owner AccountId)
	Endorsement    bool    `json:"endorsement"`    // whether as an endorser
	TotalLength    int     `json:"totalLength"`    // total length
	DNAContents    float64 `json:"dnaContents"`    // DNA contents
}

// Authorizing Authorize some patients to DNA testing smart contract
// Need to confirm if objectOfAuthorize belongs to Hospital
// Patient is initially empty
// Hospital and ObjectOfTesting as composite key, guaranteed access to all DNA sequence which stores in hospital through 'Hospital'
type Authorizing struct {
	ObjectOfAuthorize string  `json:"objectOfAuthorize"` // Authorize DNA test object (DNA in Authorization:RealSequenceID)
	Hospital          string  `json:"hospital"`          // Hospitals that preserve dna(Hospital AccountId)
	Patient           string  `json:"patient"`           // Patients involved in testing DNA data(Patient AccountId)
	Price             float64 `json:"price"`             // Authorizing price
	CreateTime        string  `json:"createTime"`        // Create time
	AuthorizePeriod   int     `json:"authorizePeriod"`   // the validity of the smart contract(in days)
	AuthorizingStatus string  `json:"authorizingStatus"` // authorize status
}

// AuthorizationStatusConstant Authorization Status
var AuthorizationStatusConstant = func() map[string]string {
	return map[string]string{
		"publish":   "Publish",     // publish, wait for Patient continue to deal with it
		"cancelled": "Cancelled",   // cancellation by the hospital to cancel the public or the patient refund due to cancellation
		"expired":   "Expired",     // public expiration
		"delivery":  "In Delivery", // Patient use some research fund to pay for this research, if the hospital has not accepted the delegation, patient can cancel it
		"done":      "Finish",      // Hospital confirm accepting research fund, transaction finished
	}
}

// Appointing appoint patients to research DNA sequence
// Authorization object can not be patients
// Patient and CreateTime as composite key, guaranteed access to all DNA sequence which has been appointed for patients
type Appointing struct {
	Patient     string      `json:"patient"`     // patient appointed(Patient AccountId)
	CreateTime  string      `json:"createTime"`  // create time
	Authorizing Authorizing `json:"authorizing"` // authorizing object
}

const (
	AccountKey      = "account-key"
	RealSequenceKey = "real-sequence-key"
	AuthorizingKey  = "authorizing-key"
	AppointingKey   = "appointing-key"
)

package model

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

package model

// Account virtual admin and many users
type Account struct {
	AccountId string  `json:"accountId"` // Account ID
	UserName  string  `json:"userName"`  // Account Name
	Balance   float64 `json:"balance"`   // Account Balance9
}

// RealSequence Endorsement is true when dna sequence is checked, verified or valid as security, default status:false
// Only when Endorsement is false, dna sequence can be transfer
// Owner and RealSequenceID as composite key, guaranteed access to DNA sequence which belongs to themselves through 'Owner'
type RealSequence struct {
	RealSequenceID string `json:"realSequenceId"` // DNA Sequence ID
	Owner          string `json:"owner"`          // Owner(DNA Holders)(Owner AccountId)
	Endorsement    bool   `json:"endorsement"`    // whether as an endorser
	TotalLength    int    `json:"totalLength"`    // total length
	DNAContents    string `json:"dnaContents"`    // DNA contents
}

const (
	AccountKey      = "account-key"
	RealSequenceKey = "real-sequence-key"
)

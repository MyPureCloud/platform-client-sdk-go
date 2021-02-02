package platformclientv2
import (
	"encoding/json"
)

// Disconnectreason
type Disconnectreason struct { 
	// VarType - Disconnect reason protocol type.
	VarType *string `json:"type,omitempty"`


	// Code - Protocol specific reason code. See the Q.850 and SIP specs.
	Code *int `json:"code,omitempty"`


	// Phrase - Human readable English description of the disconnect reason.
	Phrase *string `json:"phrase,omitempty"`

}

// String returns a JSON representation of the model
func (o *Disconnectreason) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

package platformclientv2
import (
	"encoding/json"
)

// Conversationeventtopicdisconnectreason
type Conversationeventtopicdisconnectreason struct { 
	// VarType
	VarType *string `json:"type,omitempty"`


	// Code
	Code *int32 `json:"code,omitempty"`


	// Phrase
	Phrase *string `json:"phrase,omitempty"`

}

// String returns a JSON representation of the model
func (o *Conversationeventtopicdisconnectreason) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

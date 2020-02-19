package platformclientv2
import (
	"encoding/json"
)

// Queueconversationsocialexpressioneventtopicdisconnectreason
type Queueconversationsocialexpressioneventtopicdisconnectreason struct { 
	// VarType
	VarType *string `json:"type,omitempty"`


	// Code
	Code *int32 `json:"code,omitempty"`


	// Phrase
	Phrase *string `json:"phrase,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queueconversationsocialexpressioneventtopicdisconnectreason) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

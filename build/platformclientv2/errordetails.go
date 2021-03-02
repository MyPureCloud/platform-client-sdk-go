package platformclientv2
import (
	"encoding/json"
)

// Errordetails
type Errordetails struct { 
	// Status
	Status *int `json:"status,omitempty"`


	// Message
	Message *string `json:"message,omitempty"`


	// MessageWithParams
	MessageWithParams *string `json:"messageWithParams,omitempty"`


	// MessageParams
	MessageParams *map[string]string `json:"messageParams,omitempty"`


	// Code
	Code *string `json:"code,omitempty"`


	// ContextId
	ContextId *string `json:"contextId,omitempty"`


	// Nested
	Nested **Errordetails `json:"nested,omitempty"`


	// Details
	Details *string `json:"details,omitempty"`

}

// String returns a JSON representation of the model
func (o *Errordetails) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

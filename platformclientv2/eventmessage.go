package platformclientv2
import (
	"encoding/json"
)

// Eventmessage
type Eventmessage struct { 
	// Code
	Code *string `json:"code,omitempty"`


	// Message
	Message *string `json:"message,omitempty"`


	// MessageWithParams
	MessageWithParams *string `json:"messageWithParams,omitempty"`


	// MessageParams
	MessageParams *map[string]interface{} `json:"messageParams,omitempty"`


	// DocumentationUri
	DocumentationUri *string `json:"documentationUri,omitempty"`


	// ResourceURIs
	ResourceURIs *[]string `json:"resourceURIs,omitempty"`

}

// String returns a JSON representation of the model
func (o *Eventmessage) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

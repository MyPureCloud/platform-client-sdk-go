package platformclientv2
import (
	"encoding/json"
)

// Queueconversationvideoeventtopicerrordetails
type Queueconversationvideoeventtopicerrordetails struct { 
	// Status
	Status *int32 `json:"status,omitempty"`


	// Code
	Code *string `json:"code,omitempty"`


	// Message
	Message *string `json:"message,omitempty"`


	// MessageWithParams
	MessageWithParams *string `json:"messageWithParams,omitempty"`


	// MessageParams
	MessageParams *map[string]string `json:"messageParams,omitempty"`


	// ContextId
	ContextId *string `json:"contextId,omitempty"`


	// Uri
	Uri *string `json:"uri,omitempty"`


	// AdditionalProperties
	AdditionalProperties *map[string]interface{} `json:"additionalProperties,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queueconversationvideoeventtopicerrordetails) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

package platformclientv2
import (
	"encoding/json"
)

// Conversationsocialexpressioneventtopicerrorbody
type Conversationsocialexpressioneventtopicerrorbody struct { 
	// Status
	Status *int `json:"status,omitempty"`


	// Code
	Code *string `json:"code,omitempty"`


	// EntityId
	EntityId *string `json:"entityId,omitempty"`


	// EntityName
	EntityName *string `json:"entityName,omitempty"`


	// Message
	Message *string `json:"message,omitempty"`


	// MessageWithParams
	MessageWithParams *string `json:"messageWithParams,omitempty"`


	// MessageParams
	MessageParams *map[string]string `json:"messageParams,omitempty"`


	// ContextId
	ContextId *string `json:"contextId,omitempty"`


	// Details
	Details *[]Conversationsocialexpressioneventtopicdetail `json:"details,omitempty"`


	// Errors
	Errors *[]Conversationsocialexpressioneventtopicerrorbody `json:"errors,omitempty"`

}

// String returns a JSON representation of the model
func (o *Conversationsocialexpressioneventtopicerrorbody) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

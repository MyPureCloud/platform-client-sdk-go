package platformclientv2
import (
	"encoding/json"
)

// Errorbody
type Errorbody struct { 
	// Message
	Message *string `json:"message,omitempty"`


	// Code
	Code *string `json:"code,omitempty"`


	// Status
	Status *int `json:"status,omitempty"`


	// EntityId
	EntityId *string `json:"entityId,omitempty"`


	// EntityName
	EntityName *string `json:"entityName,omitempty"`


	// MessageWithParams
	MessageWithParams *string `json:"messageWithParams,omitempty"`


	// MessageParams
	MessageParams *map[string]string `json:"messageParams,omitempty"`


	// ContextId
	ContextId *string `json:"contextId,omitempty"`


	// Details
	Details *[]Detail `json:"details,omitempty"`


	// Errors
	Errors *[]Errorbody `json:"errors,omitempty"`

}

// String returns a JSON representation of the model
func (o *Errorbody) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

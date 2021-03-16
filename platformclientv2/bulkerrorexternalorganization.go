package platformclientv2
import (
	"encoding/json"
)

// Bulkerrorexternalorganization
type Bulkerrorexternalorganization struct { 
	// Code
	Code *string `json:"code,omitempty"`


	// Message
	Message *string `json:"message,omitempty"`


	// Status
	Status *int `json:"status,omitempty"`


	// Retryable
	Retryable *bool `json:"retryable,omitempty"`


	// Entity
	Entity *Externalorganization `json:"entity,omitempty"`


	// Details
	Details *[]Bulkerrordetail `json:"details,omitempty"`

}

// String returns a JSON representation of the model
func (o *Bulkerrorexternalorganization) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

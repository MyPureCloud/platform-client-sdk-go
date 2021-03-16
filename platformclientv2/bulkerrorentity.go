package platformclientv2
import (
	"encoding/json"
)

// Bulkerrorentity
type Bulkerrorentity struct { 
	// Code
	Code *string `json:"code,omitempty"`


	// Message
	Message *string `json:"message,omitempty"`


	// Status
	Status *int `json:"status,omitempty"`


	// Retryable
	Retryable *bool `json:"retryable,omitempty"`


	// Entity
	Entity *Entity `json:"entity,omitempty"`


	// Details
	Details *[]Bulkerrordetail `json:"details,omitempty"`

}

// String returns a JSON representation of the model
func (o *Bulkerrorentity) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

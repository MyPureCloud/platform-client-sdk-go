package platformclientv2
import (
	"encoding/json"
)

// Presenceexpand
type Presenceexpand struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Presences - An array of user presences
	Presences *[]Userpresence `json:"presences,omitempty"`


	// OutOfOffices - An array of out of office statuses
	OutOfOffices *[]Outofoffice `json:"outOfOffices,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Presenceexpand) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

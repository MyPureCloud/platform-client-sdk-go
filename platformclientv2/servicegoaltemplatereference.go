package platformclientv2
import (
	"encoding/json"
)

// Servicegoaltemplatereference
type Servicegoaltemplatereference struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Servicegoaltemplatereference) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

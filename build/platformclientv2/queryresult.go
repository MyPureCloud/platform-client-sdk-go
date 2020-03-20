package platformclientv2
import (
	"encoding/json"
)

// Queryresult
type Queryresult struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Body
	Body *Domainentity `json:"body,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queryresult) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

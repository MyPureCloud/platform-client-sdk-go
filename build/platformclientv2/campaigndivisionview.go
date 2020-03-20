package platformclientv2
import (
	"encoding/json"
)

// Campaigndivisionview
type Campaigndivisionview struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Division - The division to which this entity belongs.
	Division *Division `json:"division,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Campaigndivisionview) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

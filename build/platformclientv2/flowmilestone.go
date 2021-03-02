package platformclientv2
import (
	"encoding/json"
)

// Flowmilestone
type Flowmilestone struct { 
	// Id - The flow milestone identifier
	Id *string `json:"id,omitempty"`


	// Name - The flow milestone name.
	Name *string `json:"name,omitempty"`


	// Description - The flow milestone description.
	Description *string `json:"description,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Flowmilestone) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

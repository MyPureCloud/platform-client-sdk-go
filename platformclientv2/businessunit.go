package platformclientv2
import (
	"encoding/json"
)

// Businessunit
type Businessunit struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Division - The division to which this entity belongs.
	Division *Division `json:"division,omitempty"`


	// Settings - Settings for this business unit
	Settings *Businessunitsettings `json:"settings,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Businessunit) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

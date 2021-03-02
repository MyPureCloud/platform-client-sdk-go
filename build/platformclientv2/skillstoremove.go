package platformclientv2
import (
	"encoding/json"
)

// Skillstoremove
type Skillstoremove struct { 
	// Name
	Name *string `json:"name,omitempty"`


	// Id
	Id *string `json:"id,omitempty"`


	// SelfUri
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Skillstoremove) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

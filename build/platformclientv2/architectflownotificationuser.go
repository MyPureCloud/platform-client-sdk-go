package platformclientv2
import (
	"encoding/json"
)

// Architectflownotificationuser
type Architectflownotificationuser struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// HomeOrg
	HomeOrg *Architectflownotificationhomeorganization `json:"homeOrg,omitempty"`

}

// String returns a JSON representation of the model
func (o *Architectflownotificationuser) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

package platformclientv2
import (
	"encoding/json"
)

// Architectflowoutcomenotificationuser
type Architectflowoutcomenotificationuser struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// HomeOrg
	HomeOrg *Architectflowoutcomenotificationhomeorganization `json:"homeOrg,omitempty"`

}

// String returns a JSON representation of the model
func (o *Architectflowoutcomenotificationuser) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

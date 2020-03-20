package platformclientv2
import (
	"encoding/json"
)

// Domainrole
type Domainrole struct { 
	// Id - The ID of the role
	Id *string `json:"id,omitempty"`


	// Name - The name of the role
	Name *string `json:"name,omitempty"`

}

// String returns a JSON representation of the model
func (o *Domainrole) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

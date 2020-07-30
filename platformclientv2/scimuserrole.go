package platformclientv2
import (
	"encoding/json"
)

// Scimuserrole - Defines a user role.
type Scimuserrole struct { 
	// Value - The role of the Genesys Cloud user.
	Value *string `json:"value,omitempty"`

}

// String returns a JSON representation of the model
func (o *Scimuserrole) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

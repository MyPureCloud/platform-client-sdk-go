package platformclientv2
import (
	"encoding/json"
)

// Emailaddress
type Emailaddress struct { 
	// Email
	Email *string `json:"email,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`

}

// String returns a JSON representation of the model
func (o *Emailaddress) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

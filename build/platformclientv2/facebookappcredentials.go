package platformclientv2
import (
	"encoding/json"
)

// Facebookappcredentials
type Facebookappcredentials struct { 
	// Id - Genesys Cloud Facebook App Id
	Id *string `json:"id,omitempty"`

}

// String returns a JSON representation of the model
func (o *Facebookappcredentials) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

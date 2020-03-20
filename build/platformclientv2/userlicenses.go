package platformclientv2
import (
	"encoding/json"
)

// Userlicenses
type Userlicenses struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Licenses
	Licenses *[]string `json:"licenses,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Userlicenses) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

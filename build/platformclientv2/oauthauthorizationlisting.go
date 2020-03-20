package platformclientv2
import (
	"encoding/json"
)

// Oauthauthorizationlisting
type Oauthauthorizationlisting struct { 
	// Entities
	Entities *[]Oauthauthorization `json:"entities,omitempty"`


	// SelfUri
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Oauthauthorizationlisting) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

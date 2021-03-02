package platformclientv2
import (
	"encoding/json"
)

// Smsphonenumberref
type Smsphonenumberref struct { 
	// PhoneNumber - A phone number provisioned for SMS communications in E.164 format. E.g. +13175555555 or +34234234234
	PhoneNumber *string `json:"phoneNumber,omitempty"`


	// SelfUri
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Smsphonenumberref) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

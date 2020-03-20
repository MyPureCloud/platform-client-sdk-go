package platformclientv2
import (
	"time"
	"encoding/json"
)

// Callforwarding
type Callforwarding struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// User
	User *User `json:"user,omitempty"`


	// Enabled - Whether or not CallForwarding is enabled
	Enabled *bool `json:"enabled,omitempty"`


	// PhoneNumber - This property is deprecated. Please use the calls property
	PhoneNumber *string `json:"phoneNumber,omitempty"`


	// Calls - An ordered list of CallRoutes to be executed when CallForwarding is enabled
	Calls *[]Callroute `json:"calls,omitempty"`


	// Voicemail - The type of voicemail to use with the callForwarding configuration
	Voicemail *string `json:"voicemail,omitempty"`


	// ModifiedDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	ModifiedDate *time.Time `json:"modifiedDate,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Callforwarding) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

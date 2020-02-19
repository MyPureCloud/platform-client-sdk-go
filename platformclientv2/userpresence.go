package platformclientv2
import (
	"time"
	"encoding/json"
)

// Userpresence
type Userpresence struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Source - Represents the source where the Presence was set. Some examples are: PURECLOUD, LYNC, OUTLOOK, etc.
	Source *string `json:"source,omitempty"`


	// Primary - A boolean used to tell whether or not to set this presence source as the primary on a PATCH
	Primary *bool `json:"primary,omitempty"`


	// PresenceDefinition
	PresenceDefinition *Presencedefinition `json:"presenceDefinition,omitempty"`


	// Message
	Message *string `json:"message,omitempty"`


	// ModifiedDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	ModifiedDate *time.Time `json:"modifiedDate,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Userpresence) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

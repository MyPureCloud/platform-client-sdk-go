package platformclientv2
import (
	"time"
	"encoding/json"
)

// Historyentry
type Historyentry struct { 
	// Action - The action performed
	Action *string `json:"action,omitempty"`


	// Resource - For actions performed not on the item itself, but on a sub-item, this field identifies the sub-item by name.  For example, for actions performed on prompt resources, this will be the prompt resource name.
	Resource *string `json:"resource,omitempty"`


	// Timestamp - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	Timestamp *time.Time `json:"timestamp,omitempty"`


	// User - User associated with this entry.
	User *User `json:"user,omitempty"`


	// Client - OAuth client associated with this entry.
	Client *Domainentityref `json:"client,omitempty"`


	// Version
	Version *string `json:"version,omitempty"`


	// Secure
	Secure *bool `json:"secure,omitempty"`

}

// String returns a JSON representation of the model
func (o *Historyentry) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

package platformclientv2
import (
	"time"
	"encoding/json"
)

// Dependencystatus
type Dependencystatus struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// User - User that initiated the build.
	User *User `json:"user,omitempty"`


	// Client - OAuth client that initiated the build.
	Client *Domainentityref `json:"client,omitempty"`


	// BuildId
	BuildId *string `json:"buildId,omitempty"`


	// DateStarted - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateStarted *time.Time `json:"dateStarted,omitempty"`


	// DateCompleted - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCompleted *time.Time `json:"dateCompleted,omitempty"`


	// Status
	Status *string `json:"status,omitempty"`


	// FailedObjects
	FailedObjects *[]Failedobject `json:"failedObjects,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Dependencystatus) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

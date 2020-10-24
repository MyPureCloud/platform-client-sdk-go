package platformclientv2
import (
	"time"
	"encoding/json"
)

// Messagingintegration
type Messagingintegration struct { 
	// Id - A unique Integration Id
	Id *string `json:"id,omitempty"`


	// Name - The name of the Integration
	Name *string `json:"name,omitempty"`


	// Status - The status of the Integration
	Status *string `json:"status,omitempty"`


	// MessengerType - The type of Messaging Integration
	MessengerType *string `json:"messengerType,omitempty"`


	// Recipient - The recipient associated to the Integration. This recipient is used to associate a flow to an integration
	Recipient *Domainentityref `json:"recipient,omitempty"`


	// DateCreated - Date this Integration was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateModified - Date this Integration was modified. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`


	// CreatedBy - User reference that created this Integration
	CreatedBy *Domainentityref `json:"createdBy,omitempty"`


	// ModifiedBy - User reference that last modified this Integration
	ModifiedBy *Domainentityref `json:"modifiedBy,omitempty"`


	// Version - Version number required for updates.
	Version *int32 `json:"version,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Messagingintegration) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

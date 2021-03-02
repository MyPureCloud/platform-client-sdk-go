package platformclientv2
import (
	"time"
	"encoding/json"
)

// Auditlogmessage
type Auditlogmessage struct { 
	// Id - Id of the audit message.
	Id *string `json:"id,omitempty"`


	// UserHomeOrgId - Home Organization Id associated with this audit message.
	UserHomeOrgId *string `json:"userHomeOrgId,omitempty"`


	// User - User associated with this audit message.
	User *Domainentityref `json:"user,omitempty"`


	// Client - Client associated with this audit message.
	Client *Addressableentityref `json:"client,omitempty"`


	// RemoteIp - List of IP addresses of systems that originated or handled the request.
	RemoteIp *[]string `json:"remoteIp,omitempty"`


	// ServiceName - Name of the service that logged this audit message.
	ServiceName *string `json:"serviceName,omitempty"`


	// EventDate - Date and time of when the audit message was logged. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	EventDate *time.Time `json:"eventDate,omitempty"`


	// Message - Message describing the event being audited.
	Message *Messageinfo `json:"message,omitempty"`


	// Action - Action that took place.
	Action *string `json:"action,omitempty"`


	// Entity - Entity that was impacted.
	Entity *Domainentityref `json:"entity,omitempty"`


	// EntityType - Type of the entity that was impacted.
	EntityType *string `json:"entityType,omitempty"`


	// PropertyChanges - List of properties that were changed and changes made to those properties.
	PropertyChanges *[]Propertychange `json:"propertyChanges,omitempty"`


	// Context - Additional context for this message.
	Context *map[string]string `json:"context,omitempty"`

}

// String returns a JSON representation of the model
func (o *Auditlogmessage) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

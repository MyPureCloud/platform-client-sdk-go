package platformclientv2
import (
	"time"
	"encoding/json"
)

// Eventlog
type Eventlog struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// ErrorEntity
	ErrorEntity *Domainentityref `json:"errorEntity,omitempty"`


	// RelatedEntity
	RelatedEntity *Domainentityref `json:"relatedEntity,omitempty"`


	// Timestamp - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	Timestamp *time.Time `json:"timestamp,omitempty"`


	// Level
	Level *string `json:"level,omitempty"`


	// Category
	Category *string `json:"category,omitempty"`


	// CorrelationId
	CorrelationId *string `json:"correlationId,omitempty"`


	// EventMessage
	EventMessage *Eventmessage `json:"eventMessage,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Eventlog) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

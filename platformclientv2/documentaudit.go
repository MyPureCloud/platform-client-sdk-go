package platformclientv2
import (
	"time"
	"encoding/json"
)

// Documentaudit
type Documentaudit struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// User
	User *Domainentityref `json:"user,omitempty"`


	// Workspace
	Workspace *Domainentityref `json:"workspace,omitempty"`


	// TransactionId
	TransactionId *string `json:"transactionId,omitempty"`


	// TransactionInitiator
	TransactionInitiator *bool `json:"transactionInitiator,omitempty"`


	// Application
	Application *string `json:"application,omitempty"`


	// ServiceName
	ServiceName *string `json:"serviceName,omitempty"`


	// Level
	Level *string `json:"level,omitempty"`


	// Timestamp - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	Timestamp *time.Time `json:"timestamp,omitempty"`


	// Status
	Status *string `json:"status,omitempty"`


	// ActionContext
	ActionContext *string `json:"actionContext,omitempty"`


	// Action
	Action *string `json:"action,omitempty"`


	// Entity
	Entity *Auditentityreference `json:"entity,omitempty"`


	// Changes
	Changes *[]Auditchange `json:"changes,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Documentaudit) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

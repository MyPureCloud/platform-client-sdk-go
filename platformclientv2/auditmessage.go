package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Auditmessage
type Auditmessage struct { 
	// Id - AuditMessage ID.
	Id *string `json:"id,omitempty"`


	// User
	User *Audituser `json:"user,omitempty"`


	// CorrelationId - Correlation ID.
	CorrelationId *string `json:"correlationId,omitempty"`


	// TransactionId - Transaction ID.
	TransactionId *string `json:"transactionId,omitempty"`


	// TransactionInitiator - Whether or not this audit can be considered the initiator of the transaction it is a part of.
	TransactionInitiator *bool `json:"transactionInitiator,omitempty"`


	// Application - The application through which the action of this AuditMessage was initiated.
	Application *string `json:"application,omitempty"`


	// ServiceName - The name of the service which sent this AuditMessage.
	ServiceName *string `json:"serviceName,omitempty"`


	// Level - The level of this audit. USER or SYSTEM.
	Level *string `json:"level,omitempty"`


	// Timestamp - The time at which the action of this AuditMessage was initiated.
	Timestamp *string `json:"timestamp,omitempty"`


	// ReceivedTimestamp - The time at which this AuditMessage was received.
	ReceivedTimestamp *string `json:"receivedTimestamp,omitempty"`


	// Status - The status of the action of this AuditMessage
	Status *string `json:"status,omitempty"`


	// ActionContext - The context of a system-level action
	ActionContext *string `json:"actionContext,omitempty"`


	// Action - A string representing the action that took place
	Action *string `json:"action,omitempty"`


	// Changes - Details about any changes that occurred in this audit
	Changes *[]Change `json:"changes,omitempty"`


	// Entity
	Entity *Auditentity `json:"entity,omitempty"`


	// ServiceContext - The service-specific context associated with this AuditMessage.
	ServiceContext *Servicecontext `json:"serviceContext,omitempty"`

}

// String returns a JSON representation of the model
func (o *Auditmessage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

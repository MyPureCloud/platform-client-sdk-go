package platformclientv2
import (
	"github.com/leekchan/timeutil"
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

func (o *Auditmessage) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Auditmessage
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		User *Audituser `json:"user,omitempty"`
		
		CorrelationId *string `json:"correlationId,omitempty"`
		
		TransactionId *string `json:"transactionId,omitempty"`
		
		TransactionInitiator *bool `json:"transactionInitiator,omitempty"`
		
		Application *string `json:"application,omitempty"`
		
		ServiceName *string `json:"serviceName,omitempty"`
		
		Level *string `json:"level,omitempty"`
		
		Timestamp *string `json:"timestamp,omitempty"`
		
		ReceivedTimestamp *string `json:"receivedTimestamp,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		ActionContext *string `json:"actionContext,omitempty"`
		
		Action *string `json:"action,omitempty"`
		
		Changes *[]Change `json:"changes,omitempty"`
		
		Entity *Auditentity `json:"entity,omitempty"`
		
		ServiceContext *Servicecontext `json:"serviceContext,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		User: o.User,
		
		CorrelationId: o.CorrelationId,
		
		TransactionId: o.TransactionId,
		
		TransactionInitiator: o.TransactionInitiator,
		
		Application: o.Application,
		
		ServiceName: o.ServiceName,
		
		Level: o.Level,
		
		Timestamp: o.Timestamp,
		
		ReceivedTimestamp: o.ReceivedTimestamp,
		
		Status: o.Status,
		
		ActionContext: o.ActionContext,
		
		Action: o.Action,
		
		Changes: o.Changes,
		
		Entity: o.Entity,
		
		ServiceContext: o.ServiceContext,
		Alias:    (*Alias)(o),
	})
}

func (o *Auditmessage) UnmarshalJSON(b []byte) error {
	var AuditmessageMap map[string]interface{}
	err := json.Unmarshal(b, &AuditmessageMap)
	if err != nil {
		return err
	}
	
	if Id, ok := AuditmessageMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if User, ok := AuditmessageMap["user"].(map[string]interface{}); ok {
		UserString, _ := json.Marshal(User)
		json.Unmarshal(UserString, &o.User)
	}
	
	if CorrelationId, ok := AuditmessageMap["correlationId"].(string); ok {
		o.CorrelationId = &CorrelationId
	}
	
	if TransactionId, ok := AuditmessageMap["transactionId"].(string); ok {
		o.TransactionId = &TransactionId
	}
	
	if TransactionInitiator, ok := AuditmessageMap["transactionInitiator"].(bool); ok {
		o.TransactionInitiator = &TransactionInitiator
	}
	
	if Application, ok := AuditmessageMap["application"].(string); ok {
		o.Application = &Application
	}
	
	if ServiceName, ok := AuditmessageMap["serviceName"].(string); ok {
		o.ServiceName = &ServiceName
	}
	
	if Level, ok := AuditmessageMap["level"].(string); ok {
		o.Level = &Level
	}
	
	if Timestamp, ok := AuditmessageMap["timestamp"].(string); ok {
		o.Timestamp = &Timestamp
	}
	
	if ReceivedTimestamp, ok := AuditmessageMap["receivedTimestamp"].(string); ok {
		o.ReceivedTimestamp = &ReceivedTimestamp
	}
	
	if Status, ok := AuditmessageMap["status"].(string); ok {
		o.Status = &Status
	}
	
	if ActionContext, ok := AuditmessageMap["actionContext"].(string); ok {
		o.ActionContext = &ActionContext
	}
	
	if Action, ok := AuditmessageMap["action"].(string); ok {
		o.Action = &Action
	}
	
	if Changes, ok := AuditmessageMap["changes"].([]interface{}); ok {
		ChangesString, _ := json.Marshal(Changes)
		json.Unmarshal(ChangesString, &o.Changes)
	}
	
	if Entity, ok := AuditmessageMap["entity"].(map[string]interface{}); ok {
		EntityString, _ := json.Marshal(Entity)
		json.Unmarshal(EntityString, &o.Entity)
	}
	
	if ServiceContext, ok := AuditmessageMap["serviceContext"].(map[string]interface{}); ok {
		ServiceContextString, _ := json.Marshal(ServiceContext)
		json.Unmarshal(ServiceContextString, &o.ServiceContext)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Auditmessage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

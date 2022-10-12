package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
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


	// Level - Level of this audit message, USER or SYSTEM.
	Level *string `json:"level,omitempty"`


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


	// Status - Status of the event being audited
	Status *string `json:"status,omitempty"`


	// Application - Name of the application used to perform the audit's action
	Application *string `json:"application,omitempty"`


	// InitiatingAction - Id and action of the audit initiating the transaction
	InitiatingAction *Initiatingaction `json:"initiatingAction,omitempty"`


	// TransactionInitiator - Whether the current audit is the initiator of the transaction
	TransactionInitiator *bool `json:"transactionInitiator,omitempty"`


	// PropertyChanges - List of properties that were changed and changes made to those properties.
	PropertyChanges *[]Propertychange `json:"propertyChanges,omitempty"`


	// Context - Additional context for this message.
	Context *map[string]string `json:"context,omitempty"`


	// EntityChanges - List of entities that were changed and changes made to those entities.
	EntityChanges *[]Entitychange `json:"entityChanges,omitempty"`

}

func (o *Auditlogmessage) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Auditlogmessage
	
	EventDate := new(string)
	if o.EventDate != nil {
		
		*EventDate = timeutil.Strftime(o.EventDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		EventDate = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		UserHomeOrgId *string `json:"userHomeOrgId,omitempty"`
		
		User *Domainentityref `json:"user,omitempty"`
		
		Client *Addressableentityref `json:"client,omitempty"`
		
		RemoteIp *[]string `json:"remoteIp,omitempty"`
		
		ServiceName *string `json:"serviceName,omitempty"`
		
		Level *string `json:"level,omitempty"`
		
		EventDate *string `json:"eventDate,omitempty"`
		
		Message *Messageinfo `json:"message,omitempty"`
		
		Action *string `json:"action,omitempty"`
		
		Entity *Domainentityref `json:"entity,omitempty"`
		
		EntityType *string `json:"entityType,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		Application *string `json:"application,omitempty"`
		
		InitiatingAction *Initiatingaction `json:"initiatingAction,omitempty"`
		
		TransactionInitiator *bool `json:"transactionInitiator,omitempty"`
		
		PropertyChanges *[]Propertychange `json:"propertyChanges,omitempty"`
		
		Context *map[string]string `json:"context,omitempty"`
		
		EntityChanges *[]Entitychange `json:"entityChanges,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		UserHomeOrgId: o.UserHomeOrgId,
		
		User: o.User,
		
		Client: o.Client,
		
		RemoteIp: o.RemoteIp,
		
		ServiceName: o.ServiceName,
		
		Level: o.Level,
		
		EventDate: EventDate,
		
		Message: o.Message,
		
		Action: o.Action,
		
		Entity: o.Entity,
		
		EntityType: o.EntityType,
		
		Status: o.Status,
		
		Application: o.Application,
		
		InitiatingAction: o.InitiatingAction,
		
		TransactionInitiator: o.TransactionInitiator,
		
		PropertyChanges: o.PropertyChanges,
		
		Context: o.Context,
		
		EntityChanges: o.EntityChanges,
		Alias:    (*Alias)(o),
	})
}

func (o *Auditlogmessage) UnmarshalJSON(b []byte) error {
	var AuditlogmessageMap map[string]interface{}
	err := json.Unmarshal(b, &AuditlogmessageMap)
	if err != nil {
		return err
	}
	
	if Id, ok := AuditlogmessageMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if UserHomeOrgId, ok := AuditlogmessageMap["userHomeOrgId"].(string); ok {
		o.UserHomeOrgId = &UserHomeOrgId
	}
    
	if User, ok := AuditlogmessageMap["user"].(map[string]interface{}); ok {
		UserString, _ := json.Marshal(User)
		json.Unmarshal(UserString, &o.User)
	}
	
	if Client, ok := AuditlogmessageMap["client"].(map[string]interface{}); ok {
		ClientString, _ := json.Marshal(Client)
		json.Unmarshal(ClientString, &o.Client)
	}
	
	if RemoteIp, ok := AuditlogmessageMap["remoteIp"].([]interface{}); ok {
		RemoteIpString, _ := json.Marshal(RemoteIp)
		json.Unmarshal(RemoteIpString, &o.RemoteIp)
	}
	
	if ServiceName, ok := AuditlogmessageMap["serviceName"].(string); ok {
		o.ServiceName = &ServiceName
	}
    
	if Level, ok := AuditlogmessageMap["level"].(string); ok {
		o.Level = &Level
	}
    
	if eventDateString, ok := AuditlogmessageMap["eventDate"].(string); ok {
		EventDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", eventDateString)
		o.EventDate = &EventDate
	}
	
	if Message, ok := AuditlogmessageMap["message"].(map[string]interface{}); ok {
		MessageString, _ := json.Marshal(Message)
		json.Unmarshal(MessageString, &o.Message)
	}
	
	if Action, ok := AuditlogmessageMap["action"].(string); ok {
		o.Action = &Action
	}
    
	if Entity, ok := AuditlogmessageMap["entity"].(map[string]interface{}); ok {
		EntityString, _ := json.Marshal(Entity)
		json.Unmarshal(EntityString, &o.Entity)
	}
	
	if EntityType, ok := AuditlogmessageMap["entityType"].(string); ok {
		o.EntityType = &EntityType
	}
    
	if Status, ok := AuditlogmessageMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if Application, ok := AuditlogmessageMap["application"].(string); ok {
		o.Application = &Application
	}
    
	if InitiatingAction, ok := AuditlogmessageMap["initiatingAction"].(map[string]interface{}); ok {
		InitiatingActionString, _ := json.Marshal(InitiatingAction)
		json.Unmarshal(InitiatingActionString, &o.InitiatingAction)
	}
	
	if TransactionInitiator, ok := AuditlogmessageMap["transactionInitiator"].(bool); ok {
		o.TransactionInitiator = &TransactionInitiator
	}
    
	if PropertyChanges, ok := AuditlogmessageMap["propertyChanges"].([]interface{}); ok {
		PropertyChangesString, _ := json.Marshal(PropertyChanges)
		json.Unmarshal(PropertyChangesString, &o.PropertyChanges)
	}
	
	if Context, ok := AuditlogmessageMap["context"].(map[string]interface{}); ok {
		ContextString, _ := json.Marshal(Context)
		json.Unmarshal(ContextString, &o.Context)
	}
	
	if EntityChanges, ok := AuditlogmessageMap["entityChanges"].([]interface{}); ok {
		EntityChangesString, _ := json.Marshal(EntityChanges)
		json.Unmarshal(EntityChangesString, &o.EntityChanges)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Auditlogmessage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

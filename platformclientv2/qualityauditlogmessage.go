package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Qualityauditlogmessage
type Qualityauditlogmessage struct { 
	// Id - Id of the audit message.
	Id *string `json:"id,omitempty"`


	// UserHomeOrgId - Home Organization Id associated with this audit message.
	UserHomeOrgId *string `json:"userHomeOrgId,omitempty"`


	// UserTrusteeOrgId - Trustee Organization Id if this audit message is from trustee access.
	UserTrusteeOrgId *string `json:"userTrusteeOrgId,omitempty"`


	// User - User associated with this audit message.
	User *Domainentityref `json:"user,omitempty"`


	// Client - Client associated with this audit message.
	Client *Addressableentityref `json:"client,omitempty"`


	// RemoteIps - List of IP addresses of systems that originated or handled the request.
	RemoteIps *[]string `json:"remoteIps,omitempty"`


	// ServiceName - Name of the service that logged this audit message.
	ServiceName *string `json:"serviceName,omitempty"`


	// Level - The level of this audit message.
	Level *string `json:"level,omitempty"`


	// Status - The status of the action of this audit message.
	Status *string `json:"status,omitempty"`


	// EventDate - Date and time of when the audit message was logged. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	EventDate *time.Time `json:"eventDate,omitempty"`


	// MessageInfo - Message describing the event being audited.
	MessageInfo *Messageinfo `json:"messageInfo,omitempty"`


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

func (o *Qualityauditlogmessage) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Qualityauditlogmessage
	
	EventDate := new(string)
	if o.EventDate != nil {
		
		*EventDate = timeutil.Strftime(o.EventDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		EventDate = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		UserHomeOrgId *string `json:"userHomeOrgId,omitempty"`
		
		UserTrusteeOrgId *string `json:"userTrusteeOrgId,omitempty"`
		
		User *Domainentityref `json:"user,omitempty"`
		
		Client *Addressableentityref `json:"client,omitempty"`
		
		RemoteIps *[]string `json:"remoteIps,omitempty"`
		
		ServiceName *string `json:"serviceName,omitempty"`
		
		Level *string `json:"level,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		EventDate *string `json:"eventDate,omitempty"`
		
		MessageInfo *Messageinfo `json:"messageInfo,omitempty"`
		
		Action *string `json:"action,omitempty"`
		
		Entity *Domainentityref `json:"entity,omitempty"`
		
		EntityType *string `json:"entityType,omitempty"`
		
		PropertyChanges *[]Propertychange `json:"propertyChanges,omitempty"`
		
		Context *map[string]string `json:"context,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		UserHomeOrgId: o.UserHomeOrgId,
		
		UserTrusteeOrgId: o.UserTrusteeOrgId,
		
		User: o.User,
		
		Client: o.Client,
		
		RemoteIps: o.RemoteIps,
		
		ServiceName: o.ServiceName,
		
		Level: o.Level,
		
		Status: o.Status,
		
		EventDate: EventDate,
		
		MessageInfo: o.MessageInfo,
		
		Action: o.Action,
		
		Entity: o.Entity,
		
		EntityType: o.EntityType,
		
		PropertyChanges: o.PropertyChanges,
		
		Context: o.Context,
		Alias:    (*Alias)(o),
	})
}

func (o *Qualityauditlogmessage) UnmarshalJSON(b []byte) error {
	var QualityauditlogmessageMap map[string]interface{}
	err := json.Unmarshal(b, &QualityauditlogmessageMap)
	if err != nil {
		return err
	}
	
	if Id, ok := QualityauditlogmessageMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if UserHomeOrgId, ok := QualityauditlogmessageMap["userHomeOrgId"].(string); ok {
		o.UserHomeOrgId = &UserHomeOrgId
	}
    
	if UserTrusteeOrgId, ok := QualityauditlogmessageMap["userTrusteeOrgId"].(string); ok {
		o.UserTrusteeOrgId = &UserTrusteeOrgId
	}
    
	if User, ok := QualityauditlogmessageMap["user"].(map[string]interface{}); ok {
		UserString, _ := json.Marshal(User)
		json.Unmarshal(UserString, &o.User)
	}
	
	if Client, ok := QualityauditlogmessageMap["client"].(map[string]interface{}); ok {
		ClientString, _ := json.Marshal(Client)
		json.Unmarshal(ClientString, &o.Client)
	}
	
	if RemoteIps, ok := QualityauditlogmessageMap["remoteIps"].([]interface{}); ok {
		RemoteIpsString, _ := json.Marshal(RemoteIps)
		json.Unmarshal(RemoteIpsString, &o.RemoteIps)
	}
	
	if ServiceName, ok := QualityauditlogmessageMap["serviceName"].(string); ok {
		o.ServiceName = &ServiceName
	}
    
	if Level, ok := QualityauditlogmessageMap["level"].(string); ok {
		o.Level = &Level
	}
    
	if Status, ok := QualityauditlogmessageMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if eventDateString, ok := QualityauditlogmessageMap["eventDate"].(string); ok {
		EventDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", eventDateString)
		o.EventDate = &EventDate
	}
	
	if MessageInfo, ok := QualityauditlogmessageMap["messageInfo"].(map[string]interface{}); ok {
		MessageInfoString, _ := json.Marshal(MessageInfo)
		json.Unmarshal(MessageInfoString, &o.MessageInfo)
	}
	
	if Action, ok := QualityauditlogmessageMap["action"].(string); ok {
		o.Action = &Action
	}
    
	if Entity, ok := QualityauditlogmessageMap["entity"].(map[string]interface{}); ok {
		EntityString, _ := json.Marshal(Entity)
		json.Unmarshal(EntityString, &o.Entity)
	}
	
	if EntityType, ok := QualityauditlogmessageMap["entityType"].(string); ok {
		o.EntityType = &EntityType
	}
    
	if PropertyChanges, ok := QualityauditlogmessageMap["propertyChanges"].([]interface{}); ok {
		PropertyChangesString, _ := json.Marshal(PropertyChanges)
		json.Unmarshal(PropertyChangesString, &o.PropertyChanges)
	}
	
	if Context, ok := QualityauditlogmessageMap["context"].(map[string]interface{}); ok {
		ContextString, _ := json.Marshal(Context)
		json.Unmarshal(ContextString, &o.Context)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Qualityauditlogmessage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

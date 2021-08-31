package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Audittopicauditlogmessage
type Audittopicauditlogmessage struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// UserId
	UserId *string `json:"userId,omitempty"`


	// UserHomeOrgId
	UserHomeOrgId *string `json:"userHomeOrgId,omitempty"`


	// Username
	Username *Audittopicdomainentityref `json:"username,omitempty"`


	// UserDisplay
	UserDisplay *string `json:"userDisplay,omitempty"`


	// ClientId
	ClientId *Audittopicaddressableentityref `json:"clientId,omitempty"`


	// RemoteIp
	RemoteIp *[]string `json:"remoteIp,omitempty"`


	// ServiceName
	ServiceName *string `json:"serviceName,omitempty"`


	// EventTime
	EventTime *time.Time `json:"eventTime,omitempty"`


	// Message
	Message *Audittopicmessageinfo `json:"message,omitempty"`


	// Action
	Action *string `json:"action,omitempty"`


	// EntityType
	EntityType *string `json:"entityType,omitempty"`


	// Entity
	Entity *Audittopicdomainentityref `json:"entity,omitempty"`


	// PropertyChanges
	PropertyChanges *[]Audittopicpropertychange `json:"propertyChanges,omitempty"`


	// Context
	Context *map[string]string `json:"context,omitempty"`

}

func (o *Audittopicauditlogmessage) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Audittopicauditlogmessage
	
	EventTime := new(string)
	if o.EventTime != nil {
		
		*EventTime = timeutil.Strftime(o.EventTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		EventTime = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		UserId *string `json:"userId,omitempty"`
		
		UserHomeOrgId *string `json:"userHomeOrgId,omitempty"`
		
		Username *Audittopicdomainentityref `json:"username,omitempty"`
		
		UserDisplay *string `json:"userDisplay,omitempty"`
		
		ClientId *Audittopicaddressableentityref `json:"clientId,omitempty"`
		
		RemoteIp *[]string `json:"remoteIp,omitempty"`
		
		ServiceName *string `json:"serviceName,omitempty"`
		
		EventTime *string `json:"eventTime,omitempty"`
		
		Message *Audittopicmessageinfo `json:"message,omitempty"`
		
		Action *string `json:"action,omitempty"`
		
		EntityType *string `json:"entityType,omitempty"`
		
		Entity *Audittopicdomainentityref `json:"entity,omitempty"`
		
		PropertyChanges *[]Audittopicpropertychange `json:"propertyChanges,omitempty"`
		
		Context *map[string]string `json:"context,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		UserId: o.UserId,
		
		UserHomeOrgId: o.UserHomeOrgId,
		
		Username: o.Username,
		
		UserDisplay: o.UserDisplay,
		
		ClientId: o.ClientId,
		
		RemoteIp: o.RemoteIp,
		
		ServiceName: o.ServiceName,
		
		EventTime: EventTime,
		
		Message: o.Message,
		
		Action: o.Action,
		
		EntityType: o.EntityType,
		
		Entity: o.Entity,
		
		PropertyChanges: o.PropertyChanges,
		
		Context: o.Context,
		Alias:    (*Alias)(o),
	})
}

func (o *Audittopicauditlogmessage) UnmarshalJSON(b []byte) error {
	var AudittopicauditlogmessageMap map[string]interface{}
	err := json.Unmarshal(b, &AudittopicauditlogmessageMap)
	if err != nil {
		return err
	}
	
	if Id, ok := AudittopicauditlogmessageMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if UserId, ok := AudittopicauditlogmessageMap["userId"].(string); ok {
		o.UserId = &UserId
	}
	
	if UserHomeOrgId, ok := AudittopicauditlogmessageMap["userHomeOrgId"].(string); ok {
		o.UserHomeOrgId = &UserHomeOrgId
	}
	
	if Username, ok := AudittopicauditlogmessageMap["username"].(map[string]interface{}); ok {
		UsernameString, _ := json.Marshal(Username)
		json.Unmarshal(UsernameString, &o.Username)
	}
	
	if UserDisplay, ok := AudittopicauditlogmessageMap["userDisplay"].(string); ok {
		o.UserDisplay = &UserDisplay
	}
	
	if ClientId, ok := AudittopicauditlogmessageMap["clientId"].(map[string]interface{}); ok {
		ClientIdString, _ := json.Marshal(ClientId)
		json.Unmarshal(ClientIdString, &o.ClientId)
	}
	
	if RemoteIp, ok := AudittopicauditlogmessageMap["remoteIp"].([]interface{}); ok {
		RemoteIpString, _ := json.Marshal(RemoteIp)
		json.Unmarshal(RemoteIpString, &o.RemoteIp)
	}
	
	if ServiceName, ok := AudittopicauditlogmessageMap["serviceName"].(string); ok {
		o.ServiceName = &ServiceName
	}
	
	if eventTimeString, ok := AudittopicauditlogmessageMap["eventTime"].(string); ok {
		EventTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", eventTimeString)
		o.EventTime = &EventTime
	}
	
	if Message, ok := AudittopicauditlogmessageMap["message"].(map[string]interface{}); ok {
		MessageString, _ := json.Marshal(Message)
		json.Unmarshal(MessageString, &o.Message)
	}
	
	if Action, ok := AudittopicauditlogmessageMap["action"].(string); ok {
		o.Action = &Action
	}
	
	if EntityType, ok := AudittopicauditlogmessageMap["entityType"].(string); ok {
		o.EntityType = &EntityType
	}
	
	if Entity, ok := AudittopicauditlogmessageMap["entity"].(map[string]interface{}); ok {
		EntityString, _ := json.Marshal(Entity)
		json.Unmarshal(EntityString, &o.Entity)
	}
	
	if PropertyChanges, ok := AudittopicauditlogmessageMap["propertyChanges"].([]interface{}); ok {
		PropertyChangesString, _ := json.Marshal(PropertyChanges)
		json.Unmarshal(PropertyChangesString, &o.PropertyChanges)
	}
	
	if Context, ok := AudittopicauditlogmessageMap["context"].(map[string]interface{}); ok {
		ContextString, _ := json.Marshal(Context)
		json.Unmarshal(ContextString, &o.Context)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Audittopicauditlogmessage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

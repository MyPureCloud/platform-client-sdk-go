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

func (u *Audittopicauditlogmessage) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Audittopicauditlogmessage

	
	EventTime := new(string)
	if u.EventTime != nil {
		
		*EventTime = timeutil.Strftime(u.EventTime, "%Y-%m-%dT%H:%M:%S.%fZ")
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
		Id: u.Id,
		
		UserId: u.UserId,
		
		UserHomeOrgId: u.UserHomeOrgId,
		
		Username: u.Username,
		
		UserDisplay: u.UserDisplay,
		
		ClientId: u.ClientId,
		
		RemoteIp: u.RemoteIp,
		
		ServiceName: u.ServiceName,
		
		EventTime: EventTime,
		
		Message: u.Message,
		
		Action: u.Action,
		
		EntityType: u.EntityType,
		
		Entity: u.Entity,
		
		PropertyChanges: u.PropertyChanges,
		
		Context: u.Context,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Audittopicauditlogmessage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

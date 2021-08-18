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

func (u *Qualityauditlogmessage) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Qualityauditlogmessage

	
	EventDate := new(string)
	if u.EventDate != nil {
		
		*EventDate = timeutil.Strftime(u.EventDate, "%Y-%m-%dT%H:%M:%S.%fZ")
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
		Id: u.Id,
		
		UserHomeOrgId: u.UserHomeOrgId,
		
		UserTrusteeOrgId: u.UserTrusteeOrgId,
		
		User: u.User,
		
		Client: u.Client,
		
		RemoteIps: u.RemoteIps,
		
		ServiceName: u.ServiceName,
		
		Level: u.Level,
		
		Status: u.Status,
		
		EventDate: EventDate,
		
		MessageInfo: u.MessageInfo,
		
		Action: u.Action,
		
		Entity: u.Entity,
		
		EntityType: u.EntityType,
		
		PropertyChanges: u.PropertyChanges,
		
		Context: u.Context,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Qualityauditlogmessage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

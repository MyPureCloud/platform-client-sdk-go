package platformclientv2
import (
	"time"
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

// String returns a JSON representation of the model
func (o *Audittopicauditlogmessage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

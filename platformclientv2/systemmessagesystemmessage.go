package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Systemmessagesystemmessage
type Systemmessagesystemmessage struct { 
	// ChannelId
	ChannelId *string `json:"channelId,omitempty"`


	// SystemTopicType
	SystemTopicType *string `json:"systemTopicType,omitempty"`


	// CorrelationId
	CorrelationId *string `json:"correlationId,omitempty"`


	// OrganizationId
	OrganizationId *string `json:"organizationId,omitempty"`


	// UserId
	UserId *string `json:"userId,omitempty"`


	// OauthClientId
	OauthClientId *string `json:"oauthClientId,omitempty"`


	// Reason
	Reason *string `json:"reason,omitempty"`


	// Message
	Message *string `json:"message,omitempty"`


	// Data
	Data *interface{} `json:"data,omitempty"`

}

func (u *Systemmessagesystemmessage) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Systemmessagesystemmessage

	

	return json.Marshal(&struct { 
		ChannelId *string `json:"channelId,omitempty"`
		
		SystemTopicType *string `json:"systemTopicType,omitempty"`
		
		CorrelationId *string `json:"correlationId,omitempty"`
		
		OrganizationId *string `json:"organizationId,omitempty"`
		
		UserId *string `json:"userId,omitempty"`
		
		OauthClientId *string `json:"oauthClientId,omitempty"`
		
		Reason *string `json:"reason,omitempty"`
		
		Message *string `json:"message,omitempty"`
		
		Data *interface{} `json:"data,omitempty"`
		*Alias
	}{ 
		ChannelId: u.ChannelId,
		
		SystemTopicType: u.SystemTopicType,
		
		CorrelationId: u.CorrelationId,
		
		OrganizationId: u.OrganizationId,
		
		UserId: u.UserId,
		
		OauthClientId: u.OauthClientId,
		
		Reason: u.Reason,
		
		Message: u.Message,
		
		Data: u.Data,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Systemmessagesystemmessage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

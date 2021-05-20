package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Systemmessagesystemmessage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

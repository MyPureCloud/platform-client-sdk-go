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

func (o *Systemmessagesystemmessage) MarshalJSON() ([]byte, error) {
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
		ChannelId: o.ChannelId,
		
		SystemTopicType: o.SystemTopicType,
		
		CorrelationId: o.CorrelationId,
		
		OrganizationId: o.OrganizationId,
		
		UserId: o.UserId,
		
		OauthClientId: o.OauthClientId,
		
		Reason: o.Reason,
		
		Message: o.Message,
		
		Data: o.Data,
		Alias:    (*Alias)(o),
	})
}

func (o *Systemmessagesystemmessage) UnmarshalJSON(b []byte) error {
	var SystemmessagesystemmessageMap map[string]interface{}
	err := json.Unmarshal(b, &SystemmessagesystemmessageMap)
	if err != nil {
		return err
	}
	
	if ChannelId, ok := SystemmessagesystemmessageMap["channelId"].(string); ok {
		o.ChannelId = &ChannelId
	}
	
	if SystemTopicType, ok := SystemmessagesystemmessageMap["systemTopicType"].(string); ok {
		o.SystemTopicType = &SystemTopicType
	}
	
	if CorrelationId, ok := SystemmessagesystemmessageMap["correlationId"].(string); ok {
		o.CorrelationId = &CorrelationId
	}
	
	if OrganizationId, ok := SystemmessagesystemmessageMap["organizationId"].(string); ok {
		o.OrganizationId = &OrganizationId
	}
	
	if UserId, ok := SystemmessagesystemmessageMap["userId"].(string); ok {
		o.UserId = &UserId
	}
	
	if OauthClientId, ok := SystemmessagesystemmessageMap["oauthClientId"].(string); ok {
		o.OauthClientId = &OauthClientId
	}
	
	if Reason, ok := SystemmessagesystemmessageMap["reason"].(string); ok {
		o.Reason = &Reason
	}
	
	if Message, ok := SystemmessagesystemmessageMap["message"].(string); ok {
		o.Message = &Message
	}
	
	if Data, ok := SystemmessagesystemmessageMap["data"].(map[string]interface{}); ok {
		DataString, _ := json.Marshal(Data)
		json.Unmarshal(DataString, &o.Data)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Systemmessagesystemmessage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

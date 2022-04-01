package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// V2conversationmessagetypingeventforworkflowtopicconversationmessagingchannel
type V2conversationmessagetypingeventforworkflowtopicconversationmessagingchannel struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Platform
	Platform *string `json:"platform,omitempty"`


	// MessageId
	MessageId *string `json:"messageId,omitempty"`


	// To
	To *V2conversationmessagetypingeventforworkflowtopicconversationmessagingtorecipient `json:"to,omitempty"`


	// From
	From *V2conversationmessagetypingeventforworkflowtopicconversationmessagingfromrecipient `json:"from,omitempty"`


	// Time
	Time *time.Time `json:"time,omitempty"`


	// Metadata
	Metadata *V2conversationmessagetypingeventforworkflowtopicconversationmessagingchannelmetadata `json:"metadata,omitempty"`

}

func (o *V2conversationmessagetypingeventforworkflowtopicconversationmessagingchannel) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias V2conversationmessagetypingeventforworkflowtopicconversationmessagingchannel
	
	Time := new(string)
	if o.Time != nil {
		
		*Time = timeutil.Strftime(o.Time, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		Time = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Platform *string `json:"platform,omitempty"`
		
		MessageId *string `json:"messageId,omitempty"`
		
		To *V2conversationmessagetypingeventforworkflowtopicconversationmessagingtorecipient `json:"to,omitempty"`
		
		From *V2conversationmessagetypingeventforworkflowtopicconversationmessagingfromrecipient `json:"from,omitempty"`
		
		Time *string `json:"time,omitempty"`
		
		Metadata *V2conversationmessagetypingeventforworkflowtopicconversationmessagingchannelmetadata `json:"metadata,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Platform: o.Platform,
		
		MessageId: o.MessageId,
		
		To: o.To,
		
		From: o.From,
		
		Time: Time,
		
		Metadata: o.Metadata,
		Alias:    (*Alias)(o),
	})
}

func (o *V2conversationmessagetypingeventforworkflowtopicconversationmessagingchannel) UnmarshalJSON(b []byte) error {
	var V2conversationmessagetypingeventforworkflowtopicconversationmessagingchannelMap map[string]interface{}
	err := json.Unmarshal(b, &V2conversationmessagetypingeventforworkflowtopicconversationmessagingchannelMap)
	if err != nil {
		return err
	}
	
	if Id, ok := V2conversationmessagetypingeventforworkflowtopicconversationmessagingchannelMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Platform, ok := V2conversationmessagetypingeventforworkflowtopicconversationmessagingchannelMap["platform"].(string); ok {
		o.Platform = &Platform
	}
	
	if MessageId, ok := V2conversationmessagetypingeventforworkflowtopicconversationmessagingchannelMap["messageId"].(string); ok {
		o.MessageId = &MessageId
	}
	
	if To, ok := V2conversationmessagetypingeventforworkflowtopicconversationmessagingchannelMap["to"].(map[string]interface{}); ok {
		ToString, _ := json.Marshal(To)
		json.Unmarshal(ToString, &o.To)
	}
	
	if From, ok := V2conversationmessagetypingeventforworkflowtopicconversationmessagingchannelMap["from"].(map[string]interface{}); ok {
		FromString, _ := json.Marshal(From)
		json.Unmarshal(FromString, &o.From)
	}
	
	if timeString, ok := V2conversationmessagetypingeventforworkflowtopicconversationmessagingchannelMap["time"].(string); ok {
		Time, _ := time.Parse("2006-01-02T15:04:05.999999Z", timeString)
		o.Time = &Time
	}
	
	if Metadata, ok := V2conversationmessagetypingeventforworkflowtopicconversationmessagingchannelMap["metadata"].(map[string]interface{}); ok {
		MetadataString, _ := json.Marshal(Metadata)
		json.Unmarshal(MetadataString, &o.Metadata)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *V2conversationmessagetypingeventforworkflowtopicconversationmessagingchannel) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

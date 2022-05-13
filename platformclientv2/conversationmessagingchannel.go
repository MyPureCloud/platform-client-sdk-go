package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationmessagingchannel - Channel-specific information that describes the message and the message channel/provider.
type Conversationmessagingchannel struct { 
	// Id - The integration ID.
	Id *string `json:"id,omitempty"`


	// Platform - The provider type.
	Platform *string `json:"platform,omitempty"`


	// MessageId - Unique provider ID of the message such as a Facebook message ID.
	MessageId *string `json:"messageId,omitempty"`


	// To - Information about the recipient the message is sent to.
	To *Conversationmessagingtorecipient `json:"to,omitempty"`


	// From - Information about the recipient the message is received from.
	From *Conversationmessagingfromrecipient `json:"from,omitempty"`


	// Time - Original time of the event. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	Time *time.Time `json:"time,omitempty"`


	// DateModified - Time the message was edited. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`


	// DateDeleted - Time the message was deleted. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateDeleted *time.Time `json:"dateDeleted,omitempty"`

}

func (o *Conversationmessagingchannel) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationmessagingchannel
	
	Time := new(string)
	if o.Time != nil {
		
		*Time = timeutil.Strftime(o.Time, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		Time = nil
	}
	
	DateModified := new(string)
	if o.DateModified != nil {
		
		*DateModified = timeutil.Strftime(o.DateModified, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateModified = nil
	}
	
	DateDeleted := new(string)
	if o.DateDeleted != nil {
		
		*DateDeleted = timeutil.Strftime(o.DateDeleted, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateDeleted = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Platform *string `json:"platform,omitempty"`
		
		MessageId *string `json:"messageId,omitempty"`
		
		To *Conversationmessagingtorecipient `json:"to,omitempty"`
		
		From *Conversationmessagingfromrecipient `json:"from,omitempty"`
		
		Time *string `json:"time,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		DateDeleted *string `json:"dateDeleted,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Platform: o.Platform,
		
		MessageId: o.MessageId,
		
		To: o.To,
		
		From: o.From,
		
		Time: Time,
		
		DateModified: DateModified,
		
		DateDeleted: DateDeleted,
		Alias:    (*Alias)(o),
	})
}

func (o *Conversationmessagingchannel) UnmarshalJSON(b []byte) error {
	var ConversationmessagingchannelMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationmessagingchannelMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ConversationmessagingchannelMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Platform, ok := ConversationmessagingchannelMap["platform"].(string); ok {
		o.Platform = &Platform
	}
    
	if MessageId, ok := ConversationmessagingchannelMap["messageId"].(string); ok {
		o.MessageId = &MessageId
	}
    
	if To, ok := ConversationmessagingchannelMap["to"].(map[string]interface{}); ok {
		ToString, _ := json.Marshal(To)
		json.Unmarshal(ToString, &o.To)
	}
	
	if From, ok := ConversationmessagingchannelMap["from"].(map[string]interface{}); ok {
		FromString, _ := json.Marshal(From)
		json.Unmarshal(FromString, &o.From)
	}
	
	if timeString, ok := ConversationmessagingchannelMap["time"].(string); ok {
		Time, _ := time.Parse("2006-01-02T15:04:05.999999Z", timeString)
		o.Time = &Time
	}
	
	if dateModifiedString, ok := ConversationmessagingchannelMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if dateDeletedString, ok := ConversationmessagingchannelMap["dateDeleted"].(string); ok {
		DateDeleted, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateDeletedString)
		o.DateDeleted = &DateDeleted
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationmessagingchannel) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Webmessagingchannel - Channel-specific information that describes the message and the message channel/provider.
type Webmessagingchannel struct { 
	// From - Information about the recipient the message is received from.
	From *Webmessagingrecipient `json:"from,omitempty"`


	// To - Information about the recipient the message is sent to.
	To *Webmessagingrecipient `json:"to,omitempty"`


	// Time - When the message was processed by Genesys Cloud. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	Time *time.Time `json:"time,omitempty"`


	// MessageId - Unique provider ID of the message.
	MessageId *string `json:"messageId,omitempty"`

}

func (o *Webmessagingchannel) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Webmessagingchannel
	
	Time := new(string)
	if o.Time != nil {
		
		*Time = timeutil.Strftime(o.Time, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		Time = nil
	}
	
	return json.Marshal(&struct { 
		From *Webmessagingrecipient `json:"from,omitempty"`
		
		To *Webmessagingrecipient `json:"to,omitempty"`
		
		Time *string `json:"time,omitempty"`
		
		MessageId *string `json:"messageId,omitempty"`
		*Alias
	}{ 
		From: o.From,
		
		To: o.To,
		
		Time: Time,
		
		MessageId: o.MessageId,
		Alias:    (*Alias)(o),
	})
}

func (o *Webmessagingchannel) UnmarshalJSON(b []byte) error {
	var WebmessagingchannelMap map[string]interface{}
	err := json.Unmarshal(b, &WebmessagingchannelMap)
	if err != nil {
		return err
	}
	
	if From, ok := WebmessagingchannelMap["from"].(map[string]interface{}); ok {
		FromString, _ := json.Marshal(From)
		json.Unmarshal(FromString, &o.From)
	}
	
	if To, ok := WebmessagingchannelMap["to"].(map[string]interface{}); ok {
		ToString, _ := json.Marshal(To)
		json.Unmarshal(ToString, &o.To)
	}
	
	if timeString, ok := WebmessagingchannelMap["time"].(string); ok {
		Time, _ := time.Parse("2006-01-02T15:04:05.999999Z", timeString)
		o.Time = &Time
	}
	
	if MessageId, ok := WebmessagingchannelMap["messageId"].(string); ok {
		o.MessageId = &MessageId
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Webmessagingchannel) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

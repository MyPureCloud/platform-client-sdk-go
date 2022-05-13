package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Openmessagingchannel - Channel-specific information that describes the message and the message channel/provider.
type Openmessagingchannel struct { 
	// Id - The Messaging Platform integration ID.
	Id *string `json:"id,omitempty"`


	// Platform - The provider type.
	Platform *string `json:"platform,omitempty"`


	// VarType - Specifies if this message is part of a private or public conversation.
	VarType *string `json:"type,omitempty"`


	// MessageId - Unique provider ID of the message such as a Facebook message ID.
	MessageId *string `json:"messageId,omitempty"`


	// To - Information about the recipient the message is sent to.
	To *Openmessagingtorecipient `json:"to,omitempty"`


	// From - Information about the recipient the message is received from.
	From *Openmessagingfromrecipient `json:"from,omitempty"`


	// Time - Original time of the event. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	Time *time.Time `json:"time,omitempty"`


	// Metadata - Information about the channel.
	Metadata *interface{} `json:"metadata,omitempty"`

}

func (o *Openmessagingchannel) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Openmessagingchannel
	
	Time := new(string)
	if o.Time != nil {
		
		*Time = timeutil.Strftime(o.Time, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		Time = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Platform *string `json:"platform,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		MessageId *string `json:"messageId,omitempty"`
		
		To *Openmessagingtorecipient `json:"to,omitempty"`
		
		From *Openmessagingfromrecipient `json:"from,omitempty"`
		
		Time *string `json:"time,omitempty"`
		
		Metadata *interface{} `json:"metadata,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Platform: o.Platform,
		
		VarType: o.VarType,
		
		MessageId: o.MessageId,
		
		To: o.To,
		
		From: o.From,
		
		Time: Time,
		
		Metadata: o.Metadata,
		Alias:    (*Alias)(o),
	})
}

func (o *Openmessagingchannel) UnmarshalJSON(b []byte) error {
	var OpenmessagingchannelMap map[string]interface{}
	err := json.Unmarshal(b, &OpenmessagingchannelMap)
	if err != nil {
		return err
	}
	
	if Id, ok := OpenmessagingchannelMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Platform, ok := OpenmessagingchannelMap["platform"].(string); ok {
		o.Platform = &Platform
	}
    
	if VarType, ok := OpenmessagingchannelMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if MessageId, ok := OpenmessagingchannelMap["messageId"].(string); ok {
		o.MessageId = &MessageId
	}
    
	if To, ok := OpenmessagingchannelMap["to"].(map[string]interface{}); ok {
		ToString, _ := json.Marshal(To)
		json.Unmarshal(ToString, &o.To)
	}
	
	if From, ok := OpenmessagingchannelMap["from"].(map[string]interface{}); ok {
		FromString, _ := json.Marshal(From)
		json.Unmarshal(FromString, &o.From)
	}
	
	if timeString, ok := OpenmessagingchannelMap["time"].(string); ok {
		Time, _ := time.Parse("2006-01-02T15:04:05.999999Z", timeString)
		o.Time = &Time
	}
	
	if Metadata, ok := OpenmessagingchannelMap["metadata"].(map[string]interface{}); ok {
		MetadataString, _ := json.Marshal(Metadata)
		json.Unmarshal(MetadataString, &o.Metadata)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Openmessagingchannel) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

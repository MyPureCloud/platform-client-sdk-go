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

func (u *Conversationmessagingchannel) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationmessagingchannel

	
	Time := new(string)
	if u.Time != nil {
		
		*Time = timeutil.Strftime(u.Time, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		Time = nil
	}
	
	DateModified := new(string)
	if u.DateModified != nil {
		
		*DateModified = timeutil.Strftime(u.DateModified, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateModified = nil
	}
	
	DateDeleted := new(string)
	if u.DateDeleted != nil {
		
		*DateDeleted = timeutil.Strftime(u.DateDeleted, "%Y-%m-%dT%H:%M:%S.%fZ")
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
		Id: u.Id,
		
		Platform: u.Platform,
		
		MessageId: u.MessageId,
		
		To: u.To,
		
		From: u.From,
		
		Time: Time,
		
		DateModified: DateModified,
		
		DateDeleted: DateDeleted,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Conversationmessagingchannel) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

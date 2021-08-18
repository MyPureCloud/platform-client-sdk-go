package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationmessageeventtopicmessagedetails
type Conversationmessageeventtopicmessagedetails struct { 
	// Message
	Message *Conversationmessageeventtopicurireference `json:"message,omitempty"`


	// MessageTime
	MessageTime *time.Time `json:"messageTime,omitempty"`


	// MessageSegmentCount
	MessageSegmentCount *int `json:"messageSegmentCount,omitempty"`


	// MessageStatus
	MessageStatus *string `json:"messageStatus,omitempty"`


	// Media
	Media *[]Conversationmessageeventtopicmessagemedia `json:"media,omitempty"`


	// Stickers
	Stickers *[]Conversationmessageeventtopicmessagesticker `json:"stickers,omitempty"`

}

func (u *Conversationmessageeventtopicmessagedetails) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationmessageeventtopicmessagedetails

	
	MessageTime := new(string)
	if u.MessageTime != nil {
		
		*MessageTime = timeutil.Strftime(u.MessageTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		MessageTime = nil
	}
	

	return json.Marshal(&struct { 
		Message *Conversationmessageeventtopicurireference `json:"message,omitempty"`
		
		MessageTime *string `json:"messageTime,omitempty"`
		
		MessageSegmentCount *int `json:"messageSegmentCount,omitempty"`
		
		MessageStatus *string `json:"messageStatus,omitempty"`
		
		Media *[]Conversationmessageeventtopicmessagemedia `json:"media,omitempty"`
		
		Stickers *[]Conversationmessageeventtopicmessagesticker `json:"stickers,omitempty"`
		*Alias
	}{ 
		Message: u.Message,
		
		MessageTime: MessageTime,
		
		MessageSegmentCount: u.MessageSegmentCount,
		
		MessageStatus: u.MessageStatus,
		
		Media: u.Media,
		
		Stickers: u.Stickers,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Conversationmessageeventtopicmessagedetails) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

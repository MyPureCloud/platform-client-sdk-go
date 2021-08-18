package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationmessageeventtopicmessagedetails
type Queueconversationmessageeventtopicmessagedetails struct { 
	// Message
	Message *Queueconversationmessageeventtopicurireference `json:"message,omitempty"`


	// MessageTime
	MessageTime *time.Time `json:"messageTime,omitempty"`


	// MessageSegmentCount
	MessageSegmentCount *int `json:"messageSegmentCount,omitempty"`


	// MessageStatus
	MessageStatus *string `json:"messageStatus,omitempty"`


	// Media
	Media *[]Queueconversationmessageeventtopicmessagemedia `json:"media,omitempty"`


	// Stickers
	Stickers *[]Queueconversationmessageeventtopicmessagesticker `json:"stickers,omitempty"`

}

func (u *Queueconversationmessageeventtopicmessagedetails) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queueconversationmessageeventtopicmessagedetails

	
	MessageTime := new(string)
	if u.MessageTime != nil {
		
		*MessageTime = timeutil.Strftime(u.MessageTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		MessageTime = nil
	}
	

	return json.Marshal(&struct { 
		Message *Queueconversationmessageeventtopicurireference `json:"message,omitempty"`
		
		MessageTime *string `json:"messageTime,omitempty"`
		
		MessageSegmentCount *int `json:"messageSegmentCount,omitempty"`
		
		MessageStatus *string `json:"messageStatus,omitempty"`
		
		Media *[]Queueconversationmessageeventtopicmessagemedia `json:"media,omitempty"`
		
		Stickers *[]Queueconversationmessageeventtopicmessagesticker `json:"stickers,omitempty"`
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
func (o *Queueconversationmessageeventtopicmessagedetails) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationvideoeventtopicmessagedetails
type Queueconversationvideoeventtopicmessagedetails struct { 
	// MessageId
	MessageId *string `json:"messageId,omitempty"`


	// MessageTime
	MessageTime *time.Time `json:"messageTime,omitempty"`


	// MessageStatus
	MessageStatus *string `json:"messageStatus,omitempty"`


	// MessageSegmentCount
	MessageSegmentCount *int `json:"messageSegmentCount,omitempty"`


	// Media
	Media *[]Queueconversationvideoeventtopicmessagemedia `json:"media,omitempty"`


	// Stickers
	Stickers *[]Queueconversationvideoeventtopicmessagesticker `json:"stickers,omitempty"`

}

func (u *Queueconversationvideoeventtopicmessagedetails) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queueconversationvideoeventtopicmessagedetails

	
	MessageTime := new(string)
	if u.MessageTime != nil {
		
		*MessageTime = timeutil.Strftime(u.MessageTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		MessageTime = nil
	}
	

	return json.Marshal(&struct { 
		MessageId *string `json:"messageId,omitempty"`
		
		MessageTime *string `json:"messageTime,omitempty"`
		
		MessageStatus *string `json:"messageStatus,omitempty"`
		
		MessageSegmentCount *int `json:"messageSegmentCount,omitempty"`
		
		Media *[]Queueconversationvideoeventtopicmessagemedia `json:"media,omitempty"`
		
		Stickers *[]Queueconversationvideoeventtopicmessagesticker `json:"stickers,omitempty"`
		*Alias
	}{ 
		MessageId: u.MessageId,
		
		MessageTime: MessageTime,
		
		MessageStatus: u.MessageStatus,
		
		MessageSegmentCount: u.MessageSegmentCount,
		
		Media: u.Media,
		
		Stickers: u.Stickers,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Queueconversationvideoeventtopicmessagedetails) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

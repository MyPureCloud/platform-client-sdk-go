package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationsocialexpressioneventtopicmessagedetails
type Queueconversationsocialexpressioneventtopicmessagedetails struct { 
	// MessageId
	MessageId *string `json:"messageId,omitempty"`


	// MessageTime
	MessageTime *time.Time `json:"messageTime,omitempty"`


	// MessageStatus
	MessageStatus *string `json:"messageStatus,omitempty"`


	// MessageSegmentCount
	MessageSegmentCount *int `json:"messageSegmentCount,omitempty"`


	// Media
	Media *[]Queueconversationsocialexpressioneventtopicmessagemedia `json:"media,omitempty"`


	// Stickers
	Stickers *[]Queueconversationsocialexpressioneventtopicmessagesticker `json:"stickers,omitempty"`

}

func (u *Queueconversationsocialexpressioneventtopicmessagedetails) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queueconversationsocialexpressioneventtopicmessagedetails

	
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
		
		Media *[]Queueconversationsocialexpressioneventtopicmessagemedia `json:"media,omitempty"`
		
		Stickers *[]Queueconversationsocialexpressioneventtopicmessagesticker `json:"stickers,omitempty"`
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
func (o *Queueconversationsocialexpressioneventtopicmessagedetails) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

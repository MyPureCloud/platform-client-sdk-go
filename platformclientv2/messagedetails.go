package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Messagedetails
type Messagedetails struct { 
	// MessageId - UUID identifying the message media.
	MessageId *string `json:"messageId,omitempty"`


	// MessageURI - A URI for this message entity.
	MessageURI *string `json:"messageURI,omitempty"`


	// MessageStatus - Indicates the delivery status of the message.
	MessageStatus *string `json:"messageStatus,omitempty"`


	// MessageSegmentCount - The message segment count, greater than 1 if the message content was split into multiple parts for this message type, e.g. SMS character limits.
	MessageSegmentCount *int `json:"messageSegmentCount,omitempty"`


	// MessageTime - The time when the message was sent or received. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	MessageTime *time.Time `json:"messageTime,omitempty"`


	// Media - The media (images, files, etc) associated with this message, if any
	Media *[]Messagemedia `json:"media,omitempty"`


	// Stickers - One or more stickers associated with this message, if any
	Stickers *[]Messagesticker `json:"stickers,omitempty"`

}

func (u *Messagedetails) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Messagedetails

	
	MessageTime := new(string)
	if u.MessageTime != nil {
		
		*MessageTime = timeutil.Strftime(u.MessageTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		MessageTime = nil
	}
	

	return json.Marshal(&struct { 
		MessageId *string `json:"messageId,omitempty"`
		
		MessageURI *string `json:"messageURI,omitempty"`
		
		MessageStatus *string `json:"messageStatus,omitempty"`
		
		MessageSegmentCount *int `json:"messageSegmentCount,omitempty"`
		
		MessageTime *string `json:"messageTime,omitempty"`
		
		Media *[]Messagemedia `json:"media,omitempty"`
		
		Stickers *[]Messagesticker `json:"stickers,omitempty"`
		*Alias
	}{ 
		MessageId: u.MessageId,
		
		MessageURI: u.MessageURI,
		
		MessageStatus: u.MessageStatus,
		
		MessageSegmentCount: u.MessageSegmentCount,
		
		MessageTime: MessageTime,
		
		Media: u.Media,
		
		Stickers: u.Stickers,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Messagedetails) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

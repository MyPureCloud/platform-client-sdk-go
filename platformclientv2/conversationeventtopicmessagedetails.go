package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationeventtopicmessagedetails
type Conversationeventtopicmessagedetails struct { 
	// MessageId
	MessageId *string `json:"messageId,omitempty"`


	// MessageTime
	MessageTime *time.Time `json:"messageTime,omitempty"`


	// MessageStatus
	MessageStatus *string `json:"messageStatus,omitempty"`


	// MessageSegmentCount
	MessageSegmentCount *int `json:"messageSegmentCount,omitempty"`


	// Media
	Media *[]Conversationeventtopicmessagemedia `json:"media,omitempty"`


	// Stickers
	Stickers *[]Conversationeventtopicmessagesticker `json:"stickers,omitempty"`

}

// String returns a JSON representation of the model
func (o *Conversationeventtopicmessagedetails) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationeventtopicmessagedetails
type Queueconversationeventtopicmessagedetails struct { 
	// MessageId
	MessageId *string `json:"messageId,omitempty"`


	// MessageTime
	MessageTime *time.Time `json:"messageTime,omitempty"`


	// MessageStatus
	MessageStatus *string `json:"messageStatus,omitempty"`


	// MessageSegmentCount
	MessageSegmentCount *int `json:"messageSegmentCount,omitempty"`


	// Media
	Media *[]Queueconversationeventtopicmessagemedia `json:"media,omitempty"`


	// Stickers
	Stickers *[]Queueconversationeventtopicmessagesticker `json:"stickers,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queueconversationeventtopicmessagedetails) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

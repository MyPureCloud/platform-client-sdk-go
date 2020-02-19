package platformclientv2
import (
	"time"
	"encoding/json"
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
	MessageSegmentCount *int32 `json:"messageSegmentCount,omitempty"`


	// Media
	Media *[]Queueconversationsocialexpressioneventtopicmessagemedia `json:"media,omitempty"`


	// Stickers
	Stickers *[]Queueconversationsocialexpressioneventtopicmessagesticker `json:"stickers,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queueconversationsocialexpressioneventtopicmessagedetails) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

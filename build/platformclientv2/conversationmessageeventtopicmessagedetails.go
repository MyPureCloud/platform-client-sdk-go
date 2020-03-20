package platformclientv2
import (
	"time"
	"encoding/json"
)

// Conversationmessageeventtopicmessagedetails
type Conversationmessageeventtopicmessagedetails struct { 
	// Message
	Message *Conversationmessageeventtopicurireference `json:"message,omitempty"`


	// MessageTime
	MessageTime *time.Time `json:"messageTime,omitempty"`


	// MessageSegmentCount
	MessageSegmentCount *int32 `json:"messageSegmentCount,omitempty"`


	// MessageStatus
	MessageStatus *string `json:"messageStatus,omitempty"`


	// Media
	Media *[]Conversationmessageeventtopicmessagemedia `json:"media,omitempty"`


	// Stickers
	Stickers *[]Conversationmessageeventtopicmessagesticker `json:"stickers,omitempty"`

}

// String returns a JSON representation of the model
func (o *Conversationmessageeventtopicmessagedetails) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

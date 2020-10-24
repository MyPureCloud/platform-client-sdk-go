package platformclientv2
import (
	"time"
	"encoding/json"
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
	MessageSegmentCount *int32 `json:"messageSegmentCount,omitempty"`


	// MessageTime - The time when the message was sent or received. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	MessageTime *time.Time `json:"messageTime,omitempty"`


	// Media - The media (images, files, etc) associated with this message, if any
	Media *[]Messagemedia `json:"media,omitempty"`


	// Stickers - One or more stickers associated with this message, if any
	Stickers *[]Messagesticker `json:"stickers,omitempty"`

}

// String returns a JSON representation of the model
func (o *Messagedetails) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

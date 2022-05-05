package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationeventtopicmessagedetails
type Queueconversationeventtopicmessagedetails struct { 
	// MessageId - UUID identifying the message media.
	MessageId *string `json:"messageId,omitempty"`


	// MessageTime - The time when the message was sent or received.
	MessageTime *time.Time `json:"messageTime,omitempty"`


	// MessageStatus - Indicates the delivery status of the message.
	MessageStatus *string `json:"messageStatus,omitempty"`


	// MessageSegmentCount - The message segment count, greater than 1 if the message content was split into multiple parts for this message type, e.g. SMS character limits.
	MessageSegmentCount *int `json:"messageSegmentCount,omitempty"`


	// Media - The media (images, files, etc) associated with this message, if any
	Media *[]Queueconversationeventtopicmessagemedia `json:"media,omitempty"`


	// ErrorInfo - Detailed information about an error response.
	ErrorInfo *Queueconversationeventtopicerrordetails `json:"errorInfo,omitempty"`


	// Stickers - A list of stickers included in the message
	Stickers *[]Queueconversationeventtopicmessagesticker `json:"stickers,omitempty"`

}

func (o *Queueconversationeventtopicmessagedetails) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queueconversationeventtopicmessagedetails
	
	MessageTime := new(string)
	if o.MessageTime != nil {
		
		*MessageTime = timeutil.Strftime(o.MessageTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		MessageTime = nil
	}
	
	return json.Marshal(&struct { 
		MessageId *string `json:"messageId,omitempty"`
		
		MessageTime *string `json:"messageTime,omitempty"`
		
		MessageStatus *string `json:"messageStatus,omitempty"`
		
		MessageSegmentCount *int `json:"messageSegmentCount,omitempty"`
		
		Media *[]Queueconversationeventtopicmessagemedia `json:"media,omitempty"`
		
		ErrorInfo *Queueconversationeventtopicerrordetails `json:"errorInfo,omitempty"`
		
		Stickers *[]Queueconversationeventtopicmessagesticker `json:"stickers,omitempty"`
		*Alias
	}{ 
		MessageId: o.MessageId,
		
		MessageTime: MessageTime,
		
		MessageStatus: o.MessageStatus,
		
		MessageSegmentCount: o.MessageSegmentCount,
		
		Media: o.Media,
		
		ErrorInfo: o.ErrorInfo,
		
		Stickers: o.Stickers,
		Alias:    (*Alias)(o),
	})
}

func (o *Queueconversationeventtopicmessagedetails) UnmarshalJSON(b []byte) error {
	var QueueconversationeventtopicmessagedetailsMap map[string]interface{}
	err := json.Unmarshal(b, &QueueconversationeventtopicmessagedetailsMap)
	if err != nil {
		return err
	}
	
	if MessageId, ok := QueueconversationeventtopicmessagedetailsMap["messageId"].(string); ok {
		o.MessageId = &MessageId
	}
	
	if messageTimeString, ok := QueueconversationeventtopicmessagedetailsMap["messageTime"].(string); ok {
		MessageTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", messageTimeString)
		o.MessageTime = &MessageTime
	}
	
	if MessageStatus, ok := QueueconversationeventtopicmessagedetailsMap["messageStatus"].(string); ok {
		o.MessageStatus = &MessageStatus
	}
	
	if MessageSegmentCount, ok := QueueconversationeventtopicmessagedetailsMap["messageSegmentCount"].(float64); ok {
		MessageSegmentCountInt := int(MessageSegmentCount)
		o.MessageSegmentCount = &MessageSegmentCountInt
	}
	
	if Media, ok := QueueconversationeventtopicmessagedetailsMap["media"].([]interface{}); ok {
		MediaString, _ := json.Marshal(Media)
		json.Unmarshal(MediaString, &o.Media)
	}
	
	if ErrorInfo, ok := QueueconversationeventtopicmessagedetailsMap["errorInfo"].(map[string]interface{}); ok {
		ErrorInfoString, _ := json.Marshal(ErrorInfo)
		json.Unmarshal(ErrorInfoString, &o.ErrorInfo)
	}
	
	if Stickers, ok := QueueconversationeventtopicmessagedetailsMap["stickers"].([]interface{}); ok {
		StickersString, _ := json.Marshal(Stickers)
		json.Unmarshal(StickersString, &o.Stickers)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Queueconversationeventtopicmessagedetails) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

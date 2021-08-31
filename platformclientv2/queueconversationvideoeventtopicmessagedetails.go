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

func (o *Queueconversationvideoeventtopicmessagedetails) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queueconversationvideoeventtopicmessagedetails
	
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
		
		Media *[]Queueconversationvideoeventtopicmessagemedia `json:"media,omitempty"`
		
		Stickers *[]Queueconversationvideoeventtopicmessagesticker `json:"stickers,omitempty"`
		*Alias
	}{ 
		MessageId: o.MessageId,
		
		MessageTime: MessageTime,
		
		MessageStatus: o.MessageStatus,
		
		MessageSegmentCount: o.MessageSegmentCount,
		
		Media: o.Media,
		
		Stickers: o.Stickers,
		Alias:    (*Alias)(o),
	})
}

func (o *Queueconversationvideoeventtopicmessagedetails) UnmarshalJSON(b []byte) error {
	var QueueconversationvideoeventtopicmessagedetailsMap map[string]interface{}
	err := json.Unmarshal(b, &QueueconversationvideoeventtopicmessagedetailsMap)
	if err != nil {
		return err
	}
	
	if MessageId, ok := QueueconversationvideoeventtopicmessagedetailsMap["messageId"].(string); ok {
		o.MessageId = &MessageId
	}
	
	if messageTimeString, ok := QueueconversationvideoeventtopicmessagedetailsMap["messageTime"].(string); ok {
		MessageTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", messageTimeString)
		o.MessageTime = &MessageTime
	}
	
	if MessageStatus, ok := QueueconversationvideoeventtopicmessagedetailsMap["messageStatus"].(string); ok {
		o.MessageStatus = &MessageStatus
	}
	
	if MessageSegmentCount, ok := QueueconversationvideoeventtopicmessagedetailsMap["messageSegmentCount"].(float64); ok {
		MessageSegmentCountInt := int(MessageSegmentCount)
		o.MessageSegmentCount = &MessageSegmentCountInt
	}
	
	if Media, ok := QueueconversationvideoeventtopicmessagedetailsMap["media"].([]interface{}); ok {
		MediaString, _ := json.Marshal(Media)
		json.Unmarshal(MediaString, &o.Media)
	}
	
	if Stickers, ok := QueueconversationvideoeventtopicmessagedetailsMap["stickers"].([]interface{}); ok {
		StickersString, _ := json.Marshal(Stickers)
		json.Unmarshal(StickersString, &o.Stickers)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Queueconversationvideoeventtopicmessagedetails) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

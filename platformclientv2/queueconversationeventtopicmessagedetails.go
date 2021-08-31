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
		
		Stickers *[]Queueconversationeventtopicmessagesticker `json:"stickers,omitempty"`
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

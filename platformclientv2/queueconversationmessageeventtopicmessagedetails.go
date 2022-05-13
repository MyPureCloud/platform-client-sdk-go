package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationmessageeventtopicmessagedetails
type Queueconversationmessageeventtopicmessagedetails struct { 
	// Message
	Message *Queueconversationmessageeventtopicurireference `json:"message,omitempty"`


	// MessageTime
	MessageTime *time.Time `json:"messageTime,omitempty"`


	// MessageSegmentCount
	MessageSegmentCount *int `json:"messageSegmentCount,omitempty"`


	// MessageStatus
	MessageStatus *string `json:"messageStatus,omitempty"`


	// Media
	Media *[]Queueconversationmessageeventtopicmessagemedia `json:"media,omitempty"`


	// Stickers
	Stickers *[]Queueconversationmessageeventtopicmessagesticker `json:"stickers,omitempty"`


	// ErrorInfo
	ErrorInfo *Queueconversationmessageeventtopicerrordetails `json:"errorInfo,omitempty"`

}

func (o *Queueconversationmessageeventtopicmessagedetails) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queueconversationmessageeventtopicmessagedetails
	
	MessageTime := new(string)
	if o.MessageTime != nil {
		
		*MessageTime = timeutil.Strftime(o.MessageTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		MessageTime = nil
	}
	
	return json.Marshal(&struct { 
		Message *Queueconversationmessageeventtopicurireference `json:"message,omitempty"`
		
		MessageTime *string `json:"messageTime,omitempty"`
		
		MessageSegmentCount *int `json:"messageSegmentCount,omitempty"`
		
		MessageStatus *string `json:"messageStatus,omitempty"`
		
		Media *[]Queueconversationmessageeventtopicmessagemedia `json:"media,omitempty"`
		
		Stickers *[]Queueconversationmessageeventtopicmessagesticker `json:"stickers,omitempty"`
		
		ErrorInfo *Queueconversationmessageeventtopicerrordetails `json:"errorInfo,omitempty"`
		*Alias
	}{ 
		Message: o.Message,
		
		MessageTime: MessageTime,
		
		MessageSegmentCount: o.MessageSegmentCount,
		
		MessageStatus: o.MessageStatus,
		
		Media: o.Media,
		
		Stickers: o.Stickers,
		
		ErrorInfo: o.ErrorInfo,
		Alias:    (*Alias)(o),
	})
}

func (o *Queueconversationmessageeventtopicmessagedetails) UnmarshalJSON(b []byte) error {
	var QueueconversationmessageeventtopicmessagedetailsMap map[string]interface{}
	err := json.Unmarshal(b, &QueueconversationmessageeventtopicmessagedetailsMap)
	if err != nil {
		return err
	}
	
	if Message, ok := QueueconversationmessageeventtopicmessagedetailsMap["message"].(map[string]interface{}); ok {
		MessageString, _ := json.Marshal(Message)
		json.Unmarshal(MessageString, &o.Message)
	}
	
	if messageTimeString, ok := QueueconversationmessageeventtopicmessagedetailsMap["messageTime"].(string); ok {
		MessageTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", messageTimeString)
		o.MessageTime = &MessageTime
	}
	
	if MessageSegmentCount, ok := QueueconversationmessageeventtopicmessagedetailsMap["messageSegmentCount"].(float64); ok {
		MessageSegmentCountInt := int(MessageSegmentCount)
		o.MessageSegmentCount = &MessageSegmentCountInt
	}
	
	if MessageStatus, ok := QueueconversationmessageeventtopicmessagedetailsMap["messageStatus"].(string); ok {
		o.MessageStatus = &MessageStatus
	}
    
	if Media, ok := QueueconversationmessageeventtopicmessagedetailsMap["media"].([]interface{}); ok {
		MediaString, _ := json.Marshal(Media)
		json.Unmarshal(MediaString, &o.Media)
	}
	
	if Stickers, ok := QueueconversationmessageeventtopicmessagedetailsMap["stickers"].([]interface{}); ok {
		StickersString, _ := json.Marshal(Stickers)
		json.Unmarshal(StickersString, &o.Stickers)
	}
	
	if ErrorInfo, ok := QueueconversationmessageeventtopicmessagedetailsMap["errorInfo"].(map[string]interface{}); ok {
		ErrorInfoString, _ := json.Marshal(ErrorInfo)
		json.Unmarshal(ErrorInfoString, &o.ErrorInfo)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Queueconversationmessageeventtopicmessagedetails) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationmessageeventtopicmessagedetails
type Conversationmessageeventtopicmessagedetails struct { 
	// Message
	Message *Conversationmessageeventtopicurireference `json:"message,omitempty"`


	// MessageTime
	MessageTime *time.Time `json:"messageTime,omitempty"`


	// MessageSegmentCount
	MessageSegmentCount *int `json:"messageSegmentCount,omitempty"`


	// MessageStatus
	MessageStatus *string `json:"messageStatus,omitempty"`


	// Media
	Media *[]Conversationmessageeventtopicmessagemedia `json:"media,omitempty"`


	// Stickers
	Stickers *[]Conversationmessageeventtopicmessagesticker `json:"stickers,omitempty"`


	// ErrorInfo
	ErrorInfo *Conversationmessageeventtopicerrordetails `json:"errorInfo,omitempty"`

}

func (o *Conversationmessageeventtopicmessagedetails) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationmessageeventtopicmessagedetails
	
	MessageTime := new(string)
	if o.MessageTime != nil {
		
		*MessageTime = timeutil.Strftime(o.MessageTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		MessageTime = nil
	}
	
	return json.Marshal(&struct { 
		Message *Conversationmessageeventtopicurireference `json:"message,omitempty"`
		
		MessageTime *string `json:"messageTime,omitempty"`
		
		MessageSegmentCount *int `json:"messageSegmentCount,omitempty"`
		
		MessageStatus *string `json:"messageStatus,omitempty"`
		
		Media *[]Conversationmessageeventtopicmessagemedia `json:"media,omitempty"`
		
		Stickers *[]Conversationmessageeventtopicmessagesticker `json:"stickers,omitempty"`
		
		ErrorInfo *Conversationmessageeventtopicerrordetails `json:"errorInfo,omitempty"`
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

func (o *Conversationmessageeventtopicmessagedetails) UnmarshalJSON(b []byte) error {
	var ConversationmessageeventtopicmessagedetailsMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationmessageeventtopicmessagedetailsMap)
	if err != nil {
		return err
	}
	
	if Message, ok := ConversationmessageeventtopicmessagedetailsMap["message"].(map[string]interface{}); ok {
		MessageString, _ := json.Marshal(Message)
		json.Unmarshal(MessageString, &o.Message)
	}
	
	if messageTimeString, ok := ConversationmessageeventtopicmessagedetailsMap["messageTime"].(string); ok {
		MessageTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", messageTimeString)
		o.MessageTime = &MessageTime
	}
	
	if MessageSegmentCount, ok := ConversationmessageeventtopicmessagedetailsMap["messageSegmentCount"].(float64); ok {
		MessageSegmentCountInt := int(MessageSegmentCount)
		o.MessageSegmentCount = &MessageSegmentCountInt
	}
	
	if MessageStatus, ok := ConversationmessageeventtopicmessagedetailsMap["messageStatus"].(string); ok {
		o.MessageStatus = &MessageStatus
	}
	
	if Media, ok := ConversationmessageeventtopicmessagedetailsMap["media"].([]interface{}); ok {
		MediaString, _ := json.Marshal(Media)
		json.Unmarshal(MediaString, &o.Media)
	}
	
	if Stickers, ok := ConversationmessageeventtopicmessagedetailsMap["stickers"].([]interface{}); ok {
		StickersString, _ := json.Marshal(Stickers)
		json.Unmarshal(StickersString, &o.Stickers)
	}
	
	if ErrorInfo, ok := ConversationmessageeventtopicmessagedetailsMap["errorInfo"].(map[string]interface{}); ok {
		ErrorInfoString, _ := json.Marshal(ErrorInfo)
		json.Unmarshal(ErrorInfoString, &o.ErrorInfo)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationmessageeventtopicmessagedetails) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

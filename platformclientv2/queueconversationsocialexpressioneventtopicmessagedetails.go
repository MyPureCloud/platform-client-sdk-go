package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
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
	MessageSegmentCount *int `json:"messageSegmentCount,omitempty"`


	// Media
	Media *[]Queueconversationsocialexpressioneventtopicmessagemedia `json:"media,omitempty"`


	// Stickers
	Stickers *[]Queueconversationsocialexpressioneventtopicmessagesticker `json:"stickers,omitempty"`

}

func (o *Queueconversationsocialexpressioneventtopicmessagedetails) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queueconversationsocialexpressioneventtopicmessagedetails
	
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
		
		Media *[]Queueconversationsocialexpressioneventtopicmessagemedia `json:"media,omitempty"`
		
		Stickers *[]Queueconversationsocialexpressioneventtopicmessagesticker `json:"stickers,omitempty"`
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

func (o *Queueconversationsocialexpressioneventtopicmessagedetails) UnmarshalJSON(b []byte) error {
	var QueueconversationsocialexpressioneventtopicmessagedetailsMap map[string]interface{}
	err := json.Unmarshal(b, &QueueconversationsocialexpressioneventtopicmessagedetailsMap)
	if err != nil {
		return err
	}
	
	if MessageId, ok := QueueconversationsocialexpressioneventtopicmessagedetailsMap["messageId"].(string); ok {
		o.MessageId = &MessageId
	}
	
	if messageTimeString, ok := QueueconversationsocialexpressioneventtopicmessagedetailsMap["messageTime"].(string); ok {
		MessageTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", messageTimeString)
		o.MessageTime = &MessageTime
	}
	
	if MessageStatus, ok := QueueconversationsocialexpressioneventtopicmessagedetailsMap["messageStatus"].(string); ok {
		o.MessageStatus = &MessageStatus
	}
	
	if MessageSegmentCount, ok := QueueconversationsocialexpressioneventtopicmessagedetailsMap["messageSegmentCount"].(float64); ok {
		MessageSegmentCountInt := int(MessageSegmentCount)
		o.MessageSegmentCount = &MessageSegmentCountInt
	}
	
	if Media, ok := QueueconversationsocialexpressioneventtopicmessagedetailsMap["media"].([]interface{}); ok {
		MediaString, _ := json.Marshal(Media)
		json.Unmarshal(MediaString, &o.Media)
	}
	
	if Stickers, ok := QueueconversationsocialexpressioneventtopicmessagedetailsMap["stickers"].([]interface{}); ok {
		StickersString, _ := json.Marshal(Stickers)
		json.Unmarshal(StickersString, &o.Stickers)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Queueconversationsocialexpressioneventtopicmessagedetails) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

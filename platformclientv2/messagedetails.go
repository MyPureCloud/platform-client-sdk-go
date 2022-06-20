package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
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
	MessageSegmentCount *int `json:"messageSegmentCount,omitempty"`


	// MessageTime - The time when the message was sent or received. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	MessageTime *time.Time `json:"messageTime,omitempty"`


	// Media - The media (images, files, etc) associated with this message, if any
	Media *[]Messagemedia `json:"media,omitempty"`


	// Stickers - One or more stickers associated with this message, if any
	Stickers *[]Messagesticker `json:"stickers,omitempty"`


	// MessageMetadata - Information that describes the content of the message, if any
	MessageMetadata *Conversationmessagemetadata `json:"messageMetadata,omitempty"`


	// ErrorInfo - Provider specific error information for a communication.
	ErrorInfo *Errorbody `json:"errorInfo,omitempty"`

}

func (o *Messagedetails) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Messagedetails
	
	MessageTime := new(string)
	if o.MessageTime != nil {
		
		*MessageTime = timeutil.Strftime(o.MessageTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		MessageTime = nil
	}
	
	return json.Marshal(&struct { 
		MessageId *string `json:"messageId,omitempty"`
		
		MessageURI *string `json:"messageURI,omitempty"`
		
		MessageStatus *string `json:"messageStatus,omitempty"`
		
		MessageSegmentCount *int `json:"messageSegmentCount,omitempty"`
		
		MessageTime *string `json:"messageTime,omitempty"`
		
		Media *[]Messagemedia `json:"media,omitempty"`
		
		Stickers *[]Messagesticker `json:"stickers,omitempty"`
		
		MessageMetadata *Conversationmessagemetadata `json:"messageMetadata,omitempty"`
		
		ErrorInfo *Errorbody `json:"errorInfo,omitempty"`
		*Alias
	}{ 
		MessageId: o.MessageId,
		
		MessageURI: o.MessageURI,
		
		MessageStatus: o.MessageStatus,
		
		MessageSegmentCount: o.MessageSegmentCount,
		
		MessageTime: MessageTime,
		
		Media: o.Media,
		
		Stickers: o.Stickers,
		
		MessageMetadata: o.MessageMetadata,
		
		ErrorInfo: o.ErrorInfo,
		Alias:    (*Alias)(o),
	})
}

func (o *Messagedetails) UnmarshalJSON(b []byte) error {
	var MessagedetailsMap map[string]interface{}
	err := json.Unmarshal(b, &MessagedetailsMap)
	if err != nil {
		return err
	}
	
	if MessageId, ok := MessagedetailsMap["messageId"].(string); ok {
		o.MessageId = &MessageId
	}
    
	if MessageURI, ok := MessagedetailsMap["messageURI"].(string); ok {
		o.MessageURI = &MessageURI
	}
    
	if MessageStatus, ok := MessagedetailsMap["messageStatus"].(string); ok {
		o.MessageStatus = &MessageStatus
	}
    
	if MessageSegmentCount, ok := MessagedetailsMap["messageSegmentCount"].(float64); ok {
		MessageSegmentCountInt := int(MessageSegmentCount)
		o.MessageSegmentCount = &MessageSegmentCountInt
	}
	
	if messageTimeString, ok := MessagedetailsMap["messageTime"].(string); ok {
		MessageTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", messageTimeString)
		o.MessageTime = &MessageTime
	}
	
	if Media, ok := MessagedetailsMap["media"].([]interface{}); ok {
		MediaString, _ := json.Marshal(Media)
		json.Unmarshal(MediaString, &o.Media)
	}
	
	if Stickers, ok := MessagedetailsMap["stickers"].([]interface{}); ok {
		StickersString, _ := json.Marshal(Stickers)
		json.Unmarshal(StickersString, &o.Stickers)
	}
	
	if MessageMetadata, ok := MessagedetailsMap["messageMetadata"].(map[string]interface{}); ok {
		MessageMetadataString, _ := json.Marshal(MessageMetadata)
		json.Unmarshal(MessageMetadataString, &o.MessageMetadata)
	}
	
	if ErrorInfo, ok := MessagedetailsMap["errorInfo"].(map[string]interface{}); ok {
		ErrorInfoString, _ := json.Marshal(ErrorInfo)
		json.Unmarshal(ErrorInfoString, &o.ErrorInfo)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Messagedetails) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

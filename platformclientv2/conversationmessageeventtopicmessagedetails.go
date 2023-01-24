package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationmessageeventtopicmessagedetails
type Conversationmessageeventtopicmessagedetails struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
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

	// MessageMetadata
	MessageMetadata *Conversationmessageeventtopicmessagemetadata `json:"messageMetadata,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Conversationmessageeventtopicmessagedetails) SetField(field string, fieldValue interface{}) {
	// Get Value object for field
	target := reflect.ValueOf(o)
	targetField := reflect.Indirect(target).FieldByName(field)

	// Set value
	if fieldValue != nil {
		targetField.Set(reflect.ValueOf(fieldValue))
	} else {
		// Must create a new Value (creates **type) then get its element (*type), which will be nil pointer of the appropriate type
		x := reflect.Indirect(reflect.New(targetField.Type()))
		targetField.Set(x)
	}

	// Add field to set field names list
	if o.SetFieldNames == nil {
		o.SetFieldNames = make(map[string]bool)
	}
	o.SetFieldNames[field] = true
}

func (o Conversationmessageeventtopicmessagedetails) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "MessageTime", }
		localDateTimeFields := []string{  }
		dateFields := []string{  }

		// Construct object
		newObj := make(map[string]interface{})
		for fieldName := range o.SetFieldNames {
			// Get initial field value
			fieldValue := val.FieldByName(fieldName).Interface()

			// Apply value formatting overrides
			if contains(dateTimeFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%dT%H:%M:%S.%fZ")
			} else if contains(localDateTimeFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%dT%H:%M:%S.%f")
			} else if contains(dateFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%d")
			}

			// Assign value to field using JSON tag name
			newObj[getFieldName(reflect.TypeOf(&o), fieldName)] = fieldValue
		}

		// Marshal and return dynamically constructed interface
		return json.Marshal(newObj)
	}

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
		
		MessageMetadata *Conversationmessageeventtopicmessagemetadata `json:"messageMetadata,omitempty"`
		Alias
	}{ 
		Message: o.Message,
		
		MessageTime: MessageTime,
		
		MessageSegmentCount: o.MessageSegmentCount,
		
		MessageStatus: o.MessageStatus,
		
		Media: o.Media,
		
		Stickers: o.Stickers,
		
		ErrorInfo: o.ErrorInfo,
		
		MessageMetadata: o.MessageMetadata,
		Alias:    (Alias)(o),
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
	
	if MessageMetadata, ok := ConversationmessageeventtopicmessagedetailsMap["messageMetadata"].(map[string]interface{}); ok {
		MessageMetadataString, _ := json.Marshal(MessageMetadata)
		json.Unmarshal(MessageMetadataString, &o.MessageMetadata)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationmessageeventtopicmessagedetails) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

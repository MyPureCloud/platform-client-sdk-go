package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationmessageeventtopicmessagedetails
type Queueconversationmessageeventtopicmessagedetails struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
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

	// MessageMetadata
	MessageMetadata *Queueconversationmessageeventtopicmessagemetadata `json:"messageMetadata,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Queueconversationmessageeventtopicmessagedetails) SetField(field string, fieldValue interface{}) {
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

func (o Queueconversationmessageeventtopicmessagedetails) MarshalJSON() ([]byte, error) {
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
			if fieldValue == nil || reflect.ValueOf(fieldValue).IsNil()  {
				// Do nothing. Just catching this case to avoid trying to custom serialize a nil value
			} else if contains(dateTimeFields, fieldName) {
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
		
		MessageMetadata *Queueconversationmessageeventtopicmessagemetadata `json:"messageMetadata,omitempty"`
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
	
	if MessageMetadata, ok := QueueconversationmessageeventtopicmessagedetailsMap["messageMetadata"].(map[string]interface{}); ok {
		MessageMetadataString, _ := json.Marshal(MessageMetadata)
		json.Unmarshal(MessageMetadataString, &o.MessageMetadata)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Queueconversationmessageeventtopicmessagedetails) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

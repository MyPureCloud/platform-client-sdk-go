package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationsocialexpressioneventtopicmessagedetails
type Queueconversationsocialexpressioneventtopicmessagedetails struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// MessageId - UUID identifying the message media.
	MessageId *string `json:"messageId,omitempty"`

	// MessageTime - The time when the message was sent or received.
	MessageTime *time.Time `json:"messageTime,omitempty"`

	// MessageStatus - Indicates the delivery status of the message.
	MessageStatus *string `json:"messageStatus,omitempty"`

	// MessageSegmentCount - The message segment count, greater than 1 if the message content was split into multiple parts for this message type, e.g. SMS character limits.
	MessageSegmentCount *int `json:"messageSegmentCount,omitempty"`

	// Media - The media (images, files, etc) associated with this message, if any
	Media *[]Queueconversationsocialexpressioneventtopicmessagemedia `json:"media,omitempty"`

	// ErrorInfo - Detailed information about an error response.
	ErrorInfo *Queueconversationsocialexpressioneventtopicerrordetails `json:"errorInfo,omitempty"`

	// Stickers - A list of stickers included in the message
	Stickers *[]Queueconversationsocialexpressioneventtopicmessagesticker `json:"stickers,omitempty"`

	// MessageMetadata
	MessageMetadata *Queueconversationsocialexpressioneventtopicmessagemetadata `json:"messageMetadata,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Queueconversationsocialexpressioneventtopicmessagedetails) SetField(field string, fieldValue interface{}) {
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

func (o Queueconversationsocialexpressioneventtopicmessagedetails) MarshalJSON() ([]byte, error) {
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
		
		ErrorInfo *Queueconversationsocialexpressioneventtopicerrordetails `json:"errorInfo,omitempty"`
		
		Stickers *[]Queueconversationsocialexpressioneventtopicmessagesticker `json:"stickers,omitempty"`
		
		MessageMetadata *Queueconversationsocialexpressioneventtopicmessagemetadata `json:"messageMetadata,omitempty"`
		Alias
	}{ 
		MessageId: o.MessageId,
		
		MessageTime: MessageTime,
		
		MessageStatus: o.MessageStatus,
		
		MessageSegmentCount: o.MessageSegmentCount,
		
		Media: o.Media,
		
		ErrorInfo: o.ErrorInfo,
		
		Stickers: o.Stickers,
		
		MessageMetadata: o.MessageMetadata,
		Alias:    (Alias)(o),
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
	
	if ErrorInfo, ok := QueueconversationsocialexpressioneventtopicmessagedetailsMap["errorInfo"].(map[string]interface{}); ok {
		ErrorInfoString, _ := json.Marshal(ErrorInfo)
		json.Unmarshal(ErrorInfoString, &o.ErrorInfo)
	}
	
	if Stickers, ok := QueueconversationsocialexpressioneventtopicmessagedetailsMap["stickers"].([]interface{}); ok {
		StickersString, _ := json.Marshal(Stickers)
		json.Unmarshal(StickersString, &o.Stickers)
	}
	
	if MessageMetadata, ok := QueueconversationsocialexpressioneventtopicmessagedetailsMap["messageMetadata"].(map[string]interface{}); ok {
		MessageMetadataString, _ := json.Marshal(MessageMetadata)
		json.Unmarshal(MessageMetadataString, &o.MessageMetadata)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Queueconversationsocialexpressioneventtopicmessagedetails) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

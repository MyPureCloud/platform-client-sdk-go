package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Messagedetails
type Messagedetails struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
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

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Messagedetails) SetField(field string, fieldValue interface{}) {
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

func (o Messagedetails) MarshalJSON() ([]byte, error) {
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
		Alias
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
		Alias:    (Alias)(o),
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

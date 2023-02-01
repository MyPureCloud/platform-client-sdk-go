package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationmessagingchannel - Channel-specific information that describes the message and the message channel/provider.
type Conversationmessagingchannel struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The integration ID.
	Id *string `json:"id,omitempty"`

	// Platform - The provider type.
	Platform *string `json:"platform,omitempty"`

	// MessageId - Unique provider ID of the message such as a Facebook message ID.
	MessageId *string `json:"messageId,omitempty"`

	// To - Information about the recipient the message is sent to.
	To *Conversationmessagingtorecipient `json:"to,omitempty"`

	// From - Information about the recipient the message is received from.
	From *Conversationmessagingfromrecipient `json:"from,omitempty"`

	// Time - Original time of the event. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	Time *time.Time `json:"time,omitempty"`

	// DateModified - Time the message was edited. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`

	// DateDeleted - Time the message was deleted. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateDeleted *time.Time `json:"dateDeleted,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Conversationmessagingchannel) SetField(field string, fieldValue interface{}) {
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

func (o Conversationmessagingchannel) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "Time","DateModified","DateDeleted", }
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
	type Alias Conversationmessagingchannel
	
	Time := new(string)
	if o.Time != nil {
		
		*Time = timeutil.Strftime(o.Time, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		Time = nil
	}
	
	DateModified := new(string)
	if o.DateModified != nil {
		
		*DateModified = timeutil.Strftime(o.DateModified, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateModified = nil
	}
	
	DateDeleted := new(string)
	if o.DateDeleted != nil {
		
		*DateDeleted = timeutil.Strftime(o.DateDeleted, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateDeleted = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Platform *string `json:"platform,omitempty"`
		
		MessageId *string `json:"messageId,omitempty"`
		
		To *Conversationmessagingtorecipient `json:"to,omitempty"`
		
		From *Conversationmessagingfromrecipient `json:"from,omitempty"`
		
		Time *string `json:"time,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		DateDeleted *string `json:"dateDeleted,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Platform: o.Platform,
		
		MessageId: o.MessageId,
		
		To: o.To,
		
		From: o.From,
		
		Time: Time,
		
		DateModified: DateModified,
		
		DateDeleted: DateDeleted,
		Alias:    (Alias)(o),
	})
}

func (o *Conversationmessagingchannel) UnmarshalJSON(b []byte) error {
	var ConversationmessagingchannelMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationmessagingchannelMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ConversationmessagingchannelMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Platform, ok := ConversationmessagingchannelMap["platform"].(string); ok {
		o.Platform = &Platform
	}
    
	if MessageId, ok := ConversationmessagingchannelMap["messageId"].(string); ok {
		o.MessageId = &MessageId
	}
    
	if To, ok := ConversationmessagingchannelMap["to"].(map[string]interface{}); ok {
		ToString, _ := json.Marshal(To)
		json.Unmarshal(ToString, &o.To)
	}
	
	if From, ok := ConversationmessagingchannelMap["from"].(map[string]interface{}); ok {
		FromString, _ := json.Marshal(From)
		json.Unmarshal(FromString, &o.From)
	}
	
	if timeString, ok := ConversationmessagingchannelMap["time"].(string); ok {
		Time, _ := time.Parse("2006-01-02T15:04:05.999999Z", timeString)
		o.Time = &Time
	}
	
	if dateModifiedString, ok := ConversationmessagingchannelMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if dateDeletedString, ok := ConversationmessagingchannelMap["dateDeleted"].(string); ok {
		DateDeleted, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateDeletedString)
		o.DateDeleted = &DateDeleted
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationmessagingchannel) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

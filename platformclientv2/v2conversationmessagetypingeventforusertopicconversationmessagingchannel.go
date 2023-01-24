package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// V2conversationmessagetypingeventforusertopicconversationmessagingchannel
type V2conversationmessagetypingeventforusertopicconversationmessagingchannel struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id
	Id *string `json:"id,omitempty"`

	// Platform
	Platform *string `json:"platform,omitempty"`

	// MessageId
	MessageId *string `json:"messageId,omitempty"`

	// To
	To *V2conversationmessagetypingeventforusertopicconversationmessagingtorecipient `json:"to,omitempty"`

	// From
	From *V2conversationmessagetypingeventforusertopicconversationmessagingfromrecipient `json:"from,omitempty"`

	// Time
	Time *time.Time `json:"time,omitempty"`

	// DateModified
	DateModified *time.Time `json:"dateModified,omitempty"`

	// DateDeleted
	DateDeleted *time.Time `json:"dateDeleted,omitempty"`

	// Metadata
	Metadata *V2conversationmessagetypingeventforusertopicconversationmessagingchannelmetadata `json:"metadata,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *V2conversationmessagetypingeventforusertopicconversationmessagingchannel) SetField(field string, fieldValue interface{}) {
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

func (o V2conversationmessagetypingeventforusertopicconversationmessagingchannel) MarshalJSON() ([]byte, error) {
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
	type Alias V2conversationmessagetypingeventforusertopicconversationmessagingchannel
	
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
		
		To *V2conversationmessagetypingeventforusertopicconversationmessagingtorecipient `json:"to,omitempty"`
		
		From *V2conversationmessagetypingeventforusertopicconversationmessagingfromrecipient `json:"from,omitempty"`
		
		Time *string `json:"time,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		DateDeleted *string `json:"dateDeleted,omitempty"`
		
		Metadata *V2conversationmessagetypingeventforusertopicconversationmessagingchannelmetadata `json:"metadata,omitempty"`
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
		
		Metadata: o.Metadata,
		Alias:    (Alias)(o),
	})
}

func (o *V2conversationmessagetypingeventforusertopicconversationmessagingchannel) UnmarshalJSON(b []byte) error {
	var V2conversationmessagetypingeventforusertopicconversationmessagingchannelMap map[string]interface{}
	err := json.Unmarshal(b, &V2conversationmessagetypingeventforusertopicconversationmessagingchannelMap)
	if err != nil {
		return err
	}
	
	if Id, ok := V2conversationmessagetypingeventforusertopicconversationmessagingchannelMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Platform, ok := V2conversationmessagetypingeventforusertopicconversationmessagingchannelMap["platform"].(string); ok {
		o.Platform = &Platform
	}
    
	if MessageId, ok := V2conversationmessagetypingeventforusertopicconversationmessagingchannelMap["messageId"].(string); ok {
		o.MessageId = &MessageId
	}
    
	if To, ok := V2conversationmessagetypingeventforusertopicconversationmessagingchannelMap["to"].(map[string]interface{}); ok {
		ToString, _ := json.Marshal(To)
		json.Unmarshal(ToString, &o.To)
	}
	
	if From, ok := V2conversationmessagetypingeventforusertopicconversationmessagingchannelMap["from"].(map[string]interface{}); ok {
		FromString, _ := json.Marshal(From)
		json.Unmarshal(FromString, &o.From)
	}
	
	if timeString, ok := V2conversationmessagetypingeventforusertopicconversationmessagingchannelMap["time"].(string); ok {
		Time, _ := time.Parse("2006-01-02T15:04:05.999999Z", timeString)
		o.Time = &Time
	}
	
	if dateModifiedString, ok := V2conversationmessagetypingeventforusertopicconversationmessagingchannelMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if dateDeletedString, ok := V2conversationmessagetypingeventforusertopicconversationmessagingchannelMap["dateDeleted"].(string); ok {
		DateDeleted, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateDeletedString)
		o.DateDeleted = &DateDeleted
	}
	
	if Metadata, ok := V2conversationmessagetypingeventforusertopicconversationmessagingchannelMap["metadata"].(map[string]interface{}); ok {
		MetadataString, _ := json.Marshal(Metadata)
		json.Unmarshal(MetadataString, &o.Metadata)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *V2conversationmessagetypingeventforusertopicconversationmessagingchannel) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

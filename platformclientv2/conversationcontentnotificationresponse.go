package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationcontentnotificationresponse - Inbound response to a notification, such as an Apple Invitations acceptance.
type Conversationcontentnotificationresponse struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// OriginatingMessageId - Reference to the ID of the original outbound notification message this response is for (e.g. the Apple requestIdentifier).
	OriginatingMessageId *string `json:"originatingMessageId,omitempty"`

	// ReferenceId - The business context reference associated with the notification (e.g. order ID, case ID). May be empty if the provider does not return it.
	ReferenceId *string `json:"referenceId,omitempty"`

	// NotificationStatus - The status of the notification response.
	NotificationStatus *string `json:"notificationStatus,omitempty"`

	// NotificationText - The localized display text of the user's response (e.g. \"Yes\").
	NotificationText *string `json:"notificationText,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Conversationcontentnotificationresponse) SetField(field string, fieldValue interface{}) {
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

func (o Conversationcontentnotificationresponse) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{  }
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
	type Alias Conversationcontentnotificationresponse
	
	return json.Marshal(&struct { 
		OriginatingMessageId *string `json:"originatingMessageId,omitempty"`
		
		ReferenceId *string `json:"referenceId,omitempty"`
		
		NotificationStatus *string `json:"notificationStatus,omitempty"`
		
		NotificationText *string `json:"notificationText,omitempty"`
		Alias
	}{ 
		OriginatingMessageId: o.OriginatingMessageId,
		
		ReferenceId: o.ReferenceId,
		
		NotificationStatus: o.NotificationStatus,
		
		NotificationText: o.NotificationText,
		Alias:    (Alias)(o),
	})
}

func (o *Conversationcontentnotificationresponse) UnmarshalJSON(b []byte) error {
	var ConversationcontentnotificationresponseMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationcontentnotificationresponseMap)
	if err != nil {
		return err
	}
	
	if OriginatingMessageId, ok := ConversationcontentnotificationresponseMap["originatingMessageId"].(string); ok {
		o.OriginatingMessageId = &OriginatingMessageId
	}
    
	if ReferenceId, ok := ConversationcontentnotificationresponseMap["referenceId"].(string); ok {
		o.ReferenceId = &ReferenceId
	}
    
	if NotificationStatus, ok := ConversationcontentnotificationresponseMap["notificationStatus"].(string); ok {
		o.NotificationStatus = &NotificationStatus
	}
    
	if NotificationText, ok := ConversationcontentnotificationresponseMap["notificationText"].(string); ok {
		o.NotificationText = &NotificationText
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationcontentnotificationresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

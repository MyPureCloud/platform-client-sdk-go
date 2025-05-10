package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationcontentlistpicker - List Picker object for presenting multiple sections of selectable items.
type Conversationcontentlistpicker struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Sections - An array of sections in the List Picker.
	Sections *[]Conversationcontentlistpickersection `json:"sections,omitempty"`

	// ReplyMessage - The reply message after the user has selected the options from the List Picker.
	ReplyMessage *Conversationcontentreceivedreplymessage `json:"replyMessage,omitempty"`

	// ReceivedMessage - The message prompt to select options in the List Picker sections.
	ReceivedMessage *Conversationcontentreceivedreplymessage `json:"receivedMessage,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Conversationcontentlistpicker) SetField(field string, fieldValue interface{}) {
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

func (o Conversationcontentlistpicker) MarshalJSON() ([]byte, error) {
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
	type Alias Conversationcontentlistpicker
	
	return json.Marshal(&struct { 
		Sections *[]Conversationcontentlistpickersection `json:"sections,omitempty"`
		
		ReplyMessage *Conversationcontentreceivedreplymessage `json:"replyMessage,omitempty"`
		
		ReceivedMessage *Conversationcontentreceivedreplymessage `json:"receivedMessage,omitempty"`
		Alias
	}{ 
		Sections: o.Sections,
		
		ReplyMessage: o.ReplyMessage,
		
		ReceivedMessage: o.ReceivedMessage,
		Alias:    (Alias)(o),
	})
}

func (o *Conversationcontentlistpicker) UnmarshalJSON(b []byte) error {
	var ConversationcontentlistpickerMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationcontentlistpickerMap)
	if err != nil {
		return err
	}
	
	if Sections, ok := ConversationcontentlistpickerMap["sections"].([]interface{}); ok {
		SectionsString, _ := json.Marshal(Sections)
		json.Unmarshal(SectionsString, &o.Sections)
	}
	
	if ReplyMessage, ok := ConversationcontentlistpickerMap["replyMessage"].(map[string]interface{}); ok {
		ReplyMessageString, _ := json.Marshal(ReplyMessage)
		json.Unmarshal(ReplyMessageString, &o.ReplyMessage)
	}
	
	if ReceivedMessage, ok := ConversationcontentlistpickerMap["receivedMessage"].(map[string]interface{}); ok {
		ReceivedMessageString, _ := json.Marshal(ReceivedMessage)
		json.Unmarshal(ReceivedMessageString, &o.ReceivedMessage)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationcontentlistpicker) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

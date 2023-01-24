package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Webchattyping
type Webchattyping struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The event identifier of this typing indicator event (useful to guard against event re-delivery
	Id *string `json:"id,omitempty"`

	// Conversation - The identifier of the conversation
	Conversation *Webchatconversation `json:"conversation,omitempty"`

	// Sender - The member who sent the message
	Sender *Webchatmemberinfo `json:"sender,omitempty"`

	// Timestamp - The timestamp of the message, in ISO-8601 format
	Timestamp *time.Time `json:"timestamp,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Webchattyping) SetField(field string, fieldValue interface{}) {
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

func (o Webchattyping) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "Timestamp", }
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
	type Alias Webchattyping
	
	Timestamp := new(string)
	if o.Timestamp != nil {
		
		*Timestamp = timeutil.Strftime(o.Timestamp, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		Timestamp = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Conversation *Webchatconversation `json:"conversation,omitempty"`
		
		Sender *Webchatmemberinfo `json:"sender,omitempty"`
		
		Timestamp *string `json:"timestamp,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Conversation: o.Conversation,
		
		Sender: o.Sender,
		
		Timestamp: Timestamp,
		Alias:    (Alias)(o),
	})
}

func (o *Webchattyping) UnmarshalJSON(b []byte) error {
	var WebchattypingMap map[string]interface{}
	err := json.Unmarshal(b, &WebchattypingMap)
	if err != nil {
		return err
	}
	
	if Id, ok := WebchattypingMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Conversation, ok := WebchattypingMap["conversation"].(map[string]interface{}); ok {
		ConversationString, _ := json.Marshal(Conversation)
		json.Unmarshal(ConversationString, &o.Conversation)
	}
	
	if Sender, ok := WebchattypingMap["sender"].(map[string]interface{}); ok {
		SenderString, _ := json.Marshal(Sender)
		json.Unmarshal(SenderString, &o.Sender)
	}
	
	if timestampString, ok := WebchattypingMap["timestamp"].(string); ok {
		Timestamp, _ := time.Parse("2006-01-02T15:04:05.999999Z", timestampString)
		o.Timestamp = &Timestamp
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Webchattyping) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

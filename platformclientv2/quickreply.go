package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Quickreply
type Quickreply struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Text - Text to show inside the quick reply. This is also used as the response text after clicking on the quick reply.
	Text *string `json:"text,omitempty"`

	// Payload - Content of the textback payload after clicking a quick reply
	Payload *string `json:"payload,omitempty"`

	// Url - The location of the image file associated with quick reply
	Url *string `json:"url,omitempty"`

	// Action - Specifies the type of action that is triggered upon clicking the quick reply. Currently, the only supported action is \"Message\" which sends a message using the quick reply text.
	Action *string `json:"action,omitempty"`

	// IsSelected - Indicates if the quick reply option is selected by end customer
	IsSelected *bool `json:"isSelected,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Quickreply) SetField(field string, fieldValue interface{}) {
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

func (o Quickreply) MarshalJSON() ([]byte, error) {
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
	type Alias Quickreply
	
	return json.Marshal(&struct { 
		Text *string `json:"text,omitempty"`
		
		Payload *string `json:"payload,omitempty"`
		
		Url *string `json:"url,omitempty"`
		
		Action *string `json:"action,omitempty"`
		
		IsSelected *bool `json:"isSelected,omitempty"`
		Alias
	}{ 
		Text: o.Text,
		
		Payload: o.Payload,
		
		Url: o.Url,
		
		Action: o.Action,
		
		IsSelected: o.IsSelected,
		Alias:    (Alias)(o),
	})
}

func (o *Quickreply) UnmarshalJSON(b []byte) error {
	var QuickreplyMap map[string]interface{}
	err := json.Unmarshal(b, &QuickreplyMap)
	if err != nil {
		return err
	}
	
	if Text, ok := QuickreplyMap["text"].(string); ok {
		o.Text = &Text
	}
    
	if Payload, ok := QuickreplyMap["payload"].(string); ok {
		o.Payload = &Payload
	}
    
	if Url, ok := QuickreplyMap["url"].(string); ok {
		o.Url = &Url
	}
    
	if Action, ok := QuickreplyMap["action"].(string); ok {
		o.Action = &Action
	}
    
	if IsSelected, ok := QuickreplyMap["isSelected"].(bool); ok {
		o.IsSelected = &IsSelected
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Quickreply) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

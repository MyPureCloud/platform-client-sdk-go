package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Cardaction - A card action that a user can take.
type Cardaction struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// VarType - Describes the type of action.
	VarType *string `json:"type,omitempty"`

	// Text - The response text from the button click.
	Text *string `json:"text,omitempty"`

	// Payload - Content of the textback payload after clicking a button.
	Payload *string `json:"payload,omitempty"`

	// Url - The location of the image file associated with action.
	Url *string `json:"url,omitempty"`

	// IsSelected - Indicates if the card option is selected by end customer.
	IsSelected *bool `json:"isSelected,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Cardaction) SetField(field string, fieldValue interface{}) {
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

func (o Cardaction) MarshalJSON() ([]byte, error) {
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
	type Alias Cardaction
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		Text *string `json:"text,omitempty"`
		
		Payload *string `json:"payload,omitempty"`
		
		Url *string `json:"url,omitempty"`
		
		IsSelected *bool `json:"isSelected,omitempty"`
		Alias
	}{ 
		VarType: o.VarType,
		
		Text: o.Text,
		
		Payload: o.Payload,
		
		Url: o.Url,
		
		IsSelected: o.IsSelected,
		Alias:    (Alias)(o),
	})
}

func (o *Cardaction) UnmarshalJSON(b []byte) error {
	var CardactionMap map[string]interface{}
	err := json.Unmarshal(b, &CardactionMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := CardactionMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Text, ok := CardactionMap["text"].(string); ok {
		o.Text = &Text
	}
    
	if Payload, ok := CardactionMap["payload"].(string); ok {
		o.Payload = &Payload
	}
    
	if Url, ok := CardactionMap["url"].(string); ok {
		o.Url = &Url
	}
    
	if IsSelected, ok := CardactionMap["isSelected"].(bool); ok {
		o.IsSelected = &IsSelected
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Cardaction) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

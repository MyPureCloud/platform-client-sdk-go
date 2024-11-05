package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Recordingtemplatebutton
type Recordingtemplatebutton struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// VarType - Specifies the type of the button.
	VarType *string `json:"type,omitempty"`

	// Text - Button text message.
	Text *string `json:"text,omitempty"`

	// Index - Index of the button in the list.
	Index *int `json:"index,omitempty"`

	// PhoneNumber - Button phone number.
	PhoneNumber *string `json:"phoneNumber,omitempty"`

	// Url - Button URL link.
	Url *string `json:"url,omitempty"`

	// IsSelected - Indicates if the button is selected by end customer.
	IsSelected *bool `json:"isSelected,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Recordingtemplatebutton) SetField(field string, fieldValue interface{}) {
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

func (o Recordingtemplatebutton) MarshalJSON() ([]byte, error) {
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
	type Alias Recordingtemplatebutton
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		Text *string `json:"text,omitempty"`
		
		Index *int `json:"index,omitempty"`
		
		PhoneNumber *string `json:"phoneNumber,omitempty"`
		
		Url *string `json:"url,omitempty"`
		
		IsSelected *bool `json:"isSelected,omitempty"`
		Alias
	}{ 
		VarType: o.VarType,
		
		Text: o.Text,
		
		Index: o.Index,
		
		PhoneNumber: o.PhoneNumber,
		
		Url: o.Url,
		
		IsSelected: o.IsSelected,
		Alias:    (Alias)(o),
	})
}

func (o *Recordingtemplatebutton) UnmarshalJSON(b []byte) error {
	var RecordingtemplatebuttonMap map[string]interface{}
	err := json.Unmarshal(b, &RecordingtemplatebuttonMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := RecordingtemplatebuttonMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Text, ok := RecordingtemplatebuttonMap["text"].(string); ok {
		o.Text = &Text
	}
    
	if Index, ok := RecordingtemplatebuttonMap["index"].(float64); ok {
		IndexInt := int(Index)
		o.Index = &IndexInt
	}
	
	if PhoneNumber, ok := RecordingtemplatebuttonMap["phoneNumber"].(string); ok {
		o.PhoneNumber = &PhoneNumber
	}
    
	if Url, ok := RecordingtemplatebuttonMap["url"].(string); ok {
		o.Url = &Url
	}
    
	if IsSelected, ok := RecordingtemplatebuttonMap["isSelected"].(bool); ok {
		o.IsSelected = &IsSelected
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Recordingtemplatebutton) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

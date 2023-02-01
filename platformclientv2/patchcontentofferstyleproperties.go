package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Patchcontentofferstyleproperties
type Patchcontentofferstyleproperties struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Padding - Padding of the offer. (eg. 10px)
	Padding *string `json:"padding,omitempty"`

	// Color - Text color of the offer. (eg. #FF0000)
	Color *string `json:"color,omitempty"`

	// BackgroundColor - Background color of the offer. (eg. #000000)
	BackgroundColor *string `json:"backgroundColor,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Patchcontentofferstyleproperties) SetField(field string, fieldValue interface{}) {
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

func (o Patchcontentofferstyleproperties) MarshalJSON() ([]byte, error) {
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
	type Alias Patchcontentofferstyleproperties
	
	return json.Marshal(&struct { 
		Padding *string `json:"padding,omitempty"`
		
		Color *string `json:"color,omitempty"`
		
		BackgroundColor *string `json:"backgroundColor,omitempty"`
		Alias
	}{ 
		Padding: o.Padding,
		
		Color: o.Color,
		
		BackgroundColor: o.BackgroundColor,
		Alias:    (Alias)(o),
	})
}

func (o *Patchcontentofferstyleproperties) UnmarshalJSON(b []byte) error {
	var PatchcontentofferstylepropertiesMap map[string]interface{}
	err := json.Unmarshal(b, &PatchcontentofferstylepropertiesMap)
	if err != nil {
		return err
	}
	
	if Padding, ok := PatchcontentofferstylepropertiesMap["padding"].(string); ok {
		o.Padding = &Padding
	}
    
	if Color, ok := PatchcontentofferstylepropertiesMap["color"].(string); ok {
		o.Color = &Color
	}
    
	if BackgroundColor, ok := PatchcontentofferstylepropertiesMap["backgroundColor"].(string); ok {
		o.BackgroundColor = &BackgroundColor
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Patchcontentofferstyleproperties) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

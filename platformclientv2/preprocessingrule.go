package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Preprocessingrule
type Preprocessingrule struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Find - The regular expression to which file lines are to be matched.
	Find *string `json:"find,omitempty"`

	// ReplaceWith - The string to be substituted for each match.
	ReplaceWith *string `json:"replaceWith,omitempty"`

	// Global - Replaces all matching substrings in every line.
	Global *bool `json:"global,omitempty"`

	// IgnoreCase - Enables case-insensitive matching
	IgnoreCase *bool `json:"ignoreCase,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Preprocessingrule) SetField(field string, fieldValue interface{}) {
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

func (o Preprocessingrule) MarshalJSON() ([]byte, error) {
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
	type Alias Preprocessingrule
	
	return json.Marshal(&struct { 
		Find *string `json:"find,omitempty"`
		
		ReplaceWith *string `json:"replaceWith,omitempty"`
		
		Global *bool `json:"global,omitempty"`
		
		IgnoreCase *bool `json:"ignoreCase,omitempty"`
		Alias
	}{ 
		Find: o.Find,
		
		ReplaceWith: o.ReplaceWith,
		
		Global: o.Global,
		
		IgnoreCase: o.IgnoreCase,
		Alias:    (Alias)(o),
	})
}

func (o *Preprocessingrule) UnmarshalJSON(b []byte) error {
	var PreprocessingruleMap map[string]interface{}
	err := json.Unmarshal(b, &PreprocessingruleMap)
	if err != nil {
		return err
	}
	
	if Find, ok := PreprocessingruleMap["find"].(string); ok {
		o.Find = &Find
	}
    
	if ReplaceWith, ok := PreprocessingruleMap["replaceWith"].(string); ok {
		o.ReplaceWith = &ReplaceWith
	}
    
	if Global, ok := PreprocessingruleMap["global"].(bool); ok {
		o.Global = &Global
	}
    
	if IgnoreCase, ok := PreprocessingruleMap["ignoreCase"].(bool); ok {
		o.IgnoreCase = &IgnoreCase
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Preprocessingrule) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

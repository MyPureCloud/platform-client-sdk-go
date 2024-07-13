package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Webeventbrowser
type Webeventbrowser struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Family - Browser family (e.g. Chrome, Safari, Firefox).
	Family *string `json:"family,omitempty"`

	// Version - Browser version (e.g. 68.0.3440.84).
	Version *string `json:"version,omitempty"`

	// Lang - Language the browser is set to. Must conform to BCP 47.
	Lang *string `json:"lang,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Webeventbrowser) SetField(field string, fieldValue interface{}) {
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

func (o Webeventbrowser) MarshalJSON() ([]byte, error) {
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
	type Alias Webeventbrowser
	
	return json.Marshal(&struct { 
		Family *string `json:"family,omitempty"`
		
		Version *string `json:"version,omitempty"`
		
		Lang *string `json:"lang,omitempty"`
		Alias
	}{ 
		Family: o.Family,
		
		Version: o.Version,
		
		Lang: o.Lang,
		Alias:    (Alias)(o),
	})
}

func (o *Webeventbrowser) UnmarshalJSON(b []byte) error {
	var WebeventbrowserMap map[string]interface{}
	err := json.Unmarshal(b, &WebeventbrowserMap)
	if err != nil {
		return err
	}
	
	if Family, ok := WebeventbrowserMap["family"].(string); ok {
		o.Family = &Family
	}
    
	if Version, ok := WebeventbrowserMap["version"].(string); ok {
		o.Version = &Version
	}
    
	if Lang, ok := WebeventbrowserMap["lang"].(string); ok {
		o.Lang = &Lang
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Webeventbrowser) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

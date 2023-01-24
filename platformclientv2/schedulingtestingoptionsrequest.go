package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Schedulingtestingoptionsrequest
type Schedulingtestingoptionsrequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// FastScheduling - Whether to enable fast scheduling
	FastScheduling *bool `json:"fastScheduling,omitempty"`

	// DelayScheduling - Whether to force delayed scheduling
	DelayScheduling *bool `json:"delayScheduling,omitempty"`

	// FailScheduling - Whether to force scheduling to fail
	FailScheduling *bool `json:"failScheduling,omitempty"`

	// PopulateWarnings - Whether to populate warnings in the generated schedule
	PopulateWarnings *bool `json:"populateWarnings,omitempty"`

	// PopulateDeprecatedWarnings - Whether to populate deprecated warnings in the generated schedule
	PopulateDeprecatedWarnings *bool `json:"populateDeprecatedWarnings,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Schedulingtestingoptionsrequest) SetField(field string, fieldValue interface{}) {
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

func (o Schedulingtestingoptionsrequest) MarshalJSON() ([]byte, error) {
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
	type Alias Schedulingtestingoptionsrequest
	
	return json.Marshal(&struct { 
		FastScheduling *bool `json:"fastScheduling,omitempty"`
		
		DelayScheduling *bool `json:"delayScheduling,omitempty"`
		
		FailScheduling *bool `json:"failScheduling,omitempty"`
		
		PopulateWarnings *bool `json:"populateWarnings,omitempty"`
		
		PopulateDeprecatedWarnings *bool `json:"populateDeprecatedWarnings,omitempty"`
		Alias
	}{ 
		FastScheduling: o.FastScheduling,
		
		DelayScheduling: o.DelayScheduling,
		
		FailScheduling: o.FailScheduling,
		
		PopulateWarnings: o.PopulateWarnings,
		
		PopulateDeprecatedWarnings: o.PopulateDeprecatedWarnings,
		Alias:    (Alias)(o),
	})
}

func (o *Schedulingtestingoptionsrequest) UnmarshalJSON(b []byte) error {
	var SchedulingtestingoptionsrequestMap map[string]interface{}
	err := json.Unmarshal(b, &SchedulingtestingoptionsrequestMap)
	if err != nil {
		return err
	}
	
	if FastScheduling, ok := SchedulingtestingoptionsrequestMap["fastScheduling"].(bool); ok {
		o.FastScheduling = &FastScheduling
	}
    
	if DelayScheduling, ok := SchedulingtestingoptionsrequestMap["delayScheduling"].(bool); ok {
		o.DelayScheduling = &DelayScheduling
	}
    
	if FailScheduling, ok := SchedulingtestingoptionsrequestMap["failScheduling"].(bool); ok {
		o.FailScheduling = &FailScheduling
	}
    
	if PopulateWarnings, ok := SchedulingtestingoptionsrequestMap["populateWarnings"].(bool); ok {
		o.PopulateWarnings = &PopulateWarnings
	}
    
	if PopulateDeprecatedWarnings, ok := SchedulingtestingoptionsrequestMap["populateDeprecatedWarnings"].(bool); ok {
		o.PopulateDeprecatedWarnings = &PopulateDeprecatedWarnings
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Schedulingtestingoptionsrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Workplanconfigurationviolationmessage
type Workplanconfigurationviolationmessage struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// VarType - Type of configuration violation message for this work plan
	VarType *string `json:"type,omitempty"`

	// Arguments - Arguments of the message that provide information about the misconfigured value or the threshold that is exceeded by the misconfigured value
	Arguments *[]Workplanvalidationmessageargument `json:"arguments,omitempty"`

	// Severity - Severity of the message. A message with Error severity indicates the scheduler won't be able to produce schedules and thus the work plan is invalid.
	Severity *string `json:"severity,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Workplanconfigurationviolationmessage) SetField(field string, fieldValue interface{}) {
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

func (o Workplanconfigurationviolationmessage) MarshalJSON() ([]byte, error) {
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
	type Alias Workplanconfigurationviolationmessage
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		Arguments *[]Workplanvalidationmessageargument `json:"arguments,omitempty"`
		
		Severity *string `json:"severity,omitempty"`
		Alias
	}{ 
		VarType: o.VarType,
		
		Arguments: o.Arguments,
		
		Severity: o.Severity,
		Alias:    (Alias)(o),
	})
}

func (o *Workplanconfigurationviolationmessage) UnmarshalJSON(b []byte) error {
	var WorkplanconfigurationviolationmessageMap map[string]interface{}
	err := json.Unmarshal(b, &WorkplanconfigurationviolationmessageMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := WorkplanconfigurationviolationmessageMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Arguments, ok := WorkplanconfigurationviolationmessageMap["arguments"].([]interface{}); ok {
		ArgumentsString, _ := json.Marshal(Arguments)
		json.Unmarshal(ArgumentsString, &o.Arguments)
	}
	
	if Severity, ok := WorkplanconfigurationviolationmessageMap["severity"].(string); ok {
		o.Severity = &Severity
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Workplanconfigurationviolationmessage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

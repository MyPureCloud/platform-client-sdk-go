package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Evaluationsearchcriteriadto
type Evaluationsearchcriteriadto struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// VarType - The type of query Operation to Perform
	VarType *string `json:"type,omitempty"`

	// Field - Field name to search against
	Field *string `json:"field,omitempty"`

	// EndValue - The end value of the range. This field is used for range search types. Date values must be in the format yyyy-MM-dd'T'HH:mm:ss.SSS'Z'.
	EndValue *string `json:"endValue,omitempty"`

	// Values - A list of values for the search to match against. Only for Type: EXACT
	Values *[]string `json:"values,omitempty"`

	// StartValue - The start value of the range. This field is used for range search types. Date values must be in the format yyyy-MM-dd'T'HH:mm:ss.SSS'Z'.
	StartValue *string `json:"startValue,omitempty"`

	// Value - A value for the search to match against
	Value *string `json:"value,omitempty"`

	// Operator - How to apply this search criteria against other criteria
	Operator *string `json:"operator,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Evaluationsearchcriteriadto) SetField(field string, fieldValue interface{}) {
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

func (o Evaluationsearchcriteriadto) MarshalJSON() ([]byte, error) {
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
	type Alias Evaluationsearchcriteriadto
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		Field *string `json:"field,omitempty"`
		
		EndValue *string `json:"endValue,omitempty"`
		
		Values *[]string `json:"values,omitempty"`
		
		StartValue *string `json:"startValue,omitempty"`
		
		Value *string `json:"value,omitempty"`
		
		Operator *string `json:"operator,omitempty"`
		Alias
	}{ 
		VarType: o.VarType,
		
		Field: o.Field,
		
		EndValue: o.EndValue,
		
		Values: o.Values,
		
		StartValue: o.StartValue,
		
		Value: o.Value,
		
		Operator: o.Operator,
		Alias:    (Alias)(o),
	})
}

func (o *Evaluationsearchcriteriadto) UnmarshalJSON(b []byte) error {
	var EvaluationsearchcriteriadtoMap map[string]interface{}
	err := json.Unmarshal(b, &EvaluationsearchcriteriadtoMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := EvaluationsearchcriteriadtoMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Field, ok := EvaluationsearchcriteriadtoMap["field"].(string); ok {
		o.Field = &Field
	}
    
	if EndValue, ok := EvaluationsearchcriteriadtoMap["endValue"].(string); ok {
		o.EndValue = &EndValue
	}
    
	if Values, ok := EvaluationsearchcriteriadtoMap["values"].([]interface{}); ok {
		ValuesString, _ := json.Marshal(Values)
		json.Unmarshal(ValuesString, &o.Values)
	}
	
	if StartValue, ok := EvaluationsearchcriteriadtoMap["startValue"].(string); ok {
		o.StartValue = &StartValue
	}
    
	if Value, ok := EvaluationsearchcriteriadtoMap["value"].(string); ok {
		o.Value = &Value
	}
    
	if Operator, ok := EvaluationsearchcriteriadtoMap["operator"].(string); ok {
		o.Operator = &Operator
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Evaluationsearchcriteriadto) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

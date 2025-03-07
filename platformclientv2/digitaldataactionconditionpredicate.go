package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Digitaldataactionconditionpredicate
type Digitaldataactionconditionpredicate struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// OutputField - The name of an output field from the data action's output to use for this condition
	OutputField *string `json:"outputField,omitempty"`

	// OutputOperator - The operation with which to evaluate this condition
	OutputOperator *string `json:"outputOperator,omitempty"`

	// ComparisonValue - The value to compare against for this condition
	ComparisonValue *string `json:"comparisonValue,omitempty"`

	// Inverted - If true, inverts the result of evaluating this Predicate. Default is false.
	Inverted *bool `json:"inverted,omitempty"`

	// OutputFieldMissingResolution - The result of this predicate if the requested output field is missing from the data action's result
	OutputFieldMissingResolution *bool `json:"outputFieldMissingResolution,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Digitaldataactionconditionpredicate) SetField(field string, fieldValue interface{}) {
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

func (o Digitaldataactionconditionpredicate) MarshalJSON() ([]byte, error) {
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
	type Alias Digitaldataactionconditionpredicate
	
	return json.Marshal(&struct { 
		OutputField *string `json:"outputField,omitempty"`
		
		OutputOperator *string `json:"outputOperator,omitempty"`
		
		ComparisonValue *string `json:"comparisonValue,omitempty"`
		
		Inverted *bool `json:"inverted,omitempty"`
		
		OutputFieldMissingResolution *bool `json:"outputFieldMissingResolution,omitempty"`
		Alias
	}{ 
		OutputField: o.OutputField,
		
		OutputOperator: o.OutputOperator,
		
		ComparisonValue: o.ComparisonValue,
		
		Inverted: o.Inverted,
		
		OutputFieldMissingResolution: o.OutputFieldMissingResolution,
		Alias:    (Alias)(o),
	})
}

func (o *Digitaldataactionconditionpredicate) UnmarshalJSON(b []byte) error {
	var DigitaldataactionconditionpredicateMap map[string]interface{}
	err := json.Unmarshal(b, &DigitaldataactionconditionpredicateMap)
	if err != nil {
		return err
	}
	
	if OutputField, ok := DigitaldataactionconditionpredicateMap["outputField"].(string); ok {
		o.OutputField = &OutputField
	}
    
	if OutputOperator, ok := DigitaldataactionconditionpredicateMap["outputOperator"].(string); ok {
		o.OutputOperator = &OutputOperator
	}
    
	if ComparisonValue, ok := DigitaldataactionconditionpredicateMap["comparisonValue"].(string); ok {
		o.ComparisonValue = &ComparisonValue
	}
    
	if Inverted, ok := DigitaldataactionconditionpredicateMap["inverted"].(bool); ok {
		o.Inverted = &Inverted
	}
    
	if OutputFieldMissingResolution, ok := DigitaldataactionconditionpredicateMap["outputFieldMissingResolution"].(bool); ok {
		o.OutputFieldMissingResolution = &OutputFieldMissingResolution
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Digitaldataactionconditionpredicate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

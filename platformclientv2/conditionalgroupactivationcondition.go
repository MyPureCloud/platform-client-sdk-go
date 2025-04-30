package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Conditionalgroupactivationcondition
type Conditionalgroupactivationcondition struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// SimpleMetric - Instructs this condition to evaluate a simple queue-level metric
	SimpleMetric *Conditionalgroupactivationsimplemetric `json:"simpleMetric,omitempty"`

	// Operator - The operator used to compare the actual value against the threshold value
	Operator *string `json:"operator,omitempty"`

	// Value - The threshold value, beyond which a rule evaluates as true
	Value *float64 `json:"value,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Conditionalgroupactivationcondition) SetField(field string, fieldValue interface{}) {
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

func (o Conditionalgroupactivationcondition) MarshalJSON() ([]byte, error) {
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
	type Alias Conditionalgroupactivationcondition
	
	return json.Marshal(&struct { 
		SimpleMetric *Conditionalgroupactivationsimplemetric `json:"simpleMetric,omitempty"`
		
		Operator *string `json:"operator,omitempty"`
		
		Value *float64 `json:"value,omitempty"`
		Alias
	}{ 
		SimpleMetric: o.SimpleMetric,
		
		Operator: o.Operator,
		
		Value: o.Value,
		Alias:    (Alias)(o),
	})
}

func (o *Conditionalgroupactivationcondition) UnmarshalJSON(b []byte) error {
	var ConditionalgroupactivationconditionMap map[string]interface{}
	err := json.Unmarshal(b, &ConditionalgroupactivationconditionMap)
	if err != nil {
		return err
	}
	
	if SimpleMetric, ok := ConditionalgroupactivationconditionMap["simpleMetric"].(map[string]interface{}); ok {
		SimpleMetricString, _ := json.Marshal(SimpleMetric)
		json.Unmarshal(SimpleMetricString, &o.SimpleMetric)
	}
	
	if Operator, ok := ConditionalgroupactivationconditionMap["operator"].(string); ok {
		o.Operator = &Operator
	}
    
	if Value, ok := ConditionalgroupactivationconditionMap["value"].(float64); ok {
		o.Value = &Value
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Conditionalgroupactivationcondition) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

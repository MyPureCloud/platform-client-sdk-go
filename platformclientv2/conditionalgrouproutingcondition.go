package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Conditionalgrouproutingcondition
type Conditionalgrouproutingcondition struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Queue - The queue being evaluated for this rule.  If null, the current queue will be used.
	Queue *Domainentityref `json:"queue,omitempty"`

	// Metric - The queue metric being evaluated
	Metric *string `json:"metric,omitempty"`

	// Operator - The operator that compares the actual value against the condition value
	Operator *string `json:"operator,omitempty"`

	// Value - The limit value, beyond which a rule evaluates as true
	Value *float64 `json:"value,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Conditionalgrouproutingcondition) SetField(field string, fieldValue interface{}) {
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

func (o Conditionalgrouproutingcondition) MarshalJSON() ([]byte, error) {
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
	type Alias Conditionalgrouproutingcondition
	
	return json.Marshal(&struct { 
		Queue *Domainentityref `json:"queue,omitempty"`
		
		Metric *string `json:"metric,omitempty"`
		
		Operator *string `json:"operator,omitempty"`
		
		Value *float64 `json:"value,omitempty"`
		Alias
	}{ 
		Queue: o.Queue,
		
		Metric: o.Metric,
		
		Operator: o.Operator,
		
		Value: o.Value,
		Alias:    (Alias)(o),
	})
}

func (o *Conditionalgrouproutingcondition) UnmarshalJSON(b []byte) error {
	var ConditionalgrouproutingconditionMap map[string]interface{}
	err := json.Unmarshal(b, &ConditionalgrouproutingconditionMap)
	if err != nil {
		return err
	}
	
	if Queue, ok := ConditionalgrouproutingconditionMap["queue"].(map[string]interface{}); ok {
		QueueString, _ := json.Marshal(Queue)
		json.Unmarshal(QueueString, &o.Queue)
	}
	
	if Metric, ok := ConditionalgrouproutingconditionMap["metric"].(string); ok {
		o.Metric = &Metric
	}
    
	if Operator, ok := ConditionalgrouproutingconditionMap["operator"].(string); ok {
		o.Operator = &Operator
	}
    
	if Value, ok := ConditionalgrouproutingconditionMap["value"].(float64); ok {
		o.Value = &Value
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Conditionalgrouproutingcondition) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Conditionalgrouproutingrule
type Conditionalgrouproutingrule struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Queue - The queue being evaluated for this rule.  If null, the current queue will be used.
	Queue *Domainentityref `json:"queue,omitempty"`

	// Metric - The queue metric being evaluated
	Metric *string `json:"metric,omitempty"`

	// Operator - The operator that compares the actual value against the condition value
	Operator *string `json:"operator,omitempty"`

	// ConditionValue - The limit value, beyond which a rule evaluates as true
	ConditionValue *float64 `json:"conditionValue,omitempty"`

	// Groups - The group(s) to activate if the rule evaluates as true
	Groups *[]Membergroup `json:"groups,omitempty"`

	// WaitSeconds - The number of seconds to wait in this rule, if it evaluates as true, before evaluating the next rule.  For the final rule, this is ignored, so need not be specified.
	WaitSeconds *int `json:"waitSeconds,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Conditionalgrouproutingrule) SetField(field string, fieldValue interface{}) {
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

func (o Conditionalgrouproutingrule) MarshalJSON() ([]byte, error) {
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
	type Alias Conditionalgrouproutingrule
	
	return json.Marshal(&struct { 
		Queue *Domainentityref `json:"queue,omitempty"`
		
		Metric *string `json:"metric,omitempty"`
		
		Operator *string `json:"operator,omitempty"`
		
		ConditionValue *float64 `json:"conditionValue,omitempty"`
		
		Groups *[]Membergroup `json:"groups,omitempty"`
		
		WaitSeconds *int `json:"waitSeconds,omitempty"`
		Alias
	}{ 
		Queue: o.Queue,
		
		Metric: o.Metric,
		
		Operator: o.Operator,
		
		ConditionValue: o.ConditionValue,
		
		Groups: o.Groups,
		
		WaitSeconds: o.WaitSeconds,
		Alias:    (Alias)(o),
	})
}

func (o *Conditionalgrouproutingrule) UnmarshalJSON(b []byte) error {
	var ConditionalgrouproutingruleMap map[string]interface{}
	err := json.Unmarshal(b, &ConditionalgrouproutingruleMap)
	if err != nil {
		return err
	}
	
	if Queue, ok := ConditionalgrouproutingruleMap["queue"].(map[string]interface{}); ok {
		QueueString, _ := json.Marshal(Queue)
		json.Unmarshal(QueueString, &o.Queue)
	}
	
	if Metric, ok := ConditionalgrouproutingruleMap["metric"].(string); ok {
		o.Metric = &Metric
	}
    
	if Operator, ok := ConditionalgrouproutingruleMap["operator"].(string); ok {
		o.Operator = &Operator
	}
    
	if ConditionValue, ok := ConditionalgrouproutingruleMap["conditionValue"].(float64); ok {
		o.ConditionValue = &ConditionValue
	}
    
	if Groups, ok := ConditionalgrouproutingruleMap["groups"].([]interface{}); ok {
		GroupsString, _ := json.Marshal(Groups)
		json.Unmarshal(GroupsString, &o.Groups)
	}
	
	if WaitSeconds, ok := ConditionalgrouproutingruleMap["waitSeconds"].(float64); ok {
		WaitSecondsInt := int(WaitSeconds)
		o.WaitSeconds = &WaitSecondsInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conditionalgrouproutingrule) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

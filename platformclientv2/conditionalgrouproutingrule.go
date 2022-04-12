package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conditionalgrouproutingrule
type Conditionalgrouproutingrule struct { 
	// Queue - The queue being evaluated for this rule.  For rule 1, this is always the current queue.
	Queue *Domainentityref `json:"queue,omitempty"`


	// Metric - The queue metric being evaluated
	Metric *string `json:"metric,omitempty"`


	// Operator - The operator that compares the actual value against the condition value
	Operator *string `json:"operator,omitempty"`


	// ConditionValue - The limit value, beyond which a rule evaluates as true
	ConditionValue *float64 `json:"conditionValue,omitempty"`


	// Groups - The group(s) to activate if the rule evaluates as true
	Groups *[]Membergroup `json:"groups,omitempty"`


	// WaitSeconds - The number of seconds to wait in this rule, if it evaluates as true, before evaluating the next rule
	WaitSeconds *int `json:"waitSeconds,omitempty"`

}

func (o *Conditionalgrouproutingrule) MarshalJSON() ([]byte, error) {
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
		*Alias
	}{ 
		Queue: o.Queue,
		
		Metric: o.Metric,
		
		Operator: o.Operator,
		
		ConditionValue: o.ConditionValue,
		
		Groups: o.Groups,
		
		WaitSeconds: o.WaitSeconds,
		Alias:    (*Alias)(o),
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

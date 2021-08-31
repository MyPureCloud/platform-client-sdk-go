package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Callmediapolicy
type Callmediapolicy struct { 
	// Actions - Actions applied when specified conditions are met
	Actions *Policyactions `json:"actions,omitempty"`


	// Conditions - Conditions for when actions should be applied
	Conditions *Callmediapolicyconditions `json:"conditions,omitempty"`

}

func (o *Callmediapolicy) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Callmediapolicy
	
	return json.Marshal(&struct { 
		Actions *Policyactions `json:"actions,omitempty"`
		
		Conditions *Callmediapolicyconditions `json:"conditions,omitempty"`
		*Alias
	}{ 
		Actions: o.Actions,
		
		Conditions: o.Conditions,
		Alias:    (*Alias)(o),
	})
}

func (o *Callmediapolicy) UnmarshalJSON(b []byte) error {
	var CallmediapolicyMap map[string]interface{}
	err := json.Unmarshal(b, &CallmediapolicyMap)
	if err != nil {
		return err
	}
	
	if Actions, ok := CallmediapolicyMap["actions"].(map[string]interface{}); ok {
		ActionsString, _ := json.Marshal(Actions)
		json.Unmarshal(ActionsString, &o.Actions)
	}
	
	if Conditions, ok := CallmediapolicyMap["conditions"].(map[string]interface{}); ok {
		ConditionsString, _ := json.Marshal(Conditions)
		json.Unmarshal(ConditionsString, &o.Conditions)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Callmediapolicy) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

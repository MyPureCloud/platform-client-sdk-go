package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Crossplatformemailmediapolicy
type Crossplatformemailmediapolicy struct { 
	// Actions - Actions applied when specified conditions are met
	Actions *Crossplatformpolicyactions `json:"actions,omitempty"`


	// Conditions - Conditions for when actions should be applied
	Conditions *Emailmediapolicyconditions `json:"conditions,omitempty"`

}

func (o *Crossplatformemailmediapolicy) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Crossplatformemailmediapolicy
	
	return json.Marshal(&struct { 
		Actions *Crossplatformpolicyactions `json:"actions,omitempty"`
		
		Conditions *Emailmediapolicyconditions `json:"conditions,omitempty"`
		*Alias
	}{ 
		Actions: o.Actions,
		
		Conditions: o.Conditions,
		Alias:    (*Alias)(o),
	})
}

func (o *Crossplatformemailmediapolicy) UnmarshalJSON(b []byte) error {
	var CrossplatformemailmediapolicyMap map[string]interface{}
	err := json.Unmarshal(b, &CrossplatformemailmediapolicyMap)
	if err != nil {
		return err
	}
	
	if Actions, ok := CrossplatformemailmediapolicyMap["actions"].(map[string]interface{}); ok {
		ActionsString, _ := json.Marshal(Actions)
		json.Unmarshal(ActionsString, &o.Actions)
	}
	
	if Conditions, ok := CrossplatformemailmediapolicyMap["conditions"].(map[string]interface{}); ok {
		ConditionsString, _ := json.Marshal(Conditions)
		json.Unmarshal(ConditionsString, &o.Conditions)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Crossplatformemailmediapolicy) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

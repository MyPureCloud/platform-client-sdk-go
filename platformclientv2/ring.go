package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Ring
type Ring struct { 
	// ExpansionCriteria - The conditions that will trigger conversations to move to the next bullseye ring.
	ExpansionCriteria *[]Expansioncriterium `json:"expansionCriteria,omitempty"`


	// Actions - The actions that will be performed just before moving conversations to the next bullseye ring.
	Actions *Actions `json:"actions,omitempty"`

}

func (o *Ring) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Ring
	
	return json.Marshal(&struct { 
		ExpansionCriteria *[]Expansioncriterium `json:"expansionCriteria,omitempty"`
		
		Actions *Actions `json:"actions,omitempty"`
		*Alias
	}{ 
		ExpansionCriteria: o.ExpansionCriteria,
		
		Actions: o.Actions,
		Alias:    (*Alias)(o),
	})
}

func (o *Ring) UnmarshalJSON(b []byte) error {
	var RingMap map[string]interface{}
	err := json.Unmarshal(b, &RingMap)
	if err != nil {
		return err
	}
	
	if ExpansionCriteria, ok := RingMap["expansionCriteria"].([]interface{}); ok {
		ExpansionCriteriaString, _ := json.Marshal(ExpansionCriteria)
		json.Unmarshal(ExpansionCriteriaString, &o.ExpansionCriteria)
	}
	
	if Actions, ok := RingMap["actions"].(map[string]interface{}); ok {
		ActionsString, _ := json.Marshal(Actions)
		json.Unmarshal(ActionsString, &o.Actions)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Ring) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

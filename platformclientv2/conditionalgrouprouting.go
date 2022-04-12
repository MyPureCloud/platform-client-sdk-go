package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conditionalgrouprouting
type Conditionalgrouprouting struct { 
	// Rules - The set of rules that defines Conditional Group Routing for this queue
	Rules *[]Conditionalgrouproutingrule `json:"rules,omitempty"`

}

func (o *Conditionalgrouprouting) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conditionalgrouprouting
	
	return json.Marshal(&struct { 
		Rules *[]Conditionalgrouproutingrule `json:"rules,omitempty"`
		*Alias
	}{ 
		Rules: o.Rules,
		Alias:    (*Alias)(o),
	})
}

func (o *Conditionalgrouprouting) UnmarshalJSON(b []byte) error {
	var ConditionalgrouproutingMap map[string]interface{}
	err := json.Unmarshal(b, &ConditionalgrouproutingMap)
	if err != nil {
		return err
	}
	
	if Rules, ok := ConditionalgrouproutingMap["rules"].([]interface{}); ok {
		RulesString, _ := json.Marshal(Rules)
		json.Unmarshal(RulesString, &o.Rules)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conditionalgrouprouting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

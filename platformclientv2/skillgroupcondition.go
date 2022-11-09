package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Skillgroupcondition
type Skillgroupcondition struct { 
	// RoutingSkillConditions - Routing skill conditions that will be used for building the query
	RoutingSkillConditions *[]Skillgrouproutingcondition `json:"routingSkillConditions,omitempty"`


	// LanguageSkillConditions - Routing skill conditions that will be used for building the query
	LanguageSkillConditions *[]Skillgrouplanguagecondition `json:"languageSkillConditions,omitempty"`


	// Operation - Operator that will be applied to the conditions
	Operation *string `json:"operation,omitempty"`

}

func (o *Skillgroupcondition) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Skillgroupcondition
	
	return json.Marshal(&struct { 
		RoutingSkillConditions *[]Skillgrouproutingcondition `json:"routingSkillConditions,omitempty"`
		
		LanguageSkillConditions *[]Skillgrouplanguagecondition `json:"languageSkillConditions,omitempty"`
		
		Operation *string `json:"operation,omitempty"`
		*Alias
	}{ 
		RoutingSkillConditions: o.RoutingSkillConditions,
		
		LanguageSkillConditions: o.LanguageSkillConditions,
		
		Operation: o.Operation,
		Alias:    (*Alias)(o),
	})
}

func (o *Skillgroupcondition) UnmarshalJSON(b []byte) error {
	var SkillgroupconditionMap map[string]interface{}
	err := json.Unmarshal(b, &SkillgroupconditionMap)
	if err != nil {
		return err
	}
	
	if RoutingSkillConditions, ok := SkillgroupconditionMap["routingSkillConditions"].([]interface{}); ok {
		RoutingSkillConditionsString, _ := json.Marshal(RoutingSkillConditions)
		json.Unmarshal(RoutingSkillConditionsString, &o.RoutingSkillConditions)
	}
	
	if LanguageSkillConditions, ok := SkillgroupconditionMap["languageSkillConditions"].([]interface{}); ok {
		LanguageSkillConditionsString, _ := json.Marshal(LanguageSkillConditions)
		json.Unmarshal(LanguageSkillConditionsString, &o.LanguageSkillConditions)
	}
	
	if Operation, ok := SkillgroupconditionMap["operation"].(string); ok {
		o.Operation = &Operation
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Skillgroupcondition) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

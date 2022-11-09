package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Skillgrouproutingcondition
type Skillgrouproutingcondition struct { 
	// RoutingSkill - The routing skill to be used in the skill condition query
	RoutingSkill *string `json:"routingSkill,omitempty"`


	// Comparator - Comparator that will be applied to the proficiency
	Comparator *string `json:"comparator,omitempty"`


	// Proficiency - The skill proficiency that will be used for the routing skill. Integer range 0-5
	Proficiency *int `json:"proficiency,omitempty"`


	// ChildConditions - Nested conditions to be applied to this skill condition
	ChildConditions *[]Skillgroupcondition `json:"childConditions,omitempty"`

}

func (o *Skillgrouproutingcondition) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Skillgrouproutingcondition
	
	return json.Marshal(&struct { 
		RoutingSkill *string `json:"routingSkill,omitempty"`
		
		Comparator *string `json:"comparator,omitempty"`
		
		Proficiency *int `json:"proficiency,omitempty"`
		
		ChildConditions *[]Skillgroupcondition `json:"childConditions,omitempty"`
		*Alias
	}{ 
		RoutingSkill: o.RoutingSkill,
		
		Comparator: o.Comparator,
		
		Proficiency: o.Proficiency,
		
		ChildConditions: o.ChildConditions,
		Alias:    (*Alias)(o),
	})
}

func (o *Skillgrouproutingcondition) UnmarshalJSON(b []byte) error {
	var SkillgrouproutingconditionMap map[string]interface{}
	err := json.Unmarshal(b, &SkillgrouproutingconditionMap)
	if err != nil {
		return err
	}
	
	if RoutingSkill, ok := SkillgrouproutingconditionMap["routingSkill"].(string); ok {
		o.RoutingSkill = &RoutingSkill
	}
    
	if Comparator, ok := SkillgrouproutingconditionMap["comparator"].(string); ok {
		o.Comparator = &Comparator
	}
    
	if Proficiency, ok := SkillgrouproutingconditionMap["proficiency"].(float64); ok {
		ProficiencyInt := int(Proficiency)
		o.Proficiency = &ProficiencyInt
	}
	
	if ChildConditions, ok := SkillgrouproutingconditionMap["childConditions"].([]interface{}); ok {
		ChildConditionsString, _ := json.Marshal(ChildConditions)
		json.Unmarshal(ChildConditionsString, &o.ChildConditions)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Skillgrouproutingcondition) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

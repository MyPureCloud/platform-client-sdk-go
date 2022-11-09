package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Skillgrouplanguagecondition
type Skillgrouplanguagecondition struct { 
	// LanguageSkill - The language skill to be used in the skill condition query
	LanguageSkill *string `json:"languageSkill,omitempty"`


	// Comparator - Comparator that will be applied to the proficiency
	Comparator *string `json:"comparator,omitempty"`


	// Proficiency - The skill proficiency that will be used for the language skill. Integer range 0-5
	Proficiency *int `json:"proficiency,omitempty"`


	// ChildConditions - Nested conditions to be applied to this skill condition
	ChildConditions *[]Skillgroupcondition `json:"childConditions,omitempty"`

}

func (o *Skillgrouplanguagecondition) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Skillgrouplanguagecondition
	
	return json.Marshal(&struct { 
		LanguageSkill *string `json:"languageSkill,omitempty"`
		
		Comparator *string `json:"comparator,omitempty"`
		
		Proficiency *int `json:"proficiency,omitempty"`
		
		ChildConditions *[]Skillgroupcondition `json:"childConditions,omitempty"`
		*Alias
	}{ 
		LanguageSkill: o.LanguageSkill,
		
		Comparator: o.Comparator,
		
		Proficiency: o.Proficiency,
		
		ChildConditions: o.ChildConditions,
		Alias:    (*Alias)(o),
	})
}

func (o *Skillgrouplanguagecondition) UnmarshalJSON(b []byte) error {
	var SkillgrouplanguageconditionMap map[string]interface{}
	err := json.Unmarshal(b, &SkillgrouplanguageconditionMap)
	if err != nil {
		return err
	}
	
	if LanguageSkill, ok := SkillgrouplanguageconditionMap["languageSkill"].(string); ok {
		o.LanguageSkill = &LanguageSkill
	}
    
	if Comparator, ok := SkillgrouplanguageconditionMap["comparator"].(string); ok {
		o.Comparator = &Comparator
	}
    
	if Proficiency, ok := SkillgrouplanguageconditionMap["proficiency"].(float64); ok {
		ProficiencyInt := int(Proficiency)
		o.Proficiency = &ProficiencyInt
	}
	
	if ChildConditions, ok := SkillgrouplanguageconditionMap["childConditions"].([]interface{}); ok {
		ChildConditionsString, _ := json.Marshal(ChildConditions)
		json.Unmarshal(ChildConditionsString, &o.ChildConditions)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Skillgrouplanguagecondition) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

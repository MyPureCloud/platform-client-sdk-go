package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Routingconversationattributesresponse
type Routingconversationattributesresponse struct { 
	// Priority - Current priority value on in-queue conversation. Range:[-25000000, 25000000]
	Priority *int `json:"priority,omitempty"`


	// Skills - Current routing skills on in-queue conversation
	Skills *[]Routingskill `json:"skills,omitempty"`


	// Language - Current language on in-queue conversation
	Language *Language `json:"language,omitempty"`


	// ScoredAgents - Current scored agents on in-queue conversation
	ScoredAgents *[]Scoredagent `json:"scoredAgents,omitempty"`

}

func (o *Routingconversationattributesresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Routingconversationattributesresponse
	
	return json.Marshal(&struct { 
		Priority *int `json:"priority,omitempty"`
		
		Skills *[]Routingskill `json:"skills,omitempty"`
		
		Language *Language `json:"language,omitempty"`
		
		ScoredAgents *[]Scoredagent `json:"scoredAgents,omitempty"`
		*Alias
	}{ 
		Priority: o.Priority,
		
		Skills: o.Skills,
		
		Language: o.Language,
		
		ScoredAgents: o.ScoredAgents,
		Alias:    (*Alias)(o),
	})
}

func (o *Routingconversationattributesresponse) UnmarshalJSON(b []byte) error {
	var RoutingconversationattributesresponseMap map[string]interface{}
	err := json.Unmarshal(b, &RoutingconversationattributesresponseMap)
	if err != nil {
		return err
	}
	
	if Priority, ok := RoutingconversationattributesresponseMap["priority"].(float64); ok {
		PriorityInt := int(Priority)
		o.Priority = &PriorityInt
	}
	
	if Skills, ok := RoutingconversationattributesresponseMap["skills"].([]interface{}); ok {
		SkillsString, _ := json.Marshal(Skills)
		json.Unmarshal(SkillsString, &o.Skills)
	}
	
	if Language, ok := RoutingconversationattributesresponseMap["language"].(map[string]interface{}); ok {
		LanguageString, _ := json.Marshal(Language)
		json.Unmarshal(LanguageString, &o.Language)
	}
	
	if ScoredAgents, ok := RoutingconversationattributesresponseMap["scoredAgents"].([]interface{}); ok {
		ScoredAgentsString, _ := json.Marshal(ScoredAgents)
		json.Unmarshal(ScoredAgentsString, &o.ScoredAgents)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Routingconversationattributesresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

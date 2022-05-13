package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Userroutingskill - Represents an organization skill assigned to a user. When assigning to a user specify the organization skill id as the id.
type Userroutingskill struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Proficiency - A rating from 0.0 to 5.0 that indicates how adept an agent is at a particular skill. When \"Best available skills\" is enabled for a queue in Genesys Cloud, ACD interactions in that queue are routed to agents with higher proficiency ratings.
	Proficiency *float64 `json:"proficiency,omitempty"`


	// State - Activate or deactivate this routing skill.
	State *string `json:"state,omitempty"`


	// SkillUri - URI to the organization skill used by this user skill.
	SkillUri *string `json:"skillUri,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Userroutingskill) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Userroutingskill
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Proficiency *float64 `json:"proficiency,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		SkillUri *string `json:"skillUri,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Proficiency: o.Proficiency,
		
		State: o.State,
		
		SkillUri: o.SkillUri,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Userroutingskill) UnmarshalJSON(b []byte) error {
	var UserroutingskillMap map[string]interface{}
	err := json.Unmarshal(b, &UserroutingskillMap)
	if err != nil {
		return err
	}
	
	if Id, ok := UserroutingskillMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := UserroutingskillMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Proficiency, ok := UserroutingskillMap["proficiency"].(float64); ok {
		o.Proficiency = &Proficiency
	}
    
	if State, ok := UserroutingskillMap["state"].(string); ok {
		o.State = &State
	}
    
	if SkillUri, ok := UserroutingskillMap["skillUri"].(string); ok {
		o.SkillUri = &SkillUri
	}
    
	if SelfUri, ok := UserroutingskillMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Userroutingskill) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

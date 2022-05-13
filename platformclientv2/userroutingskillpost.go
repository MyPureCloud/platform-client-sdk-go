package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Userroutingskillpost - Represents an organization skill assigned to a user. When assigning to a user specify the organization skill id as the id.
type Userroutingskillpost struct { 
	// Id - The id of the existing routing skill to add to the user
	Id *string `json:"id,omitempty"`


	// Proficiency - Proficiency is a rating from 0.0 to 5.0 on how competent an agent is for a particular skill. It is used when a queue is set to \"Best available skills\" mode to allow acd interactions to target agents with higher proficiency ratings.
	Proficiency *float64 `json:"proficiency,omitempty"`


	// SkillUri - URI to the organization skill used by this user skill.
	SkillUri *string `json:"skillUri,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Userroutingskillpost) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Userroutingskillpost
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Proficiency *float64 `json:"proficiency,omitempty"`
		
		SkillUri *string `json:"skillUri,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Proficiency: o.Proficiency,
		
		SkillUri: o.SkillUri,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Userroutingskillpost) UnmarshalJSON(b []byte) error {
	var UserroutingskillpostMap map[string]interface{}
	err := json.Unmarshal(b, &UserroutingskillpostMap)
	if err != nil {
		return err
	}
	
	if Id, ok := UserroutingskillpostMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Proficiency, ok := UserroutingskillpostMap["proficiency"].(float64); ok {
		o.Proficiency = &Proficiency
	}
    
	if SkillUri, ok := UserroutingskillpostMap["skillUri"].(string); ok {
		o.SkillUri = &SkillUri
	}
    
	if SelfUri, ok := UserroutingskillpostMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Userroutingskillpost) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

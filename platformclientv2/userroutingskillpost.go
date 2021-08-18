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

func (u *Userroutingskillpost) MarshalJSON() ([]byte, error) {
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
		Id: u.Id,
		
		Proficiency: u.Proficiency,
		
		SkillUri: u.SkillUri,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Userroutingskillpost) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Actions
type Actions struct { 
	// SkillsToRemove
	SkillsToRemove *[]Skillstoremove `json:"skillsToRemove,omitempty"`

}

func (u *Actions) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Actions

	

	return json.Marshal(&struct { 
		SkillsToRemove *[]Skillstoremove `json:"skillsToRemove,omitempty"`
		*Alias
	}{ 
		SkillsToRemove: u.SkillsToRemove,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Actions) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

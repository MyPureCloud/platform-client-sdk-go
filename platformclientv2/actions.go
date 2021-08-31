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

func (o *Actions) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Actions
	
	return json.Marshal(&struct { 
		SkillsToRemove *[]Skillstoremove `json:"skillsToRemove,omitempty"`
		*Alias
	}{ 
		SkillsToRemove: o.SkillsToRemove,
		Alias:    (*Alias)(o),
	})
}

func (o *Actions) UnmarshalJSON(b []byte) error {
	var ActionsMap map[string]interface{}
	err := json.Unmarshal(b, &ActionsMap)
	if err != nil {
		return err
	}
	
	if SkillsToRemove, ok := ActionsMap["skillsToRemove"].([]interface{}); ok {
		SkillsToRemoveString, _ := json.Marshal(SkillsToRemove)
		json.Unmarshal(SkillsToRemoveString, &o.SkillsToRemove)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Actions) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

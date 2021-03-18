package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Actions
type Actions struct { 
	// SkillsToRemove
	SkillsToRemove *[]Skillstoremove `json:"skillsToRemove,omitempty"`

}

// String returns a JSON representation of the model
func (o *Actions) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

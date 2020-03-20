package platformclientv2
import (
	"encoding/json"
)

// Actions
type Actions struct { 
	// SkillsToRemove
	SkillsToRemove *[]Skillstoremove `json:"skillsToRemove,omitempty"`

}

// String returns a JSON representation of the model
func (o *Actions) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

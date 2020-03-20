package platformclientv2
import (
	"encoding/json"
)

// Contactcentersettings
type Contactcentersettings struct { 
	// RemoveSkillsFromBlindTransfer - Strip skills from transfer
	RemoveSkillsFromBlindTransfer *bool `json:"removeSkillsFromBlindTransfer,omitempty"`

}

// String returns a JSON representation of the model
func (o *Contactcentersettings) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

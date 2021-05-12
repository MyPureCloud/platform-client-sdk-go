package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Contactcentersettings
type Contactcentersettings struct { 
	// RemoveSkillsFromBlindTransfer - Strip skills from transfer
	RemoveSkillsFromBlindTransfer *bool `json:"removeSkillsFromBlindTransfer,omitempty"`

}

// String returns a JSON representation of the model
func (o *Contactcentersettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Contactcentersettings
type Contactcentersettings struct { 
	// RemoveSkillsFromBlindTransfer - Strip skills from transfer
	RemoveSkillsFromBlindTransfer *bool `json:"removeSkillsFromBlindTransfer,omitempty"`

}

func (u *Contactcentersettings) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Contactcentersettings

	

	return json.Marshal(&struct { 
		RemoveSkillsFromBlindTransfer *bool `json:"removeSkillsFromBlindTransfer,omitempty"`
		*Alias
	}{ 
		RemoveSkillsFromBlindTransfer: u.RemoveSkillsFromBlindTransfer,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Contactcentersettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

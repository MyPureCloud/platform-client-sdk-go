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

func (o *Contactcentersettings) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Contactcentersettings
	
	return json.Marshal(&struct { 
		RemoveSkillsFromBlindTransfer *bool `json:"removeSkillsFromBlindTransfer,omitempty"`
		*Alias
	}{ 
		RemoveSkillsFromBlindTransfer: o.RemoveSkillsFromBlindTransfer,
		Alias:    (*Alias)(o),
	})
}

func (o *Contactcentersettings) UnmarshalJSON(b []byte) error {
	var ContactcentersettingsMap map[string]interface{}
	err := json.Unmarshal(b, &ContactcentersettingsMap)
	if err != nil {
		return err
	}
	
	if RemoveSkillsFromBlindTransfer, ok := ContactcentersettingsMap["removeSkillsFromBlindTransfer"].(bool); ok {
		o.RemoveSkillsFromBlindTransfer = &RemoveSkillsFromBlindTransfer
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Contactcentersettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

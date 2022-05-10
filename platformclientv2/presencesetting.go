package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Presencesetting
type Presencesetting struct { 
	// Join - Should Presence Events be sent
	Join *Settingdirection `json:"join,omitempty"`

}

func (o *Presencesetting) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Presencesetting
	
	return json.Marshal(&struct { 
		Join *Settingdirection `json:"join,omitempty"`
		*Alias
	}{ 
		Join: o.Join,
		Alias:    (*Alias)(o),
	})
}

func (o *Presencesetting) UnmarshalJSON(b []byte) error {
	var PresencesettingMap map[string]interface{}
	err := json.Unmarshal(b, &PresencesettingMap)
	if err != nil {
		return err
	}
	
	if Join, ok := PresencesettingMap["join"].(map[string]interface{}); ok {
		JoinString, _ := json.Marshal(Join)
		json.Unmarshal(JoinString, &o.Join)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Presencesetting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

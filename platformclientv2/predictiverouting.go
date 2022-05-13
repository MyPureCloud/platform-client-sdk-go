package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Predictiverouting
type Predictiverouting struct { 
	// RespectSkills - A switch used to determine if agent skills will be considered.
	RespectSkills *bool `json:"respectSkills,omitempty"`

}

func (o *Predictiverouting) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Predictiverouting
	
	return json.Marshal(&struct { 
		RespectSkills *bool `json:"respectSkills,omitempty"`
		*Alias
	}{ 
		RespectSkills: o.RespectSkills,
		Alias:    (*Alias)(o),
	})
}

func (o *Predictiverouting) UnmarshalJSON(b []byte) error {
	var PredictiveroutingMap map[string]interface{}
	err := json.Unmarshal(b, &PredictiveroutingMap)
	if err != nil {
		return err
	}
	
	if RespectSkills, ok := PredictiveroutingMap["respectSkills"].(bool); ok {
		o.RespectSkills = &RespectSkills
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Predictiverouting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

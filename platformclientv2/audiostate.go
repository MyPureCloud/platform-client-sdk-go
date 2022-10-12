package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Audiostate
type Audiostate struct { 
	// CanHear - Indicates that this communication's audio allows its participant to hear others.
	CanHear *bool `json:"canHear,omitempty"`


	// CanSpeak - Indicates that this communication's audio allows others to hear this participant.
	CanSpeak *bool `json:"canSpeak,omitempty"`

}

func (o *Audiostate) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Audiostate
	
	return json.Marshal(&struct { 
		CanHear *bool `json:"canHear,omitempty"`
		
		CanSpeak *bool `json:"canSpeak,omitempty"`
		*Alias
	}{ 
		CanHear: o.CanHear,
		
		CanSpeak: o.CanSpeak,
		Alias:    (*Alias)(o),
	})
}

func (o *Audiostate) UnmarshalJSON(b []byte) error {
	var AudiostateMap map[string]interface{}
	err := json.Unmarshal(b, &AudiostateMap)
	if err != nil {
		return err
	}
	
	if CanHear, ok := AudiostateMap["canHear"].(bool); ok {
		o.CanHear = &CanHear
	}
    
	if CanSpeak, ok := AudiostateMap["canSpeak"].(bool); ok {
		o.CanSpeak = &CanSpeak
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Audiostate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

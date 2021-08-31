package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Textbotuserinputevent - Settings for an input event to the bot flow indicating user input is available.
type Textbotuserinputevent struct { 
	// Mode - The input mode.
	Mode *string `json:"mode,omitempty"`


	// Alternatives - The input alternatives.
	Alternatives *[]Textbotuserinputalternative `json:"alternatives,omitempty"`

}

func (o *Textbotuserinputevent) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Textbotuserinputevent
	
	return json.Marshal(&struct { 
		Mode *string `json:"mode,omitempty"`
		
		Alternatives *[]Textbotuserinputalternative `json:"alternatives,omitempty"`
		*Alias
	}{ 
		Mode: o.Mode,
		
		Alternatives: o.Alternatives,
		Alias:    (*Alias)(o),
	})
}

func (o *Textbotuserinputevent) UnmarshalJSON(b []byte) error {
	var TextbotuserinputeventMap map[string]interface{}
	err := json.Unmarshal(b, &TextbotuserinputeventMap)
	if err != nil {
		return err
	}
	
	if Mode, ok := TextbotuserinputeventMap["mode"].(string); ok {
		o.Mode = &Mode
	}
	
	if Alternatives, ok := TextbotuserinputeventMap["alternatives"].([]interface{}); ok {
		AlternativesString, _ := json.Marshal(Alternatives)
		json.Unmarshal(AlternativesString, &o.Alternatives)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Textbotuserinputevent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

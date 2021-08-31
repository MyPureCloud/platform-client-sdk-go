package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Textbotwaitforinputaction - Settings for a next-action of waiting for additional user input and sending the data as an input action to the bot flow.
type Textbotwaitforinputaction struct { 
	// ModeConstraints - The mode constraints for the user input.
	ModeConstraints *Textbotmodeconstraints `json:"modeConstraints,omitempty"`

}

func (o *Textbotwaitforinputaction) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Textbotwaitforinputaction
	
	return json.Marshal(&struct { 
		ModeConstraints *Textbotmodeconstraints `json:"modeConstraints,omitempty"`
		*Alias
	}{ 
		ModeConstraints: o.ModeConstraints,
		Alias:    (*Alias)(o),
	})
}

func (o *Textbotwaitforinputaction) UnmarshalJSON(b []byte) error {
	var TextbotwaitforinputactionMap map[string]interface{}
	err := json.Unmarshal(b, &TextbotwaitforinputactionMap)
	if err != nil {
		return err
	}
	
	if ModeConstraints, ok := TextbotwaitforinputactionMap["modeConstraints"].(map[string]interface{}); ok {
		ModeConstraintsString, _ := json.Marshal(ModeConstraints)
		json.Unmarshal(ModeConstraintsString, &o.ModeConstraints)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Textbotwaitforinputaction) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

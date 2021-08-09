package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Textbotuserinputalternative - User input data used in a bot flow turn.
type Textbotuserinputalternative struct { 
	// Transcript - The user input transcript.
	Transcript *Textbottranscript `json:"transcript,omitempty"`

}

// String returns a JSON representation of the model
func (o *Textbotuserinputalternative) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

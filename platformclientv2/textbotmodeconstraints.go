package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Textbotmodeconstraints - Mode constraints to observe when operating on a bot flow.
type Textbotmodeconstraints struct { 
	// Text - Mode constraints that apply to text scenarios.
	Text *Textbottextmodeconstraints `json:"text,omitempty"`

}

// String returns a JSON representation of the model
func (o *Textbotmodeconstraints) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

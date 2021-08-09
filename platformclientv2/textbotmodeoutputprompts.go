package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Textbotmodeoutputprompts - Prompt information related to a bot flow turn.
type Textbotmodeoutputprompts struct { 
	// Segments - The list of prompt segments.
	Segments *[]Textbotpromptsegment `json:"segments,omitempty"`

}

// String returns a JSON representation of the model
func (o *Textbotmodeoutputprompts) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

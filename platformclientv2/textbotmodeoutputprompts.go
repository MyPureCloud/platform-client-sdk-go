package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Textbotmodeoutputprompts - Prompt information related to a bot flow turn.
type Textbotmodeoutputprompts struct { 
	// Segments - The list of prompt segments.
	Segments *[]Textbotpromptsegment `json:"segments,omitempty"`

}

func (u *Textbotmodeoutputprompts) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Textbotmodeoutputprompts

	

	return json.Marshal(&struct { 
		Segments *[]Textbotpromptsegment `json:"segments,omitempty"`
		*Alias
	}{ 
		Segments: u.Segments,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Textbotmodeoutputprompts) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

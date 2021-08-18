package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Textbotmodeconstraints - Mode constraints to observe when operating on a bot flow.
type Textbotmodeconstraints struct { 
	// Text - Mode constraints that apply to text scenarios.
	Text *Textbottextmodeconstraints `json:"text,omitempty"`

}

func (u *Textbotmodeconstraints) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Textbotmodeconstraints

	

	return json.Marshal(&struct { 
		Text *Textbottextmodeconstraints `json:"text,omitempty"`
		*Alias
	}{ 
		Text: u.Text,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Textbotmodeconstraints) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

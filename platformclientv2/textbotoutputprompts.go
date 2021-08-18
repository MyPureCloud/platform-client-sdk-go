package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Textbotoutputprompts - Prompt information related to a bot flow turn.
type Textbotoutputprompts struct { 
	// OutputLanguage - The ISO code of the output language for this prompt item.
	OutputLanguage *string `json:"outputLanguage,omitempty"`


	// TextPrompts - Text output prompts, if any.
	TextPrompts *Textbotmodeoutputprompts `json:"textPrompts,omitempty"`

}

func (u *Textbotoutputprompts) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Textbotoutputprompts

	

	return json.Marshal(&struct { 
		OutputLanguage *string `json:"outputLanguage,omitempty"`
		
		TextPrompts *Textbotmodeoutputprompts `json:"textPrompts,omitempty"`
		*Alias
	}{ 
		OutputLanguage: u.OutputLanguage,
		
		TextPrompts: u.TextPrompts,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Textbotoutputprompts) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

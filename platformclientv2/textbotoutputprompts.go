package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Textbotoutputprompts) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

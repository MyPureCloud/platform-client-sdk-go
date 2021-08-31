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

func (o *Textbotmodeconstraints) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Textbotmodeconstraints
	
	return json.Marshal(&struct { 
		Text *Textbottextmodeconstraints `json:"text,omitempty"`
		*Alias
	}{ 
		Text: o.Text,
		Alias:    (*Alias)(o),
	})
}

func (o *Textbotmodeconstraints) UnmarshalJSON(b []byte) error {
	var TextbotmodeconstraintsMap map[string]interface{}
	err := json.Unmarshal(b, &TextbotmodeconstraintsMap)
	if err != nil {
		return err
	}
	
	if Text, ok := TextbotmodeconstraintsMap["text"].(map[string]interface{}); ok {
		TextString, _ := json.Marshal(Text)
		json.Unmarshal(TextString, &o.Text)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Textbotmodeconstraints) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

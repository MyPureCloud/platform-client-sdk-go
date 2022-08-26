package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Documenttext
type Documenttext struct { 
	// Text - Text.
	Text *string `json:"text,omitempty"`


	// Marks - The unique list of marks (whether it is bold and/or underlined etc.) for the text.
	Marks *[]string `json:"marks,omitempty"`


	// Hyperlink - The URL of the page that the hyperlink goes to.
	Hyperlink *string `json:"hyperlink,omitempty"`

}

func (o *Documenttext) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Documenttext
	
	return json.Marshal(&struct { 
		Text *string `json:"text,omitempty"`
		
		Marks *[]string `json:"marks,omitempty"`
		
		Hyperlink *string `json:"hyperlink,omitempty"`
		*Alias
	}{ 
		Text: o.Text,
		
		Marks: o.Marks,
		
		Hyperlink: o.Hyperlink,
		Alias:    (*Alias)(o),
	})
}

func (o *Documenttext) UnmarshalJSON(b []byte) error {
	var DocumenttextMap map[string]interface{}
	err := json.Unmarshal(b, &DocumenttextMap)
	if err != nil {
		return err
	}
	
	if Text, ok := DocumenttextMap["text"].(string); ok {
		o.Text = &Text
	}
    
	if Marks, ok := DocumenttextMap["marks"].([]interface{}); ok {
		MarksString, _ := json.Marshal(Marks)
		json.Unmarshal(MarksString, &o.Marks)
	}
	
	if Hyperlink, ok := DocumenttextMap["hyperlink"].(string); ok {
		o.Hyperlink = &Hyperlink
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Documenttext) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

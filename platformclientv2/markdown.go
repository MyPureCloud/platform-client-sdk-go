package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Markdown
type Markdown struct { 
	// Enabled - whether or not markdown is enabled
	Enabled *bool `json:"enabled,omitempty"`

}

func (o *Markdown) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Markdown
	
	return json.Marshal(&struct { 
		Enabled *bool `json:"enabled,omitempty"`
		*Alias
	}{ 
		Enabled: o.Enabled,
		Alias:    (*Alias)(o),
	})
}

func (o *Markdown) UnmarshalJSON(b []byte) error {
	var MarkdownMap map[string]interface{}
	err := json.Unmarshal(b, &MarkdownMap)
	if err != nil {
		return err
	}
	
	if Enabled, ok := MarkdownMap["enabled"].(bool); ok {
		o.Enabled = &Enabled
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Markdown) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

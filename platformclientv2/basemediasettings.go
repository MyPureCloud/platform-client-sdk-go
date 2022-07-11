package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Basemediasettings
type Basemediasettings struct { 
	// EnableAutoAnswer - Indicates if auto-answer is enabled for the given media type or subtype (default is false).  Subtype settings take precedence over media type settings.
	EnableAutoAnswer *bool `json:"enableAutoAnswer,omitempty"`

}

func (o *Basemediasettings) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Basemediasettings
	
	return json.Marshal(&struct { 
		EnableAutoAnswer *bool `json:"enableAutoAnswer,omitempty"`
		*Alias
	}{ 
		EnableAutoAnswer: o.EnableAutoAnswer,
		Alias:    (*Alias)(o),
	})
}

func (o *Basemediasettings) UnmarshalJSON(b []byte) error {
	var BasemediasettingsMap map[string]interface{}
	err := json.Unmarshal(b, &BasemediasettingsMap)
	if err != nil {
		return err
	}
	
	if EnableAutoAnswer, ok := BasemediasettingsMap["enableAutoAnswer"].(bool); ok {
		o.EnableAutoAnswer = &EnableAutoAnswer
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Basemediasettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

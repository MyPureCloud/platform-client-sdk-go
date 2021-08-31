package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Importerror
type Importerror struct { 
	// Message
	Message *string `json:"message,omitempty"`


	// Line
	Line *int `json:"line,omitempty"`

}

func (o *Importerror) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Importerror
	
	return json.Marshal(&struct { 
		Message *string `json:"message,omitempty"`
		
		Line *int `json:"line,omitempty"`
		*Alias
	}{ 
		Message: o.Message,
		
		Line: o.Line,
		Alias:    (*Alias)(o),
	})
}

func (o *Importerror) UnmarshalJSON(b []byte) error {
	var ImporterrorMap map[string]interface{}
	err := json.Unmarshal(b, &ImporterrorMap)
	if err != nil {
		return err
	}
	
	if Message, ok := ImporterrorMap["message"].(string); ok {
		o.Message = &Message
	}
	
	if Line, ok := ImporterrorMap["line"].(float64); ok {
		LineInt := int(Line)
		o.Line = &LineInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Importerror) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

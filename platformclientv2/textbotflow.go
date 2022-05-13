package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Textbotflow - Description of the Bot Flow.
type Textbotflow struct { 
	// Id - The Bot Flow ID.
	Id *string `json:"id,omitempty"`

}

func (o *Textbotflow) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Textbotflow
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		Alias:    (*Alias)(o),
	})
}

func (o *Textbotflow) UnmarshalJSON(b []byte) error {
	var TextbotflowMap map[string]interface{}
	err := json.Unmarshal(b, &TextbotflowMap)
	if err != nil {
		return err
	}
	
	if Id, ok := TextbotflowMap["id"].(string); ok {
		o.Id = &Id
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Textbotflow) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

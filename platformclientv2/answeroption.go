package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Answeroption
type Answeroption struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Text
	Text *string `json:"text,omitempty"`


	// Value
	Value *int `json:"value,omitempty"`

}

func (u *Answeroption) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Answeroption

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Text *string `json:"text,omitempty"`
		
		Value *int `json:"value,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Text: u.Text,
		
		Value: u.Value,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Answeroption) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

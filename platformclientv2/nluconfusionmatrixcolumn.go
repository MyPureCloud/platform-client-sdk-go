package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Nluconfusionmatrixcolumn
type Nluconfusionmatrixcolumn struct { 
	// Name - The name of the intent for the column.
	Name *string `json:"name,omitempty"`


	// Value - The confusion value between the intents
	Value *float32 `json:"value,omitempty"`

}

func (u *Nluconfusionmatrixcolumn) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Nluconfusionmatrixcolumn

	

	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		Value *float32 `json:"value,omitempty"`
		*Alias
	}{ 
		Name: u.Name,
		
		Value: u.Value,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Nluconfusionmatrixcolumn) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

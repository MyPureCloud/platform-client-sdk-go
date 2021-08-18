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

func (u *Importerror) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Importerror

	

	return json.Marshal(&struct { 
		Message *string `json:"message,omitempty"`
		
		Line *int `json:"line,omitempty"`
		*Alias
	}{ 
		Message: u.Message,
		
		Line: u.Line,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Importerror) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

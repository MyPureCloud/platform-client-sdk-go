package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Bullseye
type Bullseye struct { 
	// Rings
	Rings *[]Ring `json:"rings,omitempty"`

}

func (u *Bullseye) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Bullseye

	

	return json.Marshal(&struct { 
		Rings *[]Ring `json:"rings,omitempty"`
		*Alias
	}{ 
		Rings: u.Rings,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Bullseye) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

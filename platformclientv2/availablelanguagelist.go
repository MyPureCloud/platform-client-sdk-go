package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Availablelanguagelist
type Availablelanguagelist struct { 
	// Languages
	Languages *[]string `json:"languages,omitempty"`

}

func (u *Availablelanguagelist) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Availablelanguagelist

	

	return json.Marshal(&struct { 
		Languages *[]string `json:"languages,omitempty"`
		*Alias
	}{ 
		Languages: u.Languages,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Availablelanguagelist) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

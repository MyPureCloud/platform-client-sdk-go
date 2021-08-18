package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Exportscriptresponse
type Exportscriptresponse struct { 
	// Url
	Url *string `json:"url,omitempty"`

}

func (u *Exportscriptresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Exportscriptresponse

	

	return json.Marshal(&struct { 
		Url *string `json:"url,omitempty"`
		*Alias
	}{ 
		Url: u.Url,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Exportscriptresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

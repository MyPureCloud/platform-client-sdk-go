package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Forecastabandonrateresponse
type Forecastabandonrateresponse struct { 
	// Percent - The target percent abandon rate goal
	Percent *int `json:"percent,omitempty"`

}

func (u *Forecastabandonrateresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Forecastabandonrateresponse

	

	return json.Marshal(&struct { 
		Percent *int `json:"percent,omitempty"`
		*Alias
	}{ 
		Percent: u.Percent,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Forecastabandonrateresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

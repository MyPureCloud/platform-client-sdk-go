package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Bushorttermforecastingsettings
type Bushorttermforecastingsettings struct { 
	// DefaultHistoryWeeks - The number of historical weeks to consider when creating a forecast. This setting is only used for legacy weighted average forecasts
	DefaultHistoryWeeks *int `json:"defaultHistoryWeeks,omitempty"`

}

func (u *Bushorttermforecastingsettings) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Bushorttermforecastingsettings

	

	return json.Marshal(&struct { 
		DefaultHistoryWeeks *int `json:"defaultHistoryWeeks,omitempty"`
		*Alias
	}{ 
		DefaultHistoryWeeks: u.DefaultHistoryWeeks,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Bushorttermforecastingsettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

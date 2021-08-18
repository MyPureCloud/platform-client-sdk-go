package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Shorttermforecastingsettings - Short Term Forecasting Settings
type Shorttermforecastingsettings struct { 
	// DefaultHistoryWeeks - The number of weeks to consider by default when generating a volume forecast
	DefaultHistoryWeeks *int `json:"defaultHistoryWeeks,omitempty"`

}

func (u *Shorttermforecastingsettings) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Shorttermforecastingsettings

	

	return json.Marshal(&struct { 
		DefaultHistoryWeeks *int `json:"defaultHistoryWeeks,omitempty"`
		*Alias
	}{ 
		DefaultHistoryWeeks: u.DefaultHistoryWeeks,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Shorttermforecastingsettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

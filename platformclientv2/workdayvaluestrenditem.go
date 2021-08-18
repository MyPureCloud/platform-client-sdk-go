package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Workdayvaluestrenditem
type Workdayvaluestrenditem struct { 
	// DateWorkday - The workday for the metric value. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	DateWorkday *time.Time `json:"dateWorkday,omitempty"`


	// Value - The metric value
	Value *float64 `json:"value,omitempty"`

}

func (u *Workdayvaluestrenditem) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Workdayvaluestrenditem

	
	DateWorkday := new(string)
	if u.DateWorkday != nil {
		*DateWorkday = timeutil.Strftime(u.DateWorkday, "%Y-%m-%d")
	} else {
		DateWorkday = nil
	}
	

	return json.Marshal(&struct { 
		DateWorkday *string `json:"dateWorkday,omitempty"`
		
		Value *float64 `json:"value,omitempty"`
		*Alias
	}{ 
		DateWorkday: DateWorkday,
		
		Value: u.Value,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Workdayvaluestrenditem) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

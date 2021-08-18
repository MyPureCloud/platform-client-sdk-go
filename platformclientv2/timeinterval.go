package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Timeinterval
type Timeinterval struct { 
	// Months
	Months *int `json:"months,omitempty"`


	// Weeks
	Weeks *int `json:"weeks,omitempty"`


	// Days
	Days *int `json:"days,omitempty"`


	// Hours
	Hours *int `json:"hours,omitempty"`

}

func (u *Timeinterval) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Timeinterval

	

	return json.Marshal(&struct { 
		Months *int `json:"months,omitempty"`
		
		Weeks *int `json:"weeks,omitempty"`
		
		Days *int `json:"days,omitempty"`
		
		Hours *int `json:"hours,omitempty"`
		*Alias
	}{ 
		Months: u.Months,
		
		Weeks: u.Weeks,
		
		Days: u.Days,
		
		Hours: u.Hours,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Timeinterval) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

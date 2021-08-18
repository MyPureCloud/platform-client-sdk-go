package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Copybuforecastrequest
type Copybuforecastrequest struct { 
	// Description - The description for the forecast
	Description *string `json:"description,omitempty"`


	// WeekDate - The start date of the new forecast to create from the existing forecast. Must correspond to the start day of week for the business unit. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	WeekDate *time.Time `json:"weekDate,omitempty"`

}

func (u *Copybuforecastrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Copybuforecastrequest

	
	WeekDate := new(string)
	if u.WeekDate != nil {
		*WeekDate = timeutil.Strftime(u.WeekDate, "%Y-%m-%d")
	} else {
		WeekDate = nil
	}
	

	return json.Marshal(&struct { 
		Description *string `json:"description,omitempty"`
		
		WeekDate *string `json:"weekDate,omitempty"`
		*Alias
	}{ 
		Description: u.Description,
		
		WeekDate: WeekDate,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Copybuforecastrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

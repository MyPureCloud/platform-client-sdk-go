package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Bucreateblankschedulerequest
type Bucreateblankschedulerequest struct { 
	// Description - The description for the schedule
	Description *string `json:"description,omitempty"`


	// ShortTermForecast - The forecast to use when generating the schedule.  Note that the forecast must fully encompass the schedule's start week + week count
	ShortTermForecast *Bushorttermforecastreference `json:"shortTermForecast,omitempty"`


	// WeekCount - The number of weeks in the schedule. One extra day is added at the end
	WeekCount *int `json:"weekCount,omitempty"`

}

func (u *Bucreateblankschedulerequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Bucreateblankschedulerequest

	

	return json.Marshal(&struct { 
		Description *string `json:"description,omitempty"`
		
		ShortTermForecast *Bushorttermforecastreference `json:"shortTermForecast,omitempty"`
		
		WeekCount *int `json:"weekCount,omitempty"`
		*Alias
	}{ 
		Description: u.Description,
		
		ShortTermForecast: u.ShortTermForecast,
		
		WeekCount: u.WeekCount,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Bucreateblankschedulerequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

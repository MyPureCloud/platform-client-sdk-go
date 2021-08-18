package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Buforecastresult
type Buforecastresult struct { 
	// ReferenceStartDate - The reference start date for interval-based data for this forecast. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ReferenceStartDate *time.Time `json:"referenceStartDate,omitempty"`


	// PlanningGroups - The forecast data broken up by planning group
	PlanningGroups *[]Forecastplanninggroupdata `json:"planningGroups,omitempty"`


	// WeekNumber - The week number represented by this response
	WeekNumber *int `json:"weekNumber,omitempty"`


	// WeekCount - The number of weeks in this forecast
	WeekCount *int `json:"weekCount,omitempty"`

}

func (u *Buforecastresult) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Buforecastresult

	
	ReferenceStartDate := new(string)
	if u.ReferenceStartDate != nil {
		
		*ReferenceStartDate = timeutil.Strftime(u.ReferenceStartDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ReferenceStartDate = nil
	}
	

	return json.Marshal(&struct { 
		ReferenceStartDate *string `json:"referenceStartDate,omitempty"`
		
		PlanningGroups *[]Forecastplanninggroupdata `json:"planningGroups,omitempty"`
		
		WeekNumber *int `json:"weekNumber,omitempty"`
		
		WeekCount *int `json:"weekCount,omitempty"`
		*Alias
	}{ 
		ReferenceStartDate: ReferenceStartDate,
		
		PlanningGroups: u.PlanningGroups,
		
		WeekNumber: u.WeekNumber,
		
		WeekCount: u.WeekCount,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Buforecastresult) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

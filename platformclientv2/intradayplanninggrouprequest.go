package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Intradayplanninggrouprequest
type Intradayplanninggrouprequest struct { 
	// BusinessUnitDate - Requested date in yyyy-MM-dd format. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	BusinessUnitDate *time.Time `json:"businessUnitDate,omitempty"`


	// Categories - The metric categories
	Categories *[]string `json:"categories,omitempty"`


	// PlanningGroupIds - The IDs of the planning groups for which to fetch data.  Omitting or passing an empty list will return all available planning groups
	PlanningGroupIds *[]string `json:"planningGroupIds,omitempty"`


	// IntervalLengthMinutes - The period/interval in minutes for which to aggregate the data. Required, defaults to 15
	IntervalLengthMinutes *int `json:"intervalLengthMinutes,omitempty"`

}

func (u *Intradayplanninggrouprequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Intradayplanninggrouprequest

	
	BusinessUnitDate := new(string)
	if u.BusinessUnitDate != nil {
		*BusinessUnitDate = timeutil.Strftime(u.BusinessUnitDate, "%Y-%m-%d")
	} else {
		BusinessUnitDate = nil
	}
	

	return json.Marshal(&struct { 
		BusinessUnitDate *string `json:"businessUnitDate,omitempty"`
		
		Categories *[]string `json:"categories,omitempty"`
		
		PlanningGroupIds *[]string `json:"planningGroupIds,omitempty"`
		
		IntervalLengthMinutes *int `json:"intervalLengthMinutes,omitempty"`
		*Alias
	}{ 
		BusinessUnitDate: BusinessUnitDate,
		
		Categories: u.Categories,
		
		PlanningGroupIds: u.PlanningGroupIds,
		
		IntervalLengthMinutes: u.IntervalLengthMinutes,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Intradayplanninggrouprequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

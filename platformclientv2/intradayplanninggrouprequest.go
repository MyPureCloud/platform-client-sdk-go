package platformclientv2
import (
	"time"
	"encoding/json"
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
	IntervalLengthMinutes *int32 `json:"intervalLengthMinutes,omitempty"`

}

// String returns a JSON representation of the model
func (o *Intradayplanninggrouprequest) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

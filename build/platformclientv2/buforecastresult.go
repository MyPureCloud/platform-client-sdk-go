package platformclientv2
import (
	"time"
	"encoding/json"
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

// String returns a JSON representation of the model
func (o *Buforecastresult) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

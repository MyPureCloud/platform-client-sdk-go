package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Longtermforecastresult
type Longtermforecastresult struct { 
	// PlanningGroups - The forecast data broken up by planning group
	PlanningGroups *[]Longtermforecastplanninggroupdata `json:"planningGroups,omitempty"`


	// ReferenceStartDate - The reference start date relative to the business unit time zone in this forecast. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	ReferenceStartDate *time.Time `json:"referenceStartDate,omitempty"`


	// WeekCount - The number of weeks in this forecast
	WeekCount *int `json:"weekCount,omitempty"`

}

// String returns a JSON representation of the model
func (o *Longtermforecastresult) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

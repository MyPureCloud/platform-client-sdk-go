package platformclientv2
import (
	"encoding/json"
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

// String returns a JSON representation of the model
func (o *Bucreateblankschedulerequest) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

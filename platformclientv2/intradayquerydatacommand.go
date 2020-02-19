package platformclientv2
import (
	"time"
	"encoding/json"
)

// Intradayquerydatacommand
type Intradayquerydatacommand struct { 
	// StartDate - Start date of the requested date range in ISO-8601 format
	StartDate *time.Time `json:"startDate,omitempty"`


	// EndDate - End date of the requested date range in ISO-8601 format.  Must be within the same 7 day schedule week as defined by the management unit's start day of week
	EndDate *time.Time `json:"endDate,omitempty"`


	// Metrics - The metrics to validate
	Metrics *[]Intradaymetric `json:"metrics,omitempty"`


	// QueueIds - The queue IDs for which to fetch data.  Omitting or passing an empty list will return all available queues
	QueueIds *[]string `json:"queueIds,omitempty"`


	// IntervalLengthMinutes - The period/interval for which to aggregate the data.  Optional, defaults to 15
	IntervalLengthMinutes *int32 `json:"intervalLengthMinutes,omitempty"`

}

// String returns a JSON representation of the model
func (o *Intradayquerydatacommand) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

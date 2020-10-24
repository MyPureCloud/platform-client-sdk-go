package platformclientv2
import (
	"time"
	"encoding/json"
)

// Buintradayresponse
type Buintradayresponse struct { 
	// StartDate - The start of the date range for which this data applies.  This is also the start reference point for the intervals represented in the various arrays. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	StartDate *time.Time `json:"startDate,omitempty"`


	// EndDate - The end of the date range for which this data applies. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	EndDate *time.Time `json:"endDate,omitempty"`


	// IntervalLengthMinutes - The aggregation period in minutes, which determines the interval duration of the returned data
	IntervalLengthMinutes *int32 `json:"intervalLengthMinutes,omitempty"`


	// NoDataReason - If not null, the reason there was no data for the request
	NoDataReason *string `json:"noDataReason,omitempty"`


	// Categories - The categories to which this data corresponds
	Categories *[]string `json:"categories,omitempty"`


	// ShortTermForecast - Short term forecast reference
	ShortTermForecast *Bushorttermforecastreference `json:"shortTermForecast,omitempty"`


	// Schedule - Schedule reference
	Schedule *Buschedulereference `json:"schedule,omitempty"`


	// IntradayDataGroupings - Intraday data grouped by a single media type and set of planning group IDs
	IntradayDataGroupings *[]Buintradaydatagroup `json:"intradayDataGroupings,omitempty"`

}

// String returns a JSON representation of the model
func (o *Buintradayresponse) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

package platformclientv2
import (
	"time"
	"encoding/json"
)

// Intradayresponse
type Intradayresponse struct { 
	// StartDate - The start of the date range for which this data applies.  This is also the start reference point for the intervals represented in the various arrays. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	StartDate *time.Time `json:"startDate,omitempty"`


	// EndDate - The end of the date range for which this data applies. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	EndDate *time.Time `json:"endDate,omitempty"`


	// IntervalLengthMinutes - The aggregation period in minutes, which determines the interval duration of the returned data
	IntervalLengthMinutes *int32 `json:"intervalLengthMinutes,omitempty"`


	// NumberOfIntervals - The total number of time intervals represented by this data
	NumberOfIntervals *int32 `json:"numberOfIntervals,omitempty"`


	// Metrics - The metrics to which this data corresponds
	Metrics *[]Intradaymetric `json:"metrics,omitempty"`


	// NoDataReason - If not null, the reason there was no data for the request
	NoDataReason *string `json:"noDataReason,omitempty"`


	// QueueIds - The IDs of the queues this data corresponds to
	QueueIds *[]string `json:"queueIds,omitempty"`


	// IntradayDataGroupings - Intraday data grouped by a single media type and set of queue IDs
	IntradayDataGroupings *[]Intradaydatagroup `json:"intradayDataGroupings,omitempty"`

}

// String returns a JSON representation of the model
func (o *Intradayresponse) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

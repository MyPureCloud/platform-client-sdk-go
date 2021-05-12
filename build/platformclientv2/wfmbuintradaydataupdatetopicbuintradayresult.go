package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmbuintradaydataupdatetopicbuintradayresult
type Wfmbuintradaydataupdatetopicbuintradayresult struct { 
	// StartDate
	StartDate *time.Time `json:"startDate,omitempty"`


	// EndDate
	EndDate *time.Time `json:"endDate,omitempty"`


	// IntervalLengthMinutes
	IntervalLengthMinutes *int `json:"intervalLengthMinutes,omitempty"`


	// IntradayDataGroupings
	IntradayDataGroupings *[]Wfmbuintradaydataupdatetopicbuintradaydatagroup `json:"intradayDataGroupings,omitempty"`


	// Categories
	Categories *[]string `json:"categories,omitempty"`


	// NoDataReason
	NoDataReason *string `json:"noDataReason,omitempty"`


	// Schedule
	Schedule *Wfmbuintradaydataupdatetopicbuschedulereference `json:"schedule,omitempty"`


	// ShortTermForecast
	ShortTermForecast *Wfmbuintradaydataupdatetopicbushorttermforecastreference `json:"shortTermForecast,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wfmbuintradaydataupdatetopicbuintradayresult) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

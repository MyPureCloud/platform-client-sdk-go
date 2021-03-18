package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmintradaydataupdatetopicintradaydataupdate
type Wfmintradaydataupdatetopicintradaydataupdate struct { 
	// StartDate
	StartDate *time.Time `json:"startDate,omitempty"`


	// EndDate
	EndDate *time.Time `json:"endDate,omitempty"`


	// IntervalLengthMinutes
	IntervalLengthMinutes *int `json:"intervalLengthMinutes,omitempty"`


	// NumberOfIntervals
	NumberOfIntervals *int `json:"numberOfIntervals,omitempty"`


	// Metrics
	Metrics *[]Wfmintradaydataupdatetopicintradaymetric `json:"metrics,omitempty"`


	// QueueIds
	QueueIds *[]string `json:"queueIds,omitempty"`


	// IntradayDataGroupings
	IntradayDataGroupings *[]Wfmintradaydataupdatetopicintradaydatagroup `json:"intradayDataGroupings,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wfmintradaydataupdatetopicintradaydataupdate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

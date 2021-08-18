package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
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

func (u *Wfmintradaydataupdatetopicintradaydataupdate) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmintradaydataupdatetopicintradaydataupdate

	
	StartDate := new(string)
	if u.StartDate != nil {
		
		*StartDate = timeutil.Strftime(u.StartDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartDate = nil
	}
	
	EndDate := new(string)
	if u.EndDate != nil {
		
		*EndDate = timeutil.Strftime(u.EndDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		EndDate = nil
	}
	

	return json.Marshal(&struct { 
		StartDate *string `json:"startDate,omitempty"`
		
		EndDate *string `json:"endDate,omitempty"`
		
		IntervalLengthMinutes *int `json:"intervalLengthMinutes,omitempty"`
		
		NumberOfIntervals *int `json:"numberOfIntervals,omitempty"`
		
		Metrics *[]Wfmintradaydataupdatetopicintradaymetric `json:"metrics,omitempty"`
		
		QueueIds *[]string `json:"queueIds,omitempty"`
		
		IntradayDataGroupings *[]Wfmintradaydataupdatetopicintradaydatagroup `json:"intradayDataGroupings,omitempty"`
		*Alias
	}{ 
		StartDate: StartDate,
		
		EndDate: EndDate,
		
		IntervalLengthMinutes: u.IntervalLengthMinutes,
		
		NumberOfIntervals: u.NumberOfIntervals,
		
		Metrics: u.Metrics,
		
		QueueIds: u.QueueIds,
		
		IntradayDataGroupings: u.IntradayDataGroupings,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Wfmintradaydataupdatetopicintradaydataupdate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

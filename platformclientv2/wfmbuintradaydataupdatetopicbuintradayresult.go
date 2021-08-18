package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
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

func (u *Wfmbuintradaydataupdatetopicbuintradayresult) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmbuintradaydataupdatetopicbuintradayresult

	
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
		
		IntradayDataGroupings *[]Wfmbuintradaydataupdatetopicbuintradaydatagroup `json:"intradayDataGroupings,omitempty"`
		
		Categories *[]string `json:"categories,omitempty"`
		
		NoDataReason *string `json:"noDataReason,omitempty"`
		
		Schedule *Wfmbuintradaydataupdatetopicbuschedulereference `json:"schedule,omitempty"`
		
		ShortTermForecast *Wfmbuintradaydataupdatetopicbushorttermforecastreference `json:"shortTermForecast,omitempty"`
		*Alias
	}{ 
		StartDate: StartDate,
		
		EndDate: EndDate,
		
		IntervalLengthMinutes: u.IntervalLengthMinutes,
		
		IntradayDataGroupings: u.IntradayDataGroupings,
		
		Categories: u.Categories,
		
		NoDataReason: u.NoDataReason,
		
		Schedule: u.Schedule,
		
		ShortTermForecast: u.ShortTermForecast,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Wfmbuintradaydataupdatetopicbuintradayresult) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

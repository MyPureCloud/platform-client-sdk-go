package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Buintradayresponse
type Buintradayresponse struct { 
	// StartDate - The start of the date range for which this data applies.  This is also the start reference point for the intervals represented in the various arrays. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	StartDate *time.Time `json:"startDate,omitempty"`


	// EndDate - The end of the date range for which this data applies. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	EndDate *time.Time `json:"endDate,omitempty"`


	// IntervalLengthMinutes - The aggregation period in minutes, which determines the interval duration of the returned data
	IntervalLengthMinutes *int `json:"intervalLengthMinutes,omitempty"`


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

func (u *Buintradayresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Buintradayresponse

	
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
		
		NoDataReason *string `json:"noDataReason,omitempty"`
		
		Categories *[]string `json:"categories,omitempty"`
		
		ShortTermForecast *Bushorttermforecastreference `json:"shortTermForecast,omitempty"`
		
		Schedule *Buschedulereference `json:"schedule,omitempty"`
		
		IntradayDataGroupings *[]Buintradaydatagroup `json:"intradayDataGroupings,omitempty"`
		*Alias
	}{ 
		StartDate: StartDate,
		
		EndDate: EndDate,
		
		IntervalLengthMinutes: u.IntervalLengthMinutes,
		
		NoDataReason: u.NoDataReason,
		
		Categories: u.Categories,
		
		ShortTermForecast: u.ShortTermForecast,
		
		Schedule: u.Schedule,
		
		IntradayDataGroupings: u.IntradayDataGroupings,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Buintradayresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

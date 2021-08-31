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

func (o *Buintradayresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Buintradayresponse
	
	StartDate := new(string)
	if o.StartDate != nil {
		
		*StartDate = timeutil.Strftime(o.StartDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartDate = nil
	}
	
	EndDate := new(string)
	if o.EndDate != nil {
		
		*EndDate = timeutil.Strftime(o.EndDate, "%Y-%m-%dT%H:%M:%S.%fZ")
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
		
		IntervalLengthMinutes: o.IntervalLengthMinutes,
		
		NoDataReason: o.NoDataReason,
		
		Categories: o.Categories,
		
		ShortTermForecast: o.ShortTermForecast,
		
		Schedule: o.Schedule,
		
		IntradayDataGroupings: o.IntradayDataGroupings,
		Alias:    (*Alias)(o),
	})
}

func (o *Buintradayresponse) UnmarshalJSON(b []byte) error {
	var BuintradayresponseMap map[string]interface{}
	err := json.Unmarshal(b, &BuintradayresponseMap)
	if err != nil {
		return err
	}
	
	if startDateString, ok := BuintradayresponseMap["startDate"].(string); ok {
		StartDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", startDateString)
		o.StartDate = &StartDate
	}
	
	if endDateString, ok := BuintradayresponseMap["endDate"].(string); ok {
		EndDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", endDateString)
		o.EndDate = &EndDate
	}
	
	if IntervalLengthMinutes, ok := BuintradayresponseMap["intervalLengthMinutes"].(float64); ok {
		IntervalLengthMinutesInt := int(IntervalLengthMinutes)
		o.IntervalLengthMinutes = &IntervalLengthMinutesInt
	}
	
	if NoDataReason, ok := BuintradayresponseMap["noDataReason"].(string); ok {
		o.NoDataReason = &NoDataReason
	}
	
	if Categories, ok := BuintradayresponseMap["categories"].([]interface{}); ok {
		CategoriesString, _ := json.Marshal(Categories)
		json.Unmarshal(CategoriesString, &o.Categories)
	}
	
	if ShortTermForecast, ok := BuintradayresponseMap["shortTermForecast"].(map[string]interface{}); ok {
		ShortTermForecastString, _ := json.Marshal(ShortTermForecast)
		json.Unmarshal(ShortTermForecastString, &o.ShortTermForecast)
	}
	
	if Schedule, ok := BuintradayresponseMap["schedule"].(map[string]interface{}); ok {
		ScheduleString, _ := json.Marshal(Schedule)
		json.Unmarshal(ScheduleString, &o.Schedule)
	}
	
	if IntradayDataGroupings, ok := BuintradayresponseMap["intradayDataGroupings"].([]interface{}); ok {
		IntradayDataGroupingsString, _ := json.Marshal(IntradayDataGroupings)
		json.Unmarshal(IntradayDataGroupingsString, &o.IntradayDataGroupings)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Buintradayresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

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

func (o *Wfmbuintradaydataupdatetopicbuintradayresult) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmbuintradaydataupdatetopicbuintradayresult
	
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
		
		IntradayDataGroupings *[]Wfmbuintradaydataupdatetopicbuintradaydatagroup `json:"intradayDataGroupings,omitempty"`
		
		Categories *[]string `json:"categories,omitempty"`
		
		NoDataReason *string `json:"noDataReason,omitempty"`
		
		Schedule *Wfmbuintradaydataupdatetopicbuschedulereference `json:"schedule,omitempty"`
		
		ShortTermForecast *Wfmbuintradaydataupdatetopicbushorttermforecastreference `json:"shortTermForecast,omitempty"`
		*Alias
	}{ 
		StartDate: StartDate,
		
		EndDate: EndDate,
		
		IntervalLengthMinutes: o.IntervalLengthMinutes,
		
		IntradayDataGroupings: o.IntradayDataGroupings,
		
		Categories: o.Categories,
		
		NoDataReason: o.NoDataReason,
		
		Schedule: o.Schedule,
		
		ShortTermForecast: o.ShortTermForecast,
		Alias:    (*Alias)(o),
	})
}

func (o *Wfmbuintradaydataupdatetopicbuintradayresult) UnmarshalJSON(b []byte) error {
	var WfmbuintradaydataupdatetopicbuintradayresultMap map[string]interface{}
	err := json.Unmarshal(b, &WfmbuintradaydataupdatetopicbuintradayresultMap)
	if err != nil {
		return err
	}
	
	if startDateString, ok := WfmbuintradaydataupdatetopicbuintradayresultMap["startDate"].(string); ok {
		StartDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", startDateString)
		o.StartDate = &StartDate
	}
	
	if endDateString, ok := WfmbuintradaydataupdatetopicbuintradayresultMap["endDate"].(string); ok {
		EndDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", endDateString)
		o.EndDate = &EndDate
	}
	
	if IntervalLengthMinutes, ok := WfmbuintradaydataupdatetopicbuintradayresultMap["intervalLengthMinutes"].(float64); ok {
		IntervalLengthMinutesInt := int(IntervalLengthMinutes)
		o.IntervalLengthMinutes = &IntervalLengthMinutesInt
	}
	
	if IntradayDataGroupings, ok := WfmbuintradaydataupdatetopicbuintradayresultMap["intradayDataGroupings"].([]interface{}); ok {
		IntradayDataGroupingsString, _ := json.Marshal(IntradayDataGroupings)
		json.Unmarshal(IntradayDataGroupingsString, &o.IntradayDataGroupings)
	}
	
	if Categories, ok := WfmbuintradaydataupdatetopicbuintradayresultMap["categories"].([]interface{}); ok {
		CategoriesString, _ := json.Marshal(Categories)
		json.Unmarshal(CategoriesString, &o.Categories)
	}
	
	if NoDataReason, ok := WfmbuintradaydataupdatetopicbuintradayresultMap["noDataReason"].(string); ok {
		o.NoDataReason = &NoDataReason
	}
    
	if Schedule, ok := WfmbuintradaydataupdatetopicbuintradayresultMap["schedule"].(map[string]interface{}); ok {
		ScheduleString, _ := json.Marshal(Schedule)
		json.Unmarshal(ScheduleString, &o.Schedule)
	}
	
	if ShortTermForecast, ok := WfmbuintradaydataupdatetopicbuintradayresultMap["shortTermForecast"].(map[string]interface{}); ok {
		ShortTermForecastString, _ := json.Marshal(ShortTermForecast)
		json.Unmarshal(ShortTermForecastString, &o.ShortTermForecast)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmbuintradaydataupdatetopicbuintradayresult) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

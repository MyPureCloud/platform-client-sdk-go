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

func (o *Wfmintradaydataupdatetopicintradaydataupdate) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmintradaydataupdatetopicintradaydataupdate
	
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
		
		NumberOfIntervals *int `json:"numberOfIntervals,omitempty"`
		
		Metrics *[]Wfmintradaydataupdatetopicintradaymetric `json:"metrics,omitempty"`
		
		QueueIds *[]string `json:"queueIds,omitempty"`
		
		IntradayDataGroupings *[]Wfmintradaydataupdatetopicintradaydatagroup `json:"intradayDataGroupings,omitempty"`
		*Alias
	}{ 
		StartDate: StartDate,
		
		EndDate: EndDate,
		
		IntervalLengthMinutes: o.IntervalLengthMinutes,
		
		NumberOfIntervals: o.NumberOfIntervals,
		
		Metrics: o.Metrics,
		
		QueueIds: o.QueueIds,
		
		IntradayDataGroupings: o.IntradayDataGroupings,
		Alias:    (*Alias)(o),
	})
}

func (o *Wfmintradaydataupdatetopicintradaydataupdate) UnmarshalJSON(b []byte) error {
	var WfmintradaydataupdatetopicintradaydataupdateMap map[string]interface{}
	err := json.Unmarshal(b, &WfmintradaydataupdatetopicintradaydataupdateMap)
	if err != nil {
		return err
	}
	
	if startDateString, ok := WfmintradaydataupdatetopicintradaydataupdateMap["startDate"].(string); ok {
		StartDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", startDateString)
		o.StartDate = &StartDate
	}
	
	if endDateString, ok := WfmintradaydataupdatetopicintradaydataupdateMap["endDate"].(string); ok {
		EndDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", endDateString)
		o.EndDate = &EndDate
	}
	
	if IntervalLengthMinutes, ok := WfmintradaydataupdatetopicintradaydataupdateMap["intervalLengthMinutes"].(float64); ok {
		IntervalLengthMinutesInt := int(IntervalLengthMinutes)
		o.IntervalLengthMinutes = &IntervalLengthMinutesInt
	}
	
	if NumberOfIntervals, ok := WfmintradaydataupdatetopicintradaydataupdateMap["numberOfIntervals"].(float64); ok {
		NumberOfIntervalsInt := int(NumberOfIntervals)
		o.NumberOfIntervals = &NumberOfIntervalsInt
	}
	
	if Metrics, ok := WfmintradaydataupdatetopicintradaydataupdateMap["metrics"].([]interface{}); ok {
		MetricsString, _ := json.Marshal(Metrics)
		json.Unmarshal(MetricsString, &o.Metrics)
	}
	
	if QueueIds, ok := WfmintradaydataupdatetopicintradaydataupdateMap["queueIds"].([]interface{}); ok {
		QueueIdsString, _ := json.Marshal(QueueIds)
		json.Unmarshal(QueueIdsString, &o.QueueIds)
	}
	
	if IntradayDataGroupings, ok := WfmintradaydataupdatetopicintradaydataupdateMap["intradayDataGroupings"].([]interface{}); ok {
		IntradayDataGroupingsString, _ := json.Marshal(IntradayDataGroupings)
		json.Unmarshal(IntradayDataGroupingsString, &o.IntradayDataGroupings)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmintradaydataupdatetopicintradaydataupdate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Historicalshrinkageaggregateresponse
type Historicalshrinkageaggregateresponse struct { 
	// ScheduledShrinkageSeconds - Aggregated shrinkage value in seconds for scheduled activities
	ScheduledShrinkageSeconds *int `json:"scheduledShrinkageSeconds,omitempty"`


	// ScheduledShrinkagePercent - Aggregated shrinkage value in percent from 0.0 to 100.0 for scheduled activities
	ScheduledShrinkagePercent *float64 `json:"scheduledShrinkagePercent,omitempty"`


	// ActualShrinkageSeconds - Aggregated actual value in seconds for scheduled activities
	ActualShrinkageSeconds *int `json:"actualShrinkageSeconds,omitempty"`


	// ActualShrinkagePercent - Aggregated actual value in percent from 0.0 to 100.0 for scheduled activities
	ActualShrinkagePercent *float64 `json:"actualShrinkagePercent,omitempty"`


	// PaidShrinkageSeconds - Aggregated shrinkage value in seconds for paid activities
	PaidShrinkageSeconds *int `json:"paidShrinkageSeconds,omitempty"`


	// UnpaidShrinkageSeconds - Aggregated shrinkage value in seconds for unpaid activities
	UnpaidShrinkageSeconds *int `json:"unpaidShrinkageSeconds,omitempty"`


	// PlannedShrinkageSeconds - Aggregated shrinkage value in seconds for planned activities
	PlannedShrinkageSeconds *int `json:"plannedShrinkageSeconds,omitempty"`


	// UnplannedShrinkageSeconds - Aggregated shrinkage value in seconds for unplanned activities
	UnplannedShrinkageSeconds *int `json:"unplannedShrinkageSeconds,omitempty"`

}

func (o *Historicalshrinkageaggregateresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Historicalshrinkageaggregateresponse
	
	return json.Marshal(&struct { 
		ScheduledShrinkageSeconds *int `json:"scheduledShrinkageSeconds,omitempty"`
		
		ScheduledShrinkagePercent *float64 `json:"scheduledShrinkagePercent,omitempty"`
		
		ActualShrinkageSeconds *int `json:"actualShrinkageSeconds,omitempty"`
		
		ActualShrinkagePercent *float64 `json:"actualShrinkagePercent,omitempty"`
		
		PaidShrinkageSeconds *int `json:"paidShrinkageSeconds,omitempty"`
		
		UnpaidShrinkageSeconds *int `json:"unpaidShrinkageSeconds,omitempty"`
		
		PlannedShrinkageSeconds *int `json:"plannedShrinkageSeconds,omitempty"`
		
		UnplannedShrinkageSeconds *int `json:"unplannedShrinkageSeconds,omitempty"`
		*Alias
	}{ 
		ScheduledShrinkageSeconds: o.ScheduledShrinkageSeconds,
		
		ScheduledShrinkagePercent: o.ScheduledShrinkagePercent,
		
		ActualShrinkageSeconds: o.ActualShrinkageSeconds,
		
		ActualShrinkagePercent: o.ActualShrinkagePercent,
		
		PaidShrinkageSeconds: o.PaidShrinkageSeconds,
		
		UnpaidShrinkageSeconds: o.UnpaidShrinkageSeconds,
		
		PlannedShrinkageSeconds: o.PlannedShrinkageSeconds,
		
		UnplannedShrinkageSeconds: o.UnplannedShrinkageSeconds,
		Alias:    (*Alias)(o),
	})
}

func (o *Historicalshrinkageaggregateresponse) UnmarshalJSON(b []byte) error {
	var HistoricalshrinkageaggregateresponseMap map[string]interface{}
	err := json.Unmarshal(b, &HistoricalshrinkageaggregateresponseMap)
	if err != nil {
		return err
	}
	
	if ScheduledShrinkageSeconds, ok := HistoricalshrinkageaggregateresponseMap["scheduledShrinkageSeconds"].(float64); ok {
		ScheduledShrinkageSecondsInt := int(ScheduledShrinkageSeconds)
		o.ScheduledShrinkageSeconds = &ScheduledShrinkageSecondsInt
	}
	
	if ScheduledShrinkagePercent, ok := HistoricalshrinkageaggregateresponseMap["scheduledShrinkagePercent"].(float64); ok {
		o.ScheduledShrinkagePercent = &ScheduledShrinkagePercent
	}
    
	if ActualShrinkageSeconds, ok := HistoricalshrinkageaggregateresponseMap["actualShrinkageSeconds"].(float64); ok {
		ActualShrinkageSecondsInt := int(ActualShrinkageSeconds)
		o.ActualShrinkageSeconds = &ActualShrinkageSecondsInt
	}
	
	if ActualShrinkagePercent, ok := HistoricalshrinkageaggregateresponseMap["actualShrinkagePercent"].(float64); ok {
		o.ActualShrinkagePercent = &ActualShrinkagePercent
	}
    
	if PaidShrinkageSeconds, ok := HistoricalshrinkageaggregateresponseMap["paidShrinkageSeconds"].(float64); ok {
		PaidShrinkageSecondsInt := int(PaidShrinkageSeconds)
		o.PaidShrinkageSeconds = &PaidShrinkageSecondsInt
	}
	
	if UnpaidShrinkageSeconds, ok := HistoricalshrinkageaggregateresponseMap["unpaidShrinkageSeconds"].(float64); ok {
		UnpaidShrinkageSecondsInt := int(UnpaidShrinkageSeconds)
		o.UnpaidShrinkageSeconds = &UnpaidShrinkageSecondsInt
	}
	
	if PlannedShrinkageSeconds, ok := HistoricalshrinkageaggregateresponseMap["plannedShrinkageSeconds"].(float64); ok {
		PlannedShrinkageSecondsInt := int(PlannedShrinkageSeconds)
		o.PlannedShrinkageSeconds = &PlannedShrinkageSecondsInt
	}
	
	if UnplannedShrinkageSeconds, ok := HistoricalshrinkageaggregateresponseMap["unplannedShrinkageSeconds"].(float64); ok {
		UnplannedShrinkageSecondsInt := int(UnplannedShrinkageSeconds)
		o.UnplannedShrinkageSeconds = &UnplannedShrinkageSecondsInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Historicalshrinkageaggregateresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

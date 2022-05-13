package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Schedulegenerationresultsummary
type Schedulegenerationresultsummary struct { 
	// Failed - Whether the schedule generation run failed
	Failed *bool `json:"failed,omitempty"`


	// RunId - The ID of the schedule generation run. Reference this when requesting support
	RunId *string `json:"runId,omitempty"`


	// MessageCount - The number of schedule generation messages for this schedule generation run
	MessageCount *int `json:"messageCount,omitempty"`


	// MessageSeverityCounts - The list of schedule generation message counts by severity for this schedule generation run
	MessageSeverityCounts *[]Schedulermessageseveritycount `json:"messageSeverityCounts,omitempty"`

}

func (o *Schedulegenerationresultsummary) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Schedulegenerationresultsummary
	
	return json.Marshal(&struct { 
		Failed *bool `json:"failed,omitempty"`
		
		RunId *string `json:"runId,omitempty"`
		
		MessageCount *int `json:"messageCount,omitempty"`
		
		MessageSeverityCounts *[]Schedulermessageseveritycount `json:"messageSeverityCounts,omitempty"`
		*Alias
	}{ 
		Failed: o.Failed,
		
		RunId: o.RunId,
		
		MessageCount: o.MessageCount,
		
		MessageSeverityCounts: o.MessageSeverityCounts,
		Alias:    (*Alias)(o),
	})
}

func (o *Schedulegenerationresultsummary) UnmarshalJSON(b []byte) error {
	var SchedulegenerationresultsummaryMap map[string]interface{}
	err := json.Unmarshal(b, &SchedulegenerationresultsummaryMap)
	if err != nil {
		return err
	}
	
	if Failed, ok := SchedulegenerationresultsummaryMap["failed"].(bool); ok {
		o.Failed = &Failed
	}
    
	if RunId, ok := SchedulegenerationresultsummaryMap["runId"].(string); ok {
		o.RunId = &RunId
	}
    
	if MessageCount, ok := SchedulegenerationresultsummaryMap["messageCount"].(float64); ok {
		MessageCountInt := int(MessageCount)
		o.MessageCount = &MessageCountInt
	}
	
	if MessageSeverityCounts, ok := SchedulegenerationresultsummaryMap["messageSeverityCounts"].([]interface{}); ok {
		MessageSeverityCountsString, _ := json.Marshal(MessageSeverityCounts)
		json.Unmarshal(MessageSeverityCountsString, &o.MessageSeverityCounts)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Schedulegenerationresultsummary) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

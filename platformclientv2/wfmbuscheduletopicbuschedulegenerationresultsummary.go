package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmbuscheduletopicbuschedulegenerationresultsummary
type Wfmbuscheduletopicbuschedulegenerationresultsummary struct { 
	// Failed
	Failed *bool `json:"failed,omitempty"`


	// RunId
	RunId *string `json:"runId,omitempty"`


	// MessageCount
	MessageCount *int `json:"messageCount,omitempty"`


	// MessageSeverityCounts
	MessageSeverityCounts *[]Wfmbuscheduletopicschedulermessageseveritycount `json:"messageSeverityCounts,omitempty"`

}

func (o *Wfmbuscheduletopicbuschedulegenerationresultsummary) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmbuscheduletopicbuschedulegenerationresultsummary
	
	return json.Marshal(&struct { 
		Failed *bool `json:"failed,omitempty"`
		
		RunId *string `json:"runId,omitempty"`
		
		MessageCount *int `json:"messageCount,omitempty"`
		
		MessageSeverityCounts *[]Wfmbuscheduletopicschedulermessageseveritycount `json:"messageSeverityCounts,omitempty"`
		*Alias
	}{ 
		Failed: o.Failed,
		
		RunId: o.RunId,
		
		MessageCount: o.MessageCount,
		
		MessageSeverityCounts: o.MessageSeverityCounts,
		Alias:    (*Alias)(o),
	})
}

func (o *Wfmbuscheduletopicbuschedulegenerationresultsummary) UnmarshalJSON(b []byte) error {
	var WfmbuscheduletopicbuschedulegenerationresultsummaryMap map[string]interface{}
	err := json.Unmarshal(b, &WfmbuscheduletopicbuschedulegenerationresultsummaryMap)
	if err != nil {
		return err
	}
	
	if Failed, ok := WfmbuscheduletopicbuschedulegenerationresultsummaryMap["failed"].(bool); ok {
		o.Failed = &Failed
	}
	
	if RunId, ok := WfmbuscheduletopicbuschedulegenerationresultsummaryMap["runId"].(string); ok {
		o.RunId = &RunId
	}
	
	if MessageCount, ok := WfmbuscheduletopicbuschedulegenerationresultsummaryMap["messageCount"].(float64); ok {
		MessageCountInt := int(MessageCount)
		o.MessageCount = &MessageCountInt
	}
	
	if MessageSeverityCounts, ok := WfmbuscheduletopicbuschedulegenerationresultsummaryMap["messageSeverityCounts"].([]interface{}); ok {
		MessageSeverityCountsString, _ := json.Marshal(MessageSeverityCounts)
		json.Unmarshal(MessageSeverityCountsString, &o.MessageSeverityCounts)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmbuscheduletopicbuschedulegenerationresultsummary) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

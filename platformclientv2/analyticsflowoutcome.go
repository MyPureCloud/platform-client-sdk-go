package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Analyticsflowoutcome
type Analyticsflowoutcome struct { 
	// FlowOutcome - Combination of unique flow outcome identifier and its value separated by colon
	FlowOutcome *string `json:"flowOutcome,omitempty"`


	// FlowOutcomeEndTimestamp - The outcome ending timestamp in ISO 8601 format. This may be null if the outcome did not succeed.
	FlowOutcomeEndTimestamp *time.Time `json:"flowOutcomeEndTimestamp,omitempty"`


	// FlowOutcomeId - Unique identifier of a flow outcome
	FlowOutcomeId *string `json:"flowOutcomeId,omitempty"`


	// FlowOutcomeStartTimestamp - The outcome starting timestamp in ISO 8601 format
	FlowOutcomeStartTimestamp *time.Time `json:"flowOutcomeStartTimestamp,omitempty"`


	// FlowOutcomeValue - Flow outcome value, e.g. SUCCESS
	FlowOutcomeValue *string `json:"flowOutcomeValue,omitempty"`

}

func (o *Analyticsflowoutcome) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Analyticsflowoutcome
	
	FlowOutcomeEndTimestamp := new(string)
	if o.FlowOutcomeEndTimestamp != nil {
		
		*FlowOutcomeEndTimestamp = timeutil.Strftime(o.FlowOutcomeEndTimestamp, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		FlowOutcomeEndTimestamp = nil
	}
	
	FlowOutcomeStartTimestamp := new(string)
	if o.FlowOutcomeStartTimestamp != nil {
		
		*FlowOutcomeStartTimestamp = timeutil.Strftime(o.FlowOutcomeStartTimestamp, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		FlowOutcomeStartTimestamp = nil
	}
	
	return json.Marshal(&struct { 
		FlowOutcome *string `json:"flowOutcome,omitempty"`
		
		FlowOutcomeEndTimestamp *string `json:"flowOutcomeEndTimestamp,omitempty"`
		
		FlowOutcomeId *string `json:"flowOutcomeId,omitempty"`
		
		FlowOutcomeStartTimestamp *string `json:"flowOutcomeStartTimestamp,omitempty"`
		
		FlowOutcomeValue *string `json:"flowOutcomeValue,omitempty"`
		*Alias
	}{ 
		FlowOutcome: o.FlowOutcome,
		
		FlowOutcomeEndTimestamp: FlowOutcomeEndTimestamp,
		
		FlowOutcomeId: o.FlowOutcomeId,
		
		FlowOutcomeStartTimestamp: FlowOutcomeStartTimestamp,
		
		FlowOutcomeValue: o.FlowOutcomeValue,
		Alias:    (*Alias)(o),
	})
}

func (o *Analyticsflowoutcome) UnmarshalJSON(b []byte) error {
	var AnalyticsflowoutcomeMap map[string]interface{}
	err := json.Unmarshal(b, &AnalyticsflowoutcomeMap)
	if err != nil {
		return err
	}
	
	if FlowOutcome, ok := AnalyticsflowoutcomeMap["flowOutcome"].(string); ok {
		o.FlowOutcome = &FlowOutcome
	}
	
	if flowOutcomeEndTimestampString, ok := AnalyticsflowoutcomeMap["flowOutcomeEndTimestamp"].(string); ok {
		FlowOutcomeEndTimestamp, _ := time.Parse("2006-01-02T15:04:05.999999Z", flowOutcomeEndTimestampString)
		o.FlowOutcomeEndTimestamp = &FlowOutcomeEndTimestamp
	}
	
	if FlowOutcomeId, ok := AnalyticsflowoutcomeMap["flowOutcomeId"].(string); ok {
		o.FlowOutcomeId = &FlowOutcomeId
	}
	
	if flowOutcomeStartTimestampString, ok := AnalyticsflowoutcomeMap["flowOutcomeStartTimestamp"].(string); ok {
		FlowOutcomeStartTimestamp, _ := time.Parse("2006-01-02T15:04:05.999999Z", flowOutcomeStartTimestampString)
		o.FlowOutcomeStartTimestamp = &FlowOutcomeStartTimestamp
	}
	
	if FlowOutcomeValue, ok := AnalyticsflowoutcomeMap["flowOutcomeValue"].(string); ok {
		o.FlowOutcomeValue = &FlowOutcomeValue
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Analyticsflowoutcome) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

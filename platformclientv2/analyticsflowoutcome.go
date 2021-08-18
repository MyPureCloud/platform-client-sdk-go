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

func (u *Analyticsflowoutcome) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Analyticsflowoutcome

	
	FlowOutcomeEndTimestamp := new(string)
	if u.FlowOutcomeEndTimestamp != nil {
		
		*FlowOutcomeEndTimestamp = timeutil.Strftime(u.FlowOutcomeEndTimestamp, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		FlowOutcomeEndTimestamp = nil
	}
	
	FlowOutcomeStartTimestamp := new(string)
	if u.FlowOutcomeStartTimestamp != nil {
		
		*FlowOutcomeStartTimestamp = timeutil.Strftime(u.FlowOutcomeStartTimestamp, "%Y-%m-%dT%H:%M:%S.%fZ")
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
		FlowOutcome: u.FlowOutcome,
		
		FlowOutcomeEndTimestamp: FlowOutcomeEndTimestamp,
		
		FlowOutcomeId: u.FlowOutcomeId,
		
		FlowOutcomeStartTimestamp: FlowOutcomeStartTimestamp,
		
		FlowOutcomeValue: u.FlowOutcomeValue,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Analyticsflowoutcome) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

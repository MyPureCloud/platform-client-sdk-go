package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Schedulingtestingoptionsrequest
type Schedulingtestingoptionsrequest struct { 
	// FastScheduling - Whether to enable fast scheduling
	FastScheduling *bool `json:"fastScheduling,omitempty"`


	// DelayScheduling - Whether to force delayed scheduling
	DelayScheduling *bool `json:"delayScheduling,omitempty"`


	// FailScheduling - Whether to force scheduling to fail
	FailScheduling *bool `json:"failScheduling,omitempty"`


	// PopulateWarnings - Whether to populate warnings in the generated schedule
	PopulateWarnings *bool `json:"populateWarnings,omitempty"`


	// PopulateDeprecatedWarnings - Whether to populate deprecated warnings in the generated schedule
	PopulateDeprecatedWarnings *bool `json:"populateDeprecatedWarnings,omitempty"`

}

func (o *Schedulingtestingoptionsrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Schedulingtestingoptionsrequest
	
	return json.Marshal(&struct { 
		FastScheduling *bool `json:"fastScheduling,omitempty"`
		
		DelayScheduling *bool `json:"delayScheduling,omitempty"`
		
		FailScheduling *bool `json:"failScheduling,omitempty"`
		
		PopulateWarnings *bool `json:"populateWarnings,omitempty"`
		
		PopulateDeprecatedWarnings *bool `json:"populateDeprecatedWarnings,omitempty"`
		*Alias
	}{ 
		FastScheduling: o.FastScheduling,
		
		DelayScheduling: o.DelayScheduling,
		
		FailScheduling: o.FailScheduling,
		
		PopulateWarnings: o.PopulateWarnings,
		
		PopulateDeprecatedWarnings: o.PopulateDeprecatedWarnings,
		Alias:    (*Alias)(o),
	})
}

func (o *Schedulingtestingoptionsrequest) UnmarshalJSON(b []byte) error {
	var SchedulingtestingoptionsrequestMap map[string]interface{}
	err := json.Unmarshal(b, &SchedulingtestingoptionsrequestMap)
	if err != nil {
		return err
	}
	
	if FastScheduling, ok := SchedulingtestingoptionsrequestMap["fastScheduling"].(bool); ok {
		o.FastScheduling = &FastScheduling
	}
    
	if DelayScheduling, ok := SchedulingtestingoptionsrequestMap["delayScheduling"].(bool); ok {
		o.DelayScheduling = &DelayScheduling
	}
    
	if FailScheduling, ok := SchedulingtestingoptionsrequestMap["failScheduling"].(bool); ok {
		o.FailScheduling = &FailScheduling
	}
    
	if PopulateWarnings, ok := SchedulingtestingoptionsrequestMap["populateWarnings"].(bool); ok {
		o.PopulateWarnings = &PopulateWarnings
	}
    
	if PopulateDeprecatedWarnings, ok := SchedulingtestingoptionsrequestMap["populateDeprecatedWarnings"].(bool); ok {
		o.PopulateDeprecatedWarnings = &PopulateDeprecatedWarnings
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Schedulingtestingoptionsrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

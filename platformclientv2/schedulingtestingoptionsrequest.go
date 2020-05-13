package platformclientv2
import (
	"encoding/json"
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

}

// String returns a JSON representation of the model
func (o *Schedulingtestingoptionsrequest) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

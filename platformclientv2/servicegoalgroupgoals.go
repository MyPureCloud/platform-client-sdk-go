package platformclientv2
import (
	"encoding/json"
)

// Servicegoalgroupgoals - Goals defined for the service goal group
type Servicegoalgroupgoals struct { 
	// ServiceLevel - Service level targets for this service goal group
	ServiceLevel *Wfmservicelevel `json:"serviceLevel,omitempty"`


	// AverageSpeedOfAnswer - Average speed of answer targets for this service goal group
	AverageSpeedOfAnswer *Wfmaveragespeedofanswer `json:"averageSpeedOfAnswer,omitempty"`


	// AbandonRate - Abandon rate targets for this service goal group
	AbandonRate *Wfmabandonrate `json:"abandonRate,omitempty"`

}

// String returns a JSON representation of the model
func (o *Servicegoalgroupgoals) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

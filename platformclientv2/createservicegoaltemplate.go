package platformclientv2
import (
	"encoding/json"
)

// Createservicegoaltemplate
type Createservicegoaltemplate struct { 
	// Name - The name of the service goal template.
	Name *string `json:"name,omitempty"`


	// ServiceLevel - Service level targets for this service goal template
	ServiceLevel *Buservicelevel `json:"serviceLevel,omitempty"`


	// AverageSpeedOfAnswer - Average speed of answer targets for this service goal template
	AverageSpeedOfAnswer *Buaveragespeedofanswer `json:"averageSpeedOfAnswer,omitempty"`


	// AbandonRate - Abandon rate targets for this service goal template
	AbandonRate *Buabandonrate `json:"abandonRate,omitempty"`

}

// String returns a JSON representation of the model
func (o *Createservicegoaltemplate) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

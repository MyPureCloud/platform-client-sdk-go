package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
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

func (u *Createservicegoaltemplate) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Createservicegoaltemplate

	

	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		ServiceLevel *Buservicelevel `json:"serviceLevel,omitempty"`
		
		AverageSpeedOfAnswer *Buaveragespeedofanswer `json:"averageSpeedOfAnswer,omitempty"`
		
		AbandonRate *Buabandonrate `json:"abandonRate,omitempty"`
		*Alias
	}{ 
		Name: u.Name,
		
		ServiceLevel: u.ServiceLevel,
		
		AverageSpeedOfAnswer: u.AverageSpeedOfAnswer,
		
		AbandonRate: u.AbandonRate,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Createservicegoaltemplate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Forecastservicegoaltemplateresponse
type Forecastservicegoaltemplateresponse struct { 
	// ServiceLevel - The service level goal for this forecast
	ServiceLevel *Forecastservicelevelresponse `json:"serviceLevel,omitempty"`


	// AverageSpeedOfAnswer - The average speed of answer goal for this forecast
	AverageSpeedOfAnswer *Forecastaveragespeedofanswerresponse `json:"averageSpeedOfAnswer,omitempty"`


	// AbandonRate - The abandon rate goal for this forecast
	AbandonRate *Forecastabandonrateresponse `json:"abandonRate,omitempty"`

}

func (u *Forecastservicegoaltemplateresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Forecastservicegoaltemplateresponse

	

	return json.Marshal(&struct { 
		ServiceLevel *Forecastservicelevelresponse `json:"serviceLevel,omitempty"`
		
		AverageSpeedOfAnswer *Forecastaveragespeedofanswerresponse `json:"averageSpeedOfAnswer,omitempty"`
		
		AbandonRate *Forecastabandonrateresponse `json:"abandonRate,omitempty"`
		*Alias
	}{ 
		ServiceLevel: u.ServiceLevel,
		
		AverageSpeedOfAnswer: u.AverageSpeedOfAnswer,
		
		AbandonRate: u.AbandonRate,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Forecastservicegoaltemplateresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

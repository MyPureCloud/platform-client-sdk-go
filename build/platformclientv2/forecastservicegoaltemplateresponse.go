package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Forecastservicegoaltemplateresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

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

func (o *Forecastservicegoaltemplateresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Forecastservicegoaltemplateresponse
	
	return json.Marshal(&struct { 
		ServiceLevel *Forecastservicelevelresponse `json:"serviceLevel,omitempty"`
		
		AverageSpeedOfAnswer *Forecastaveragespeedofanswerresponse `json:"averageSpeedOfAnswer,omitempty"`
		
		AbandonRate *Forecastabandonrateresponse `json:"abandonRate,omitempty"`
		*Alias
	}{ 
		ServiceLevel: o.ServiceLevel,
		
		AverageSpeedOfAnswer: o.AverageSpeedOfAnswer,
		
		AbandonRate: o.AbandonRate,
		Alias:    (*Alias)(o),
	})
}

func (o *Forecastservicegoaltemplateresponse) UnmarshalJSON(b []byte) error {
	var ForecastservicegoaltemplateresponseMap map[string]interface{}
	err := json.Unmarshal(b, &ForecastservicegoaltemplateresponseMap)
	if err != nil {
		return err
	}
	
	if ServiceLevel, ok := ForecastservicegoaltemplateresponseMap["serviceLevel"].(map[string]interface{}); ok {
		ServiceLevelString, _ := json.Marshal(ServiceLevel)
		json.Unmarshal(ServiceLevelString, &o.ServiceLevel)
	}
	
	if AverageSpeedOfAnswer, ok := ForecastservicegoaltemplateresponseMap["averageSpeedOfAnswer"].(map[string]interface{}); ok {
		AverageSpeedOfAnswerString, _ := json.Marshal(AverageSpeedOfAnswer)
		json.Unmarshal(AverageSpeedOfAnswerString, &o.AverageSpeedOfAnswer)
	}
	
	if AbandonRate, ok := ForecastservicegoaltemplateresponseMap["abandonRate"].(map[string]interface{}); ok {
		AbandonRateString, _ := json.Marshal(AbandonRate)
		json.Unmarshal(AbandonRateString, &o.AbandonRate)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Forecastservicegoaltemplateresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

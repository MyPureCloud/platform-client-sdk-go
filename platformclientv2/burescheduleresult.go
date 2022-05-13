package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Burescheduleresult
type Burescheduleresult struct { 
	// GenerationResults - The generation results.  Note the result will always be delivered via the downloadUrl; however the schema is included for documentation
	GenerationResults *Schedulegenerationresult `json:"generationResults,omitempty"`


	// GenerationResultsDownloadUrl - The download URL from which to fetch the generation results for the rescheduling run
	GenerationResultsDownloadUrl *string `json:"generationResultsDownloadUrl,omitempty"`


	// HeadcountForecast - The headcount forecast.  Note the result will always be delivered via the downloadUrl; however the schema is included for documentation
	HeadcountForecast *Buheadcountforecast `json:"headcountForecast,omitempty"`


	// HeadcountForecastDownloadUrl - The download URL from which to fetch the headcount forecast for the rescheduling run
	HeadcountForecastDownloadUrl *string `json:"headcountForecastDownloadUrl,omitempty"`


	// AgentSchedules - List of download links for agent schedules produced by the rescheduling run
	AgentSchedules *[]Burescheduleagentscheduleresult `json:"agentSchedules,omitempty"`

}

func (o *Burescheduleresult) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Burescheduleresult
	
	return json.Marshal(&struct { 
		GenerationResults *Schedulegenerationresult `json:"generationResults,omitempty"`
		
		GenerationResultsDownloadUrl *string `json:"generationResultsDownloadUrl,omitempty"`
		
		HeadcountForecast *Buheadcountforecast `json:"headcountForecast,omitempty"`
		
		HeadcountForecastDownloadUrl *string `json:"headcountForecastDownloadUrl,omitempty"`
		
		AgentSchedules *[]Burescheduleagentscheduleresult `json:"agentSchedules,omitempty"`
		*Alias
	}{ 
		GenerationResults: o.GenerationResults,
		
		GenerationResultsDownloadUrl: o.GenerationResultsDownloadUrl,
		
		HeadcountForecast: o.HeadcountForecast,
		
		HeadcountForecastDownloadUrl: o.HeadcountForecastDownloadUrl,
		
		AgentSchedules: o.AgentSchedules,
		Alias:    (*Alias)(o),
	})
}

func (o *Burescheduleresult) UnmarshalJSON(b []byte) error {
	var BurescheduleresultMap map[string]interface{}
	err := json.Unmarshal(b, &BurescheduleresultMap)
	if err != nil {
		return err
	}
	
	if GenerationResults, ok := BurescheduleresultMap["generationResults"].(map[string]interface{}); ok {
		GenerationResultsString, _ := json.Marshal(GenerationResults)
		json.Unmarshal(GenerationResultsString, &o.GenerationResults)
	}
	
	if GenerationResultsDownloadUrl, ok := BurescheduleresultMap["generationResultsDownloadUrl"].(string); ok {
		o.GenerationResultsDownloadUrl = &GenerationResultsDownloadUrl
	}
    
	if HeadcountForecast, ok := BurescheduleresultMap["headcountForecast"].(map[string]interface{}); ok {
		HeadcountForecastString, _ := json.Marshal(HeadcountForecast)
		json.Unmarshal(HeadcountForecastString, &o.HeadcountForecast)
	}
	
	if HeadcountForecastDownloadUrl, ok := BurescheduleresultMap["headcountForecastDownloadUrl"].(string); ok {
		o.HeadcountForecastDownloadUrl = &HeadcountForecastDownloadUrl
	}
    
	if AgentSchedules, ok := BurescheduleresultMap["agentSchedules"].([]interface{}); ok {
		AgentSchedulesString, _ := json.Marshal(AgentSchedules)
		json.Unmarshal(AgentSchedulesString, &o.AgentSchedules)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Burescheduleresult) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

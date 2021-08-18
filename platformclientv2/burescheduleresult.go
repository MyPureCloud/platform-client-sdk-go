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

func (u *Burescheduleresult) MarshalJSON() ([]byte, error) {
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
		GenerationResults: u.GenerationResults,
		
		GenerationResultsDownloadUrl: u.GenerationResultsDownloadUrl,
		
		HeadcountForecast: u.HeadcountForecast,
		
		HeadcountForecastDownloadUrl: u.HeadcountForecastDownloadUrl,
		
		AgentSchedules: u.AgentSchedules,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Burescheduleresult) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

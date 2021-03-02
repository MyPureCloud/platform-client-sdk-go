package platformclientv2
import (
	"encoding/json"
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

// String returns a JSON representation of the model
func (o *Burescheduleresult) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

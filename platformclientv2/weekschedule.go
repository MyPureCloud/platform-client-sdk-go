package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Weekschedule - Week schedule information
type Weekschedule struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`


	// WeekDate - First day of this week schedule in yyyy-MM-dd format
	WeekDate *string `json:"weekDate,omitempty"`


	// Description - Description of the week schedule
	Description *string `json:"description,omitempty"`


	// Published - Whether the week schedule is published
	Published *bool `json:"published,omitempty"`


	// GenerationResults - Summary of the results from the schedule run
	GenerationResults *Weekschedulegenerationresult `json:"generationResults,omitempty"`


	// ShortTermForecast - Short term forecast associated with this schedule
	ShortTermForecast *Shorttermforecastreference `json:"shortTermForecast,omitempty"`


	// Metadata - Version metadata for this work plan
	Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`


	// UserSchedules - User schedules in the week
	UserSchedules *map[string]Userschedule `json:"userSchedules,omitempty"`


	// HeadcountForecast - Headcount information for the week schedule
	HeadcountForecast *Headcountforecast `json:"headcountForecast,omitempty"`


	// AgentSchedulesVersion - Version of agent schedules in the week schedule
	AgentSchedulesVersion *int `json:"agentSchedulesVersion,omitempty"`

}

func (u *Weekschedule) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Weekschedule

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		
		WeekDate *string `json:"weekDate,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		Published *bool `json:"published,omitempty"`
		
		GenerationResults *Weekschedulegenerationresult `json:"generationResults,omitempty"`
		
		ShortTermForecast *Shorttermforecastreference `json:"shortTermForecast,omitempty"`
		
		Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`
		
		UserSchedules *map[string]Userschedule `json:"userSchedules,omitempty"`
		
		HeadcountForecast *Headcountforecast `json:"headcountForecast,omitempty"`
		
		AgentSchedulesVersion *int `json:"agentSchedulesVersion,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		SelfUri: u.SelfUri,
		
		WeekDate: u.WeekDate,
		
		Description: u.Description,
		
		Published: u.Published,
		
		GenerationResults: u.GenerationResults,
		
		ShortTermForecast: u.ShortTermForecast,
		
		Metadata: u.Metadata,
		
		UserSchedules: u.UserSchedules,
		
		HeadcountForecast: u.HeadcountForecast,
		
		AgentSchedulesVersion: u.AgentSchedulesVersion,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Weekschedule) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

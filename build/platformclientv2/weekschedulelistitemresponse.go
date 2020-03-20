package platformclientv2
import (
	"encoding/json"
)

// Weekschedulelistitemresponse
type Weekschedulelistitemresponse struct { 
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

}

// String returns a JSON representation of the model
func (o *Weekschedulelistitemresponse) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

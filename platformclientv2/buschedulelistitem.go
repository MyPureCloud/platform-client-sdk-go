package platformclientv2
import (
	"time"
	"encoding/json"
)

// Buschedulelistitem
type Buschedulelistitem struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// WeekDate - The start week date for this schedule. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	WeekDate *time.Time `json:"weekDate,omitempty"`


	// WeekCount - The number of weeks spanned by this schedule
	WeekCount *int32 `json:"weekCount,omitempty"`


	// Description - The description of this schedule
	Description *string `json:"description,omitempty"`


	// Published - Whether this schedule is published
	Published *bool `json:"published,omitempty"`


	// ShortTermForecast - The forecast used for this schedule, if applicable
	ShortTermForecast *Bushorttermforecastreference `json:"shortTermForecast,omitempty"`


	// GenerationResults - Generation result summary for this schedule, if applicable
	GenerationResults *Schedulegenerationresultsummary `json:"generationResults,omitempty"`


	// Metadata - Version metadata for this schedule
	Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Buschedulelistitem) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

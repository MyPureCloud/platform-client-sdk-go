package platformclientv2
import (
	"time"
	"encoding/json"
)

// Bushorttermforecast
type Bushorttermforecast struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// WeekDate - The start week date of this forecast in yyyy-MM-dd.  Must fall on the start day of week for the associated business unit. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	WeekDate *time.Time `json:"weekDate,omitempty"`


	// WeekCount - The number of weeks this forecast covers
	WeekCount *int `json:"weekCount,omitempty"`


	// CreationMethod - The method by which this forecast was created
	CreationMethod *string `json:"creationMethod,omitempty"`


	// Description - The description of this forecast
	Description *string `json:"description,omitempty"`


	// Legacy - Whether this forecast contains modifications on legacy metrics
	Legacy *bool `json:"legacy,omitempty"`


	// Metadata - Metadata for this forecast
	Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`


	// ReferenceStartDate - The reference start date for interval-based data for this forecast. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ReferenceStartDate *time.Time `json:"referenceStartDate,omitempty"`


	// SourceDays - The source day pointers for this forecast
	SourceDays *[]Forecastsourcedaypointer `json:"sourceDays,omitempty"`


	// Modifications - Any manual modifications applied to this forecast
	Modifications *[]Buforecastmodification `json:"modifications,omitempty"`


	// GenerationResults - Generation result metadata
	GenerationResults *Buforecastgenerationresult `json:"generationResults,omitempty"`


	// TimeZone - The time zone for this forecast
	TimeZone *string `json:"timeZone,omitempty"`


	// PlanningGroupsVersion - The version of the planning groups that was used for this forecast
	PlanningGroupsVersion *int `json:"planningGroupsVersion,omitempty"`


	// PlanningGroups - A snapshot of the planning groups used for this forecast as of the version number indicated
	PlanningGroups *Forecastplanninggroupsresponse `json:"planningGroups,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Bushorttermforecast) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

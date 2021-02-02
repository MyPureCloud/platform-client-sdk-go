package platformclientv2
import (
	"time"
	"encoding/json"
)

// Wfmbushorttermforecastupdatecompletetopicbushorttermforecast
type Wfmbushorttermforecastupdatecompletetopicbushorttermforecast struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// WeekDate
	WeekDate *string `json:"weekDate,omitempty"`


	// CreationMethod
	CreationMethod *string `json:"creationMethod,omitempty"`


	// Description
	Description *string `json:"description,omitempty"`


	// Legacy
	Legacy *bool `json:"legacy,omitempty"`


	// ReferenceStartDate
	ReferenceStartDate *time.Time `json:"referenceStartDate,omitempty"`


	// SourceDays
	SourceDays *[]Wfmbushorttermforecastupdatecompletetopicforecastsourcedaypointer `json:"sourceDays,omitempty"`


	// Modifications
	Modifications *[]Wfmbushorttermforecastupdatecompletetopicbuforecastmodification `json:"modifications,omitempty"`


	// TimeZone
	TimeZone *string `json:"timeZone,omitempty"`


	// PlanningGroupsVersion
	PlanningGroupsVersion *int `json:"planningGroupsVersion,omitempty"`


	// WeekCount
	WeekCount *int `json:"weekCount,omitempty"`


	// Metadata
	Metadata *Wfmbushorttermforecastupdatecompletetopicwfmversionedentitymetadata `json:"metadata,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wfmbushorttermforecastupdatecompletetopicbushorttermforecast) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

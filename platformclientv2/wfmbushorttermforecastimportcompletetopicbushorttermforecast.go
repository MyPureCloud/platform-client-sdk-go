package platformclientv2
import (
	"time"
	"encoding/json"
)

// Wfmbushorttermforecastimportcompletetopicbushorttermforecast
type Wfmbushorttermforecastimportcompletetopicbushorttermforecast struct { 
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
	SourceDays *[]Wfmbushorttermforecastimportcompletetopicforecastsourcedaypointer `json:"sourceDays,omitempty"`


	// Modifications
	Modifications *[]Wfmbushorttermforecastimportcompletetopicbuforecastmodification `json:"modifications,omitempty"`


	// TimeZone
	TimeZone *string `json:"timeZone,omitempty"`


	// PlanningGroupsVersion
	PlanningGroupsVersion *int32 `json:"planningGroupsVersion,omitempty"`


	// WeekCount
	WeekCount *int32 `json:"weekCount,omitempty"`


	// Metadata
	Metadata *Wfmbushorttermforecastimportcompletetopicwfmversionedentitymetadata `json:"metadata,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wfmbushorttermforecastimportcompletetopicbushorttermforecast) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

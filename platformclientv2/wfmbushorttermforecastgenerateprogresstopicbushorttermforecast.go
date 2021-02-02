package platformclientv2
import (
	"time"
	"encoding/json"
)

// Wfmbushorttermforecastgenerateprogresstopicbushorttermforecast
type Wfmbushorttermforecastgenerateprogresstopicbushorttermforecast struct { 
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
	SourceDays *[]Wfmbushorttermforecastgenerateprogresstopicforecastsourcedaypointer `json:"sourceDays,omitempty"`


	// Modifications
	Modifications *[]Wfmbushorttermforecastgenerateprogresstopicbuforecastmodification `json:"modifications,omitempty"`


	// TimeZone
	TimeZone *string `json:"timeZone,omitempty"`


	// PlanningGroupsVersion
	PlanningGroupsVersion *int `json:"planningGroupsVersion,omitempty"`


	// WeekCount
	WeekCount *int `json:"weekCount,omitempty"`


	// Metadata
	Metadata *Wfmbushorttermforecastgenerateprogresstopicwfmversionedentitymetadata `json:"metadata,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wfmbushorttermforecastgenerateprogresstopicbushorttermforecast) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

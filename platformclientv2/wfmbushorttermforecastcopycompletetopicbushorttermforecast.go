package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmbushorttermforecastcopycompletetopicbushorttermforecast
type Wfmbushorttermforecastcopycompletetopicbushorttermforecast struct { 
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
	SourceDays *[]Wfmbushorttermforecastcopycompletetopicforecastsourcedaypointer `json:"sourceDays,omitempty"`


	// Modifications
	Modifications *[]Wfmbushorttermforecastcopycompletetopicbuforecastmodification `json:"modifications,omitempty"`


	// TimeZone
	TimeZone *string `json:"timeZone,omitempty"`


	// PlanningGroupsVersion
	PlanningGroupsVersion *int `json:"planningGroupsVersion,omitempty"`


	// WeekCount
	WeekCount *int `json:"weekCount,omitempty"`


	// Metadata
	Metadata *Wfmbushorttermforecastcopycompletetopicwfmversionedentitymetadata `json:"metadata,omitempty"`


	// CanUseForScheduling
	CanUseForScheduling *bool `json:"canUseForScheduling,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wfmbushorttermforecastcopycompletetopicbushorttermforecast) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

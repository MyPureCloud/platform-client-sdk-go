package platformclientv2
import (
	"time"
	"encoding/json"
)

// Wfmhistoricaladherencequeryforusers
type Wfmhistoricaladherencequeryforusers struct { 
	// StartDate - Beginning of the date range to query in ISO-8601 format
	StartDate *time.Time `json:"startDate,omitempty"`


	// EndDate - End of the date range to query in ISO-8601 format. If it is not set, end date will be set to current time
	EndDate *time.Time `json:"endDate,omitempty"`


	// TimeZone - The time zone to use for returned results in olson format. If it is not set, the business unit time zone will be used to compute adherence
	TimeZone *string `json:"timeZone,omitempty"`


	// UserIds - The userIds to report on
	UserIds *[]string `json:"userIds,omitempty"`


	// IncludeExceptions - Whether user exceptions should be returned as part of the results
	IncludeExceptions *bool `json:"includeExceptions,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wfmhistoricaladherencequeryforusers) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

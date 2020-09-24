package platformclientv2
import (
	"time"
	"encoding/json"
)

// Wfmhistoricaladherencequery
type Wfmhistoricaladherencequery struct { 
	// StartDate - Beginning of the date range to query in ISO-8601 format
	StartDate *time.Time `json:"startDate,omitempty"`


	// EndDate - End of the date range to query in ISO-8601 format. If it is not set, end date will be set to current time
	EndDate *time.Time `json:"endDate,omitempty"`


	// TimeZone - The time zone to use for returned results in olson format. If it is not set, the business unit time zone will be used to compute adherence
	TimeZone *string `json:"timeZone,omitempty"`


	// UserIds - The userIds to report on. If null or not set, adherence will be computed for all the users in management unit or requested teamIds. Note: Only one of [teamIds, userIds] can be requested
	UserIds *[]string `json:"userIds,omitempty"`


	// TeamIds - The teamIds to report on. If null or not set, adherence will be computed for requested users if applicable or otherwise all users in the management unit. Note: Only one of [teamIds, userIds] can be requested
	TeamIds *[]string `json:"teamIds,omitempty"`


	// IncludeExceptions - Whether user exceptions should be returned as part of the results
	IncludeExceptions *bool `json:"includeExceptions,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wfmhistoricaladherencequery) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

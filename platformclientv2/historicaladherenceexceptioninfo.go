package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Historicaladherenceexceptioninfo
type Historicaladherenceexceptioninfo struct { 
	// StartOffsetSeconds - Exception start offset in seconds relative to query start time
	StartOffsetSeconds *int `json:"startOffsetSeconds,omitempty"`


	// EndOffsetSeconds - Exception end offset in seconds relative to query start time
	EndOffsetSeconds *int `json:"endOffsetSeconds,omitempty"`


	// ScheduledActivityCodeId - The ID of the scheduled activity for this user
	ScheduledActivityCodeId *string `json:"scheduledActivityCodeId,omitempty"`


	// ScheduledActivityCategory - Activity for which the user is scheduled
	ScheduledActivityCategory *string `json:"scheduledActivityCategory,omitempty"`


	// ActualActivityCategory - Activity for which the user is actually engaged
	ActualActivityCategory *string `json:"actualActivityCategory,omitempty"`


	// SystemPresence - Actual underlying system presence value
	SystemPresence *string `json:"systemPresence,omitempty"`


	// RoutingStatus - Actual underlying routing status, used to determine whether a user is actually in adherence when OnQueue
	RoutingStatus *Routingstatus `json:"routingStatus,omitempty"`


	// Impact - The impact of the current adherence state for this user
	Impact *string `json:"impact,omitempty"`


	// SecondaryPresenceLookupId - The lookup ID used to retrieve secondary status from map of lookup ID to corresponding secondary presence ID
	SecondaryPresenceLookupId *string `json:"secondaryPresenceLookupId,omitempty"`

}

// String returns a JSON representation of the model
func (o *Historicaladherenceexceptioninfo) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

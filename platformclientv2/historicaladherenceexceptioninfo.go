package platformclientv2
import (
	"github.com/leekchan/timeutil"
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


	// ScheduledActivityCodeId - The ID of the scheduled activity code for this user
	ScheduledActivityCodeId *string `json:"scheduledActivityCodeId,omitempty"`


	// ScheduledActivityCategory - Activity for which the user is scheduled
	ScheduledActivityCategory *string `json:"scheduledActivityCategory,omitempty"`


	// ScheduledSecondaryPresenceLookupIds - The lookup IDs used to retrieve the scheduled secondary statuses from map of lookup ID to corresponding secondary presence ID
	ScheduledSecondaryPresenceLookupIds *[]string `json:"scheduledSecondaryPresenceLookupIds,omitempty"`


	// ActualActivityCodeId - The ID of the actual activity code for this user
	ActualActivityCodeId *string `json:"actualActivityCodeId,omitempty"`


	// ActualActivityCategory - Activity for which the user is actually engaged
	ActualActivityCategory *string `json:"actualActivityCategory,omitempty"`


	// SystemPresence - Actual underlying system presence value
	SystemPresence *string `json:"systemPresence,omitempty"`


	// RoutingStatus - Actual underlying routing status, used to determine whether a user is actually in adherence when OnQueue
	RoutingStatus *string `json:"routingStatus,omitempty"`


	// Impact - The impact of the current adherence state for this user
	Impact *string `json:"impact,omitempty"`


	// SecondaryPresenceLookupId - The lookup ID used to retrieve the actual secondary status from map of lookup ID to corresponding secondary presence ID
	SecondaryPresenceLookupId *string `json:"secondaryPresenceLookupId,omitempty"`

}

func (o *Historicaladherenceexceptioninfo) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Historicaladherenceexceptioninfo
	
	return json.Marshal(&struct { 
		StartOffsetSeconds *int `json:"startOffsetSeconds,omitempty"`
		
		EndOffsetSeconds *int `json:"endOffsetSeconds,omitempty"`
		
		ScheduledActivityCodeId *string `json:"scheduledActivityCodeId,omitempty"`
		
		ScheduledActivityCategory *string `json:"scheduledActivityCategory,omitempty"`
		
		ScheduledSecondaryPresenceLookupIds *[]string `json:"scheduledSecondaryPresenceLookupIds,omitempty"`
		
		ActualActivityCodeId *string `json:"actualActivityCodeId,omitempty"`
		
		ActualActivityCategory *string `json:"actualActivityCategory,omitempty"`
		
		SystemPresence *string `json:"systemPresence,omitempty"`
		
		RoutingStatus *string `json:"routingStatus,omitempty"`
		
		Impact *string `json:"impact,omitempty"`
		
		SecondaryPresenceLookupId *string `json:"secondaryPresenceLookupId,omitempty"`
		*Alias
	}{ 
		StartOffsetSeconds: o.StartOffsetSeconds,
		
		EndOffsetSeconds: o.EndOffsetSeconds,
		
		ScheduledActivityCodeId: o.ScheduledActivityCodeId,
		
		ScheduledActivityCategory: o.ScheduledActivityCategory,
		
		ScheduledSecondaryPresenceLookupIds: o.ScheduledSecondaryPresenceLookupIds,
		
		ActualActivityCodeId: o.ActualActivityCodeId,
		
		ActualActivityCategory: o.ActualActivityCategory,
		
		SystemPresence: o.SystemPresence,
		
		RoutingStatus: o.RoutingStatus,
		
		Impact: o.Impact,
		
		SecondaryPresenceLookupId: o.SecondaryPresenceLookupId,
		Alias:    (*Alias)(o),
	})
}

func (o *Historicaladherenceexceptioninfo) UnmarshalJSON(b []byte) error {
	var HistoricaladherenceexceptioninfoMap map[string]interface{}
	err := json.Unmarshal(b, &HistoricaladherenceexceptioninfoMap)
	if err != nil {
		return err
	}
	
	if StartOffsetSeconds, ok := HistoricaladherenceexceptioninfoMap["startOffsetSeconds"].(float64); ok {
		StartOffsetSecondsInt := int(StartOffsetSeconds)
		o.StartOffsetSeconds = &StartOffsetSecondsInt
	}
	
	if EndOffsetSeconds, ok := HistoricaladherenceexceptioninfoMap["endOffsetSeconds"].(float64); ok {
		EndOffsetSecondsInt := int(EndOffsetSeconds)
		o.EndOffsetSeconds = &EndOffsetSecondsInt
	}
	
	if ScheduledActivityCodeId, ok := HistoricaladherenceexceptioninfoMap["scheduledActivityCodeId"].(string); ok {
		o.ScheduledActivityCodeId = &ScheduledActivityCodeId
	}
    
	if ScheduledActivityCategory, ok := HistoricaladherenceexceptioninfoMap["scheduledActivityCategory"].(string); ok {
		o.ScheduledActivityCategory = &ScheduledActivityCategory
	}
    
	if ScheduledSecondaryPresenceLookupIds, ok := HistoricaladherenceexceptioninfoMap["scheduledSecondaryPresenceLookupIds"].([]interface{}); ok {
		ScheduledSecondaryPresenceLookupIdsString, _ := json.Marshal(ScheduledSecondaryPresenceLookupIds)
		json.Unmarshal(ScheduledSecondaryPresenceLookupIdsString, &o.ScheduledSecondaryPresenceLookupIds)
	}
	
	if ActualActivityCodeId, ok := HistoricaladherenceexceptioninfoMap["actualActivityCodeId"].(string); ok {
		o.ActualActivityCodeId = &ActualActivityCodeId
	}
    
	if ActualActivityCategory, ok := HistoricaladherenceexceptioninfoMap["actualActivityCategory"].(string); ok {
		o.ActualActivityCategory = &ActualActivityCategory
	}
    
	if SystemPresence, ok := HistoricaladherenceexceptioninfoMap["systemPresence"].(string); ok {
		o.SystemPresence = &SystemPresence
	}
    
	if RoutingStatus, ok := HistoricaladherenceexceptioninfoMap["routingStatus"].(string); ok {
		o.RoutingStatus = &RoutingStatus
	}
    
	if Impact, ok := HistoricaladherenceexceptioninfoMap["impact"].(string); ok {
		o.Impact = &Impact
	}
    
	if SecondaryPresenceLookupId, ok := HistoricaladherenceexceptioninfoMap["secondaryPresenceLookupId"].(string); ok {
		o.SecondaryPresenceLookupId = &SecondaryPresenceLookupId
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Historicaladherenceexceptioninfo) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

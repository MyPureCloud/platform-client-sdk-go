package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Campaignstats
type Campaignstats struct { 
	// ContactRate - Information regarding the campaign's connect rate
	ContactRate *Connectrate `json:"contactRate,omitempty"`


	// IdleAgents - Number of available agents not currently being utilized
	IdleAgents *int `json:"idleAgents,omitempty"`


	// EffectiveIdleAgents - Number of effective available agents not currently being utilized
	EffectiveIdleAgents *float64 `json:"effectiveIdleAgents,omitempty"`


	// AdjustedCallsPerAgent - Calls per agent adjusted by pace
	AdjustedCallsPerAgent *float64 `json:"adjustedCallsPerAgent,omitempty"`


	// OutstandingCalls - Number of campaign calls currently ongoing
	OutstandingCalls *int `json:"outstandingCalls,omitempty"`


	// ScheduledCalls - Number of campaign calls currently scheduled
	ScheduledCalls *int `json:"scheduledCalls,omitempty"`

}

// String returns a JSON representation of the model
func (o *Campaignstats) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

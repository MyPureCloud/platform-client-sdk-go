package platformclientv2
import (
	"github.com/leekchan/timeutil"
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

func (u *Campaignstats) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Campaignstats

	

	return json.Marshal(&struct { 
		ContactRate *Connectrate `json:"contactRate,omitempty"`
		
		IdleAgents *int `json:"idleAgents,omitempty"`
		
		EffectiveIdleAgents *float64 `json:"effectiveIdleAgents,omitempty"`
		
		AdjustedCallsPerAgent *float64 `json:"adjustedCallsPerAgent,omitempty"`
		
		OutstandingCalls *int `json:"outstandingCalls,omitempty"`
		
		ScheduledCalls *int `json:"scheduledCalls,omitempty"`
		*Alias
	}{ 
		ContactRate: u.ContactRate,
		
		IdleAgents: u.IdleAgents,
		
		EffectiveIdleAgents: u.EffectiveIdleAgents,
		
		AdjustedCallsPerAgent: u.AdjustedCallsPerAgent,
		
		OutstandingCalls: u.OutstandingCalls,
		
		ScheduledCalls: u.ScheduledCalls,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Campaignstats) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

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

func (o *Campaignstats) MarshalJSON() ([]byte, error) {
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
		ContactRate: o.ContactRate,
		
		IdleAgents: o.IdleAgents,
		
		EffectiveIdleAgents: o.EffectiveIdleAgents,
		
		AdjustedCallsPerAgent: o.AdjustedCallsPerAgent,
		
		OutstandingCalls: o.OutstandingCalls,
		
		ScheduledCalls: o.ScheduledCalls,
		Alias:    (*Alias)(o),
	})
}

func (o *Campaignstats) UnmarshalJSON(b []byte) error {
	var CampaignstatsMap map[string]interface{}
	err := json.Unmarshal(b, &CampaignstatsMap)
	if err != nil {
		return err
	}
	
	if ContactRate, ok := CampaignstatsMap["contactRate"].(map[string]interface{}); ok {
		ContactRateString, _ := json.Marshal(ContactRate)
		json.Unmarshal(ContactRateString, &o.ContactRate)
	}
	
	if IdleAgents, ok := CampaignstatsMap["idleAgents"].(float64); ok {
		IdleAgentsInt := int(IdleAgents)
		o.IdleAgents = &IdleAgentsInt
	}
	
	if EffectiveIdleAgents, ok := CampaignstatsMap["effectiveIdleAgents"].(float64); ok {
		o.EffectiveIdleAgents = &EffectiveIdleAgents
	}
    
	if AdjustedCallsPerAgent, ok := CampaignstatsMap["adjustedCallsPerAgent"].(float64); ok {
		o.AdjustedCallsPerAgent = &AdjustedCallsPerAgent
	}
    
	if OutstandingCalls, ok := CampaignstatsMap["outstandingCalls"].(float64); ok {
		OutstandingCallsInt := int(OutstandingCalls)
		o.OutstandingCalls = &OutstandingCallsInt
	}
	
	if ScheduledCalls, ok := CampaignstatsMap["scheduledCalls"].(float64); ok {
		ScheduledCallsInt := int(ScheduledCalls)
		o.ScheduledCalls = &ScheduledCallsInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Campaignstats) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

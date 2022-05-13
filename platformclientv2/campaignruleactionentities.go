package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Campaignruleactionentities
type Campaignruleactionentities struct { 
	// Campaigns - The list of campaigns for a CampaignRule to monitor. Required if the CampaignRule has any conditions that run on a campaign.
	Campaigns *[]Domainentityref `json:"campaigns,omitempty"`


	// Sequences - The list of sequences for a CampaignRule to monitor. Required if the CampaignRule has any conditions that run on a sequence.
	Sequences *[]Domainentityref `json:"sequences,omitempty"`


	// UseTriggeringEntity - If true, the CampaignRuleAction will apply to the same entity that triggered the CampaignRuleCondition.
	UseTriggeringEntity *bool `json:"useTriggeringEntity,omitempty"`

}

func (o *Campaignruleactionentities) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Campaignruleactionentities
	
	return json.Marshal(&struct { 
		Campaigns *[]Domainentityref `json:"campaigns,omitempty"`
		
		Sequences *[]Domainentityref `json:"sequences,omitempty"`
		
		UseTriggeringEntity *bool `json:"useTriggeringEntity,omitempty"`
		*Alias
	}{ 
		Campaigns: o.Campaigns,
		
		Sequences: o.Sequences,
		
		UseTriggeringEntity: o.UseTriggeringEntity,
		Alias:    (*Alias)(o),
	})
}

func (o *Campaignruleactionentities) UnmarshalJSON(b []byte) error {
	var CampaignruleactionentitiesMap map[string]interface{}
	err := json.Unmarshal(b, &CampaignruleactionentitiesMap)
	if err != nil {
		return err
	}
	
	if Campaigns, ok := CampaignruleactionentitiesMap["campaigns"].([]interface{}); ok {
		CampaignsString, _ := json.Marshal(Campaigns)
		json.Unmarshal(CampaignsString, &o.Campaigns)
	}
	
	if Sequences, ok := CampaignruleactionentitiesMap["sequences"].([]interface{}); ok {
		SequencesString, _ := json.Marshal(Sequences)
		json.Unmarshal(SequencesString, &o.Sequences)
	}
	
	if UseTriggeringEntity, ok := CampaignruleactionentitiesMap["useTriggeringEntity"].(bool); ok {
		o.UseTriggeringEntity = &UseTriggeringEntity
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Campaignruleactionentities) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

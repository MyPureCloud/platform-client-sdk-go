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

func (u *Campaignruleactionentities) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Campaignruleactionentities

	

	return json.Marshal(&struct { 
		Campaigns *[]Domainentityref `json:"campaigns,omitempty"`
		
		Sequences *[]Domainentityref `json:"sequences,omitempty"`
		
		UseTriggeringEntity *bool `json:"useTriggeringEntity,omitempty"`
		*Alias
	}{ 
		Campaigns: u.Campaigns,
		
		Sequences: u.Sequences,
		
		UseTriggeringEntity: u.UseTriggeringEntity,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Campaignruleactionentities) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

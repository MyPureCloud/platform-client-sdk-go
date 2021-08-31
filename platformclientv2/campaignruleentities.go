package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Campaignruleentities
type Campaignruleentities struct { 
	// Campaigns - The list of campaigns for a CampaignRule to monitor. Required if the CampaignRule has any conditions that run on a campaign.
	Campaigns *[]Domainentityref `json:"campaigns,omitempty"`


	// Sequences - The list of sequences for a CampaignRule to monitor. Required if the CampaignRule has any conditions that run on a sequence.
	Sequences *[]Domainentityref `json:"sequences,omitempty"`

}

func (o *Campaignruleentities) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Campaignruleentities
	
	return json.Marshal(&struct { 
		Campaigns *[]Domainentityref `json:"campaigns,omitempty"`
		
		Sequences *[]Domainentityref `json:"sequences,omitempty"`
		*Alias
	}{ 
		Campaigns: o.Campaigns,
		
		Sequences: o.Sequences,
		Alias:    (*Alias)(o),
	})
}

func (o *Campaignruleentities) UnmarshalJSON(b []byte) error {
	var CampaignruleentitiesMap map[string]interface{}
	err := json.Unmarshal(b, &CampaignruleentitiesMap)
	if err != nil {
		return err
	}
	
	if Campaigns, ok := CampaignruleentitiesMap["campaigns"].([]interface{}); ok {
		CampaignsString, _ := json.Marshal(Campaigns)
		json.Unmarshal(CampaignsString, &o.Campaigns)
	}
	
	if Sequences, ok := CampaignruleentitiesMap["sequences"].([]interface{}); ok {
		SequencesString, _ := json.Marshal(Sequences)
		json.Unmarshal(SequencesString, &o.Sequences)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Campaignruleentities) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

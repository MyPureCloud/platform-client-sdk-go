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

func (u *Campaignruleentities) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Campaignruleentities

	

	return json.Marshal(&struct { 
		Campaigns *[]Domainentityref `json:"campaigns,omitempty"`
		
		Sequences *[]Domainentityref `json:"sequences,omitempty"`
		*Alias
	}{ 
		Campaigns: u.Campaigns,
		
		Sequences: u.Sequences,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Campaignruleentities) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

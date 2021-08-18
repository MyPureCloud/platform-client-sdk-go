package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Dialercampaignruleconfigchangecampaignruleentities
type Dialercampaignruleconfigchangecampaignruleentities struct { 
	// Campaigns
	Campaigns *[]Dialercampaignruleconfigchangeurireference `json:"campaigns,omitempty"`


	// Sequences
	Sequences *[]Dialercampaignruleconfigchangeurireference `json:"sequences,omitempty"`

}

func (u *Dialercampaignruleconfigchangecampaignruleentities) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dialercampaignruleconfigchangecampaignruleentities

	

	return json.Marshal(&struct { 
		Campaigns *[]Dialercampaignruleconfigchangeurireference `json:"campaigns,omitempty"`
		
		Sequences *[]Dialercampaignruleconfigchangeurireference `json:"sequences,omitempty"`
		*Alias
	}{ 
		Campaigns: u.Campaigns,
		
		Sequences: u.Sequences,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Dialercampaignruleconfigchangecampaignruleentities) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

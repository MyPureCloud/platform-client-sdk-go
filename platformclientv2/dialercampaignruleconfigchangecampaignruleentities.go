package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Dialercampaignruleconfigchangecampaignruleentities) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

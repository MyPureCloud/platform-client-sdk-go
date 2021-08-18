package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Dialercampaignruleconfigchangecampaignruleactionentities
type Dialercampaignruleconfigchangecampaignruleactionentities struct { 
	// Campaigns
	Campaigns *[]Dialercampaignruleconfigchangeurireference `json:"campaigns,omitempty"`


	// Sequences
	Sequences *[]Dialercampaignruleconfigchangeurireference `json:"sequences,omitempty"`


	// UseTriggeringEntity
	UseTriggeringEntity *bool `json:"useTriggeringEntity,omitempty"`


	// AdditionalProperties
	AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`

}

func (u *Dialercampaignruleconfigchangecampaignruleactionentities) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dialercampaignruleconfigchangecampaignruleactionentities

	

	return json.Marshal(&struct { 
		Campaigns *[]Dialercampaignruleconfigchangeurireference `json:"campaigns,omitempty"`
		
		Sequences *[]Dialercampaignruleconfigchangeurireference `json:"sequences,omitempty"`
		
		UseTriggeringEntity *bool `json:"useTriggeringEntity,omitempty"`
		
		AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`
		*Alias
	}{ 
		Campaigns: u.Campaigns,
		
		Sequences: u.Sequences,
		
		UseTriggeringEntity: u.UseTriggeringEntity,
		
		AdditionalProperties: u.AdditionalProperties,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Dialercampaignruleconfigchangecampaignruleactionentities) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

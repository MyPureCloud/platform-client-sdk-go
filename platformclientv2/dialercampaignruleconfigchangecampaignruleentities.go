package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Dialercampaignruleconfigchangecampaignruleentities
type Dialercampaignruleconfigchangecampaignruleentities struct { 
	// Campaigns - A list of campaignIds to act on
	Campaigns *[]Dialercampaignruleconfigchangeurireference `json:"campaigns,omitempty"`


	// Sequences - A list of sequenceIds to act on
	Sequences *[]Dialercampaignruleconfigchangeurireference `json:"sequences,omitempty"`

}

func (o *Dialercampaignruleconfigchangecampaignruleentities) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dialercampaignruleconfigchangecampaignruleentities
	
	return json.Marshal(&struct { 
		Campaigns *[]Dialercampaignruleconfigchangeurireference `json:"campaigns,omitempty"`
		
		Sequences *[]Dialercampaignruleconfigchangeurireference `json:"sequences,omitempty"`
		*Alias
	}{ 
		Campaigns: o.Campaigns,
		
		Sequences: o.Sequences,
		Alias:    (*Alias)(o),
	})
}

func (o *Dialercampaignruleconfigchangecampaignruleentities) UnmarshalJSON(b []byte) error {
	var DialercampaignruleconfigchangecampaignruleentitiesMap map[string]interface{}
	err := json.Unmarshal(b, &DialercampaignruleconfigchangecampaignruleentitiesMap)
	if err != nil {
		return err
	}
	
	if Campaigns, ok := DialercampaignruleconfigchangecampaignruleentitiesMap["campaigns"].([]interface{}); ok {
		CampaignsString, _ := json.Marshal(Campaigns)
		json.Unmarshal(CampaignsString, &o.Campaigns)
	}
	
	if Sequences, ok := DialercampaignruleconfigchangecampaignruleentitiesMap["sequences"].([]interface{}); ok {
		SequencesString, _ := json.Marshal(Sequences)
		json.Unmarshal(SequencesString, &o.Sequences)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Dialercampaignruleconfigchangecampaignruleentities) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

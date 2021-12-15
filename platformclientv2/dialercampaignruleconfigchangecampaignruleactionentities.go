package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Dialercampaignruleconfigchangecampaignruleactionentities - the campaign/sequence entities that this action acts on
type Dialercampaignruleconfigchangecampaignruleactionentities struct { 
	// UseTriggeringEntity - Whether this action should act on the entity that triggered it
	UseTriggeringEntity *bool `json:"useTriggeringEntity,omitempty"`


	// Campaigns - A list of campaignIds to act on
	Campaigns *[]Dialercampaignruleconfigchangeurireference `json:"campaigns,omitempty"`


	// Sequences - A list of sequenceIds to act on
	Sequences *[]Dialercampaignruleconfigchangeurireference `json:"sequences,omitempty"`

}

func (o *Dialercampaignruleconfigchangecampaignruleactionentities) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dialercampaignruleconfigchangecampaignruleactionentities
	
	return json.Marshal(&struct { 
		UseTriggeringEntity *bool `json:"useTriggeringEntity,omitempty"`
		
		Campaigns *[]Dialercampaignruleconfigchangeurireference `json:"campaigns,omitempty"`
		
		Sequences *[]Dialercampaignruleconfigchangeurireference `json:"sequences,omitempty"`
		*Alias
	}{ 
		UseTriggeringEntity: o.UseTriggeringEntity,
		
		Campaigns: o.Campaigns,
		
		Sequences: o.Sequences,
		Alias:    (*Alias)(o),
	})
}

func (o *Dialercampaignruleconfigchangecampaignruleactionentities) UnmarshalJSON(b []byte) error {
	var DialercampaignruleconfigchangecampaignruleactionentitiesMap map[string]interface{}
	err := json.Unmarshal(b, &DialercampaignruleconfigchangecampaignruleactionentitiesMap)
	if err != nil {
		return err
	}
	
	if UseTriggeringEntity, ok := DialercampaignruleconfigchangecampaignruleactionentitiesMap["useTriggeringEntity"].(bool); ok {
		o.UseTriggeringEntity = &UseTriggeringEntity
	}
	
	if Campaigns, ok := DialercampaignruleconfigchangecampaignruleactionentitiesMap["campaigns"].([]interface{}); ok {
		CampaignsString, _ := json.Marshal(Campaigns)
		json.Unmarshal(CampaignsString, &o.Campaigns)
	}
	
	if Sequences, ok := DialercampaignruleconfigchangecampaignruleactionentitiesMap["sequences"].([]interface{}); ok {
		SequencesString, _ := json.Marshal(Sequences)
		json.Unmarshal(SequencesString, &o.Sequences)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Dialercampaignruleconfigchangecampaignruleactionentities) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

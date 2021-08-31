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

func (o *Dialercampaignruleconfigchangecampaignruleactionentities) MarshalJSON() ([]byte, error) {
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
		Campaigns: o.Campaigns,
		
		Sequences: o.Sequences,
		
		UseTriggeringEntity: o.UseTriggeringEntity,
		
		AdditionalProperties: o.AdditionalProperties,
		Alias:    (*Alias)(o),
	})
}

func (o *Dialercampaignruleconfigchangecampaignruleactionentities) UnmarshalJSON(b []byte) error {
	var DialercampaignruleconfigchangecampaignruleactionentitiesMap map[string]interface{}
	err := json.Unmarshal(b, &DialercampaignruleconfigchangecampaignruleactionentitiesMap)
	if err != nil {
		return err
	}
	
	if Campaigns, ok := DialercampaignruleconfigchangecampaignruleactionentitiesMap["campaigns"].([]interface{}); ok {
		CampaignsString, _ := json.Marshal(Campaigns)
		json.Unmarshal(CampaignsString, &o.Campaigns)
	}
	
	if Sequences, ok := DialercampaignruleconfigchangecampaignruleactionentitiesMap["sequences"].([]interface{}); ok {
		SequencesString, _ := json.Marshal(Sequences)
		json.Unmarshal(SequencesString, &o.Sequences)
	}
	
	if UseTriggeringEntity, ok := DialercampaignruleconfigchangecampaignruleactionentitiesMap["useTriggeringEntity"].(bool); ok {
		o.UseTriggeringEntity = &UseTriggeringEntity
	}
	
	if AdditionalProperties, ok := DialercampaignruleconfigchangecampaignruleactionentitiesMap["additionalProperties"].(map[string]interface{}); ok {
		AdditionalPropertiesString, _ := json.Marshal(AdditionalProperties)
		json.Unmarshal(AdditionalPropertiesString, &o.AdditionalProperties)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Dialercampaignruleconfigchangecampaignruleactionentities) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

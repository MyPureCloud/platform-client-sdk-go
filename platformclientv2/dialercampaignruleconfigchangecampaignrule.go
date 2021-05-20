package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Dialercampaignruleconfigchangecampaignrule
type Dialercampaignruleconfigchangecampaignrule struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// DateCreated
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateModified
	DateModified *time.Time `json:"dateModified,omitempty"`


	// Version
	Version *int `json:"version,omitempty"`


	// CampaignRuleEntities
	CampaignRuleEntities *Dialercampaignruleconfigchangecampaignruleentities `json:"campaignRuleEntities,omitempty"`


	// CampaignRuleConditions
	CampaignRuleConditions *[]Dialercampaignruleconfigchangecampaignrulecondition `json:"campaignRuleConditions,omitempty"`


	// CampaignRuleActions
	CampaignRuleActions *[]Dialercampaignruleconfigchangecampaignruleaction `json:"campaignRuleActions,omitempty"`


	// MatchAnyConditions
	MatchAnyConditions *bool `json:"matchAnyConditions,omitempty"`


	// Enabled
	Enabled *bool `json:"enabled,omitempty"`


	// AdditionalProperties
	AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`

}

// String returns a JSON representation of the model
func (o *Dialercampaignruleconfigchangecampaignrule) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
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

func (u *Dialercampaignruleconfigchangecampaignrule) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dialercampaignruleconfigchangecampaignrule

	
	DateCreated := new(string)
	if u.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(u.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	DateModified := new(string)
	if u.DateModified != nil {
		
		*DateModified = timeutil.Strftime(u.DateModified, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateModified = nil
	}
	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		Version *int `json:"version,omitempty"`
		
		CampaignRuleEntities *Dialercampaignruleconfigchangecampaignruleentities `json:"campaignRuleEntities,omitempty"`
		
		CampaignRuleConditions *[]Dialercampaignruleconfigchangecampaignrulecondition `json:"campaignRuleConditions,omitempty"`
		
		CampaignRuleActions *[]Dialercampaignruleconfigchangecampaignruleaction `json:"campaignRuleActions,omitempty"`
		
		MatchAnyConditions *bool `json:"matchAnyConditions,omitempty"`
		
		Enabled *bool `json:"enabled,omitempty"`
		
		AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		Version: u.Version,
		
		CampaignRuleEntities: u.CampaignRuleEntities,
		
		CampaignRuleConditions: u.CampaignRuleConditions,
		
		CampaignRuleActions: u.CampaignRuleActions,
		
		MatchAnyConditions: u.MatchAnyConditions,
		
		Enabled: u.Enabled,
		
		AdditionalProperties: u.AdditionalProperties,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Dialercampaignruleconfigchangecampaignrule) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

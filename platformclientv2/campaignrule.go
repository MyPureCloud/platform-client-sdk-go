package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Campaignrule
type Campaignrule struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name - The name of the CampaignRule.
	Name *string `json:"name,omitempty"`


	// DateCreated - Creation time of the entity. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateModified - Last modified time of the entity. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`


	// Version - Required for updates, must match the version number of the most recent update
	Version *int `json:"version,omitempty"`


	// CampaignRuleEntities - The list of entities that this CampaignRule monitors.
	CampaignRuleEntities *Campaignruleentities `json:"campaignRuleEntities,omitempty"`


	// CampaignRuleConditions - The list of conditions that are evaluated on the entities.
	CampaignRuleConditions *[]Campaignrulecondition `json:"campaignRuleConditions,omitempty"`


	// CampaignRuleActions - The list of actions that are executed if the conditions are satisfied.
	CampaignRuleActions *[]Campaignruleaction `json:"campaignRuleActions,omitempty"`


	// MatchAnyConditions
	MatchAnyConditions *bool `json:"matchAnyConditions,omitempty"`


	// Enabled - Whether or not this CampaignRule is currently enabled. Required on updates.
	Enabled *bool `json:"enabled,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (u *Campaignrule) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Campaignrule

	
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
		
		CampaignRuleEntities *Campaignruleentities `json:"campaignRuleEntities,omitempty"`
		
		CampaignRuleConditions *[]Campaignrulecondition `json:"campaignRuleConditions,omitempty"`
		
		CampaignRuleActions *[]Campaignruleaction `json:"campaignRuleActions,omitempty"`
		
		MatchAnyConditions *bool `json:"matchAnyConditions,omitempty"`
		
		Enabled *bool `json:"enabled,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
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
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Campaignrule) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

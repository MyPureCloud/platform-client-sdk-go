package platformclientv2
import (
	"time"
	"encoding/json"
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

// String returns a JSON representation of the model
func (o *Campaignrule) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

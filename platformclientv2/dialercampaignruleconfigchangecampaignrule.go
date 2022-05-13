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
	// CampaignRuleEntities
	CampaignRuleEntities *Dialercampaignruleconfigchangecampaignruleentities `json:"campaignRuleEntities,omitempty"`


	// CampaignRuleConditions - The list of conditions that will trigger this Campaign Rule
	CampaignRuleConditions *[]Dialercampaignruleconfigchangecampaignrulecondition `json:"campaignRuleConditions,omitempty"`


	// CampaignRuleActions - The list of actions that will be taken when this Campaign Rule's conditions are met
	CampaignRuleActions *[]Dialercampaignruleconfigchangecampaignruleaction `json:"campaignRuleActions,omitempty"`


	// MatchAnyConditions - Whether this Campaign Rule should match any conditions (inclusive OR) or match all conditions (ALL)
	MatchAnyConditions *bool `json:"matchAnyConditions,omitempty"`


	// Enabled - Whether this campaign rule is enabled
	Enabled *bool `json:"enabled,omitempty"`


	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name - The UI-visible name of the object
	Name *string `json:"name,omitempty"`


	// DateCreated - Creation time of the entity
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateModified - Last modified time of the entity
	DateModified *time.Time `json:"dateModified,omitempty"`


	// Version - Required for updates, must match the version number of the most recent update
	Version *int `json:"version,omitempty"`

}

func (o *Dialercampaignruleconfigchangecampaignrule) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dialercampaignruleconfigchangecampaignrule
	
	DateCreated := new(string)
	if o.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(o.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	DateModified := new(string)
	if o.DateModified != nil {
		
		*DateModified = timeutil.Strftime(o.DateModified, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateModified = nil
	}
	
	return json.Marshal(&struct { 
		CampaignRuleEntities *Dialercampaignruleconfigchangecampaignruleentities `json:"campaignRuleEntities,omitempty"`
		
		CampaignRuleConditions *[]Dialercampaignruleconfigchangecampaignrulecondition `json:"campaignRuleConditions,omitempty"`
		
		CampaignRuleActions *[]Dialercampaignruleconfigchangecampaignruleaction `json:"campaignRuleActions,omitempty"`
		
		MatchAnyConditions *bool `json:"matchAnyConditions,omitempty"`
		
		Enabled *bool `json:"enabled,omitempty"`
		
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		Version *int `json:"version,omitempty"`
		*Alias
	}{ 
		CampaignRuleEntities: o.CampaignRuleEntities,
		
		CampaignRuleConditions: o.CampaignRuleConditions,
		
		CampaignRuleActions: o.CampaignRuleActions,
		
		MatchAnyConditions: o.MatchAnyConditions,
		
		Enabled: o.Enabled,
		
		Id: o.Id,
		
		Name: o.Name,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		Version: o.Version,
		Alias:    (*Alias)(o),
	})
}

func (o *Dialercampaignruleconfigchangecampaignrule) UnmarshalJSON(b []byte) error {
	var DialercampaignruleconfigchangecampaignruleMap map[string]interface{}
	err := json.Unmarshal(b, &DialercampaignruleconfigchangecampaignruleMap)
	if err != nil {
		return err
	}
	
	if CampaignRuleEntities, ok := DialercampaignruleconfigchangecampaignruleMap["campaignRuleEntities"].(map[string]interface{}); ok {
		CampaignRuleEntitiesString, _ := json.Marshal(CampaignRuleEntities)
		json.Unmarshal(CampaignRuleEntitiesString, &o.CampaignRuleEntities)
	}
	
	if CampaignRuleConditions, ok := DialercampaignruleconfigchangecampaignruleMap["campaignRuleConditions"].([]interface{}); ok {
		CampaignRuleConditionsString, _ := json.Marshal(CampaignRuleConditions)
		json.Unmarshal(CampaignRuleConditionsString, &o.CampaignRuleConditions)
	}
	
	if CampaignRuleActions, ok := DialercampaignruleconfigchangecampaignruleMap["campaignRuleActions"].([]interface{}); ok {
		CampaignRuleActionsString, _ := json.Marshal(CampaignRuleActions)
		json.Unmarshal(CampaignRuleActionsString, &o.CampaignRuleActions)
	}
	
	if MatchAnyConditions, ok := DialercampaignruleconfigchangecampaignruleMap["matchAnyConditions"].(bool); ok {
		o.MatchAnyConditions = &MatchAnyConditions
	}
    
	if Enabled, ok := DialercampaignruleconfigchangecampaignruleMap["enabled"].(bool); ok {
		o.Enabled = &Enabled
	}
    
	if Id, ok := DialercampaignruleconfigchangecampaignruleMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := DialercampaignruleconfigchangecampaignruleMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if dateCreatedString, ok := DialercampaignruleconfigchangecampaignruleMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := DialercampaignruleconfigchangecampaignruleMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if Version, ok := DialercampaignruleconfigchangecampaignruleMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Dialercampaignruleconfigchangecampaignrule) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

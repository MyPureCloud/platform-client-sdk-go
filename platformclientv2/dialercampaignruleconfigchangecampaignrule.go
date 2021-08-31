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
		Id: o.Id,
		
		Name: o.Name,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		Version: o.Version,
		
		CampaignRuleEntities: o.CampaignRuleEntities,
		
		CampaignRuleConditions: o.CampaignRuleConditions,
		
		CampaignRuleActions: o.CampaignRuleActions,
		
		MatchAnyConditions: o.MatchAnyConditions,
		
		Enabled: o.Enabled,
		
		AdditionalProperties: o.AdditionalProperties,
		Alias:    (*Alias)(o),
	})
}

func (o *Dialercampaignruleconfigchangecampaignrule) UnmarshalJSON(b []byte) error {
	var DialercampaignruleconfigchangecampaignruleMap map[string]interface{}
	err := json.Unmarshal(b, &DialercampaignruleconfigchangecampaignruleMap)
	if err != nil {
		return err
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
	
	if AdditionalProperties, ok := DialercampaignruleconfigchangecampaignruleMap["additionalProperties"].(map[string]interface{}); ok {
		AdditionalPropertiesString, _ := json.Marshal(AdditionalProperties)
		json.Unmarshal(AdditionalPropertiesString, &o.AdditionalProperties)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Dialercampaignruleconfigchangecampaignrule) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

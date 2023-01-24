package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Campaignrule
type Campaignrule struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
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

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Campaignrule) SetField(field string, fieldValue interface{}) {
	// Get Value object for field
	target := reflect.ValueOf(o)
	targetField := reflect.Indirect(target).FieldByName(field)

	// Set value
	if fieldValue != nil {
		targetField.Set(reflect.ValueOf(fieldValue))
	} else {
		// Must create a new Value (creates **type) then get its element (*type), which will be nil pointer of the appropriate type
		x := reflect.Indirect(reflect.New(targetField.Type()))
		targetField.Set(x)
	}

	// Add field to set field names list
	if o.SetFieldNames == nil {
		o.SetFieldNames = make(map[string]bool)
	}
	o.SetFieldNames[field] = true
}

func (o Campaignrule) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateCreated","DateModified", }
		localDateTimeFields := []string{  }
		dateFields := []string{  }

		// Construct object
		newObj := make(map[string]interface{})
		for fieldName := range o.SetFieldNames {
			// Get initial field value
			fieldValue := val.FieldByName(fieldName).Interface()

			// Apply value formatting overrides
			if contains(dateTimeFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%dT%H:%M:%S.%fZ")
			} else if contains(localDateTimeFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%dT%H:%M:%S.%f")
			} else if contains(dateFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%d")
			}

			// Assign value to field using JSON tag name
			newObj[getFieldName(reflect.TypeOf(&o), fieldName)] = fieldValue
		}

		// Marshal and return dynamically constructed interface
		return json.Marshal(newObj)
	}

	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Campaignrule
	
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
		
		CampaignRuleEntities *Campaignruleentities `json:"campaignRuleEntities,omitempty"`
		
		CampaignRuleConditions *[]Campaignrulecondition `json:"campaignRuleConditions,omitempty"`
		
		CampaignRuleActions *[]Campaignruleaction `json:"campaignRuleActions,omitempty"`
		
		MatchAnyConditions *bool `json:"matchAnyConditions,omitempty"`
		
		Enabled *bool `json:"enabled,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
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
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Campaignrule) UnmarshalJSON(b []byte) error {
	var CampaignruleMap map[string]interface{}
	err := json.Unmarshal(b, &CampaignruleMap)
	if err != nil {
		return err
	}
	
	if Id, ok := CampaignruleMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := CampaignruleMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if dateCreatedString, ok := CampaignruleMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := CampaignruleMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if Version, ok := CampaignruleMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	
	if CampaignRuleEntities, ok := CampaignruleMap["campaignRuleEntities"].(map[string]interface{}); ok {
		CampaignRuleEntitiesString, _ := json.Marshal(CampaignRuleEntities)
		json.Unmarshal(CampaignRuleEntitiesString, &o.CampaignRuleEntities)
	}
	
	if CampaignRuleConditions, ok := CampaignruleMap["campaignRuleConditions"].([]interface{}); ok {
		CampaignRuleConditionsString, _ := json.Marshal(CampaignRuleConditions)
		json.Unmarshal(CampaignRuleConditionsString, &o.CampaignRuleConditions)
	}
	
	if CampaignRuleActions, ok := CampaignruleMap["campaignRuleActions"].([]interface{}); ok {
		CampaignRuleActionsString, _ := json.Marshal(CampaignRuleActions)
		json.Unmarshal(CampaignRuleActionsString, &o.CampaignRuleActions)
	}
	
	if MatchAnyConditions, ok := CampaignruleMap["matchAnyConditions"].(bool); ok {
		o.MatchAnyConditions = &MatchAnyConditions
	}
    
	if Enabled, ok := CampaignruleMap["enabled"].(bool); ok {
		o.Enabled = &Enabled
	}
    
	if SelfUri, ok := CampaignruleMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Campaignrule) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

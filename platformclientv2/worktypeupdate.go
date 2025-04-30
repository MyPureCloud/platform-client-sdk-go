package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Worktypeupdate
type Worktypeupdate struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Name - The name of the Worktype. Valid length between 3 and 256 characters.
	Name *string `json:"name,omitempty"`

	// DefaultWorkbinId - The ID of the default Workbin for Workitems created from the Worktype.
	DefaultWorkbinId *string `json:"defaultWorkbinId,omitempty"`

	// DefaultDurationSeconds - The default duration in seconds for Workitems created from the Worktype. Maximum of 365 days.
	DefaultDurationSeconds *int `json:"defaultDurationSeconds,omitempty"`

	// DefaultExpirationSeconds - The default expiration time in seconds for Workitems created from the Worktype. Maximum of 365 days.
	DefaultExpirationSeconds *int `json:"defaultExpirationSeconds,omitempty"`

	// DefaultDueDurationSeconds - The default due duration in seconds for Workitems created from the Worktype. Maximum of 365 days.
	DefaultDueDurationSeconds *int `json:"defaultDueDurationSeconds,omitempty"`

	// DefaultPriority - The default priority for Workitems created from the Worktype. The valid range is between -25,000,000 and 25,000,000.
	DefaultPriority *int `json:"defaultPriority,omitempty"`

	// DefaultTtlSeconds - The default time to time to live in seconds for Workitems created from the Worktype. The valid range is between 1 and 365 days.
	DefaultTtlSeconds *int `json:"defaultTtlSeconds,omitempty"`

	// AssignmentEnabled - When set to true, Workitems will be sent to the queue of the Worktype as they are created. Default value is false.
	AssignmentEnabled *bool `json:"assignmentEnabled,omitempty"`

	// SchemaId - The ID of the custom attribute schema for Workitems created from the Worktype. Must be a valid UUID.
	SchemaId *string `json:"schemaId,omitempty"`

	// ServiceLevelTarget - The target service level for Workitems created from the Worktype. The default value is 100.
	ServiceLevelTarget *int `json:"serviceLevelTarget,omitempty"`

	// RuleSettings - Settings for the worktypes rules.
	RuleSettings *Workitemrulesettings `json:"ruleSettings,omitempty"`

	// Description - The description of the Worktype. Maximum length of 512 characters.
	Description *string `json:"description,omitempty"`

	// DefaultStatusId - The ID of the default status for Workitems created from the Worktype. Must be a valid UUID.
	DefaultStatusId *string `json:"defaultStatusId,omitempty"`

	// SchemaVersion - The version of the Worktypes custom attribute schema. The latest schema version will be used if this property is not set.
	SchemaVersion *int `json:"schemaVersion,omitempty"`

	// DefaultLanguageId - The ID of the default language for Workitems created from the Worktype. Must be a valid UUID.
	DefaultLanguageId *string `json:"defaultLanguageId,omitempty"`

	// DefaultSkillIds - The IDs of the default skills for Workitems created from the Worktype. Must be valid UUIDs. Maximum of 20 IDs
	DefaultSkillIds *[]string `json:"defaultSkillIds,omitempty"`

	// DefaultQueueId - The ID of the default queue for Workitems created from the Worktype. Must be a valid UUID.
	DefaultQueueId *string `json:"defaultQueueId,omitempty"`

	// DefaultScriptId - The default script for Workitems created from the Worktype. Must be a valid UUID.
	DefaultScriptId *string `json:"defaultScriptId,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Worktypeupdate) SetField(field string, fieldValue interface{}) {
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

func (o Worktypeupdate) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{  }
		localDateTimeFields := []string{  }
		dateFields := []string{  }

		// Construct object
		newObj := make(map[string]interface{})
		for fieldName := range o.SetFieldNames {
			// Get initial field value
			fieldValue := val.FieldByName(fieldName).Interface()

			// Apply value formatting overrides
			if fieldValue == nil || reflect.ValueOf(fieldValue).IsNil()  {
				// Do nothing. Just catching this case to avoid trying to custom serialize a nil value
			} else if contains(dateTimeFields, fieldName) {
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
	type Alias Worktypeupdate
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		DefaultWorkbinId *string `json:"defaultWorkbinId,omitempty"`
		
		DefaultDurationSeconds *int `json:"defaultDurationSeconds,omitempty"`
		
		DefaultExpirationSeconds *int `json:"defaultExpirationSeconds,omitempty"`
		
		DefaultDueDurationSeconds *int `json:"defaultDueDurationSeconds,omitempty"`
		
		DefaultPriority *int `json:"defaultPriority,omitempty"`
		
		DefaultTtlSeconds *int `json:"defaultTtlSeconds,omitempty"`
		
		AssignmentEnabled *bool `json:"assignmentEnabled,omitempty"`
		
		SchemaId *string `json:"schemaId,omitempty"`
		
		ServiceLevelTarget *int `json:"serviceLevelTarget,omitempty"`
		
		RuleSettings *Workitemrulesettings `json:"ruleSettings,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		DefaultStatusId *string `json:"defaultStatusId,omitempty"`
		
		SchemaVersion *int `json:"schemaVersion,omitempty"`
		
		DefaultLanguageId *string `json:"defaultLanguageId,omitempty"`
		
		DefaultSkillIds *[]string `json:"defaultSkillIds,omitempty"`
		
		DefaultQueueId *string `json:"defaultQueueId,omitempty"`
		
		DefaultScriptId *string `json:"defaultScriptId,omitempty"`
		Alias
	}{ 
		Name: o.Name,
		
		DefaultWorkbinId: o.DefaultWorkbinId,
		
		DefaultDurationSeconds: o.DefaultDurationSeconds,
		
		DefaultExpirationSeconds: o.DefaultExpirationSeconds,
		
		DefaultDueDurationSeconds: o.DefaultDueDurationSeconds,
		
		DefaultPriority: o.DefaultPriority,
		
		DefaultTtlSeconds: o.DefaultTtlSeconds,
		
		AssignmentEnabled: o.AssignmentEnabled,
		
		SchemaId: o.SchemaId,
		
		ServiceLevelTarget: o.ServiceLevelTarget,
		
		RuleSettings: o.RuleSettings,
		
		Description: o.Description,
		
		DefaultStatusId: o.DefaultStatusId,
		
		SchemaVersion: o.SchemaVersion,
		
		DefaultLanguageId: o.DefaultLanguageId,
		
		DefaultSkillIds: o.DefaultSkillIds,
		
		DefaultQueueId: o.DefaultQueueId,
		
		DefaultScriptId: o.DefaultScriptId,
		Alias:    (Alias)(o),
	})
}

func (o *Worktypeupdate) UnmarshalJSON(b []byte) error {
	var WorktypeupdateMap map[string]interface{}
	err := json.Unmarshal(b, &WorktypeupdateMap)
	if err != nil {
		return err
	}
	
	if Name, ok := WorktypeupdateMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if DefaultWorkbinId, ok := WorktypeupdateMap["defaultWorkbinId"].(string); ok {
		o.DefaultWorkbinId = &DefaultWorkbinId
	}
    
	if DefaultDurationSeconds, ok := WorktypeupdateMap["defaultDurationSeconds"].(float64); ok {
		DefaultDurationSecondsInt := int(DefaultDurationSeconds)
		o.DefaultDurationSeconds = &DefaultDurationSecondsInt
	}
	
	if DefaultExpirationSeconds, ok := WorktypeupdateMap["defaultExpirationSeconds"].(float64); ok {
		DefaultExpirationSecondsInt := int(DefaultExpirationSeconds)
		o.DefaultExpirationSeconds = &DefaultExpirationSecondsInt
	}
	
	if DefaultDueDurationSeconds, ok := WorktypeupdateMap["defaultDueDurationSeconds"].(float64); ok {
		DefaultDueDurationSecondsInt := int(DefaultDueDurationSeconds)
		o.DefaultDueDurationSeconds = &DefaultDueDurationSecondsInt
	}
	
	if DefaultPriority, ok := WorktypeupdateMap["defaultPriority"].(float64); ok {
		DefaultPriorityInt := int(DefaultPriority)
		o.DefaultPriority = &DefaultPriorityInt
	}
	
	if DefaultTtlSeconds, ok := WorktypeupdateMap["defaultTtlSeconds"].(float64); ok {
		DefaultTtlSecondsInt := int(DefaultTtlSeconds)
		o.DefaultTtlSeconds = &DefaultTtlSecondsInt
	}
	
	if AssignmentEnabled, ok := WorktypeupdateMap["assignmentEnabled"].(bool); ok {
		o.AssignmentEnabled = &AssignmentEnabled
	}
    
	if SchemaId, ok := WorktypeupdateMap["schemaId"].(string); ok {
		o.SchemaId = &SchemaId
	}
    
	if ServiceLevelTarget, ok := WorktypeupdateMap["serviceLevelTarget"].(float64); ok {
		ServiceLevelTargetInt := int(ServiceLevelTarget)
		o.ServiceLevelTarget = &ServiceLevelTargetInt
	}
	
	if RuleSettings, ok := WorktypeupdateMap["ruleSettings"].(map[string]interface{}); ok {
		RuleSettingsString, _ := json.Marshal(RuleSettings)
		json.Unmarshal(RuleSettingsString, &o.RuleSettings)
	}
	
	if Description, ok := WorktypeupdateMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if DefaultStatusId, ok := WorktypeupdateMap["defaultStatusId"].(string); ok {
		o.DefaultStatusId = &DefaultStatusId
	}
    
	if SchemaVersion, ok := WorktypeupdateMap["schemaVersion"].(float64); ok {
		SchemaVersionInt := int(SchemaVersion)
		o.SchemaVersion = &SchemaVersionInt
	}
	
	if DefaultLanguageId, ok := WorktypeupdateMap["defaultLanguageId"].(string); ok {
		o.DefaultLanguageId = &DefaultLanguageId
	}
    
	if DefaultSkillIds, ok := WorktypeupdateMap["defaultSkillIds"].([]interface{}); ok {
		DefaultSkillIdsString, _ := json.Marshal(DefaultSkillIds)
		json.Unmarshal(DefaultSkillIdsString, &o.DefaultSkillIds)
	}
	
	if DefaultQueueId, ok := WorktypeupdateMap["defaultQueueId"].(string); ok {
		o.DefaultQueueId = &DefaultQueueId
	}
    
	if DefaultScriptId, ok := WorktypeupdateMap["defaultScriptId"].(string); ok {
		o.DefaultScriptId = &DefaultScriptId
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Worktypeupdate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

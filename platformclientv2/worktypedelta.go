package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Worktypedelta
type Worktypedelta struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Name
	Name *Workitemsattributechangestring `json:"name,omitempty"`

	// Description
	Description *Workitemsattributechangestring `json:"description,omitempty"`

	// Statuses
	Statuses *Workitemsattributechangelist `json:"statuses,omitempty"`

	// DefaultWorkbinId
	DefaultWorkbinId *Workitemsattributechangestring `json:"defaultWorkbinId,omitempty"`

	// DefaultDurationSeconds
	DefaultDurationSeconds *Workitemsattributechangeinteger `json:"defaultDurationSeconds,omitempty"`

	// DefaultExpirationSeconds
	DefaultExpirationSeconds *Workitemsattributechangeinteger `json:"defaultExpirationSeconds,omitempty"`

	// DefaultDueDurationSeconds
	DefaultDueDurationSeconds *Workitemsattributechangeinteger `json:"defaultDueDurationSeconds,omitempty"`

	// DefaultPriority
	DefaultPriority *Workitemsattributechangeinteger `json:"defaultPriority,omitempty"`

	// DefaultSkillIds
	DefaultSkillIds *Workitemsattributechangelist `json:"defaultSkillIds,omitempty"`

	// DefaultStatusId
	DefaultStatusId *Workitemsattributechangestring `json:"defaultStatusId,omitempty"`

	// DefaultLanguageId
	DefaultLanguageId *Workitemsattributechangestring `json:"defaultLanguageId,omitempty"`

	// DefaultTtlSeconds
	DefaultTtlSeconds *Workitemsattributechangeinteger `json:"defaultTtlSeconds,omitempty"`

	// AssignmentEnabled
	AssignmentEnabled *Workitemsattributechangeboolean `json:"assignmentEnabled,omitempty"`

	// DefaultQueueId
	DefaultQueueId *Workitemsattributechangestring `json:"defaultQueueId,omitempty"`

	// SchemaId
	SchemaId *Workitemsattributechangestring `json:"schemaId,omitempty"`

	// SchemaVersion
	SchemaVersion *Workitemsattributechangestring `json:"schemaVersion,omitempty"`

	// DateModified
	DateModified *Workitemsattributechangeinstant `json:"dateModified,omitempty"`

	// ModifiedBy
	ModifiedBy *Workitemsattributechangestring `json:"modifiedBy,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Worktypedelta) SetField(field string, fieldValue interface{}) {
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

func (o Worktypedelta) MarshalJSON() ([]byte, error) {
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
	type Alias Worktypedelta
	
	return json.Marshal(&struct { 
		Name *Workitemsattributechangestring `json:"name,omitempty"`
		
		Description *Workitemsattributechangestring `json:"description,omitempty"`
		
		Statuses *Workitemsattributechangelist `json:"statuses,omitempty"`
		
		DefaultWorkbinId *Workitemsattributechangestring `json:"defaultWorkbinId,omitempty"`
		
		DefaultDurationSeconds *Workitemsattributechangeinteger `json:"defaultDurationSeconds,omitempty"`
		
		DefaultExpirationSeconds *Workitemsattributechangeinteger `json:"defaultExpirationSeconds,omitempty"`
		
		DefaultDueDurationSeconds *Workitemsattributechangeinteger `json:"defaultDueDurationSeconds,omitempty"`
		
		DefaultPriority *Workitemsattributechangeinteger `json:"defaultPriority,omitempty"`
		
		DefaultSkillIds *Workitemsattributechangelist `json:"defaultSkillIds,omitempty"`
		
		DefaultStatusId *Workitemsattributechangestring `json:"defaultStatusId,omitempty"`
		
		DefaultLanguageId *Workitemsattributechangestring `json:"defaultLanguageId,omitempty"`
		
		DefaultTtlSeconds *Workitemsattributechangeinteger `json:"defaultTtlSeconds,omitempty"`
		
		AssignmentEnabled *Workitemsattributechangeboolean `json:"assignmentEnabled,omitempty"`
		
		DefaultQueueId *Workitemsattributechangestring `json:"defaultQueueId,omitempty"`
		
		SchemaId *Workitemsattributechangestring `json:"schemaId,omitempty"`
		
		SchemaVersion *Workitemsattributechangestring `json:"schemaVersion,omitempty"`
		
		DateModified *Workitemsattributechangeinstant `json:"dateModified,omitempty"`
		
		ModifiedBy *Workitemsattributechangestring `json:"modifiedBy,omitempty"`
		Alias
	}{ 
		Name: o.Name,
		
		Description: o.Description,
		
		Statuses: o.Statuses,
		
		DefaultWorkbinId: o.DefaultWorkbinId,
		
		DefaultDurationSeconds: o.DefaultDurationSeconds,
		
		DefaultExpirationSeconds: o.DefaultExpirationSeconds,
		
		DefaultDueDurationSeconds: o.DefaultDueDurationSeconds,
		
		DefaultPriority: o.DefaultPriority,
		
		DefaultSkillIds: o.DefaultSkillIds,
		
		DefaultStatusId: o.DefaultStatusId,
		
		DefaultLanguageId: o.DefaultLanguageId,
		
		DefaultTtlSeconds: o.DefaultTtlSeconds,
		
		AssignmentEnabled: o.AssignmentEnabled,
		
		DefaultQueueId: o.DefaultQueueId,
		
		SchemaId: o.SchemaId,
		
		SchemaVersion: o.SchemaVersion,
		
		DateModified: o.DateModified,
		
		ModifiedBy: o.ModifiedBy,
		Alias:    (Alias)(o),
	})
}

func (o *Worktypedelta) UnmarshalJSON(b []byte) error {
	var WorktypedeltaMap map[string]interface{}
	err := json.Unmarshal(b, &WorktypedeltaMap)
	if err != nil {
		return err
	}
	
	if Name, ok := WorktypedeltaMap["name"].(map[string]interface{}); ok {
		NameString, _ := json.Marshal(Name)
		json.Unmarshal(NameString, &o.Name)
	}
	
	if Description, ok := WorktypedeltaMap["description"].(map[string]interface{}); ok {
		DescriptionString, _ := json.Marshal(Description)
		json.Unmarshal(DescriptionString, &o.Description)
	}
	
	if Statuses, ok := WorktypedeltaMap["statuses"].(map[string]interface{}); ok {
		StatusesString, _ := json.Marshal(Statuses)
		json.Unmarshal(StatusesString, &o.Statuses)
	}
	
	if DefaultWorkbinId, ok := WorktypedeltaMap["defaultWorkbinId"].(map[string]interface{}); ok {
		DefaultWorkbinIdString, _ := json.Marshal(DefaultWorkbinId)
		json.Unmarshal(DefaultWorkbinIdString, &o.DefaultWorkbinId)
	}
	
	if DefaultDurationSeconds, ok := WorktypedeltaMap["defaultDurationSeconds"].(map[string]interface{}); ok {
		DefaultDurationSecondsString, _ := json.Marshal(DefaultDurationSeconds)
		json.Unmarshal(DefaultDurationSecondsString, &o.DefaultDurationSeconds)
	}
	
	if DefaultExpirationSeconds, ok := WorktypedeltaMap["defaultExpirationSeconds"].(map[string]interface{}); ok {
		DefaultExpirationSecondsString, _ := json.Marshal(DefaultExpirationSeconds)
		json.Unmarshal(DefaultExpirationSecondsString, &o.DefaultExpirationSeconds)
	}
	
	if DefaultDueDurationSeconds, ok := WorktypedeltaMap["defaultDueDurationSeconds"].(map[string]interface{}); ok {
		DefaultDueDurationSecondsString, _ := json.Marshal(DefaultDueDurationSeconds)
		json.Unmarshal(DefaultDueDurationSecondsString, &o.DefaultDueDurationSeconds)
	}
	
	if DefaultPriority, ok := WorktypedeltaMap["defaultPriority"].(map[string]interface{}); ok {
		DefaultPriorityString, _ := json.Marshal(DefaultPriority)
		json.Unmarshal(DefaultPriorityString, &o.DefaultPriority)
	}
	
	if DefaultSkillIds, ok := WorktypedeltaMap["defaultSkillIds"].(map[string]interface{}); ok {
		DefaultSkillIdsString, _ := json.Marshal(DefaultSkillIds)
		json.Unmarshal(DefaultSkillIdsString, &o.DefaultSkillIds)
	}
	
	if DefaultStatusId, ok := WorktypedeltaMap["defaultStatusId"].(map[string]interface{}); ok {
		DefaultStatusIdString, _ := json.Marshal(DefaultStatusId)
		json.Unmarshal(DefaultStatusIdString, &o.DefaultStatusId)
	}
	
	if DefaultLanguageId, ok := WorktypedeltaMap["defaultLanguageId"].(map[string]interface{}); ok {
		DefaultLanguageIdString, _ := json.Marshal(DefaultLanguageId)
		json.Unmarshal(DefaultLanguageIdString, &o.DefaultLanguageId)
	}
	
	if DefaultTtlSeconds, ok := WorktypedeltaMap["defaultTtlSeconds"].(map[string]interface{}); ok {
		DefaultTtlSecondsString, _ := json.Marshal(DefaultTtlSeconds)
		json.Unmarshal(DefaultTtlSecondsString, &o.DefaultTtlSeconds)
	}
	
	if AssignmentEnabled, ok := WorktypedeltaMap["assignmentEnabled"].(map[string]interface{}); ok {
		AssignmentEnabledString, _ := json.Marshal(AssignmentEnabled)
		json.Unmarshal(AssignmentEnabledString, &o.AssignmentEnabled)
	}
	
	if DefaultQueueId, ok := WorktypedeltaMap["defaultQueueId"].(map[string]interface{}); ok {
		DefaultQueueIdString, _ := json.Marshal(DefaultQueueId)
		json.Unmarshal(DefaultQueueIdString, &o.DefaultQueueId)
	}
	
	if SchemaId, ok := WorktypedeltaMap["schemaId"].(map[string]interface{}); ok {
		SchemaIdString, _ := json.Marshal(SchemaId)
		json.Unmarshal(SchemaIdString, &o.SchemaId)
	}
	
	if SchemaVersion, ok := WorktypedeltaMap["schemaVersion"].(map[string]interface{}); ok {
		SchemaVersionString, _ := json.Marshal(SchemaVersion)
		json.Unmarshal(SchemaVersionString, &o.SchemaVersion)
	}
	
	if DateModified, ok := WorktypedeltaMap["dateModified"].(map[string]interface{}); ok {
		DateModifiedString, _ := json.Marshal(DateModified)
		json.Unmarshal(DateModifiedString, &o.DateModified)
	}
	
	if ModifiedBy, ok := WorktypedeltaMap["modifiedBy"].(map[string]interface{}); ok {
		ModifiedByString, _ := json.Marshal(ModifiedBy)
		json.Unmarshal(ModifiedByString, &o.ModifiedBy)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Worktypedelta) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

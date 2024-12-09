package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Workitemdelta
type Workitemdelta struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Name
	Name *Workitemsattributechangestring `json:"name,omitempty"`

	// Description
	Description *Workitemsattributechangestring `json:"description,omitempty"`

	// LanguageId
	LanguageId *Workitemsattributechangestring `json:"languageId,omitempty"`

	// UtilizationLabelId
	UtilizationLabelId *Workitemsattributechangestring `json:"utilizationLabelId,omitempty"`

	// Priority
	Priority *Workitemsattributechangeinteger `json:"priority,omitempty"`

	// SkillIds
	SkillIds *Workitemsattributechangelist `json:"skillIds,omitempty"`

	// PreferredAgentIds
	PreferredAgentIds *Workitemsattributechangelist `json:"preferredAgentIds,omitempty"`

	// DateDue
	DateDue *Workitemsattributechangeinstant `json:"dateDue,omitempty"`

	// DateExpires
	DateExpires *Workitemsattributechangeinstant `json:"dateExpires,omitempty"`

	// DurationSeconds
	DurationSeconds *Workitemsattributechangeinteger `json:"durationSeconds,omitempty"`

	// StatusId
	StatusId *Workitemsattributechangestring `json:"statusId,omitempty"`

	// ReporterId
	ReporterId *Workitemsattributechangestring `json:"reporterId,omitempty"`

	// ExternalContactId
	ExternalContactId *Workitemsattributechangestring `json:"externalContactId,omitempty"`

	// AssigneeId
	AssigneeId *Workitemsattributechangestring `json:"assigneeId,omitempty"`

	// WorkbinId
	WorkbinId *Workitemsattributechangestring `json:"workbinId,omitempty"`

	// QueueId
	QueueId *Workitemsattributechangestring `json:"queueId,omitempty"`

	// ExternalTag
	ExternalTag *Workitemsattributechangestring `json:"externalTag,omitempty"`

	// WrapupId
	WrapupId *Workitemsattributechangestring `json:"wrapupId,omitempty"`

	// Ttl
	Ttl *Workitemsattributechangeinteger `json:"ttl,omitempty"`

	// DateClosed
	DateClosed *Workitemsattributechangeinstant `json:"dateClosed,omitempty"`

	// AssignmentState
	AssignmentState *Workitemsattributechangestring `json:"assignmentState,omitempty"`

	// AutoStatusTransition
	AutoStatusTransition *Workitemsattributechangeboolean `json:"autoStatusTransition,omitempty"`

	// CustomFields
	CustomFields *Workitemsattributechangemap `json:"customFields,omitempty"`

	// DateModified
	DateModified *Workitemsattributechangeinstant `json:"dateModified,omitempty"`

	// ModifiedBy
	ModifiedBy *Workitemsattributechangestring `json:"modifiedBy,omitempty"`

	// StatusCategory
	StatusCategory *Workitemsattributechangeworkitemstatuscategory `json:"statusCategory,omitempty"`

	// ScriptId
	ScriptId *Workitemsattributechangestring `json:"scriptId,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Workitemdelta) SetField(field string, fieldValue interface{}) {
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

func (o Workitemdelta) MarshalJSON() ([]byte, error) {
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
	type Alias Workitemdelta
	
	return json.Marshal(&struct { 
		Name *Workitemsattributechangestring `json:"name,omitempty"`
		
		Description *Workitemsattributechangestring `json:"description,omitempty"`
		
		LanguageId *Workitemsattributechangestring `json:"languageId,omitempty"`
		
		UtilizationLabelId *Workitemsattributechangestring `json:"utilizationLabelId,omitempty"`
		
		Priority *Workitemsattributechangeinteger `json:"priority,omitempty"`
		
		SkillIds *Workitemsattributechangelist `json:"skillIds,omitempty"`
		
		PreferredAgentIds *Workitemsattributechangelist `json:"preferredAgentIds,omitempty"`
		
		DateDue *Workitemsattributechangeinstant `json:"dateDue,omitempty"`
		
		DateExpires *Workitemsattributechangeinstant `json:"dateExpires,omitempty"`
		
		DurationSeconds *Workitemsattributechangeinteger `json:"durationSeconds,omitempty"`
		
		StatusId *Workitemsattributechangestring `json:"statusId,omitempty"`
		
		ReporterId *Workitemsattributechangestring `json:"reporterId,omitempty"`
		
		ExternalContactId *Workitemsattributechangestring `json:"externalContactId,omitempty"`
		
		AssigneeId *Workitemsattributechangestring `json:"assigneeId,omitempty"`
		
		WorkbinId *Workitemsattributechangestring `json:"workbinId,omitempty"`
		
		QueueId *Workitemsattributechangestring `json:"queueId,omitempty"`
		
		ExternalTag *Workitemsattributechangestring `json:"externalTag,omitempty"`
		
		WrapupId *Workitemsattributechangestring `json:"wrapupId,omitempty"`
		
		Ttl *Workitemsattributechangeinteger `json:"ttl,omitempty"`
		
		DateClosed *Workitemsattributechangeinstant `json:"dateClosed,omitempty"`
		
		AssignmentState *Workitemsattributechangestring `json:"assignmentState,omitempty"`
		
		AutoStatusTransition *Workitemsattributechangeboolean `json:"autoStatusTransition,omitempty"`
		
		CustomFields *Workitemsattributechangemap `json:"customFields,omitempty"`
		
		DateModified *Workitemsattributechangeinstant `json:"dateModified,omitempty"`
		
		ModifiedBy *Workitemsattributechangestring `json:"modifiedBy,omitempty"`
		
		StatusCategory *Workitemsattributechangeworkitemstatuscategory `json:"statusCategory,omitempty"`
		
		ScriptId *Workitemsattributechangestring `json:"scriptId,omitempty"`
		Alias
	}{ 
		Name: o.Name,
		
		Description: o.Description,
		
		LanguageId: o.LanguageId,
		
		UtilizationLabelId: o.UtilizationLabelId,
		
		Priority: o.Priority,
		
		SkillIds: o.SkillIds,
		
		PreferredAgentIds: o.PreferredAgentIds,
		
		DateDue: o.DateDue,
		
		DateExpires: o.DateExpires,
		
		DurationSeconds: o.DurationSeconds,
		
		StatusId: o.StatusId,
		
		ReporterId: o.ReporterId,
		
		ExternalContactId: o.ExternalContactId,
		
		AssigneeId: o.AssigneeId,
		
		WorkbinId: o.WorkbinId,
		
		QueueId: o.QueueId,
		
		ExternalTag: o.ExternalTag,
		
		WrapupId: o.WrapupId,
		
		Ttl: o.Ttl,
		
		DateClosed: o.DateClosed,
		
		AssignmentState: o.AssignmentState,
		
		AutoStatusTransition: o.AutoStatusTransition,
		
		CustomFields: o.CustomFields,
		
		DateModified: o.DateModified,
		
		ModifiedBy: o.ModifiedBy,
		
		StatusCategory: o.StatusCategory,
		
		ScriptId: o.ScriptId,
		Alias:    (Alias)(o),
	})
}

func (o *Workitemdelta) UnmarshalJSON(b []byte) error {
	var WorkitemdeltaMap map[string]interface{}
	err := json.Unmarshal(b, &WorkitemdeltaMap)
	if err != nil {
		return err
	}
	
	if Name, ok := WorkitemdeltaMap["name"].(map[string]interface{}); ok {
		NameString, _ := json.Marshal(Name)
		json.Unmarshal(NameString, &o.Name)
	}
	
	if Description, ok := WorkitemdeltaMap["description"].(map[string]interface{}); ok {
		DescriptionString, _ := json.Marshal(Description)
		json.Unmarshal(DescriptionString, &o.Description)
	}
	
	if LanguageId, ok := WorkitemdeltaMap["languageId"].(map[string]interface{}); ok {
		LanguageIdString, _ := json.Marshal(LanguageId)
		json.Unmarshal(LanguageIdString, &o.LanguageId)
	}
	
	if UtilizationLabelId, ok := WorkitemdeltaMap["utilizationLabelId"].(map[string]interface{}); ok {
		UtilizationLabelIdString, _ := json.Marshal(UtilizationLabelId)
		json.Unmarshal(UtilizationLabelIdString, &o.UtilizationLabelId)
	}
	
	if Priority, ok := WorkitemdeltaMap["priority"].(map[string]interface{}); ok {
		PriorityString, _ := json.Marshal(Priority)
		json.Unmarshal(PriorityString, &o.Priority)
	}
	
	if SkillIds, ok := WorkitemdeltaMap["skillIds"].(map[string]interface{}); ok {
		SkillIdsString, _ := json.Marshal(SkillIds)
		json.Unmarshal(SkillIdsString, &o.SkillIds)
	}
	
	if PreferredAgentIds, ok := WorkitemdeltaMap["preferredAgentIds"].(map[string]interface{}); ok {
		PreferredAgentIdsString, _ := json.Marshal(PreferredAgentIds)
		json.Unmarshal(PreferredAgentIdsString, &o.PreferredAgentIds)
	}
	
	if DateDue, ok := WorkitemdeltaMap["dateDue"].(map[string]interface{}); ok {
		DateDueString, _ := json.Marshal(DateDue)
		json.Unmarshal(DateDueString, &o.DateDue)
	}
	
	if DateExpires, ok := WorkitemdeltaMap["dateExpires"].(map[string]interface{}); ok {
		DateExpiresString, _ := json.Marshal(DateExpires)
		json.Unmarshal(DateExpiresString, &o.DateExpires)
	}
	
	if DurationSeconds, ok := WorkitemdeltaMap["durationSeconds"].(map[string]interface{}); ok {
		DurationSecondsString, _ := json.Marshal(DurationSeconds)
		json.Unmarshal(DurationSecondsString, &o.DurationSeconds)
	}
	
	if StatusId, ok := WorkitemdeltaMap["statusId"].(map[string]interface{}); ok {
		StatusIdString, _ := json.Marshal(StatusId)
		json.Unmarshal(StatusIdString, &o.StatusId)
	}
	
	if ReporterId, ok := WorkitemdeltaMap["reporterId"].(map[string]interface{}); ok {
		ReporterIdString, _ := json.Marshal(ReporterId)
		json.Unmarshal(ReporterIdString, &o.ReporterId)
	}
	
	if ExternalContactId, ok := WorkitemdeltaMap["externalContactId"].(map[string]interface{}); ok {
		ExternalContactIdString, _ := json.Marshal(ExternalContactId)
		json.Unmarshal(ExternalContactIdString, &o.ExternalContactId)
	}
	
	if AssigneeId, ok := WorkitemdeltaMap["assigneeId"].(map[string]interface{}); ok {
		AssigneeIdString, _ := json.Marshal(AssigneeId)
		json.Unmarshal(AssigneeIdString, &o.AssigneeId)
	}
	
	if WorkbinId, ok := WorkitemdeltaMap["workbinId"].(map[string]interface{}); ok {
		WorkbinIdString, _ := json.Marshal(WorkbinId)
		json.Unmarshal(WorkbinIdString, &o.WorkbinId)
	}
	
	if QueueId, ok := WorkitemdeltaMap["queueId"].(map[string]interface{}); ok {
		QueueIdString, _ := json.Marshal(QueueId)
		json.Unmarshal(QueueIdString, &o.QueueId)
	}
	
	if ExternalTag, ok := WorkitemdeltaMap["externalTag"].(map[string]interface{}); ok {
		ExternalTagString, _ := json.Marshal(ExternalTag)
		json.Unmarshal(ExternalTagString, &o.ExternalTag)
	}
	
	if WrapupId, ok := WorkitemdeltaMap["wrapupId"].(map[string]interface{}); ok {
		WrapupIdString, _ := json.Marshal(WrapupId)
		json.Unmarshal(WrapupIdString, &o.WrapupId)
	}
	
	if Ttl, ok := WorkitemdeltaMap["ttl"].(map[string]interface{}); ok {
		TtlString, _ := json.Marshal(Ttl)
		json.Unmarshal(TtlString, &o.Ttl)
	}
	
	if DateClosed, ok := WorkitemdeltaMap["dateClosed"].(map[string]interface{}); ok {
		DateClosedString, _ := json.Marshal(DateClosed)
		json.Unmarshal(DateClosedString, &o.DateClosed)
	}
	
	if AssignmentState, ok := WorkitemdeltaMap["assignmentState"].(map[string]interface{}); ok {
		AssignmentStateString, _ := json.Marshal(AssignmentState)
		json.Unmarshal(AssignmentStateString, &o.AssignmentState)
	}
	
	if AutoStatusTransition, ok := WorkitemdeltaMap["autoStatusTransition"].(map[string]interface{}); ok {
		AutoStatusTransitionString, _ := json.Marshal(AutoStatusTransition)
		json.Unmarshal(AutoStatusTransitionString, &o.AutoStatusTransition)
	}
	
	if CustomFields, ok := WorkitemdeltaMap["customFields"].(map[string]interface{}); ok {
		CustomFieldsString, _ := json.Marshal(CustomFields)
		json.Unmarshal(CustomFieldsString, &o.CustomFields)
	}
	
	if DateModified, ok := WorkitemdeltaMap["dateModified"].(map[string]interface{}); ok {
		DateModifiedString, _ := json.Marshal(DateModified)
		json.Unmarshal(DateModifiedString, &o.DateModified)
	}
	
	if ModifiedBy, ok := WorkitemdeltaMap["modifiedBy"].(map[string]interface{}); ok {
		ModifiedByString, _ := json.Marshal(ModifiedBy)
		json.Unmarshal(ModifiedByString, &o.ModifiedBy)
	}
	
	if StatusCategory, ok := WorkitemdeltaMap["statusCategory"].(map[string]interface{}); ok {
		StatusCategoryString, _ := json.Marshal(StatusCategory)
		json.Unmarshal(StatusCategoryString, &o.StatusCategory)
	}
	
	if ScriptId, ok := WorkitemdeltaMap["scriptId"].(map[string]interface{}); ok {
		ScriptIdString, _ := json.Marshal(ScriptId)
		json.Unmarshal(ScriptIdString, &o.ScriptId)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Workitemdelta) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

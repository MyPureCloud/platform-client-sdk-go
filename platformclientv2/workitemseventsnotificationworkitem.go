package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Workitemseventsnotificationworkitem
type Workitemseventsnotificationworkitem struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id
	Id *string `json:"id,omitempty"`

	// Name
	Name *string `json:"name,omitempty"`

	// TypeId
	TypeId *string `json:"typeId,omitempty"`

	// Description
	Description *string `json:"description,omitempty"`

	// LanguageId
	LanguageId *string `json:"languageId,omitempty"`

	// Priority
	Priority *int `json:"priority,omitempty"`

	// DateCreated
	DateCreated *string `json:"dateCreated,omitempty"`

	// DateModified
	DateModified *string `json:"dateModified,omitempty"`

	// DateDue
	DateDue *string `json:"dateDue,omitempty"`

	// DateExpires
	DateExpires *string `json:"dateExpires,omitempty"`

	// DateAssignmentStateChanged
	DateAssignmentStateChanged *string `json:"dateAssignmentStateChanged,omitempty"`

	// DurationSeconds
	DurationSeconds *int `json:"durationSeconds,omitempty"`

	// Ttl
	Ttl *int `json:"ttl,omitempty"`

	// StatusId
	StatusId *string `json:"statusId,omitempty"`

	// StatusCategory
	StatusCategory *string `json:"statusCategory,omitempty"`

	// DateClosed
	DateClosed *string `json:"dateClosed,omitempty"`

	// WorkbinId
	WorkbinId *string `json:"workbinId,omitempty"`

	// ReporterId
	ReporterId *string `json:"reporterId,omitempty"`

	// AssigneeId
	AssigneeId *string `json:"assigneeId,omitempty"`

	// ExternalContactId
	ExternalContactId *string `json:"externalContactId,omitempty"`

	// ExternalTag
	ExternalTag *string `json:"externalTag,omitempty"`

	// WrapupId
	WrapupId *string `json:"wrapupId,omitempty"`

	// ModifiedBy
	ModifiedBy *string `json:"modifiedBy,omitempty"`

	// Operation
	Operation *string `json:"operation,omitempty"`

	// Changes
	Changes *[]Workitemseventsnotificationdelta `json:"changes,omitempty"`

	// PropertyChanges
	PropertyChanges *[]Workitemseventsnotificationpropertychange `json:"propertyChanges,omitempty"`

	// AssignmentState
	AssignmentState *string `json:"assignmentState,omitempty"`

	// AssignmentId
	AssignmentId *string `json:"assignmentId,omitempty"`

	// AlertTimeoutSeconds
	AlertTimeoutSeconds *int `json:"alertTimeoutSeconds,omitempty"`

	// QueueId
	QueueId *string `json:"queueId,omitempty"`

	// CustomFields
	CustomFields *map[string]Workitemseventsnotificationcustomattribute `json:"customFields,omitempty"`

	// Wrapup
	Wrapup *Workitemseventsnotificationwrapup `json:"wrapup,omitempty"`

	// Sessions
	Sessions *[]Workitemseventsnotificationsession `json:"sessions,omitempty"`

	// SkillIds
	SkillIds *[]string `json:"skillIds,omitempty"`

	// ScriptId
	ScriptId *string `json:"scriptId,omitempty"`

	// WorkbinName
	WorkbinName *string `json:"workbinName,omitempty"`

	// TypeName
	TypeName *string `json:"typeName,omitempty"`

	// PreferredAgentIds
	PreferredAgentIds *[]string `json:"preferredAgentIds,omitempty"`

	// DivisionId
	DivisionId *string `json:"divisionId,omitempty"`

	// ScoredAgents
	ScoredAgents *[]Workitemseventsnotificationscoredagent `json:"scoredAgents,omitempty"`

	// UtilizationLabelId
	UtilizationLabelId *string `json:"utilizationLabelId,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Workitemseventsnotificationworkitem) SetField(field string, fieldValue interface{}) {
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

func (o Workitemseventsnotificationworkitem) MarshalJSON() ([]byte, error) {
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
	type Alias Workitemseventsnotificationworkitem
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		TypeId *string `json:"typeId,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		LanguageId *string `json:"languageId,omitempty"`
		
		Priority *int `json:"priority,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		DateDue *string `json:"dateDue,omitempty"`
		
		DateExpires *string `json:"dateExpires,omitempty"`
		
		DateAssignmentStateChanged *string `json:"dateAssignmentStateChanged,omitempty"`
		
		DurationSeconds *int `json:"durationSeconds,omitempty"`
		
		Ttl *int `json:"ttl,omitempty"`
		
		StatusId *string `json:"statusId,omitempty"`
		
		StatusCategory *string `json:"statusCategory,omitempty"`
		
		DateClosed *string `json:"dateClosed,omitempty"`
		
		WorkbinId *string `json:"workbinId,omitempty"`
		
		ReporterId *string `json:"reporterId,omitempty"`
		
		AssigneeId *string `json:"assigneeId,omitempty"`
		
		ExternalContactId *string `json:"externalContactId,omitempty"`
		
		ExternalTag *string `json:"externalTag,omitempty"`
		
		WrapupId *string `json:"wrapupId,omitempty"`
		
		ModifiedBy *string `json:"modifiedBy,omitempty"`
		
		Operation *string `json:"operation,omitempty"`
		
		Changes *[]Workitemseventsnotificationdelta `json:"changes,omitempty"`
		
		PropertyChanges *[]Workitemseventsnotificationpropertychange `json:"propertyChanges,omitempty"`
		
		AssignmentState *string `json:"assignmentState,omitempty"`
		
		AssignmentId *string `json:"assignmentId,omitempty"`
		
		AlertTimeoutSeconds *int `json:"alertTimeoutSeconds,omitempty"`
		
		QueueId *string `json:"queueId,omitempty"`
		
		CustomFields *map[string]Workitemseventsnotificationcustomattribute `json:"customFields,omitempty"`
		
		Wrapup *Workitemseventsnotificationwrapup `json:"wrapup,omitempty"`
		
		Sessions *[]Workitemseventsnotificationsession `json:"sessions,omitempty"`
		
		SkillIds *[]string `json:"skillIds,omitempty"`
		
		ScriptId *string `json:"scriptId,omitempty"`
		
		WorkbinName *string `json:"workbinName,omitempty"`
		
		TypeName *string `json:"typeName,omitempty"`
		
		PreferredAgentIds *[]string `json:"preferredAgentIds,omitempty"`
		
		DivisionId *string `json:"divisionId,omitempty"`
		
		ScoredAgents *[]Workitemseventsnotificationscoredagent `json:"scoredAgents,omitempty"`
		
		UtilizationLabelId *string `json:"utilizationLabelId,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		TypeId: o.TypeId,
		
		Description: o.Description,
		
		LanguageId: o.LanguageId,
		
		Priority: o.Priority,
		
		DateCreated: o.DateCreated,
		
		DateModified: o.DateModified,
		
		DateDue: o.DateDue,
		
		DateExpires: o.DateExpires,
		
		DateAssignmentStateChanged: o.DateAssignmentStateChanged,
		
		DurationSeconds: o.DurationSeconds,
		
		Ttl: o.Ttl,
		
		StatusId: o.StatusId,
		
		StatusCategory: o.StatusCategory,
		
		DateClosed: o.DateClosed,
		
		WorkbinId: o.WorkbinId,
		
		ReporterId: o.ReporterId,
		
		AssigneeId: o.AssigneeId,
		
		ExternalContactId: o.ExternalContactId,
		
		ExternalTag: o.ExternalTag,
		
		WrapupId: o.WrapupId,
		
		ModifiedBy: o.ModifiedBy,
		
		Operation: o.Operation,
		
		Changes: o.Changes,
		
		PropertyChanges: o.PropertyChanges,
		
		AssignmentState: o.AssignmentState,
		
		AssignmentId: o.AssignmentId,
		
		AlertTimeoutSeconds: o.AlertTimeoutSeconds,
		
		QueueId: o.QueueId,
		
		CustomFields: o.CustomFields,
		
		Wrapup: o.Wrapup,
		
		Sessions: o.Sessions,
		
		SkillIds: o.SkillIds,
		
		ScriptId: o.ScriptId,
		
		WorkbinName: o.WorkbinName,
		
		TypeName: o.TypeName,
		
		PreferredAgentIds: o.PreferredAgentIds,
		
		DivisionId: o.DivisionId,
		
		ScoredAgents: o.ScoredAgents,
		
		UtilizationLabelId: o.UtilizationLabelId,
		Alias:    (Alias)(o),
	})
}

func (o *Workitemseventsnotificationworkitem) UnmarshalJSON(b []byte) error {
	var WorkitemseventsnotificationworkitemMap map[string]interface{}
	err := json.Unmarshal(b, &WorkitemseventsnotificationworkitemMap)
	if err != nil {
		return err
	}
	
	if Id, ok := WorkitemseventsnotificationworkitemMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := WorkitemseventsnotificationworkitemMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if TypeId, ok := WorkitemseventsnotificationworkitemMap["typeId"].(string); ok {
		o.TypeId = &TypeId
	}
    
	if Description, ok := WorkitemseventsnotificationworkitemMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if LanguageId, ok := WorkitemseventsnotificationworkitemMap["languageId"].(string); ok {
		o.LanguageId = &LanguageId
	}
    
	if Priority, ok := WorkitemseventsnotificationworkitemMap["priority"].(float64); ok {
		PriorityInt := int(Priority)
		o.Priority = &PriorityInt
	}
	
	if DateCreated, ok := WorkitemseventsnotificationworkitemMap["dateCreated"].(string); ok {
		o.DateCreated = &DateCreated
	}
    
	if DateModified, ok := WorkitemseventsnotificationworkitemMap["dateModified"].(string); ok {
		o.DateModified = &DateModified
	}
    
	if DateDue, ok := WorkitemseventsnotificationworkitemMap["dateDue"].(string); ok {
		o.DateDue = &DateDue
	}
    
	if DateExpires, ok := WorkitemseventsnotificationworkitemMap["dateExpires"].(string); ok {
		o.DateExpires = &DateExpires
	}
    
	if DateAssignmentStateChanged, ok := WorkitemseventsnotificationworkitemMap["dateAssignmentStateChanged"].(string); ok {
		o.DateAssignmentStateChanged = &DateAssignmentStateChanged
	}
    
	if DurationSeconds, ok := WorkitemseventsnotificationworkitemMap["durationSeconds"].(float64); ok {
		DurationSecondsInt := int(DurationSeconds)
		o.DurationSeconds = &DurationSecondsInt
	}
	
	if Ttl, ok := WorkitemseventsnotificationworkitemMap["ttl"].(float64); ok {
		TtlInt := int(Ttl)
		o.Ttl = &TtlInt
	}
	
	if StatusId, ok := WorkitemseventsnotificationworkitemMap["statusId"].(string); ok {
		o.StatusId = &StatusId
	}
    
	if StatusCategory, ok := WorkitemseventsnotificationworkitemMap["statusCategory"].(string); ok {
		o.StatusCategory = &StatusCategory
	}
    
	if DateClosed, ok := WorkitemseventsnotificationworkitemMap["dateClosed"].(string); ok {
		o.DateClosed = &DateClosed
	}
    
	if WorkbinId, ok := WorkitemseventsnotificationworkitemMap["workbinId"].(string); ok {
		o.WorkbinId = &WorkbinId
	}
    
	if ReporterId, ok := WorkitemseventsnotificationworkitemMap["reporterId"].(string); ok {
		o.ReporterId = &ReporterId
	}
    
	if AssigneeId, ok := WorkitemseventsnotificationworkitemMap["assigneeId"].(string); ok {
		o.AssigneeId = &AssigneeId
	}
    
	if ExternalContactId, ok := WorkitemseventsnotificationworkitemMap["externalContactId"].(string); ok {
		o.ExternalContactId = &ExternalContactId
	}
    
	if ExternalTag, ok := WorkitemseventsnotificationworkitemMap["externalTag"].(string); ok {
		o.ExternalTag = &ExternalTag
	}
    
	if WrapupId, ok := WorkitemseventsnotificationworkitemMap["wrapupId"].(string); ok {
		o.WrapupId = &WrapupId
	}
    
	if ModifiedBy, ok := WorkitemseventsnotificationworkitemMap["modifiedBy"].(string); ok {
		o.ModifiedBy = &ModifiedBy
	}
    
	if Operation, ok := WorkitemseventsnotificationworkitemMap["operation"].(string); ok {
		o.Operation = &Operation
	}
    
	if Changes, ok := WorkitemseventsnotificationworkitemMap["changes"].([]interface{}); ok {
		ChangesString, _ := json.Marshal(Changes)
		json.Unmarshal(ChangesString, &o.Changes)
	}
	
	if PropertyChanges, ok := WorkitemseventsnotificationworkitemMap["propertyChanges"].([]interface{}); ok {
		PropertyChangesString, _ := json.Marshal(PropertyChanges)
		json.Unmarshal(PropertyChangesString, &o.PropertyChanges)
	}
	
	if AssignmentState, ok := WorkitemseventsnotificationworkitemMap["assignmentState"].(string); ok {
		o.AssignmentState = &AssignmentState
	}
    
	if AssignmentId, ok := WorkitemseventsnotificationworkitemMap["assignmentId"].(string); ok {
		o.AssignmentId = &AssignmentId
	}
    
	if AlertTimeoutSeconds, ok := WorkitemseventsnotificationworkitemMap["alertTimeoutSeconds"].(float64); ok {
		AlertTimeoutSecondsInt := int(AlertTimeoutSeconds)
		o.AlertTimeoutSeconds = &AlertTimeoutSecondsInt
	}
	
	if QueueId, ok := WorkitemseventsnotificationworkitemMap["queueId"].(string); ok {
		o.QueueId = &QueueId
	}
    
	if CustomFields, ok := WorkitemseventsnotificationworkitemMap["customFields"].(map[string]interface{}); ok {
		CustomFieldsString, _ := json.Marshal(CustomFields)
		json.Unmarshal(CustomFieldsString, &o.CustomFields)
	}
	
	if Wrapup, ok := WorkitemseventsnotificationworkitemMap["wrapup"].(map[string]interface{}); ok {
		WrapupString, _ := json.Marshal(Wrapup)
		json.Unmarshal(WrapupString, &o.Wrapup)
	}
	
	if Sessions, ok := WorkitemseventsnotificationworkitemMap["sessions"].([]interface{}); ok {
		SessionsString, _ := json.Marshal(Sessions)
		json.Unmarshal(SessionsString, &o.Sessions)
	}
	
	if SkillIds, ok := WorkitemseventsnotificationworkitemMap["skillIds"].([]interface{}); ok {
		SkillIdsString, _ := json.Marshal(SkillIds)
		json.Unmarshal(SkillIdsString, &o.SkillIds)
	}
	
	if ScriptId, ok := WorkitemseventsnotificationworkitemMap["scriptId"].(string); ok {
		o.ScriptId = &ScriptId
	}
    
	if WorkbinName, ok := WorkitemseventsnotificationworkitemMap["workbinName"].(string); ok {
		o.WorkbinName = &WorkbinName
	}
    
	if TypeName, ok := WorkitemseventsnotificationworkitemMap["typeName"].(string); ok {
		o.TypeName = &TypeName
	}
    
	if PreferredAgentIds, ok := WorkitemseventsnotificationworkitemMap["preferredAgentIds"].([]interface{}); ok {
		PreferredAgentIdsString, _ := json.Marshal(PreferredAgentIds)
		json.Unmarshal(PreferredAgentIdsString, &o.PreferredAgentIds)
	}
	
	if DivisionId, ok := WorkitemseventsnotificationworkitemMap["divisionId"].(string); ok {
		o.DivisionId = &DivisionId
	}
    
	if ScoredAgents, ok := WorkitemseventsnotificationworkitemMap["scoredAgents"].([]interface{}); ok {
		ScoredAgentsString, _ := json.Marshal(ScoredAgents)
		json.Unmarshal(ScoredAgentsString, &o.ScoredAgents)
	}
	
	if UtilizationLabelId, ok := WorkitemseventsnotificationworkitemMap["utilizationLabelId"].(string); ok {
		o.UtilizationLabelId = &UtilizationLabelId
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Workitemseventsnotificationworkitem) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

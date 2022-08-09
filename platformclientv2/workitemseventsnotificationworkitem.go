package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Workitemseventsnotificationworkitem
type Workitemseventsnotificationworkitem struct { 
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


	// DurationSeconds
	DurationSeconds *int `json:"durationSeconds,omitempty"`


	// Ttl
	Ttl *int `json:"ttl,omitempty"`


	// StatusId
	StatusId *string `json:"statusId,omitempty"`


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


	// AssignmentState
	AssignmentState *string `json:"assignmentState,omitempty"`


	// AssignmentId
	AssignmentId *string `json:"assignmentId,omitempty"`


	// AlertTimeoutSeconds
	AlertTimeoutSeconds *int `json:"alertTimeoutSeconds,omitempty"`


	// CustomFields
	CustomFields *map[string]Workitemseventsnotificationcustomattribute `json:"customFields,omitempty"`

}

func (o *Workitemseventsnotificationworkitem) MarshalJSON() ([]byte, error) {
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
		
		DurationSeconds *int `json:"durationSeconds,omitempty"`
		
		Ttl *int `json:"ttl,omitempty"`
		
		StatusId *string `json:"statusId,omitempty"`
		
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
		
		AssignmentState *string `json:"assignmentState,omitempty"`
		
		AssignmentId *string `json:"assignmentId,omitempty"`
		
		AlertTimeoutSeconds *int `json:"alertTimeoutSeconds,omitempty"`
		
		CustomFields *map[string]Workitemseventsnotificationcustomattribute `json:"customFields,omitempty"`
		*Alias
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
		
		DurationSeconds: o.DurationSeconds,
		
		Ttl: o.Ttl,
		
		StatusId: o.StatusId,
		
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
		
		AssignmentState: o.AssignmentState,
		
		AssignmentId: o.AssignmentId,
		
		AlertTimeoutSeconds: o.AlertTimeoutSeconds,
		
		CustomFields: o.CustomFields,
		Alias:    (*Alias)(o),
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
	
	if CustomFields, ok := WorkitemseventsnotificationworkitemMap["customFields"].(map[string]interface{}); ok {
		CustomFieldsString, _ := json.Marshal(CustomFields)
		json.Unmarshal(CustomFieldsString, &o.CustomFields)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Workitemseventsnotificationworkitem) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

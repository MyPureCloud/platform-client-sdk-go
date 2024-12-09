package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Workitemcreate
type Workitemcreate struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Name - The name of the Workitem. Valid length between 3 and 256 characters.
	Name *string `json:"name,omitempty"`

	// Priority - The priority of the Workitem. The valid range is between -25,000,000 and 25,000,000.
	Priority *int `json:"priority,omitempty"`

	// DateDue - The due date of the Workitem. Can not be greater than 365 days from the current time. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateDue *time.Time `json:"dateDue,omitempty"`

	// DateExpires - The expiry date of the Workitem. Can not be greater than 365 days from the current time. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateExpires *time.Time `json:"dateExpires,omitempty"`

	// DurationSeconds - The estimated duration in seconds to complete the Workitem. Maximum of 365 days.
	DurationSeconds *int `json:"durationSeconds,omitempty"`

	// Ttl - The epoch timestamp in seconds specifying the time to live for the Workitem. Can not be greater than 365 days from the current time.
	Ttl *int `json:"ttl,omitempty"`

	// StatusId - The ID of the Status of the Workitem.
	StatusId *string `json:"statusId,omitempty"`

	// WorkbinId - The ID of Workbin that contains the Workitem.
	WorkbinId *string `json:"workbinId,omitempty"`

	// AutoStatusTransition - Set it to false to disable auto status transition. By default, it is enabled.
	AutoStatusTransition *bool `json:"autoStatusTransition,omitempty"`

	// Description - The description of the Workitem. Maximum length of 512 characters.
	Description *string `json:"description,omitempty"`

	// TypeId - The ID of the Worktype of the Workitem.
	TypeId *string `json:"typeId,omitempty"`

	// CustomFields - Custom fields defined in the schema referenced by the worktype of the workitem.
	CustomFields *map[string]interface{} `json:"customFields,omitempty"`

	// QueueId - The ID of the Workitems queue. Must be a valid UUID.
	QueueId *string `json:"queueId,omitempty"`

	// AssigneeId - The ID of the assignee of the Workitem. Must be a valid UUID.
	AssigneeId *string `json:"assigneeId,omitempty"`

	// ScoredAgents - A list of scored agents for the Workitem. A workitem can have a maximum of 20 scored agents.
	ScoredAgents *[]Workitemscoredagentrequest `json:"scoredAgents,omitempty"`

	// LanguageId - The ID of language of the Workitem. Must be a valid UUID.
	LanguageId *string `json:"languageId,omitempty"`

	// ExternalContactId - The ID of the external contact of the Workitem. Must be a valid UUID.
	ExternalContactId *string `json:"externalContactId,omitempty"`

	// ExternalTag - The external tag of the Workitem.
	ExternalTag *string `json:"externalTag,omitempty"`

	// SkillIds - The skill IDs of the Workitem. Must be valid UUIDs.
	SkillIds *[]string `json:"skillIds,omitempty"`

	// PreferredAgentIds - The preferred agent IDs of the Workitem. Must be valid UUIDs.
	PreferredAgentIds *[]string `json:"preferredAgentIds,omitempty"`

	// WrapupCode - The ID of the wrapup. Must be a valid UUID.
	WrapupCode *string `json:"wrapupCode,omitempty"`

	// UtilizationLabelId - The ID of utilization label of the Workitem. Must be a valid UUID.
	UtilizationLabelId *string `json:"utilizationLabelId,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Workitemcreate) SetField(field string, fieldValue interface{}) {
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

func (o Workitemcreate) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateDue","DateExpires", }
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
	type Alias Workitemcreate
	
	DateDue := new(string)
	if o.DateDue != nil {
		
		*DateDue = timeutil.Strftime(o.DateDue, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateDue = nil
	}
	
	DateExpires := new(string)
	if o.DateExpires != nil {
		
		*DateExpires = timeutil.Strftime(o.DateExpires, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateExpires = nil
	}
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		Priority *int `json:"priority,omitempty"`
		
		DateDue *string `json:"dateDue,omitempty"`
		
		DateExpires *string `json:"dateExpires,omitempty"`
		
		DurationSeconds *int `json:"durationSeconds,omitempty"`
		
		Ttl *int `json:"ttl,omitempty"`
		
		StatusId *string `json:"statusId,omitempty"`
		
		WorkbinId *string `json:"workbinId,omitempty"`
		
		AutoStatusTransition *bool `json:"autoStatusTransition,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		TypeId *string `json:"typeId,omitempty"`
		
		CustomFields *map[string]interface{} `json:"customFields,omitempty"`
		
		QueueId *string `json:"queueId,omitempty"`
		
		AssigneeId *string `json:"assigneeId,omitempty"`
		
		ScoredAgents *[]Workitemscoredagentrequest `json:"scoredAgents,omitempty"`
		
		LanguageId *string `json:"languageId,omitempty"`
		
		ExternalContactId *string `json:"externalContactId,omitempty"`
		
		ExternalTag *string `json:"externalTag,omitempty"`
		
		SkillIds *[]string `json:"skillIds,omitempty"`
		
		PreferredAgentIds *[]string `json:"preferredAgentIds,omitempty"`
		
		WrapupCode *string `json:"wrapupCode,omitempty"`
		
		UtilizationLabelId *string `json:"utilizationLabelId,omitempty"`
		Alias
	}{ 
		Name: o.Name,
		
		Priority: o.Priority,
		
		DateDue: DateDue,
		
		DateExpires: DateExpires,
		
		DurationSeconds: o.DurationSeconds,
		
		Ttl: o.Ttl,
		
		StatusId: o.StatusId,
		
		WorkbinId: o.WorkbinId,
		
		AutoStatusTransition: o.AutoStatusTransition,
		
		Description: o.Description,
		
		TypeId: o.TypeId,
		
		CustomFields: o.CustomFields,
		
		QueueId: o.QueueId,
		
		AssigneeId: o.AssigneeId,
		
		ScoredAgents: o.ScoredAgents,
		
		LanguageId: o.LanguageId,
		
		ExternalContactId: o.ExternalContactId,
		
		ExternalTag: o.ExternalTag,
		
		SkillIds: o.SkillIds,
		
		PreferredAgentIds: o.PreferredAgentIds,
		
		WrapupCode: o.WrapupCode,
		
		UtilizationLabelId: o.UtilizationLabelId,
		Alias:    (Alias)(o),
	})
}

func (o *Workitemcreate) UnmarshalJSON(b []byte) error {
	var WorkitemcreateMap map[string]interface{}
	err := json.Unmarshal(b, &WorkitemcreateMap)
	if err != nil {
		return err
	}
	
	if Name, ok := WorkitemcreateMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Priority, ok := WorkitemcreateMap["priority"].(float64); ok {
		PriorityInt := int(Priority)
		o.Priority = &PriorityInt
	}
	
	if dateDueString, ok := WorkitemcreateMap["dateDue"].(string); ok {
		DateDue, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateDueString)
		o.DateDue = &DateDue
	}
	
	if dateExpiresString, ok := WorkitemcreateMap["dateExpires"].(string); ok {
		DateExpires, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateExpiresString)
		o.DateExpires = &DateExpires
	}
	
	if DurationSeconds, ok := WorkitemcreateMap["durationSeconds"].(float64); ok {
		DurationSecondsInt := int(DurationSeconds)
		o.DurationSeconds = &DurationSecondsInt
	}
	
	if Ttl, ok := WorkitemcreateMap["ttl"].(float64); ok {
		TtlInt := int(Ttl)
		o.Ttl = &TtlInt
	}
	
	if StatusId, ok := WorkitemcreateMap["statusId"].(string); ok {
		o.StatusId = &StatusId
	}
    
	if WorkbinId, ok := WorkitemcreateMap["workbinId"].(string); ok {
		o.WorkbinId = &WorkbinId
	}
    
	if AutoStatusTransition, ok := WorkitemcreateMap["autoStatusTransition"].(bool); ok {
		o.AutoStatusTransition = &AutoStatusTransition
	}
    
	if Description, ok := WorkitemcreateMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if TypeId, ok := WorkitemcreateMap["typeId"].(string); ok {
		o.TypeId = &TypeId
	}
    
	if CustomFields, ok := WorkitemcreateMap["customFields"].(map[string]interface{}); ok {
		CustomFieldsString, _ := json.Marshal(CustomFields)
		json.Unmarshal(CustomFieldsString, &o.CustomFields)
	}
	
	if QueueId, ok := WorkitemcreateMap["queueId"].(string); ok {
		o.QueueId = &QueueId
	}
    
	if AssigneeId, ok := WorkitemcreateMap["assigneeId"].(string); ok {
		o.AssigneeId = &AssigneeId
	}
    
	if ScoredAgents, ok := WorkitemcreateMap["scoredAgents"].([]interface{}); ok {
		ScoredAgentsString, _ := json.Marshal(ScoredAgents)
		json.Unmarshal(ScoredAgentsString, &o.ScoredAgents)
	}
	
	if LanguageId, ok := WorkitemcreateMap["languageId"].(string); ok {
		o.LanguageId = &LanguageId
	}
    
	if ExternalContactId, ok := WorkitemcreateMap["externalContactId"].(string); ok {
		o.ExternalContactId = &ExternalContactId
	}
    
	if ExternalTag, ok := WorkitemcreateMap["externalTag"].(string); ok {
		o.ExternalTag = &ExternalTag
	}
    
	if SkillIds, ok := WorkitemcreateMap["skillIds"].([]interface{}); ok {
		SkillIdsString, _ := json.Marshal(SkillIds)
		json.Unmarshal(SkillIdsString, &o.SkillIds)
	}
	
	if PreferredAgentIds, ok := WorkitemcreateMap["preferredAgentIds"].([]interface{}); ok {
		PreferredAgentIdsString, _ := json.Marshal(PreferredAgentIds)
		json.Unmarshal(PreferredAgentIdsString, &o.PreferredAgentIds)
	}
	
	if WrapupCode, ok := WorkitemcreateMap["wrapupCode"].(string); ok {
		o.WrapupCode = &WrapupCode
	}
    
	if UtilizationLabelId, ok := WorkitemcreateMap["utilizationLabelId"].(string); ok {
		o.UtilizationLabelId = &UtilizationLabelId
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Workitemcreate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Workitemupdate
type Workitemupdate struct { 
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

	// Description - The description of the Workitem. Maximum length of 4096 characters.
	Description *string `json:"description,omitempty"`

	// DateClosed - The closed date of the Workitem. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateClosed *time.Time `json:"dateClosed,omitempty"`

	// AssignmentState - The assignment state of the Workitem.
	AssignmentState *string `json:"assignmentState,omitempty"`

	// AssignmentOperation - Set this value to AgentAssignmentAlerting and supply an 'assigneeId' to assign the workitem to an agent and alert the agent of the assignment. Set this value to QueueAssignmentAlerting and supply a 'queueId' to route the workitem to an agent who is a member of the queue and alert the agent.
	AssignmentOperation *string `json:"assignmentOperation,omitempty"`

	// CustomFields - Custom fields defined in the schema referenced by the worktype of the workitem. If set to {}, the existing keys and values will be removed.
	CustomFields *map[string]interface{} `json:"customFields,omitempty"`

	// QueueId - The ID of the Workitems queue. Must be a valid UUID.
	QueueId *string `json:"queueId,omitempty"`

	// AssigneeId - The ID of the assignee of the Workitem. If supplied it must be a valid UUID.
	AssigneeId *string `json:"assigneeId,omitempty"`

	// ScoredAgents - A list of scored agents for the Workitem. A workitem can have a maximum of 20 scored agents.
	ScoredAgents *[]Workitemscoredagentrequest `json:"scoredAgents,omitempty"`

	// ExternalContactId - The ID of the external contact of the Workitem. Must be a valid UUID.
	ExternalContactId *string `json:"externalContactId,omitempty"`

	// ExternalTag - The external tag of the Workitem.
	ExternalTag *string `json:"externalTag,omitempty"`

	// SkillIds - The skill IDs of the Workitem. Must be valid UUIDs.
	SkillIds *[]string `json:"skillIds,omitempty"`

	// LanguageId - The ID of language of the Workitem. Must be a valid UUID.
	LanguageId *string `json:"languageId,omitempty"`

	// UtilizationLabelId - The ID of the utilization label of the Workitem. Must be a valid UUID.
	UtilizationLabelId *string `json:"utilizationLabelId,omitempty"`

	// PreferredAgentIds - The preferred agent IDs of the Workitem. Must be valid UUIDs.
	PreferredAgentIds *[]string `json:"preferredAgentIds,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Workitemupdate) SetField(field string, fieldValue interface{}) {
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

func (o Workitemupdate) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateDue","DateExpires","DateClosed", }
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
	type Alias Workitemupdate
	
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
	
	DateClosed := new(string)
	if o.DateClosed != nil {
		
		*DateClosed = timeutil.Strftime(o.DateClosed, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateClosed = nil
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
		
		DateClosed *string `json:"dateClosed,omitempty"`
		
		AssignmentState *string `json:"assignmentState,omitempty"`
		
		AssignmentOperation *string `json:"assignmentOperation,omitempty"`
		
		CustomFields *map[string]interface{} `json:"customFields,omitempty"`
		
		QueueId *string `json:"queueId,omitempty"`
		
		AssigneeId *string `json:"assigneeId,omitempty"`
		
		ScoredAgents *[]Workitemscoredagentrequest `json:"scoredAgents,omitempty"`
		
		ExternalContactId *string `json:"externalContactId,omitempty"`
		
		ExternalTag *string `json:"externalTag,omitempty"`
		
		SkillIds *[]string `json:"skillIds,omitempty"`
		
		LanguageId *string `json:"languageId,omitempty"`
		
		UtilizationLabelId *string `json:"utilizationLabelId,omitempty"`
		
		PreferredAgentIds *[]string `json:"preferredAgentIds,omitempty"`
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
		
		DateClosed: DateClosed,
		
		AssignmentState: o.AssignmentState,
		
		AssignmentOperation: o.AssignmentOperation,
		
		CustomFields: o.CustomFields,
		
		QueueId: o.QueueId,
		
		AssigneeId: o.AssigneeId,
		
		ScoredAgents: o.ScoredAgents,
		
		ExternalContactId: o.ExternalContactId,
		
		ExternalTag: o.ExternalTag,
		
		SkillIds: o.SkillIds,
		
		LanguageId: o.LanguageId,
		
		UtilizationLabelId: o.UtilizationLabelId,
		
		PreferredAgentIds: o.PreferredAgentIds,
		Alias:    (Alias)(o),
	})
}

func (o *Workitemupdate) UnmarshalJSON(b []byte) error {
	var WorkitemupdateMap map[string]interface{}
	err := json.Unmarshal(b, &WorkitemupdateMap)
	if err != nil {
		return err
	}
	
	if Name, ok := WorkitemupdateMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Priority, ok := WorkitemupdateMap["priority"].(float64); ok {
		PriorityInt := int(Priority)
		o.Priority = &PriorityInt
	}
	
	if dateDueString, ok := WorkitemupdateMap["dateDue"].(string); ok {
		DateDue, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateDueString)
		o.DateDue = &DateDue
	}
	
	if dateExpiresString, ok := WorkitemupdateMap["dateExpires"].(string); ok {
		DateExpires, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateExpiresString)
		o.DateExpires = &DateExpires
	}
	
	if DurationSeconds, ok := WorkitemupdateMap["durationSeconds"].(float64); ok {
		DurationSecondsInt := int(DurationSeconds)
		o.DurationSeconds = &DurationSecondsInt
	}
	
	if Ttl, ok := WorkitemupdateMap["ttl"].(float64); ok {
		TtlInt := int(Ttl)
		o.Ttl = &TtlInt
	}
	
	if StatusId, ok := WorkitemupdateMap["statusId"].(string); ok {
		o.StatusId = &StatusId
	}
    
	if WorkbinId, ok := WorkitemupdateMap["workbinId"].(string); ok {
		o.WorkbinId = &WorkbinId
	}
    
	if AutoStatusTransition, ok := WorkitemupdateMap["autoStatusTransition"].(bool); ok {
		o.AutoStatusTransition = &AutoStatusTransition
	}
    
	if Description, ok := WorkitemupdateMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if dateClosedString, ok := WorkitemupdateMap["dateClosed"].(string); ok {
		DateClosed, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateClosedString)
		o.DateClosed = &DateClosed
	}
	
	if AssignmentState, ok := WorkitemupdateMap["assignmentState"].(string); ok {
		o.AssignmentState = &AssignmentState
	}
    
	if AssignmentOperation, ok := WorkitemupdateMap["assignmentOperation"].(string); ok {
		o.AssignmentOperation = &AssignmentOperation
	}
    
	if CustomFields, ok := WorkitemupdateMap["customFields"].(map[string]interface{}); ok {
		CustomFieldsString, _ := json.Marshal(CustomFields)
		json.Unmarshal(CustomFieldsString, &o.CustomFields)
	}
	
	if QueueId, ok := WorkitemupdateMap["queueId"].(string); ok {
		o.QueueId = &QueueId
	}
    
	if AssigneeId, ok := WorkitemupdateMap["assigneeId"].(string); ok {
		o.AssigneeId = &AssigneeId
	}
    
	if ScoredAgents, ok := WorkitemupdateMap["scoredAgents"].([]interface{}); ok {
		ScoredAgentsString, _ := json.Marshal(ScoredAgents)
		json.Unmarshal(ScoredAgentsString, &o.ScoredAgents)
	}
	
	if ExternalContactId, ok := WorkitemupdateMap["externalContactId"].(string); ok {
		o.ExternalContactId = &ExternalContactId
	}
    
	if ExternalTag, ok := WorkitemupdateMap["externalTag"].(string); ok {
		o.ExternalTag = &ExternalTag
	}
    
	if SkillIds, ok := WorkitemupdateMap["skillIds"].([]interface{}); ok {
		SkillIdsString, _ := json.Marshal(SkillIds)
		json.Unmarshal(SkillIdsString, &o.SkillIds)
	}
	
	if LanguageId, ok := WorkitemupdateMap["languageId"].(string); ok {
		o.LanguageId = &LanguageId
	}
    
	if UtilizationLabelId, ok := WorkitemupdateMap["utilizationLabelId"].(string); ok {
		o.UtilizationLabelId = &UtilizationLabelId
	}
    
	if PreferredAgentIds, ok := WorkitemupdateMap["preferredAgentIds"].([]interface{}); ok {
		PreferredAgentIdsString, _ := json.Marshal(PreferredAgentIds)
		json.Unmarshal(PreferredAgentIdsString, &o.PreferredAgentIds)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Workitemupdate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

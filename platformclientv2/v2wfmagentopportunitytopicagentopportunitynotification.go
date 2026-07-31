package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// V2wfmagentopportunitytopicagentopportunitynotification
type V2wfmagentopportunitytopicagentopportunitynotification struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id
	Id *string `json:"id,omitempty"`

	// BusinessUnitId
	BusinessUnitId *string `json:"businessUnitId,omitempty"`

	// OrganizationId
	OrganizationId *string `json:"organizationId,omitempty"`

	// AgentIds
	AgentIds *[]string `json:"agentIds,omitempty"`

	// OpportunityStartDate
	OpportunityStartDate *string `json:"opportunityStartDate,omitempty"`

	// OpportunityEndDate
	OpportunityEndDate *string `json:"opportunityEndDate,omitempty"`

	// OpenDate
	OpenDate *string `json:"openDate,omitempty"`

	// ApprovalType
	ApprovalType *string `json:"approvalType,omitempty"`

	// LengthMinutes
	LengthMinutes *int `json:"lengthMinutes,omitempty"`

	// DeadlineDate
	DeadlineDate *string `json:"deadlineDate,omitempty"`

	// ActivityCodeId
	ActivityCodeId *string `json:"activityCodeId,omitempty"`

	// Name
	Name *string `json:"name,omitempty"`

	// Description
	Description *string `json:"description,omitempty"`

	// ClosedDate
	ClosedDate *string `json:"closedDate,omitempty"`

	// Status
	Status *string `json:"status,omitempty"`

	// EventType
	EventType *string `json:"eventType,omitempty"`

	// DenialCode
	DenialCode *string `json:"denialCode,omitempty"`

	// ReviewNote
	ReviewNote *string `json:"reviewNote,omitempty"`

	// RemainingSpaces
	RemainingSpaces *int `json:"remainingSpaces,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *V2wfmagentopportunitytopicagentopportunitynotification) SetField(field string, fieldValue interface{}) {
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

func (o V2wfmagentopportunitytopicagentopportunitynotification) MarshalJSON() ([]byte, error) {
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
	type Alias V2wfmagentopportunitytopicagentopportunitynotification
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		BusinessUnitId *string `json:"businessUnitId,omitempty"`
		
		OrganizationId *string `json:"organizationId,omitempty"`
		
		AgentIds *[]string `json:"agentIds,omitempty"`
		
		OpportunityStartDate *string `json:"opportunityStartDate,omitempty"`
		
		OpportunityEndDate *string `json:"opportunityEndDate,omitempty"`
		
		OpenDate *string `json:"openDate,omitempty"`
		
		ApprovalType *string `json:"approvalType,omitempty"`
		
		LengthMinutes *int `json:"lengthMinutes,omitempty"`
		
		DeadlineDate *string `json:"deadlineDate,omitempty"`
		
		ActivityCodeId *string `json:"activityCodeId,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		ClosedDate *string `json:"closedDate,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		EventType *string `json:"eventType,omitempty"`
		
		DenialCode *string `json:"denialCode,omitempty"`
		
		ReviewNote *string `json:"reviewNote,omitempty"`
		
		RemainingSpaces *int `json:"remainingSpaces,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		BusinessUnitId: o.BusinessUnitId,
		
		OrganizationId: o.OrganizationId,
		
		AgentIds: o.AgentIds,
		
		OpportunityStartDate: o.OpportunityStartDate,
		
		OpportunityEndDate: o.OpportunityEndDate,
		
		OpenDate: o.OpenDate,
		
		ApprovalType: o.ApprovalType,
		
		LengthMinutes: o.LengthMinutes,
		
		DeadlineDate: o.DeadlineDate,
		
		ActivityCodeId: o.ActivityCodeId,
		
		Name: o.Name,
		
		Description: o.Description,
		
		ClosedDate: o.ClosedDate,
		
		Status: o.Status,
		
		EventType: o.EventType,
		
		DenialCode: o.DenialCode,
		
		ReviewNote: o.ReviewNote,
		
		RemainingSpaces: o.RemainingSpaces,
		Alias:    (Alias)(o),
	})
}

func (o *V2wfmagentopportunitytopicagentopportunitynotification) UnmarshalJSON(b []byte) error {
	var V2wfmagentopportunitytopicagentopportunitynotificationMap map[string]interface{}
	err := json.Unmarshal(b, &V2wfmagentopportunitytopicagentopportunitynotificationMap)
	if err != nil {
		return err
	}
	
	if Id, ok := V2wfmagentopportunitytopicagentopportunitynotificationMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if BusinessUnitId, ok := V2wfmagentopportunitytopicagentopportunitynotificationMap["businessUnitId"].(string); ok {
		o.BusinessUnitId = &BusinessUnitId
	}
    
	if OrganizationId, ok := V2wfmagentopportunitytopicagentopportunitynotificationMap["organizationId"].(string); ok {
		o.OrganizationId = &OrganizationId
	}
    
	if AgentIds, ok := V2wfmagentopportunitytopicagentopportunitynotificationMap["agentIds"].([]interface{}); ok {
		AgentIdsString, _ := json.Marshal(AgentIds)
		json.Unmarshal(AgentIdsString, &o.AgentIds)
	}
	
	if OpportunityStartDate, ok := V2wfmagentopportunitytopicagentopportunitynotificationMap["opportunityStartDate"].(string); ok {
		o.OpportunityStartDate = &OpportunityStartDate
	}
    
	if OpportunityEndDate, ok := V2wfmagentopportunitytopicagentopportunitynotificationMap["opportunityEndDate"].(string); ok {
		o.OpportunityEndDate = &OpportunityEndDate
	}
    
	if OpenDate, ok := V2wfmagentopportunitytopicagentopportunitynotificationMap["openDate"].(string); ok {
		o.OpenDate = &OpenDate
	}
    
	if ApprovalType, ok := V2wfmagentopportunitytopicagentopportunitynotificationMap["approvalType"].(string); ok {
		o.ApprovalType = &ApprovalType
	}
    
	if LengthMinutes, ok := V2wfmagentopportunitytopicagentopportunitynotificationMap["lengthMinutes"].(float64); ok {
		LengthMinutesInt := int(LengthMinutes)
		o.LengthMinutes = &LengthMinutesInt
	}
	
	if DeadlineDate, ok := V2wfmagentopportunitytopicagentopportunitynotificationMap["deadlineDate"].(string); ok {
		o.DeadlineDate = &DeadlineDate
	}
    
	if ActivityCodeId, ok := V2wfmagentopportunitytopicagentopportunitynotificationMap["activityCodeId"].(string); ok {
		o.ActivityCodeId = &ActivityCodeId
	}
    
	if Name, ok := V2wfmagentopportunitytopicagentopportunitynotificationMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Description, ok := V2wfmagentopportunitytopicagentopportunitynotificationMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if ClosedDate, ok := V2wfmagentopportunitytopicagentopportunitynotificationMap["closedDate"].(string); ok {
		o.ClosedDate = &ClosedDate
	}
    
	if Status, ok := V2wfmagentopportunitytopicagentopportunitynotificationMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if EventType, ok := V2wfmagentopportunitytopicagentopportunitynotificationMap["eventType"].(string); ok {
		o.EventType = &EventType
	}
    
	if DenialCode, ok := V2wfmagentopportunitytopicagentopportunitynotificationMap["denialCode"].(string); ok {
		o.DenialCode = &DenialCode
	}
    
	if ReviewNote, ok := V2wfmagentopportunitytopicagentopportunitynotificationMap["reviewNote"].(string); ok {
		o.ReviewNote = &ReviewNote
	}
    
	if RemainingSpaces, ok := V2wfmagentopportunitytopicagentopportunitynotificationMap["remainingSpaces"].(float64); ok {
		RemainingSpacesInt := int(RemainingSpaces)
		o.RemainingSpaces = &RemainingSpacesInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *V2wfmagentopportunitytopicagentopportunitynotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

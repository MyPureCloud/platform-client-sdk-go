package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// V2wfmopportunitytopicopportunitynotification
type V2wfmopportunitytopicopportunitynotification struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id
	Id *string `json:"id,omitempty"`

	// BusinessUnitId
	BusinessUnitId *string `json:"businessUnitId,omitempty"`

	// OrganizationId
	OrganizationId *string `json:"organizationId,omitempty"`

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

	// CreatedDate
	CreatedDate *string `json:"createdDate,omitempty"`

	// PublishedDate
	PublishedDate *string `json:"publishedDate,omitempty"`

	// ClosedDate
	ClosedDate *string `json:"closedDate,omitempty"`

	// Status
	Status *string `json:"status,omitempty"`

	// EventType
	EventType *string `json:"eventType,omitempty"`

	// PendingCount
	PendingCount *int `json:"pendingCount,omitempty"`

	// WithdrawnCount
	WithdrawnCount *int `json:"withdrawnCount,omitempty"`

	// ApprovedCount
	ApprovedCount *int `json:"approvedCount,omitempty"`

	// DeniedCount
	DeniedCount *int `json:"deniedCount,omitempty"`

	// RemainingSpaces
	RemainingSpaces *int `json:"remainingSpaces,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *V2wfmopportunitytopicopportunitynotification) SetField(field string, fieldValue interface{}) {
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

func (o V2wfmopportunitytopicopportunitynotification) MarshalJSON() ([]byte, error) {
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
	type Alias V2wfmopportunitytopicopportunitynotification
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		BusinessUnitId *string `json:"businessUnitId,omitempty"`
		
		OrganizationId *string `json:"organizationId,omitempty"`
		
		OpportunityStartDate *string `json:"opportunityStartDate,omitempty"`
		
		OpportunityEndDate *string `json:"opportunityEndDate,omitempty"`
		
		OpenDate *string `json:"openDate,omitempty"`
		
		ApprovalType *string `json:"approvalType,omitempty"`
		
		LengthMinutes *int `json:"lengthMinutes,omitempty"`
		
		DeadlineDate *string `json:"deadlineDate,omitempty"`
		
		ActivityCodeId *string `json:"activityCodeId,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		CreatedDate *string `json:"createdDate,omitempty"`
		
		PublishedDate *string `json:"publishedDate,omitempty"`
		
		ClosedDate *string `json:"closedDate,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		EventType *string `json:"eventType,omitempty"`
		
		PendingCount *int `json:"pendingCount,omitempty"`
		
		WithdrawnCount *int `json:"withdrawnCount,omitempty"`
		
		ApprovedCount *int `json:"approvedCount,omitempty"`
		
		DeniedCount *int `json:"deniedCount,omitempty"`
		
		RemainingSpaces *int `json:"remainingSpaces,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		BusinessUnitId: o.BusinessUnitId,
		
		OrganizationId: o.OrganizationId,
		
		OpportunityStartDate: o.OpportunityStartDate,
		
		OpportunityEndDate: o.OpportunityEndDate,
		
		OpenDate: o.OpenDate,
		
		ApprovalType: o.ApprovalType,
		
		LengthMinutes: o.LengthMinutes,
		
		DeadlineDate: o.DeadlineDate,
		
		ActivityCodeId: o.ActivityCodeId,
		
		Name: o.Name,
		
		Description: o.Description,
		
		CreatedDate: o.CreatedDate,
		
		PublishedDate: o.PublishedDate,
		
		ClosedDate: o.ClosedDate,
		
		Status: o.Status,
		
		EventType: o.EventType,
		
		PendingCount: o.PendingCount,
		
		WithdrawnCount: o.WithdrawnCount,
		
		ApprovedCount: o.ApprovedCount,
		
		DeniedCount: o.DeniedCount,
		
		RemainingSpaces: o.RemainingSpaces,
		Alias:    (Alias)(o),
	})
}

func (o *V2wfmopportunitytopicopportunitynotification) UnmarshalJSON(b []byte) error {
	var V2wfmopportunitytopicopportunitynotificationMap map[string]interface{}
	err := json.Unmarshal(b, &V2wfmopportunitytopicopportunitynotificationMap)
	if err != nil {
		return err
	}
	
	if Id, ok := V2wfmopportunitytopicopportunitynotificationMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if BusinessUnitId, ok := V2wfmopportunitytopicopportunitynotificationMap["businessUnitId"].(string); ok {
		o.BusinessUnitId = &BusinessUnitId
	}
    
	if OrganizationId, ok := V2wfmopportunitytopicopportunitynotificationMap["organizationId"].(string); ok {
		o.OrganizationId = &OrganizationId
	}
    
	if OpportunityStartDate, ok := V2wfmopportunitytopicopportunitynotificationMap["opportunityStartDate"].(string); ok {
		o.OpportunityStartDate = &OpportunityStartDate
	}
    
	if OpportunityEndDate, ok := V2wfmopportunitytopicopportunitynotificationMap["opportunityEndDate"].(string); ok {
		o.OpportunityEndDate = &OpportunityEndDate
	}
    
	if OpenDate, ok := V2wfmopportunitytopicopportunitynotificationMap["openDate"].(string); ok {
		o.OpenDate = &OpenDate
	}
    
	if ApprovalType, ok := V2wfmopportunitytopicopportunitynotificationMap["approvalType"].(string); ok {
		o.ApprovalType = &ApprovalType
	}
    
	if LengthMinutes, ok := V2wfmopportunitytopicopportunitynotificationMap["lengthMinutes"].(float64); ok {
		LengthMinutesInt := int(LengthMinutes)
		o.LengthMinutes = &LengthMinutesInt
	}
	
	if DeadlineDate, ok := V2wfmopportunitytopicopportunitynotificationMap["deadlineDate"].(string); ok {
		o.DeadlineDate = &DeadlineDate
	}
    
	if ActivityCodeId, ok := V2wfmopportunitytopicopportunitynotificationMap["activityCodeId"].(string); ok {
		o.ActivityCodeId = &ActivityCodeId
	}
    
	if Name, ok := V2wfmopportunitytopicopportunitynotificationMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Description, ok := V2wfmopportunitytopicopportunitynotificationMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if CreatedDate, ok := V2wfmopportunitytopicopportunitynotificationMap["createdDate"].(string); ok {
		o.CreatedDate = &CreatedDate
	}
    
	if PublishedDate, ok := V2wfmopportunitytopicopportunitynotificationMap["publishedDate"].(string); ok {
		o.PublishedDate = &PublishedDate
	}
    
	if ClosedDate, ok := V2wfmopportunitytopicopportunitynotificationMap["closedDate"].(string); ok {
		o.ClosedDate = &ClosedDate
	}
    
	if Status, ok := V2wfmopportunitytopicopportunitynotificationMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if EventType, ok := V2wfmopportunitytopicopportunitynotificationMap["eventType"].(string); ok {
		o.EventType = &EventType
	}
    
	if PendingCount, ok := V2wfmopportunitytopicopportunitynotificationMap["pendingCount"].(float64); ok {
		PendingCountInt := int(PendingCount)
		o.PendingCount = &PendingCountInt
	}
	
	if WithdrawnCount, ok := V2wfmopportunitytopicopportunitynotificationMap["withdrawnCount"].(float64); ok {
		WithdrawnCountInt := int(WithdrawnCount)
		o.WithdrawnCount = &WithdrawnCountInt
	}
	
	if ApprovedCount, ok := V2wfmopportunitytopicopportunitynotificationMap["approvedCount"].(float64); ok {
		ApprovedCountInt := int(ApprovedCount)
		o.ApprovedCount = &ApprovedCountInt
	}
	
	if DeniedCount, ok := V2wfmopportunitytopicopportunitynotificationMap["deniedCount"].(float64); ok {
		DeniedCountInt := int(DeniedCount)
		o.DeniedCount = &DeniedCountInt
	}
	
	if RemainingSpaces, ok := V2wfmopportunitytopicopportunitynotificationMap["remainingSpaces"].(float64); ok {
		RemainingSpacesInt := int(RemainingSpaces)
		o.RemainingSpaces = &RemainingSpacesInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *V2wfmopportunitytopicopportunitynotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

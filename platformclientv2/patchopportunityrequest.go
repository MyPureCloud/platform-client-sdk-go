package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Patchopportunityrequest
type Patchopportunityrequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// StartDate - The start date and time of the opportunity in ISO-8601 format
	StartDate *time.Time `json:"startDate,omitempty"`

	// EndDate - The end date and time of the opportunity in ISO-8601 format
	EndDate *time.Time `json:"endDate,omitempty"`

	// OpenDate - The date and time when the opportunity opens for enrollment in ISO-8601 format. If not provided or in the past, it will be automatically updated to the current time when the opportunity is published
	OpenDate *Valuewrapperinstant `json:"openDate,omitempty"`

	// DeadlineDate - The deadline date and time for enrollment in the opportunity in ISO-8601 format
	DeadlineDate *time.Time `json:"deadlineDate,omitempty"`

	// Name - The name of the opportunity
	Name *string `json:"name,omitempty"`

	// Description - Additional details describing the purpose or context of this opportunity
	Description *Valuewrapperstring `json:"description,omitempty"`

	// ActivityCodeId - The ID of the activity code associated with the opportunity
	ActivityCodeId *string `json:"activityCodeId,omitempty"`

	// ApprovalType - The approval type for enrollments
	ApprovalType *string `json:"approvalType,omitempty"`

	// Capacity - The maximum capacity (enrollment slots) for this opportunity
	Capacity *int `json:"capacity,omitempty"`

	// AgentIds - The IDs of the agents that are invited to the opportunity
	AgentIds *Listwrapperstring `json:"agentIds,omitempty"`

	// Metadata - The metadata for the opportunity
	Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Patchopportunityrequest) SetField(field string, fieldValue interface{}) {
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

func (o Patchopportunityrequest) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "StartDate","EndDate","DeadlineDate", }
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
	type Alias Patchopportunityrequest
	
	StartDate := new(string)
	if o.StartDate != nil {
		
		*StartDate = timeutil.Strftime(o.StartDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartDate = nil
	}
	
	EndDate := new(string)
	if o.EndDate != nil {
		
		*EndDate = timeutil.Strftime(o.EndDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		EndDate = nil
	}
	
	DeadlineDate := new(string)
	if o.DeadlineDate != nil {
		
		*DeadlineDate = timeutil.Strftime(o.DeadlineDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DeadlineDate = nil
	}
	
	return json.Marshal(&struct { 
		StartDate *string `json:"startDate,omitempty"`
		
		EndDate *string `json:"endDate,omitempty"`
		
		OpenDate *Valuewrapperinstant `json:"openDate,omitempty"`
		
		DeadlineDate *string `json:"deadlineDate,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Description *Valuewrapperstring `json:"description,omitempty"`
		
		ActivityCodeId *string `json:"activityCodeId,omitempty"`
		
		ApprovalType *string `json:"approvalType,omitempty"`
		
		Capacity *int `json:"capacity,omitempty"`
		
		AgentIds *Listwrapperstring `json:"agentIds,omitempty"`
		
		Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`
		Alias
	}{ 
		StartDate: StartDate,
		
		EndDate: EndDate,
		
		OpenDate: o.OpenDate,
		
		DeadlineDate: DeadlineDate,
		
		Name: o.Name,
		
		Description: o.Description,
		
		ActivityCodeId: o.ActivityCodeId,
		
		ApprovalType: o.ApprovalType,
		
		Capacity: o.Capacity,
		
		AgentIds: o.AgentIds,
		
		Metadata: o.Metadata,
		Alias:    (Alias)(o),
	})
}

func (o *Patchopportunityrequest) UnmarshalJSON(b []byte) error {
	var PatchopportunityrequestMap map[string]interface{}
	err := json.Unmarshal(b, &PatchopportunityrequestMap)
	if err != nil {
		return err
	}
	
	if startDateString, ok := PatchopportunityrequestMap["startDate"].(string); ok {
		StartDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", startDateString)
		o.StartDate = &StartDate
	}
	
	if endDateString, ok := PatchopportunityrequestMap["endDate"].(string); ok {
		EndDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", endDateString)
		o.EndDate = &EndDate
	}
	
	if OpenDate, ok := PatchopportunityrequestMap["openDate"].(map[string]interface{}); ok {
		OpenDateString, _ := json.Marshal(OpenDate)
		json.Unmarshal(OpenDateString, &o.OpenDate)
	}
	
	if deadlineDateString, ok := PatchopportunityrequestMap["deadlineDate"].(string); ok {
		DeadlineDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", deadlineDateString)
		o.DeadlineDate = &DeadlineDate
	}
	
	if Name, ok := PatchopportunityrequestMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Description, ok := PatchopportunityrequestMap["description"].(map[string]interface{}); ok {
		DescriptionString, _ := json.Marshal(Description)
		json.Unmarshal(DescriptionString, &o.Description)
	}
	
	if ActivityCodeId, ok := PatchopportunityrequestMap["activityCodeId"].(string); ok {
		o.ActivityCodeId = &ActivityCodeId
	}
    
	if ApprovalType, ok := PatchopportunityrequestMap["approvalType"].(string); ok {
		o.ApprovalType = &ApprovalType
	}
    
	if Capacity, ok := PatchopportunityrequestMap["capacity"].(float64); ok {
		CapacityInt := int(Capacity)
		o.Capacity = &CapacityInt
	}
	
	if AgentIds, ok := PatchopportunityrequestMap["agentIds"].(map[string]interface{}); ok {
		AgentIdsString, _ := json.Marshal(AgentIds)
		json.Unmarshal(AgentIdsString, &o.AgentIds)
	}
	
	if Metadata, ok := PatchopportunityrequestMap["metadata"].(map[string]interface{}); ok {
		MetadataString, _ := json.Marshal(Metadata)
		json.Unmarshal(MetadataString, &o.Metadata)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Patchopportunityrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Opportunityresult
type Opportunityresult struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// StartDate - The start date and time of the opportunity in ISO-8601 format
	StartDate *time.Time `json:"startDate,omitempty"`

	// EndDate - The end date and time of the opportunity in ISO-8601 format
	EndDate *time.Time `json:"endDate,omitempty"`

	// Status - The current status of the opportunity
	Status *string `json:"status,omitempty"`

	// OpenDate - The date and time when the opportunity opens for enrollment in ISO-8601 format. If not provided or in the past, it will be automatically updated to the current time when the opportunity is published
	OpenDate *time.Time `json:"openDate,omitempty"`

	// DeadlineDate - The deadline date and time for enrollment in the opportunity in ISO-8601 format
	DeadlineDate *time.Time `json:"deadlineDate,omitempty"`

	// Name - The name of the opportunity
	Name *string `json:"name,omitempty"`

	// Description - Additional details describing the purpose or context of this opportunity
	Description *string `json:"description,omitempty"`

	// ActivityCodeId - The ID of the activity code associated with the opportunity
	ActivityCodeId *string `json:"activityCodeId,omitempty"`

	// ApprovalType - The approval type for enrollments
	ApprovalType *string `json:"approvalType,omitempty"`

	// AgentCount - The total number of agents invited to this opportunity
	AgentCount *int `json:"agentCount,omitempty"`

	// Capacity - The maximum capacity (enrollment slots) for this opportunity
	Capacity *int `json:"capacity,omitempty"`

	// EnrollmentProcessingCount - The number of enrollments currently being processed
	EnrollmentProcessingCount *int `json:"enrollmentProcessingCount,omitempty"`

	// EnrollmentCounts - The counts for enrollment statuses
	EnrollmentCounts *Opportunityenrollmentcounts `json:"enrollmentCounts,omitempty"`

	// PublishedDate - The date and time when the opportunity was published in ISO-8601 format
	PublishedDate *time.Time `json:"publishedDate,omitempty"`

	// ClosedDate - The date and time when the opportunity was closed in ISO-8601 format
	ClosedDate *time.Time `json:"closedDate,omitempty"`

	// SystemMessageCode - The system-generated message code about opportunity processing issues or validation failures
	SystemMessageCode *string `json:"systemMessageCode,omitempty"`

	// Metadata - The metadata for the opportunity
	Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Opportunityresult) SetField(field string, fieldValue interface{}) {
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

func (o Opportunityresult) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "StartDate","EndDate","OpenDate","DeadlineDate","PublishedDate","ClosedDate", }
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
	type Alias Opportunityresult
	
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
	
	OpenDate := new(string)
	if o.OpenDate != nil {
		
		*OpenDate = timeutil.Strftime(o.OpenDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		OpenDate = nil
	}
	
	DeadlineDate := new(string)
	if o.DeadlineDate != nil {
		
		*DeadlineDate = timeutil.Strftime(o.DeadlineDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DeadlineDate = nil
	}
	
	PublishedDate := new(string)
	if o.PublishedDate != nil {
		
		*PublishedDate = timeutil.Strftime(o.PublishedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		PublishedDate = nil
	}
	
	ClosedDate := new(string)
	if o.ClosedDate != nil {
		
		*ClosedDate = timeutil.Strftime(o.ClosedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ClosedDate = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		StartDate *string `json:"startDate,omitempty"`
		
		EndDate *string `json:"endDate,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		OpenDate *string `json:"openDate,omitempty"`
		
		DeadlineDate *string `json:"deadlineDate,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		ActivityCodeId *string `json:"activityCodeId,omitempty"`
		
		ApprovalType *string `json:"approvalType,omitempty"`
		
		AgentCount *int `json:"agentCount,omitempty"`
		
		Capacity *int `json:"capacity,omitempty"`
		
		EnrollmentProcessingCount *int `json:"enrollmentProcessingCount,omitempty"`
		
		EnrollmentCounts *Opportunityenrollmentcounts `json:"enrollmentCounts,omitempty"`
		
		PublishedDate *string `json:"publishedDate,omitempty"`
		
		ClosedDate *string `json:"closedDate,omitempty"`
		
		SystemMessageCode *string `json:"systemMessageCode,omitempty"`
		
		Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		StartDate: StartDate,
		
		EndDate: EndDate,
		
		Status: o.Status,
		
		OpenDate: OpenDate,
		
		DeadlineDate: DeadlineDate,
		
		Name: o.Name,
		
		Description: o.Description,
		
		ActivityCodeId: o.ActivityCodeId,
		
		ApprovalType: o.ApprovalType,
		
		AgentCount: o.AgentCount,
		
		Capacity: o.Capacity,
		
		EnrollmentProcessingCount: o.EnrollmentProcessingCount,
		
		EnrollmentCounts: o.EnrollmentCounts,
		
		PublishedDate: PublishedDate,
		
		ClosedDate: ClosedDate,
		
		SystemMessageCode: o.SystemMessageCode,
		
		Metadata: o.Metadata,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Opportunityresult) UnmarshalJSON(b []byte) error {
	var OpportunityresultMap map[string]interface{}
	err := json.Unmarshal(b, &OpportunityresultMap)
	if err != nil {
		return err
	}
	
	if Id, ok := OpportunityresultMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if startDateString, ok := OpportunityresultMap["startDate"].(string); ok {
		StartDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", startDateString)
		o.StartDate = &StartDate
	}
	
	if endDateString, ok := OpportunityresultMap["endDate"].(string); ok {
		EndDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", endDateString)
		o.EndDate = &EndDate
	}
	
	if Status, ok := OpportunityresultMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if openDateString, ok := OpportunityresultMap["openDate"].(string); ok {
		OpenDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", openDateString)
		o.OpenDate = &OpenDate
	}
	
	if deadlineDateString, ok := OpportunityresultMap["deadlineDate"].(string); ok {
		DeadlineDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", deadlineDateString)
		o.DeadlineDate = &DeadlineDate
	}
	
	if Name, ok := OpportunityresultMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Description, ok := OpportunityresultMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if ActivityCodeId, ok := OpportunityresultMap["activityCodeId"].(string); ok {
		o.ActivityCodeId = &ActivityCodeId
	}
    
	if ApprovalType, ok := OpportunityresultMap["approvalType"].(string); ok {
		o.ApprovalType = &ApprovalType
	}
    
	if AgentCount, ok := OpportunityresultMap["agentCount"].(float64); ok {
		AgentCountInt := int(AgentCount)
		o.AgentCount = &AgentCountInt
	}
	
	if Capacity, ok := OpportunityresultMap["capacity"].(float64); ok {
		CapacityInt := int(Capacity)
		o.Capacity = &CapacityInt
	}
	
	if EnrollmentProcessingCount, ok := OpportunityresultMap["enrollmentProcessingCount"].(float64); ok {
		EnrollmentProcessingCountInt := int(EnrollmentProcessingCount)
		o.EnrollmentProcessingCount = &EnrollmentProcessingCountInt
	}
	
	if EnrollmentCounts, ok := OpportunityresultMap["enrollmentCounts"].(map[string]interface{}); ok {
		EnrollmentCountsString, _ := json.Marshal(EnrollmentCounts)
		json.Unmarshal(EnrollmentCountsString, &o.EnrollmentCounts)
	}
	
	if publishedDateString, ok := OpportunityresultMap["publishedDate"].(string); ok {
		PublishedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", publishedDateString)
		o.PublishedDate = &PublishedDate
	}
	
	if closedDateString, ok := OpportunityresultMap["closedDate"].(string); ok {
		ClosedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", closedDateString)
		o.ClosedDate = &ClosedDate
	}
	
	if SystemMessageCode, ok := OpportunityresultMap["systemMessageCode"].(string); ok {
		o.SystemMessageCode = &SystemMessageCode
	}
    
	if Metadata, ok := OpportunityresultMap["metadata"].(map[string]interface{}); ok {
		MetadataString, _ := json.Marshal(Metadata)
		json.Unmarshal(MetadataString, &o.Metadata)
	}
	
	if SelfUri, ok := OpportunityresultMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Opportunityresult) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

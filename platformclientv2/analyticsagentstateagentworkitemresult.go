package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Analyticsagentstateagentworkitemresult
type Analyticsagentstateagentworkitemresult struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - Workitem Id
	Id *string `json:"id,omitempty"`

	// QueueId - The Id of the queue used for routing the workitem
	QueueId *string `json:"queueId,omitempty"`

	// DivisionId - The Id of the division the workitem belongs to
	DivisionId *string `json:"divisionId,omitempty"`

	// TypeId - The Id of the worktype of the workitem
	TypeId *string `json:"typeId,omitempty"`

	// WorkbinId - The Id of the workbin the workitem is placed in
	WorkbinId *string `json:"workbinId,omitempty"`

	// DateDue - Due datetime of the workitem. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateDue *time.Time `json:"dateDue,omitempty"`

	// DateExpires - Expiry datetime of the workitem. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateExpires *time.Time `json:"dateExpires,omitempty"`

	// StatusId - The Id of the current status of the workitem
	StatusId *string `json:"statusId,omitempty"`

	// StatusCategory - The category of the current status of the workitem
	StatusCategory *string `json:"statusCategory,omitempty"`

	// DateStatusChanged - Datetime the status of the workitem last changed. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateStatusChanged *time.Time `json:"dateStatusChanged,omitempty"`

	// Priority - Priority of the workitem
	Priority *int `json:"priority,omitempty"`

	// AssignmentState - Assignment state of the workitem
	AssignmentState *string `json:"assignmentState,omitempty"`

	// DateAssignmentStateChanged - Datetime the assignment state of the workitem last changed. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateAssignmentStateChanged *time.Time `json:"dateAssignmentStateChanged,omitempty"`

	// DateCreated - Creation datetime of the workitem. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Analyticsagentstateagentworkitemresult) SetField(field string, fieldValue interface{}) {
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

func (o Analyticsagentstateagentworkitemresult) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateDue","DateExpires","DateStatusChanged","DateAssignmentStateChanged","DateCreated", }
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
	type Alias Analyticsagentstateagentworkitemresult
	
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
	
	DateStatusChanged := new(string)
	if o.DateStatusChanged != nil {
		
		*DateStatusChanged = timeutil.Strftime(o.DateStatusChanged, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateStatusChanged = nil
	}
	
	DateAssignmentStateChanged := new(string)
	if o.DateAssignmentStateChanged != nil {
		
		*DateAssignmentStateChanged = timeutil.Strftime(o.DateAssignmentStateChanged, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateAssignmentStateChanged = nil
	}
	
	DateCreated := new(string)
	if o.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(o.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		QueueId *string `json:"queueId,omitempty"`
		
		DivisionId *string `json:"divisionId,omitempty"`
		
		TypeId *string `json:"typeId,omitempty"`
		
		WorkbinId *string `json:"workbinId,omitempty"`
		
		DateDue *string `json:"dateDue,omitempty"`
		
		DateExpires *string `json:"dateExpires,omitempty"`
		
		StatusId *string `json:"statusId,omitempty"`
		
		StatusCategory *string `json:"statusCategory,omitempty"`
		
		DateStatusChanged *string `json:"dateStatusChanged,omitempty"`
		
		Priority *int `json:"priority,omitempty"`
		
		AssignmentState *string `json:"assignmentState,omitempty"`
		
		DateAssignmentStateChanged *string `json:"dateAssignmentStateChanged,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		QueueId: o.QueueId,
		
		DivisionId: o.DivisionId,
		
		TypeId: o.TypeId,
		
		WorkbinId: o.WorkbinId,
		
		DateDue: DateDue,
		
		DateExpires: DateExpires,
		
		StatusId: o.StatusId,
		
		StatusCategory: o.StatusCategory,
		
		DateStatusChanged: DateStatusChanged,
		
		Priority: o.Priority,
		
		AssignmentState: o.AssignmentState,
		
		DateAssignmentStateChanged: DateAssignmentStateChanged,
		
		DateCreated: DateCreated,
		Alias:    (Alias)(o),
	})
}

func (o *Analyticsagentstateagentworkitemresult) UnmarshalJSON(b []byte) error {
	var AnalyticsagentstateagentworkitemresultMap map[string]interface{}
	err := json.Unmarshal(b, &AnalyticsagentstateagentworkitemresultMap)
	if err != nil {
		return err
	}
	
	if Id, ok := AnalyticsagentstateagentworkitemresultMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if QueueId, ok := AnalyticsagentstateagentworkitemresultMap["queueId"].(string); ok {
		o.QueueId = &QueueId
	}
    
	if DivisionId, ok := AnalyticsagentstateagentworkitemresultMap["divisionId"].(string); ok {
		o.DivisionId = &DivisionId
	}
    
	if TypeId, ok := AnalyticsagentstateagentworkitemresultMap["typeId"].(string); ok {
		o.TypeId = &TypeId
	}
    
	if WorkbinId, ok := AnalyticsagentstateagentworkitemresultMap["workbinId"].(string); ok {
		o.WorkbinId = &WorkbinId
	}
    
	if dateDueString, ok := AnalyticsagentstateagentworkitemresultMap["dateDue"].(string); ok {
		DateDue, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateDueString)
		o.DateDue = &DateDue
	}
	
	if dateExpiresString, ok := AnalyticsagentstateagentworkitemresultMap["dateExpires"].(string); ok {
		DateExpires, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateExpiresString)
		o.DateExpires = &DateExpires
	}
	
	if StatusId, ok := AnalyticsagentstateagentworkitemresultMap["statusId"].(string); ok {
		o.StatusId = &StatusId
	}
    
	if StatusCategory, ok := AnalyticsagentstateagentworkitemresultMap["statusCategory"].(string); ok {
		o.StatusCategory = &StatusCategory
	}
    
	if dateStatusChangedString, ok := AnalyticsagentstateagentworkitemresultMap["dateStatusChanged"].(string); ok {
		DateStatusChanged, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateStatusChangedString)
		o.DateStatusChanged = &DateStatusChanged
	}
	
	if Priority, ok := AnalyticsagentstateagentworkitemresultMap["priority"].(float64); ok {
		PriorityInt := int(Priority)
		o.Priority = &PriorityInt
	}
	
	if AssignmentState, ok := AnalyticsagentstateagentworkitemresultMap["assignmentState"].(string); ok {
		o.AssignmentState = &AssignmentState
	}
    
	if dateAssignmentStateChangedString, ok := AnalyticsagentstateagentworkitemresultMap["dateAssignmentStateChanged"].(string); ok {
		DateAssignmentStateChanged, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateAssignmentStateChangedString)
		o.DateAssignmentStateChanged = &DateAssignmentStateChanged
	}
	
	if dateCreatedString, ok := AnalyticsagentstateagentworkitemresultMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Analyticsagentstateagentworkitemresult) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

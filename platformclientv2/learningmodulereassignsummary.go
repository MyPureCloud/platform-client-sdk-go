package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Learningmodulereassignsummary - Learning module reassign summary data
type Learningmodulereassignsummary struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// TotalReassigned - The total number of users for whom assignment is reassigned
	TotalReassigned *int `json:"totalReassigned,omitempty"`

	// CompletedCount - The total number of users who have the assignment in Completed state
	CompletedCount *int `json:"completedCount,omitempty"`

	// InProgressCount - The total number of users who have the assignment in InProgress state
	InProgressCount *int `json:"inProgressCount,omitempty"`

	// AssignedCount - The total number of users who have the assignment in Assigned state
	AssignedCount *int `json:"assignedCount,omitempty"`

	// NotCompletedCount - The total number of users who have their assignment overdue
	NotCompletedCount *int `json:"notCompletedCount,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Learningmodulereassignsummary) SetField(field string, fieldValue interface{}) {
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

func (o Learningmodulereassignsummary) MarshalJSON() ([]byte, error) {
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
			if contains(dateTimeFields, fieldName) {
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
	type Alias Learningmodulereassignsummary
	
	return json.Marshal(&struct { 
		TotalReassigned *int `json:"totalReassigned,omitempty"`
		
		CompletedCount *int `json:"completedCount,omitempty"`
		
		InProgressCount *int `json:"inProgressCount,omitempty"`
		
		AssignedCount *int `json:"assignedCount,omitempty"`
		
		NotCompletedCount *int `json:"notCompletedCount,omitempty"`
		Alias
	}{ 
		TotalReassigned: o.TotalReassigned,
		
		CompletedCount: o.CompletedCount,
		
		InProgressCount: o.InProgressCount,
		
		AssignedCount: o.AssignedCount,
		
		NotCompletedCount: o.NotCompletedCount,
		Alias:    (Alias)(o),
	})
}

func (o *Learningmodulereassignsummary) UnmarshalJSON(b []byte) error {
	var LearningmodulereassignsummaryMap map[string]interface{}
	err := json.Unmarshal(b, &LearningmodulereassignsummaryMap)
	if err != nil {
		return err
	}
	
	if TotalReassigned, ok := LearningmodulereassignsummaryMap["totalReassigned"].(float64); ok {
		TotalReassignedInt := int(TotalReassigned)
		o.TotalReassigned = &TotalReassignedInt
	}
	
	if CompletedCount, ok := LearningmodulereassignsummaryMap["completedCount"].(float64); ok {
		CompletedCountInt := int(CompletedCount)
		o.CompletedCount = &CompletedCountInt
	}
	
	if InProgressCount, ok := LearningmodulereassignsummaryMap["inProgressCount"].(float64); ok {
		InProgressCountInt := int(InProgressCount)
		o.InProgressCount = &InProgressCountInt
	}
	
	if AssignedCount, ok := LearningmodulereassignsummaryMap["assignedCount"].(float64); ok {
		AssignedCountInt := int(AssignedCount)
		o.AssignedCount = &AssignedCountInt
	}
	
	if NotCompletedCount, ok := LearningmodulereassignsummaryMap["notCompletedCount"].(float64); ok {
		NotCompletedCountInt := int(NotCompletedCount)
		o.NotCompletedCount = &NotCompletedCountInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Learningmodulereassignsummary) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

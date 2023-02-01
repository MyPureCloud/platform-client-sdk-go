package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Assignment
type Assignment struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// AssignedMembers - The list of users successfully assigned to the custom performance profile
	AssignedMembers *[]Userreference `json:"assignedMembers,omitempty"`

	// RemovedMembers - The list of users successfully removed from the custom performance profile
	RemovedMembers *[]Userreference `json:"removedMembers,omitempty"`

	// AssignmentErrors - The list of users failed assignment or removal for the custom performance profile
	AssignmentErrors *[]Assignmenterror `json:"assignmentErrors,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Assignment) SetField(field string, fieldValue interface{}) {
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

func (o Assignment) MarshalJSON() ([]byte, error) {
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
	type Alias Assignment
	
	return json.Marshal(&struct { 
		AssignedMembers *[]Userreference `json:"assignedMembers,omitempty"`
		
		RemovedMembers *[]Userreference `json:"removedMembers,omitempty"`
		
		AssignmentErrors *[]Assignmenterror `json:"assignmentErrors,omitempty"`
		Alias
	}{ 
		AssignedMembers: o.AssignedMembers,
		
		RemovedMembers: o.RemovedMembers,
		
		AssignmentErrors: o.AssignmentErrors,
		Alias:    (Alias)(o),
	})
}

func (o *Assignment) UnmarshalJSON(b []byte) error {
	var AssignmentMap map[string]interface{}
	err := json.Unmarshal(b, &AssignmentMap)
	if err != nil {
		return err
	}
	
	if AssignedMembers, ok := AssignmentMap["assignedMembers"].([]interface{}); ok {
		AssignedMembersString, _ := json.Marshal(AssignedMembers)
		json.Unmarshal(AssignedMembersString, &o.AssignedMembers)
	}
	
	if RemovedMembers, ok := AssignmentMap["removedMembers"].([]interface{}); ok {
		RemovedMembersString, _ := json.Marshal(RemovedMembers)
		json.Unmarshal(RemovedMembersString, &o.RemovedMembers)
	}
	
	if AssignmentErrors, ok := AssignmentMap["assignmentErrors"].([]interface{}); ok {
		AssignmentErrorsString, _ := json.Marshal(AssignmentErrors)
		json.Unmarshal(AssignmentErrorsString, &o.AssignmentErrors)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Assignment) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

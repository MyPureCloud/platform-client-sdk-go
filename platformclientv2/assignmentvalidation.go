package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Assignmentvalidation
type Assignmentvalidation struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// MembersNotAssigned - The list of users that are not assigned to any custom performance profile
	MembersNotAssigned *[]Userreference `json:"membersNotAssigned,omitempty"`

	// MembersAlreadyAssigned - The list of users that are already assigned to the requesting custom performance profile
	MembersAlreadyAssigned *[]Userreference `json:"membersAlreadyAssigned,omitempty"`

	// MembersAlreadyAssignedToOther - The list of users that are already assigned to other custom performance profiles
	MembersAlreadyAssignedToOther *[]Otherprofileassignment `json:"membersAlreadyAssignedToOther,omitempty"`

	// InvalidMemberAssignments - The list of user id that are invalid for the gamfication service to handle
	InvalidMemberAssignments *[]Invalidassignment `json:"invalidMemberAssignments,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Assignmentvalidation) SetField(field string, fieldValue interface{}) {
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

func (o Assignmentvalidation) MarshalJSON() ([]byte, error) {
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
	type Alias Assignmentvalidation
	
	return json.Marshal(&struct { 
		MembersNotAssigned *[]Userreference `json:"membersNotAssigned,omitempty"`
		
		MembersAlreadyAssigned *[]Userreference `json:"membersAlreadyAssigned,omitempty"`
		
		MembersAlreadyAssignedToOther *[]Otherprofileassignment `json:"membersAlreadyAssignedToOther,omitempty"`
		
		InvalidMemberAssignments *[]Invalidassignment `json:"invalidMemberAssignments,omitempty"`
		Alias
	}{ 
		MembersNotAssigned: o.MembersNotAssigned,
		
		MembersAlreadyAssigned: o.MembersAlreadyAssigned,
		
		MembersAlreadyAssignedToOther: o.MembersAlreadyAssignedToOther,
		
		InvalidMemberAssignments: o.InvalidMemberAssignments,
		Alias:    (Alias)(o),
	})
}

func (o *Assignmentvalidation) UnmarshalJSON(b []byte) error {
	var AssignmentvalidationMap map[string]interface{}
	err := json.Unmarshal(b, &AssignmentvalidationMap)
	if err != nil {
		return err
	}
	
	if MembersNotAssigned, ok := AssignmentvalidationMap["membersNotAssigned"].([]interface{}); ok {
		MembersNotAssignedString, _ := json.Marshal(MembersNotAssigned)
		json.Unmarshal(MembersNotAssignedString, &o.MembersNotAssigned)
	}
	
	if MembersAlreadyAssigned, ok := AssignmentvalidationMap["membersAlreadyAssigned"].([]interface{}); ok {
		MembersAlreadyAssignedString, _ := json.Marshal(MembersAlreadyAssigned)
		json.Unmarshal(MembersAlreadyAssignedString, &o.MembersAlreadyAssigned)
	}
	
	if MembersAlreadyAssignedToOther, ok := AssignmentvalidationMap["membersAlreadyAssignedToOther"].([]interface{}); ok {
		MembersAlreadyAssignedToOtherString, _ := json.Marshal(MembersAlreadyAssignedToOther)
		json.Unmarshal(MembersAlreadyAssignedToOtherString, &o.MembersAlreadyAssignedToOther)
	}
	
	if InvalidMemberAssignments, ok := AssignmentvalidationMap["invalidMemberAssignments"].([]interface{}); ok {
		InvalidMemberAssignmentsString, _ := json.Marshal(InvalidMemberAssignments)
		json.Unmarshal(InvalidMemberAssignmentsString, &o.InvalidMemberAssignments)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Assignmentvalidation) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

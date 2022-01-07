package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Assignment
type Assignment struct { 
	// AssignedMembers - The list of users successfully assigned to the custom performance profile
	AssignedMembers *[]Userreference `json:"assignedMembers,omitempty"`


	// RemovedMembers - The list of users successfully removed from the custom performance profile
	RemovedMembers *[]Userreference `json:"removedMembers,omitempty"`


	// AssignmentErrors - The list of users failed assignment or removal for the custom performance profile
	AssignmentErrors *[]Assignmenterror `json:"assignmentErrors,omitempty"`

}

func (o *Assignment) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Assignment
	
	return json.Marshal(&struct { 
		AssignedMembers *[]Userreference `json:"assignedMembers,omitempty"`
		
		RemovedMembers *[]Userreference `json:"removedMembers,omitempty"`
		
		AssignmentErrors *[]Assignmenterror `json:"assignmentErrors,omitempty"`
		*Alias
	}{ 
		AssignedMembers: o.AssignedMembers,
		
		RemovedMembers: o.RemovedMembers,
		
		AssignmentErrors: o.AssignmentErrors,
		Alias:    (*Alias)(o),
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

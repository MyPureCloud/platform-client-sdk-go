package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Assignmentvalidation
type Assignmentvalidation struct { 
	// MembersNotAssigned - The list of users that are not assigned to any custom performance profile
	MembersNotAssigned *[]Userreference `json:"membersNotAssigned,omitempty"`


	// MembersAlreadyAssigned - The list of users that are already assigned to the requesting custom performance profile
	MembersAlreadyAssigned *[]Userreference `json:"membersAlreadyAssigned,omitempty"`


	// MembersAlreadyAssignedToOther - The list of users that are already assigned to other custom performance profiles
	MembersAlreadyAssignedToOther *[]Otherprofileassignment `json:"membersAlreadyAssignedToOther,omitempty"`


	// InvalidMemberAssignments - The list of user id that are invalid for the gamfication service to handle
	InvalidMemberAssignments *[]Invalidassignment `json:"invalidMemberAssignments,omitempty"`

}

func (o *Assignmentvalidation) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Assignmentvalidation
	
	return json.Marshal(&struct { 
		MembersNotAssigned *[]Userreference `json:"membersNotAssigned,omitempty"`
		
		MembersAlreadyAssigned *[]Userreference `json:"membersAlreadyAssigned,omitempty"`
		
		MembersAlreadyAssignedToOther *[]Otherprofileassignment `json:"membersAlreadyAssignedToOther,omitempty"`
		
		InvalidMemberAssignments *[]Invalidassignment `json:"invalidMemberAssignments,omitempty"`
		*Alias
	}{ 
		MembersNotAssigned: o.MembersNotAssigned,
		
		MembersAlreadyAssigned: o.MembersAlreadyAssigned,
		
		MembersAlreadyAssignedToOther: o.MembersAlreadyAssignedToOther,
		
		InvalidMemberAssignments: o.InvalidMemberAssignments,
		Alias:    (*Alias)(o),
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

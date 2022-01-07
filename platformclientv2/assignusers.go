package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Assignusers
type Assignusers struct { 
	// MembersToAssign - List of user ids to assign to a performance profile
	MembersToAssign *[]string `json:"membersToAssign,omitempty"`


	// MembersToRemove - List of user ids to remove from a performance profile
	MembersToRemove *[]string `json:"membersToRemove,omitempty"`

}

func (o *Assignusers) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Assignusers
	
	return json.Marshal(&struct { 
		MembersToAssign *[]string `json:"membersToAssign,omitempty"`
		
		MembersToRemove *[]string `json:"membersToRemove,omitempty"`
		*Alias
	}{ 
		MembersToAssign: o.MembersToAssign,
		
		MembersToRemove: o.MembersToRemove,
		Alias:    (*Alias)(o),
	})
}

func (o *Assignusers) UnmarshalJSON(b []byte) error {
	var AssignusersMap map[string]interface{}
	err := json.Unmarshal(b, &AssignusersMap)
	if err != nil {
		return err
	}
	
	if MembersToAssign, ok := AssignusersMap["membersToAssign"].([]interface{}); ok {
		MembersToAssignString, _ := json.Marshal(MembersToAssign)
		json.Unmarshal(MembersToAssignString, &o.MembersToAssign)
	}
	
	if MembersToRemove, ok := AssignusersMap["membersToRemove"].([]interface{}); ok {
		MembersToRemoveString, _ := json.Marshal(MembersToRemove)
		json.Unmarshal(MembersToRemoveString, &o.MembersToRemove)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Assignusers) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

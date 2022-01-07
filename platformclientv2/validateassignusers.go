package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Validateassignusers
type Validateassignusers struct { 
	// MembersToAssign - List of user ids to assign to a performance profile
	MembersToAssign *[]string `json:"membersToAssign,omitempty"`

}

func (o *Validateassignusers) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Validateassignusers
	
	return json.Marshal(&struct { 
		MembersToAssign *[]string `json:"membersToAssign,omitempty"`
		*Alias
	}{ 
		MembersToAssign: o.MembersToAssign,
		Alias:    (*Alias)(o),
	})
}

func (o *Validateassignusers) UnmarshalJSON(b []byte) error {
	var ValidateassignusersMap map[string]interface{}
	err := json.Unmarshal(b, &ValidateassignusersMap)
	if err != nil {
		return err
	}
	
	if MembersToAssign, ok := ValidateassignusersMap["membersToAssign"].([]interface{}); ok {
		MembersToAssignString, _ := json.Marshal(MembersToAssign)
		json.Unmarshal(MembersToAssignString, &o.MembersToAssign)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Validateassignusers) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

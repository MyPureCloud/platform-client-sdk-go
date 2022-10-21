package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Teammembers
type Teammembers struct { 
	// MemberIds - A list of the ids of the members to add or remove
	MemberIds *[]string `json:"memberIds,omitempty"`

}

func (o *Teammembers) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Teammembers
	
	return json.Marshal(&struct { 
		MemberIds *[]string `json:"memberIds,omitempty"`
		*Alias
	}{ 
		MemberIds: o.MemberIds,
		Alias:    (*Alias)(o),
	})
}

func (o *Teammembers) UnmarshalJSON(b []byte) error {
	var TeammembersMap map[string]interface{}
	err := json.Unmarshal(b, &TeammembersMap)
	if err != nil {
		return err
	}
	
	if MemberIds, ok := TeammembersMap["memberIds"].([]interface{}); ok {
		MemberIdsString, _ := json.Marshal(MemberIds)
		json.Unmarshal(MemberIdsString, &o.MemberIds)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Teammembers) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

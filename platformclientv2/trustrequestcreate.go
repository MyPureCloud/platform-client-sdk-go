package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Trustrequestcreate
type Trustrequestcreate struct { 
	// UserIds - The list of trustee users that are requesting access. If no users are specified, at least one group is required.
	UserIds *[]string `json:"userIds,omitempty"`


	// GroupIds - The list of trustee groups that are requesting access. If no groups are specified, at least one user is required.
	GroupIds *[]string `json:"groupIds,omitempty"`

}

func (o *Trustrequestcreate) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Trustrequestcreate
	
	return json.Marshal(&struct { 
		UserIds *[]string `json:"userIds,omitempty"`
		
		GroupIds *[]string `json:"groupIds,omitempty"`
		*Alias
	}{ 
		UserIds: o.UserIds,
		
		GroupIds: o.GroupIds,
		Alias:    (*Alias)(o),
	})
}

func (o *Trustrequestcreate) UnmarshalJSON(b []byte) error {
	var TrustrequestcreateMap map[string]interface{}
	err := json.Unmarshal(b, &TrustrequestcreateMap)
	if err != nil {
		return err
	}
	
	if UserIds, ok := TrustrequestcreateMap["userIds"].([]interface{}); ok {
		UserIdsString, _ := json.Marshal(UserIds)
		json.Unmarshal(UserIdsString, &o.UserIds)
	}
	
	if GroupIds, ok := TrustrequestcreateMap["groupIds"].([]interface{}); ok {
		GroupIdsString, _ := json.Marshal(GroupIds)
		json.Unmarshal(GroupIdsString, &o.GroupIds)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Trustrequestcreate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

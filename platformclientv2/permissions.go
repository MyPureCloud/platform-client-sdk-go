package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Permissions
type Permissions struct { 
	// Ids - List of permission ids.
	Ids *[]string `json:"ids,omitempty"`

}

func (o *Permissions) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Permissions
	
	return json.Marshal(&struct { 
		Ids *[]string `json:"ids,omitempty"`
		*Alias
	}{ 
		Ids: o.Ids,
		Alias:    (*Alias)(o),
	})
}

func (o *Permissions) UnmarshalJSON(b []byte) error {
	var PermissionsMap map[string]interface{}
	err := json.Unmarshal(b, &PermissionsMap)
	if err != nil {
		return err
	}
	
	if Ids, ok := PermissionsMap["ids"].([]interface{}); ok {
		IdsString, _ := json.Marshal(Ids)
		json.Unmarshal(IdsString, &o.Ids)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Permissions) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

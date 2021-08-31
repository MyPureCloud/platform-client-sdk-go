package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Trusteeauthorization
type Trusteeauthorization struct { 
	// Permissions - Permissions that the trustee user has in the trustor organization
	Permissions *[]string `json:"permissions,omitempty"`

}

func (o *Trusteeauthorization) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Trusteeauthorization
	
	return json.Marshal(&struct { 
		Permissions *[]string `json:"permissions,omitempty"`
		*Alias
	}{ 
		Permissions: o.Permissions,
		Alias:    (*Alias)(o),
	})
}

func (o *Trusteeauthorization) UnmarshalJSON(b []byte) error {
	var TrusteeauthorizationMap map[string]interface{}
	err := json.Unmarshal(b, &TrusteeauthorizationMap)
	if err != nil {
		return err
	}
	
	if Permissions, ok := TrusteeauthorizationMap["permissions"].([]interface{}); ok {
		PermissionsString, _ := json.Marshal(Permissions)
		json.Unmarshal(PermissionsString, &o.Permissions)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Trusteeauthorization) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

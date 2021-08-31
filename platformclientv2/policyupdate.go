package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Policyupdate
type Policyupdate struct { 
	// Enabled
	Enabled *bool `json:"enabled,omitempty"`

}

func (o *Policyupdate) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Policyupdate
	
	return json.Marshal(&struct { 
		Enabled *bool `json:"enabled,omitempty"`
		*Alias
	}{ 
		Enabled: o.Enabled,
		Alias:    (*Alias)(o),
	})
}

func (o *Policyupdate) UnmarshalJSON(b []byte) error {
	var PolicyupdateMap map[string]interface{}
	err := json.Unmarshal(b, &PolicyupdateMap)
	if err != nil {
		return err
	}
	
	if Enabled, ok := PolicyupdateMap["enabled"].(bool); ok {
		o.Enabled = &Enabled
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Policyupdate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Crossplatformpolicyupdate
type Crossplatformpolicyupdate struct { 
	// Enabled
	Enabled *bool `json:"enabled,omitempty"`

}

func (o *Crossplatformpolicyupdate) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Crossplatformpolicyupdate
	
	return json.Marshal(&struct { 
		Enabled *bool `json:"enabled,omitempty"`
		*Alias
	}{ 
		Enabled: o.Enabled,
		Alias:    (*Alias)(o),
	})
}

func (o *Crossplatformpolicyupdate) UnmarshalJSON(b []byte) error {
	var CrossplatformpolicyupdateMap map[string]interface{}
	err := json.Unmarshal(b, &CrossplatformpolicyupdateMap)
	if err != nil {
		return err
	}
	
	if Enabled, ok := CrossplatformpolicyupdateMap["enabled"].(bool); ok {
		o.Enabled = &Enabled
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Crossplatformpolicyupdate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

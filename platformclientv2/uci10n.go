package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Uci10n
type Uci10n struct { 
	// Name
	Name *string `json:"name,omitempty"`

}

func (o *Uci10n) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Uci10n
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		*Alias
	}{ 
		Name: o.Name,
		Alias:    (*Alias)(o),
	})
}

func (o *Uci10n) UnmarshalJSON(b []byte) error {
	var Uci10nMap map[string]interface{}
	err := json.Unmarshal(b, &Uci10nMap)
	if err != nil {
		return err
	}
	
	if Name, ok := Uci10nMap["name"].(string); ok {
		o.Name = &Name
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Uci10n) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

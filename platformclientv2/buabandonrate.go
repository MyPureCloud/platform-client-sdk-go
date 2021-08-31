package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Buabandonrate - Service goal abandon rate configuration
type Buabandonrate struct { 
	// Include - Whether to include abandon rate in the associated configuration
	Include *bool `json:"include,omitempty"`


	// Percent - Abandon rate percent goal. Required if include == true
	Percent *int `json:"percent,omitempty"`

}

func (o *Buabandonrate) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Buabandonrate
	
	return json.Marshal(&struct { 
		Include *bool `json:"include,omitempty"`
		
		Percent *int `json:"percent,omitempty"`
		*Alias
	}{ 
		Include: o.Include,
		
		Percent: o.Percent,
		Alias:    (*Alias)(o),
	})
}

func (o *Buabandonrate) UnmarshalJSON(b []byte) error {
	var BuabandonrateMap map[string]interface{}
	err := json.Unmarshal(b, &BuabandonrateMap)
	if err != nil {
		return err
	}
	
	if Include, ok := BuabandonrateMap["include"].(bool); ok {
		o.Include = &Include
	}
	
	if Percent, ok := BuabandonrateMap["percent"].(float64); ok {
		PercentInt := int(Percent)
		o.Percent = &PercentInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Buabandonrate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

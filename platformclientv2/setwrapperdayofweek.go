package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Setwrapperdayofweek
type Setwrapperdayofweek struct { 
	// Values
	Values *[]string `json:"values,omitempty"`

}

func (o *Setwrapperdayofweek) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Setwrapperdayofweek
	
	return json.Marshal(&struct { 
		Values *[]string `json:"values,omitempty"`
		*Alias
	}{ 
		Values: o.Values,
		Alias:    (*Alias)(o),
	})
}

func (o *Setwrapperdayofweek) UnmarshalJSON(b []byte) error {
	var SetwrapperdayofweekMap map[string]interface{}
	err := json.Unmarshal(b, &SetwrapperdayofweekMap)
	if err != nil {
		return err
	}
	
	if Values, ok := SetwrapperdayofweekMap["values"].([]interface{}); ok {
		ValuesString, _ := json.Marshal(Values)
		json.Unmarshal(ValuesString, &o.Values)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Setwrapperdayofweek) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

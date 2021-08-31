package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Listwrapperinterval
type Listwrapperinterval struct { 
	// Values
	Values *[]string `json:"values,omitempty"`

}

func (o *Listwrapperinterval) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Listwrapperinterval
	
	return json.Marshal(&struct { 
		Values *[]string `json:"values,omitempty"`
		*Alias
	}{ 
		Values: o.Values,
		Alias:    (*Alias)(o),
	})
}

func (o *Listwrapperinterval) UnmarshalJSON(b []byte) error {
	var ListwrapperintervalMap map[string]interface{}
	err := json.Unmarshal(b, &ListwrapperintervalMap)
	if err != nil {
		return err
	}
	
	if Values, ok := ListwrapperintervalMap["values"].([]interface{}); ok {
		ValuesString, _ := json.Marshal(Values)
		json.Unmarshal(ValuesString, &o.Values)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Listwrapperinterval) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

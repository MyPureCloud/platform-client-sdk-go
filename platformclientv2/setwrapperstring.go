package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Setwrapperstring
type Setwrapperstring struct { 
	// Values
	Values *[]string `json:"values,omitempty"`

}

func (o *Setwrapperstring) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Setwrapperstring
	
	return json.Marshal(&struct { 
		Values *[]string `json:"values,omitempty"`
		*Alias
	}{ 
		Values: o.Values,
		Alias:    (*Alias)(o),
	})
}

func (o *Setwrapperstring) UnmarshalJSON(b []byte) error {
	var SetwrapperstringMap map[string]interface{}
	err := json.Unmarshal(b, &SetwrapperstringMap)
	if err != nil {
		return err
	}
	
	if Values, ok := SetwrapperstringMap["values"].([]interface{}); ok {
		ValuesString, _ := json.Marshal(Values)
		json.Unmarshal(ValuesString, &o.Values)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Setwrapperstring) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

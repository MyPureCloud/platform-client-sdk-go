package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Listwrappershiftstartvariance
type Listwrappershiftstartvariance struct { 
	// Values
	Values *[]Shiftstartvariance `json:"values,omitempty"`

}

func (o *Listwrappershiftstartvariance) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Listwrappershiftstartvariance
	
	return json.Marshal(&struct { 
		Values *[]Shiftstartvariance `json:"values,omitempty"`
		*Alias
	}{ 
		Values: o.Values,
		Alias:    (*Alias)(o),
	})
}

func (o *Listwrappershiftstartvariance) UnmarshalJSON(b []byte) error {
	var ListwrappershiftstartvarianceMap map[string]interface{}
	err := json.Unmarshal(b, &ListwrappershiftstartvarianceMap)
	if err != nil {
		return err
	}
	
	if Values, ok := ListwrappershiftstartvarianceMap["values"].([]interface{}); ok {
		ValuesString, _ := json.Marshal(Values)
		json.Unmarshal(ValuesString, &o.Values)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Listwrappershiftstartvariance) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

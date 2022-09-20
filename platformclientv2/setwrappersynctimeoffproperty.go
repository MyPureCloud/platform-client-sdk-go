package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Setwrappersynctimeoffproperty
type Setwrappersynctimeoffproperty struct { 
	// Values
	Values *[]string `json:"values,omitempty"`

}

func (o *Setwrappersynctimeoffproperty) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Setwrappersynctimeoffproperty
	
	return json.Marshal(&struct { 
		Values *[]string `json:"values,omitempty"`
		*Alias
	}{ 
		Values: o.Values,
		Alias:    (*Alias)(o),
	})
}

func (o *Setwrappersynctimeoffproperty) UnmarshalJSON(b []byte) error {
	var SetwrappersynctimeoffpropertyMap map[string]interface{}
	err := json.Unmarshal(b, &SetwrappersynctimeoffpropertyMap)
	if err != nil {
		return err
	}
	
	if Values, ok := SetwrappersynctimeoffpropertyMap["values"].([]interface{}); ok {
		ValuesString, _ := json.Marshal(Values)
		json.Unmarshal(ValuesString, &o.Values)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Setwrappersynctimeoffproperty) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

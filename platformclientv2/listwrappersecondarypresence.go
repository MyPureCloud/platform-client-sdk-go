package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Listwrappersecondarypresence
type Listwrappersecondarypresence struct { 
	// Values
	Values *[]Secondarypresence `json:"values,omitempty"`

}

func (o *Listwrappersecondarypresence) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Listwrappersecondarypresence
	
	return json.Marshal(&struct { 
		Values *[]Secondarypresence `json:"values,omitempty"`
		*Alias
	}{ 
		Values: o.Values,
		Alias:    (*Alias)(o),
	})
}

func (o *Listwrappersecondarypresence) UnmarshalJSON(b []byte) error {
	var ListwrappersecondarypresenceMap map[string]interface{}
	err := json.Unmarshal(b, &ListwrappersecondarypresenceMap)
	if err != nil {
		return err
	}
	
	if Values, ok := ListwrappersecondarypresenceMap["values"].([]interface{}); ok {
		ValuesString, _ := json.Marshal(Values)
		json.Unmarshal(ValuesString, &o.Values)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Listwrappersecondarypresence) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

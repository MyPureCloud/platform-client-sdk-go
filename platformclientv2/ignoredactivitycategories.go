package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Ignoredactivitycategories
type Ignoredactivitycategories struct { 
	// Values - Activity categories list
	Values *[]string `json:"values,omitempty"`

}

func (o *Ignoredactivitycategories) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Ignoredactivitycategories
	
	return json.Marshal(&struct { 
		Values *[]string `json:"values,omitempty"`
		*Alias
	}{ 
		Values: o.Values,
		Alias:    (*Alias)(o),
	})
}

func (o *Ignoredactivitycategories) UnmarshalJSON(b []byte) error {
	var IgnoredactivitycategoriesMap map[string]interface{}
	err := json.Unmarshal(b, &IgnoredactivitycategoriesMap)
	if err != nil {
		return err
	}
	
	if Values, ok := IgnoredactivitycategoriesMap["values"].([]interface{}); ok {
		ValuesString, _ := json.Marshal(Values)
		json.Unmarshal(ValuesString, &o.Values)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Ignoredactivitycategories) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

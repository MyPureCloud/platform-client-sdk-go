package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Contextpattern
type Contextpattern struct { 
	// Criteria - A list of one or more criteria to satisfy.
	Criteria *[]Entitytypecriteria `json:"criteria,omitempty"`

}

func (o *Contextpattern) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Contextpattern
	
	return json.Marshal(&struct { 
		Criteria *[]Entitytypecriteria `json:"criteria,omitempty"`
		*Alias
	}{ 
		Criteria: o.Criteria,
		Alias:    (*Alias)(o),
	})
}

func (o *Contextpattern) UnmarshalJSON(b []byte) error {
	var ContextpatternMap map[string]interface{}
	err := json.Unmarshal(b, &ContextpatternMap)
	if err != nil {
		return err
	}
	
	if Criteria, ok := ContextpatternMap["criteria"].([]interface{}); ok {
		CriteriaString, _ := json.Marshal(Criteria)
		json.Unmarshal(CriteriaString, &o.Criteria)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Contextpattern) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

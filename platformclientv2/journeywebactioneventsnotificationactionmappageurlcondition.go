package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Journeywebactioneventsnotificationactionmappageurlcondition
type Journeywebactioneventsnotificationactionmappageurlcondition struct { 
	// Values
	Values *[]string `json:"values,omitempty"`


	// Operator
	Operator *string `json:"operator,omitempty"`

}

func (o *Journeywebactioneventsnotificationactionmappageurlcondition) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Journeywebactioneventsnotificationactionmappageurlcondition
	
	return json.Marshal(&struct { 
		Values *[]string `json:"values,omitempty"`
		
		Operator *string `json:"operator,omitempty"`
		*Alias
	}{ 
		Values: o.Values,
		
		Operator: o.Operator,
		Alias:    (*Alias)(o),
	})
}

func (o *Journeywebactioneventsnotificationactionmappageurlcondition) UnmarshalJSON(b []byte) error {
	var JourneywebactioneventsnotificationactionmappageurlconditionMap map[string]interface{}
	err := json.Unmarshal(b, &JourneywebactioneventsnotificationactionmappageurlconditionMap)
	if err != nil {
		return err
	}
	
	if Values, ok := JourneywebactioneventsnotificationactionmappageurlconditionMap["values"].([]interface{}); ok {
		ValuesString, _ := json.Marshal(Values)
		json.Unmarshal(ValuesString, &o.Values)
	}
	
	if Operator, ok := JourneywebactioneventsnotificationactionmappageurlconditionMap["operator"].(string); ok {
		o.Operator = &Operator
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Journeywebactioneventsnotificationactionmappageurlcondition) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

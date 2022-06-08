package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Journeysessioneventsnotificationcustomeventattributelist
type Journeysessioneventsnotificationcustomeventattributelist struct { 
	// Values
	Values *[]string `json:"values,omitempty"`


	// DataType
	DataType *string `json:"dataType,omitempty"`

}

func (o *Journeysessioneventsnotificationcustomeventattributelist) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Journeysessioneventsnotificationcustomeventattributelist
	
	return json.Marshal(&struct { 
		Values *[]string `json:"values,omitempty"`
		
		DataType *string `json:"dataType,omitempty"`
		*Alias
	}{ 
		Values: o.Values,
		
		DataType: o.DataType,
		Alias:    (*Alias)(o),
	})
}

func (o *Journeysessioneventsnotificationcustomeventattributelist) UnmarshalJSON(b []byte) error {
	var JourneysessioneventsnotificationcustomeventattributelistMap map[string]interface{}
	err := json.Unmarshal(b, &JourneysessioneventsnotificationcustomeventattributelistMap)
	if err != nil {
		return err
	}
	
	if Values, ok := JourneysessioneventsnotificationcustomeventattributelistMap["values"].([]interface{}); ok {
		ValuesString, _ := json.Marshal(Values)
		json.Unmarshal(ValuesString, &o.Values)
	}
	
	if DataType, ok := JourneysessioneventsnotificationcustomeventattributelistMap["dataType"].(string); ok {
		o.DataType = &DataType
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Journeysessioneventsnotificationcustomeventattributelist) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

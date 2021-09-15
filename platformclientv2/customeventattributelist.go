package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Customeventattributelist
type Customeventattributelist struct { 
	// DataType - The data type of the custom attributes.
	DataType *string `json:"dataType,omitempty"`


	// Values - The list of custom event attribute values.
	Values *[]string `json:"values,omitempty"`

}

func (o *Customeventattributelist) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Customeventattributelist
	
	return json.Marshal(&struct { 
		DataType *string `json:"dataType,omitempty"`
		
		Values *[]string `json:"values,omitempty"`
		*Alias
	}{ 
		DataType: o.DataType,
		
		Values: o.Values,
		Alias:    (*Alias)(o),
	})
}

func (o *Customeventattributelist) UnmarshalJSON(b []byte) error {
	var CustomeventattributelistMap map[string]interface{}
	err := json.Unmarshal(b, &CustomeventattributelistMap)
	if err != nil {
		return err
	}
	
	if DataType, ok := CustomeventattributelistMap["dataType"].(string); ok {
		o.DataType = &DataType
	}
	
	if Values, ok := CustomeventattributelistMap["values"].([]interface{}); ok {
		ValuesString, _ := json.Marshal(Values)
		json.Unmarshal(ValuesString, &o.Values)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Customeventattributelist) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Customeventattribute
type Customeventattribute struct { 
	// DataType - The data type of the custom attribute.
	DataType *string `json:"dataType,omitempty"`


	// Value - The value of the custom attribute.
	Value *string `json:"value,omitempty"`

}

func (o *Customeventattribute) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Customeventattribute
	
	return json.Marshal(&struct { 
		DataType *string `json:"dataType,omitempty"`
		
		Value *string `json:"value,omitempty"`
		*Alias
	}{ 
		DataType: o.DataType,
		
		Value: o.Value,
		Alias:    (*Alias)(o),
	})
}

func (o *Customeventattribute) UnmarshalJSON(b []byte) error {
	var CustomeventattributeMap map[string]interface{}
	err := json.Unmarshal(b, &CustomeventattributeMap)
	if err != nil {
		return err
	}
	
	if DataType, ok := CustomeventattributeMap["dataType"].(string); ok {
		o.DataType = &DataType
	}
	
	if Value, ok := CustomeventattributeMap["value"].(string); ok {
		o.Value = &Value
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Customeventattribute) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

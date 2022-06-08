package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Journeysessioneventsnotificationcustomeventattribute
type Journeysessioneventsnotificationcustomeventattribute struct { 
	// Value
	Value *string `json:"value,omitempty"`


	// DataType
	DataType *string `json:"dataType,omitempty"`

}

func (o *Journeysessioneventsnotificationcustomeventattribute) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Journeysessioneventsnotificationcustomeventattribute
	
	return json.Marshal(&struct { 
		Value *string `json:"value,omitempty"`
		
		DataType *string `json:"dataType,omitempty"`
		*Alias
	}{ 
		Value: o.Value,
		
		DataType: o.DataType,
		Alias:    (*Alias)(o),
	})
}

func (o *Journeysessioneventsnotificationcustomeventattribute) UnmarshalJSON(b []byte) error {
	var JourneysessioneventsnotificationcustomeventattributeMap map[string]interface{}
	err := json.Unmarshal(b, &JourneysessioneventsnotificationcustomeventattributeMap)
	if err != nil {
		return err
	}
	
	if Value, ok := JourneysessioneventsnotificationcustomeventattributeMap["value"].(string); ok {
		o.Value = &Value
	}
    
	if DataType, ok := JourneysessioneventsnotificationcustomeventattributeMap["dataType"].(string); ok {
		o.DataType = &DataType
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Journeysessioneventsnotificationcustomeventattribute) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

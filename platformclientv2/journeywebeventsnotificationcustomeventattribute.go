package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Journeywebeventsnotificationcustomeventattribute
type Journeywebeventsnotificationcustomeventattribute struct { 
	// Value
	Value *string `json:"value,omitempty"`


	// DataType
	DataType *string `json:"dataType,omitempty"`

}

func (o *Journeywebeventsnotificationcustomeventattribute) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Journeywebeventsnotificationcustomeventattribute
	
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

func (o *Journeywebeventsnotificationcustomeventattribute) UnmarshalJSON(b []byte) error {
	var JourneywebeventsnotificationcustomeventattributeMap map[string]interface{}
	err := json.Unmarshal(b, &JourneywebeventsnotificationcustomeventattributeMap)
	if err != nil {
		return err
	}
	
	if Value, ok := JourneywebeventsnotificationcustomeventattributeMap["value"].(string); ok {
		o.Value = &Value
	}
	
	if DataType, ok := JourneywebeventsnotificationcustomeventattributeMap["dataType"].(string); ok {
		o.DataType = &DataType
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Journeywebeventsnotificationcustomeventattribute) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

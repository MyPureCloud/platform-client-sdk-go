package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Journeywebeventsnotificationassociatedvalue
type Journeywebeventsnotificationassociatedvalue struct { 
	// DataType
	DataType *string `json:"dataType,omitempty"`


	// Value
	Value *float32 `json:"value,omitempty"`

}

func (o *Journeywebeventsnotificationassociatedvalue) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Journeywebeventsnotificationassociatedvalue
	
	return json.Marshal(&struct { 
		DataType *string `json:"dataType,omitempty"`
		
		Value *float32 `json:"value,omitempty"`
		*Alias
	}{ 
		DataType: o.DataType,
		
		Value: o.Value,
		Alias:    (*Alias)(o),
	})
}

func (o *Journeywebeventsnotificationassociatedvalue) UnmarshalJSON(b []byte) error {
	var JourneywebeventsnotificationassociatedvalueMap map[string]interface{}
	err := json.Unmarshal(b, &JourneywebeventsnotificationassociatedvalueMap)
	if err != nil {
		return err
	}
	
	if DataType, ok := JourneywebeventsnotificationassociatedvalueMap["dataType"].(string); ok {
		o.DataType = &DataType
	}
    
	if Value, ok := JourneywebeventsnotificationassociatedvalueMap["value"].(float64); ok {
		ValueFloat32 := float32(Value)
		o.Value = &ValueFloat32
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Journeywebeventsnotificationassociatedvalue) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

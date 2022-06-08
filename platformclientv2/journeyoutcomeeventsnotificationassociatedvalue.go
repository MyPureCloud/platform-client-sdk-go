package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Journeyoutcomeeventsnotificationassociatedvalue
type Journeyoutcomeeventsnotificationassociatedvalue struct { 
	// DataType
	DataType *string `json:"dataType,omitempty"`


	// Value
	Value *float32 `json:"value,omitempty"`

}

func (o *Journeyoutcomeeventsnotificationassociatedvalue) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Journeyoutcomeeventsnotificationassociatedvalue
	
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

func (o *Journeyoutcomeeventsnotificationassociatedvalue) UnmarshalJSON(b []byte) error {
	var JourneyoutcomeeventsnotificationassociatedvalueMap map[string]interface{}
	err := json.Unmarshal(b, &JourneyoutcomeeventsnotificationassociatedvalueMap)
	if err != nil {
		return err
	}
	
	if DataType, ok := JourneyoutcomeeventsnotificationassociatedvalueMap["dataType"].(string); ok {
		o.DataType = &DataType
	}
    
	if Value, ok := JourneyoutcomeeventsnotificationassociatedvalueMap["value"].(float64); ok {
		ValueFloat32 := float32(Value)
		o.Value = &ValueFloat32
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Journeyoutcomeeventsnotificationassociatedvalue) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

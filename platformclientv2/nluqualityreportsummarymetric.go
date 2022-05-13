package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Nluqualityreportsummarymetric
type Nluqualityreportsummarymetric struct { 
	// Name - The name of the metric. e.g. recall, f1_score
	Name *string `json:"name,omitempty"`


	// Value - The value of the metric
	Value *float32 `json:"value,omitempty"`

}

func (o *Nluqualityreportsummarymetric) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Nluqualityreportsummarymetric
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		Value *float32 `json:"value,omitempty"`
		*Alias
	}{ 
		Name: o.Name,
		
		Value: o.Value,
		Alias:    (*Alias)(o),
	})
}

func (o *Nluqualityreportsummarymetric) UnmarshalJSON(b []byte) error {
	var NluqualityreportsummarymetricMap map[string]interface{}
	err := json.Unmarshal(b, &NluqualityreportsummarymetricMap)
	if err != nil {
		return err
	}
	
	if Name, ok := NluqualityreportsummarymetricMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Value, ok := NluqualityreportsummarymetricMap["value"].(float64); ok {
		ValueFloat32 := float32(Value)
		o.Value = &ValueFloat32
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Nluqualityreportsummarymetric) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

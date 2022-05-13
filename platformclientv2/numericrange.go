package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Numericrange
type Numericrange struct { 
	// Gt - Greater than
	Gt *float32 `json:"gt,omitempty"`


	// Gte - Greater than or equal to
	Gte *float32 `json:"gte,omitempty"`


	// Lt - Less than
	Lt *float32 `json:"lt,omitempty"`


	// Lte - Less than or equal to
	Lte *float32 `json:"lte,omitempty"`

}

func (o *Numericrange) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Numericrange
	
	return json.Marshal(&struct { 
		Gt *float32 `json:"gt,omitempty"`
		
		Gte *float32 `json:"gte,omitempty"`
		
		Lt *float32 `json:"lt,omitempty"`
		
		Lte *float32 `json:"lte,omitempty"`
		*Alias
	}{ 
		Gt: o.Gt,
		
		Gte: o.Gte,
		
		Lt: o.Lt,
		
		Lte: o.Lte,
		Alias:    (*Alias)(o),
	})
}

func (o *Numericrange) UnmarshalJSON(b []byte) error {
	var NumericrangeMap map[string]interface{}
	err := json.Unmarshal(b, &NumericrangeMap)
	if err != nil {
		return err
	}
	
	if Gt, ok := NumericrangeMap["gt"].(float64); ok {
		GtFloat32 := float32(Gt)
		o.Gt = &GtFloat32
	}
    
	if Gte, ok := NumericrangeMap["gte"].(float64); ok {
		GteFloat32 := float32(Gte)
		o.Gte = &GteFloat32
	}
    
	if Lt, ok := NumericrangeMap["lt"].(float64); ok {
		LtFloat32 := float32(Lt)
		o.Lt = &LtFloat32
	}
    
	if Lte, ok := NumericrangeMap["lte"].(float64); ok {
		LteFloat32 := float32(Lte)
		o.Lte = &LteFloat32
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Numericrange) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

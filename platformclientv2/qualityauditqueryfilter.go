package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Qualityauditqueryfilter
type Qualityauditqueryfilter struct { 
	// Property - Name of the property to filter.
	Property *string `json:"property,omitempty"`


	// Value - Value of the property to filter.
	Value *string `json:"value,omitempty"`

}

func (o *Qualityauditqueryfilter) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Qualityauditqueryfilter
	
	return json.Marshal(&struct { 
		Property *string `json:"property,omitempty"`
		
		Value *string `json:"value,omitempty"`
		*Alias
	}{ 
		Property: o.Property,
		
		Value: o.Value,
		Alias:    (*Alias)(o),
	})
}

func (o *Qualityauditqueryfilter) UnmarshalJSON(b []byte) error {
	var QualityauditqueryfilterMap map[string]interface{}
	err := json.Unmarshal(b, &QualityauditqueryfilterMap)
	if err != nil {
		return err
	}
	
	if Property, ok := QualityauditqueryfilterMap["property"].(string); ok {
		o.Property = &Property
	}
    
	if Value, ok := QualityauditqueryfilterMap["value"].(string); ok {
		o.Value = &Value
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Qualityauditqueryfilter) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

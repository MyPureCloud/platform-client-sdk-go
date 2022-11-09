package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Actionaggregationview
type Actionaggregationview struct { 
	// Target - Target metric name
	Target *string `json:"target,omitempty"`


	// Name - A unique name for this view. Must be distinct from other views and built-in metric names.
	Name *string `json:"name,omitempty"`


	// Function - Type of view you wish to create
	Function *string `json:"function,omitempty"`


	// VarRange - Range of numbers for slicing up data
	VarRange *Aggregationrange `json:"range,omitempty"`

}

func (o *Actionaggregationview) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Actionaggregationview
	
	return json.Marshal(&struct { 
		Target *string `json:"target,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Function *string `json:"function,omitempty"`
		
		VarRange *Aggregationrange `json:"range,omitempty"`
		*Alias
	}{ 
		Target: o.Target,
		
		Name: o.Name,
		
		Function: o.Function,
		
		VarRange: o.VarRange,
		Alias:    (*Alias)(o),
	})
}

func (o *Actionaggregationview) UnmarshalJSON(b []byte) error {
	var ActionaggregationviewMap map[string]interface{}
	err := json.Unmarshal(b, &ActionaggregationviewMap)
	if err != nil {
		return err
	}
	
	if Target, ok := ActionaggregationviewMap["target"].(string); ok {
		o.Target = &Target
	}
    
	if Name, ok := ActionaggregationviewMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Function, ok := ActionaggregationviewMap["function"].(string); ok {
		o.Function = &Function
	}
    
	if VarRange, ok := ActionaggregationviewMap["range"].(map[string]interface{}); ok {
		VarRangeString, _ := json.Marshal(VarRange)
		json.Unmarshal(VarRangeString, &o.VarRange)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Actionaggregationview) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

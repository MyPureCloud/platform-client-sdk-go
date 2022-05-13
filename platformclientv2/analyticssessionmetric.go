package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Analyticssessionmetric
type Analyticssessionmetric struct { 
	// EmitDate - Metric emission date. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	EmitDate *time.Time `json:"emitDate,omitempty"`


	// Name - Unique name of this metric
	Name *string `json:"name,omitempty"`


	// Value - The metric value
	Value *int `json:"value,omitempty"`

}

func (o *Analyticssessionmetric) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Analyticssessionmetric
	
	EmitDate := new(string)
	if o.EmitDate != nil {
		
		*EmitDate = timeutil.Strftime(o.EmitDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		EmitDate = nil
	}
	
	return json.Marshal(&struct { 
		EmitDate *string `json:"emitDate,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Value *int `json:"value,omitempty"`
		*Alias
	}{ 
		EmitDate: EmitDate,
		
		Name: o.Name,
		
		Value: o.Value,
		Alias:    (*Alias)(o),
	})
}

func (o *Analyticssessionmetric) UnmarshalJSON(b []byte) error {
	var AnalyticssessionmetricMap map[string]interface{}
	err := json.Unmarshal(b, &AnalyticssessionmetricMap)
	if err != nil {
		return err
	}
	
	if emitDateString, ok := AnalyticssessionmetricMap["emitDate"].(string); ok {
		EmitDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", emitDateString)
		o.EmitDate = &EmitDate
	}
	
	if Name, ok := AnalyticssessionmetricMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Value, ok := AnalyticssessionmetricMap["value"].(float64); ok {
		ValueInt := int(Value)
		o.Value = &ValueInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Analyticssessionmetric) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

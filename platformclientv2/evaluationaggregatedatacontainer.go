package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Evaluationaggregatedatacontainer
type Evaluationaggregatedatacontainer struct { 
	// Group - A mapping from dimension to value
	Group *map[string]string `json:"group,omitempty"`


	// Data
	Data *[]Statisticalresponse `json:"data,omitempty"`

}

func (o *Evaluationaggregatedatacontainer) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Evaluationaggregatedatacontainer
	
	return json.Marshal(&struct { 
		Group *map[string]string `json:"group,omitempty"`
		
		Data *[]Statisticalresponse `json:"data,omitempty"`
		*Alias
	}{ 
		Group: o.Group,
		
		Data: o.Data,
		Alias:    (*Alias)(o),
	})
}

func (o *Evaluationaggregatedatacontainer) UnmarshalJSON(b []byte) error {
	var EvaluationaggregatedatacontainerMap map[string]interface{}
	err := json.Unmarshal(b, &EvaluationaggregatedatacontainerMap)
	if err != nil {
		return err
	}
	
	if Group, ok := EvaluationaggregatedatacontainerMap["group"].(map[string]interface{}); ok {
		GroupString, _ := json.Marshal(Group)
		json.Unmarshal(GroupString, &o.Group)
	}
	
	if Data, ok := EvaluationaggregatedatacontainerMap["data"].([]interface{}); ok {
		DataString, _ := json.Marshal(Data)
		json.Unmarshal(DataString, &o.Data)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Evaluationaggregatedatacontainer) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

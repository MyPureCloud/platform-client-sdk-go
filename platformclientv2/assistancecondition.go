package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Assistancecondition
type Assistancecondition struct { 
	// Operator
	Operator *string `json:"operator,omitempty"`


	// TopicIds
	TopicIds *[]string `json:"topicIds,omitempty"`

}

func (o *Assistancecondition) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Assistancecondition
	
	return json.Marshal(&struct { 
		Operator *string `json:"operator,omitempty"`
		
		TopicIds *[]string `json:"topicIds,omitempty"`
		*Alias
	}{ 
		Operator: o.Operator,
		
		TopicIds: o.TopicIds,
		Alias:    (*Alias)(o),
	})
}

func (o *Assistancecondition) UnmarshalJSON(b []byte) error {
	var AssistanceconditionMap map[string]interface{}
	err := json.Unmarshal(b, &AssistanceconditionMap)
	if err != nil {
		return err
	}
	
	if Operator, ok := AssistanceconditionMap["operator"].(string); ok {
		o.Operator = &Operator
	}
	
	if TopicIds, ok := AssistanceconditionMap["topicIds"].([]interface{}); ok {
		TopicIdsString, _ := json.Marshal(TopicIds)
		json.Unmarshal(TopicIdsString, &o.TopicIds)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Assistancecondition) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

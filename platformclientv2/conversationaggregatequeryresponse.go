package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationaggregatequeryresponse
type Conversationaggregatequeryresponse struct { 
	// Results
	Results *[]Conversationaggregatedatacontainer `json:"results,omitempty"`

}

func (o *Conversationaggregatequeryresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationaggregatequeryresponse
	
	return json.Marshal(&struct { 
		Results *[]Conversationaggregatedatacontainer `json:"results,omitempty"`
		*Alias
	}{ 
		Results: o.Results,
		Alias:    (*Alias)(o),
	})
}

func (o *Conversationaggregatequeryresponse) UnmarshalJSON(b []byte) error {
	var ConversationaggregatequeryresponseMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationaggregatequeryresponseMap)
	if err != nil {
		return err
	}
	
	if Results, ok := ConversationaggregatequeryresponseMap["results"].([]interface{}); ok {
		ResultsString, _ := json.Marshal(Results)
		json.Unmarshal(ResultsString, &o.Results)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationaggregatequeryresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

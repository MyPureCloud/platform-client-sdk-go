package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Requestscoredagent
type Requestscoredagent struct { 
	// Id - Agent's user ID
	Id *string `json:"id,omitempty"`


	// Score - Agent's score for the current conversation, from 0 - 100, higher being better
	Score *int `json:"score,omitempty"`

}

func (o *Requestscoredagent) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Requestscoredagent
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Score *int `json:"score,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Score: o.Score,
		Alias:    (*Alias)(o),
	})
}

func (o *Requestscoredagent) UnmarshalJSON(b []byte) error {
	var RequestscoredagentMap map[string]interface{}
	err := json.Unmarshal(b, &RequestscoredagentMap)
	if err != nil {
		return err
	}
	
	if Id, ok := RequestscoredagentMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Score, ok := RequestscoredagentMap["score"].(float64); ok {
		ScoreInt := int(Score)
		o.Score = &ScoreInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Requestscoredagent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

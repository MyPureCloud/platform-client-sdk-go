package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Searchupdaterequest
type Searchupdaterequest struct { 
	// SessionId - The unique identifier of this session
	SessionId *string `json:"sessionId,omitempty"`


	// Answered - Mark the search as answered/unanswered
	Answered *bool `json:"answered,omitempty"`


	// SelectedAnswer - The selected search result chosen as the answer.
	SelectedAnswer *Selectedanswer `json:"selectedAnswer,omitempty"`

}

func (o *Searchupdaterequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Searchupdaterequest
	
	return json.Marshal(&struct { 
		SessionId *string `json:"sessionId,omitempty"`
		
		Answered *bool `json:"answered,omitempty"`
		
		SelectedAnswer *Selectedanswer `json:"selectedAnswer,omitempty"`
		*Alias
	}{ 
		SessionId: o.SessionId,
		
		Answered: o.Answered,
		
		SelectedAnswer: o.SelectedAnswer,
		Alias:    (*Alias)(o),
	})
}

func (o *Searchupdaterequest) UnmarshalJSON(b []byte) error {
	var SearchupdaterequestMap map[string]interface{}
	err := json.Unmarshal(b, &SearchupdaterequestMap)
	if err != nil {
		return err
	}
	
	if SessionId, ok := SearchupdaterequestMap["sessionId"].(string); ok {
		o.SessionId = &SessionId
	}
    
	if Answered, ok := SearchupdaterequestMap["answered"].(bool); ok {
		o.Answered = &Answered
	}
    
	if SelectedAnswer, ok := SearchupdaterequestMap["selectedAnswer"].(map[string]interface{}); ok {
		SelectedAnswerString, _ := json.Marshal(SelectedAnswer)
		json.Unmarshal(SelectedAnswerString, &o.SelectedAnswer)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Searchupdaterequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

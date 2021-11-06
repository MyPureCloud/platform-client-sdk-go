package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Sentimentfeedbackentitylisting
type Sentimentfeedbackentitylisting struct { 
	// Entities
	Entities *[]Sentimentfeedback `json:"entities,omitempty"`

}

func (o *Sentimentfeedbackentitylisting) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Sentimentfeedbackentitylisting
	
	return json.Marshal(&struct { 
		Entities *[]Sentimentfeedback `json:"entities,omitempty"`
		*Alias
	}{ 
		Entities: o.Entities,
		Alias:    (*Alias)(o),
	})
}

func (o *Sentimentfeedbackentitylisting) UnmarshalJSON(b []byte) error {
	var SentimentfeedbackentitylistingMap map[string]interface{}
	err := json.Unmarshal(b, &SentimentfeedbackentitylistingMap)
	if err != nil {
		return err
	}
	
	if Entities, ok := SentimentfeedbackentitylistingMap["entities"].([]interface{}); ok {
		EntitiesString, _ := json.Marshal(Entities)
		json.Unmarshal(EntitiesString, &o.Entities)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Sentimentfeedbackentitylisting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

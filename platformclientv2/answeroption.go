package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Answeroption
type Answeroption struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Text
	Text *string `json:"text,omitempty"`


	// Value
	Value *int `json:"value,omitempty"`


	// AssistanceConditions - List of assistance conditions which are combined together with a logical AND operator. Eg ( assistanceCondtion1 && assistanceCondition2 ) wherein assistanceCondition could be ( EXISTS topic1 || topic2 || ... ) or (NOTEXISTS topic3 || topic4 || ...).
	AssistanceConditions *[]Assistancecondition `json:"assistanceConditions,omitempty"`

}

func (o *Answeroption) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Answeroption
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Text *string `json:"text,omitempty"`
		
		Value *int `json:"value,omitempty"`
		
		AssistanceConditions *[]Assistancecondition `json:"assistanceConditions,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Text: o.Text,
		
		Value: o.Value,
		
		AssistanceConditions: o.AssistanceConditions,
		Alias:    (*Alias)(o),
	})
}

func (o *Answeroption) UnmarshalJSON(b []byte) error {
	var AnsweroptionMap map[string]interface{}
	err := json.Unmarshal(b, &AnsweroptionMap)
	if err != nil {
		return err
	}
	
	if Id, ok := AnsweroptionMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Text, ok := AnsweroptionMap["text"].(string); ok {
		o.Text = &Text
	}
    
	if Value, ok := AnsweroptionMap["value"].(float64); ok {
		ValueInt := int(Value)
		o.Value = &ValueInt
	}
	
	if AssistanceConditions, ok := AnsweroptionMap["assistanceConditions"].([]interface{}); ok {
		AssistanceConditionsString, _ := json.Marshal(AssistanceConditions)
		json.Unmarshal(AssistanceConditionsString, &o.AssistanceConditions)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Answeroption) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

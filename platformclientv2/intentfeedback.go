package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Intentfeedback
type Intentfeedback struct { 
	// Name - The name of the detected intent.
	Name *string `json:"name,omitempty"`


	// Probability - The probability of the detected intent.
	Probability *float64 `json:"probability,omitempty"`


	// Entities - The collection of named entities detected.
	Entities *[]Detectednamedentity `json:"entities,omitempty"`


	// Assessment - The assessment on the detection for feedback text.
	Assessment *string `json:"assessment,omitempty"`

}

func (o *Intentfeedback) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Intentfeedback
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		Probability *float64 `json:"probability,omitempty"`
		
		Entities *[]Detectednamedentity `json:"entities,omitempty"`
		
		Assessment *string `json:"assessment,omitempty"`
		*Alias
	}{ 
		Name: o.Name,
		
		Probability: o.Probability,
		
		Entities: o.Entities,
		
		Assessment: o.Assessment,
		Alias:    (*Alias)(o),
	})
}

func (o *Intentfeedback) UnmarshalJSON(b []byte) error {
	var IntentfeedbackMap map[string]interface{}
	err := json.Unmarshal(b, &IntentfeedbackMap)
	if err != nil {
		return err
	}
	
	if Name, ok := IntentfeedbackMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if Probability, ok := IntentfeedbackMap["probability"].(float64); ok {
		o.Probability = &Probability
	}
	
	if Entities, ok := IntentfeedbackMap["entities"].([]interface{}); ok {
		EntitiesString, _ := json.Marshal(Entities)
		json.Unmarshal(EntitiesString, &o.Entities)
	}
	
	if Assessment, ok := IntentfeedbackMap["assessment"].(string); ok {
		o.Assessment = &Assessment
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Intentfeedback) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

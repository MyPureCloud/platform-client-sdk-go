package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Detectedintent
type Detectedintent struct { 
	// Name - The name of the detected intent.
	Name *string `json:"name,omitempty"`


	// Probability - The probability of the detected intent.
	Probability *float64 `json:"probability,omitempty"`


	// Entities - The collection of named entities detected.
	Entities *[]Detectednamedentity `json:"entities,omitempty"`

}

func (o *Detectedintent) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Detectedintent
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		Probability *float64 `json:"probability,omitempty"`
		
		Entities *[]Detectednamedentity `json:"entities,omitempty"`
		*Alias
	}{ 
		Name: o.Name,
		
		Probability: o.Probability,
		
		Entities: o.Entities,
		Alias:    (*Alias)(o),
	})
}

func (o *Detectedintent) UnmarshalJSON(b []byte) error {
	var DetectedintentMap map[string]interface{}
	err := json.Unmarshal(b, &DetectedintentMap)
	if err != nil {
		return err
	}
	
	if Name, ok := DetectedintentMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Probability, ok := DetectedintentMap["probability"].(float64); ok {
		o.Probability = &Probability
	}
    
	if Entities, ok := DetectedintentMap["entities"].([]interface{}); ok {
		EntitiesString, _ := json.Marshal(Entities)
		json.Unmarshal(EntitiesString, &o.Entities)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Detectedintent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

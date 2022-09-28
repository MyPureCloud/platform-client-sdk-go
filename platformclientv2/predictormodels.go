package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Predictormodels
type Predictormodels struct { 
	// Entities
	Entities *[]Predictormodel `json:"entities,omitempty"`


	// PredictorModels
	PredictorModels *[]Predictormodel `json:"predictorModels,omitempty"`

}

func (o *Predictormodels) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Predictormodels
	
	return json.Marshal(&struct { 
		Entities *[]Predictormodel `json:"entities,omitempty"`
		
		PredictorModels *[]Predictormodel `json:"predictorModels,omitempty"`
		*Alias
	}{ 
		Entities: o.Entities,
		
		PredictorModels: o.PredictorModels,
		Alias:    (*Alias)(o),
	})
}

func (o *Predictormodels) UnmarshalJSON(b []byte) error {
	var PredictormodelsMap map[string]interface{}
	err := json.Unmarshal(b, &PredictormodelsMap)
	if err != nil {
		return err
	}
	
	if Entities, ok := PredictormodelsMap["entities"].([]interface{}); ok {
		EntitiesString, _ := json.Marshal(Entities)
		json.Unmarshal(EntitiesString, &o.Entities)
	}
	
	if PredictorModels, ok := PredictormodelsMap["predictorModels"].([]interface{}); ok {
		PredictorModelsString, _ := json.Marshal(PredictorModels)
		json.Unmarshal(PredictorModelsString, &o.PredictorModels)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Predictormodels) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

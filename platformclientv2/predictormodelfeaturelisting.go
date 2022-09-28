package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Predictormodelfeaturelisting
type Predictormodelfeaturelisting struct { 
	// Entities
	Entities *[]Predictormodelfeature `json:"entities,omitempty"`

}

func (o *Predictormodelfeaturelisting) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Predictormodelfeaturelisting
	
	return json.Marshal(&struct { 
		Entities *[]Predictormodelfeature `json:"entities,omitempty"`
		*Alias
	}{ 
		Entities: o.Entities,
		Alias:    (*Alias)(o),
	})
}

func (o *Predictormodelfeaturelisting) UnmarshalJSON(b []byte) error {
	var PredictormodelfeaturelistingMap map[string]interface{}
	err := json.Unmarshal(b, &PredictormodelfeaturelistingMap)
	if err != nil {
		return err
	}
	
	if Entities, ok := PredictormodelfeaturelistingMap["entities"].([]interface{}); ok {
		EntitiesString, _ := json.Marshal(Entities)
		json.Unmarshal(EntitiesString, &o.Entities)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Predictormodelfeaturelisting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Evaluationqualityv2topiccalibration
type Evaluationqualityv2topiccalibration struct { 
	// Id
	Id *string `json:"id,omitempty"`

}

func (o *Evaluationqualityv2topiccalibration) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Evaluationqualityv2topiccalibration
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		Alias:    (*Alias)(o),
	})
}

func (o *Evaluationqualityv2topiccalibration) UnmarshalJSON(b []byte) error {
	var Evaluationqualityv2topiccalibrationMap map[string]interface{}
	err := json.Unmarshal(b, &Evaluationqualityv2topiccalibrationMap)
	if err != nil {
		return err
	}
	
	if Id, ok := Evaluationqualityv2topiccalibrationMap["id"].(string); ok {
		o.Id = &Id
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Evaluationqualityv2topiccalibration) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Forecastplanninggroupsresponse
type Forecastplanninggroupsresponse struct { 
	// Entities
	Entities *[]Forecastplanninggroupresponse `json:"entities,omitempty"`

}

func (o *Forecastplanninggroupsresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Forecastplanninggroupsresponse
	
	return json.Marshal(&struct { 
		Entities *[]Forecastplanninggroupresponse `json:"entities,omitempty"`
		*Alias
	}{ 
		Entities: o.Entities,
		Alias:    (*Alias)(o),
	})
}

func (o *Forecastplanninggroupsresponse) UnmarshalJSON(b []byte) error {
	var ForecastplanninggroupsresponseMap map[string]interface{}
	err := json.Unmarshal(b, &ForecastplanninggroupsresponseMap)
	if err != nil {
		return err
	}
	
	if Entities, ok := ForecastplanninggroupsresponseMap["entities"].([]interface{}); ok {
		EntitiesString, _ := json.Marshal(Entities)
		json.Unmarshal(EntitiesString, &o.Entities)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Forecastplanninggroupsresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

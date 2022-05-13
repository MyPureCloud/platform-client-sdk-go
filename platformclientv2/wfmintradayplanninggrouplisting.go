package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmintradayplanninggrouplisting - A list of IntradayPlanningGroup objects
type Wfmintradayplanninggrouplisting struct { 
	// Entities
	Entities *[]Forecastplanninggroupresponse `json:"entities,omitempty"`


	// NoDataReason - The reason there was no data for the request
	NoDataReason *string `json:"noDataReason,omitempty"`

}

func (o *Wfmintradayplanninggrouplisting) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmintradayplanninggrouplisting
	
	return json.Marshal(&struct { 
		Entities *[]Forecastplanninggroupresponse `json:"entities,omitempty"`
		
		NoDataReason *string `json:"noDataReason,omitempty"`
		*Alias
	}{ 
		Entities: o.Entities,
		
		NoDataReason: o.NoDataReason,
		Alias:    (*Alias)(o),
	})
}

func (o *Wfmintradayplanninggrouplisting) UnmarshalJSON(b []byte) error {
	var WfmintradayplanninggrouplistingMap map[string]interface{}
	err := json.Unmarshal(b, &WfmintradayplanninggrouplistingMap)
	if err != nil {
		return err
	}
	
	if Entities, ok := WfmintradayplanninggrouplistingMap["entities"].([]interface{}); ok {
		EntitiesString, _ := json.Marshal(Entities)
		json.Unmarshal(EntitiesString, &o.Entities)
	}
	
	if NoDataReason, ok := WfmintradayplanninggrouplistingMap["noDataReason"].(string); ok {
		o.NoDataReason = &NoDataReason
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmintradayplanninggrouplisting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

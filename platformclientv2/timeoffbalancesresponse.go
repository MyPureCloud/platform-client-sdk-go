package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Timeoffbalancesresponse
type Timeoffbalancesresponse struct { 
	// Job - The asynchronous job handling the query
	Job *Timeoffbalancejobreference `json:"job,omitempty"`


	// Entities - The list of time off balances. May come via notification
	Entities *[]Timeoffbalanceresponse `json:"entities,omitempty"`

}

func (o *Timeoffbalancesresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Timeoffbalancesresponse
	
	return json.Marshal(&struct { 
		Job *Timeoffbalancejobreference `json:"job,omitempty"`
		
		Entities *[]Timeoffbalanceresponse `json:"entities,omitempty"`
		*Alias
	}{ 
		Job: o.Job,
		
		Entities: o.Entities,
		Alias:    (*Alias)(o),
	})
}

func (o *Timeoffbalancesresponse) UnmarshalJSON(b []byte) error {
	var TimeoffbalancesresponseMap map[string]interface{}
	err := json.Unmarshal(b, &TimeoffbalancesresponseMap)
	if err != nil {
		return err
	}
	
	if Job, ok := TimeoffbalancesresponseMap["job"].(map[string]interface{}); ok {
		JobString, _ := json.Marshal(Job)
		json.Unmarshal(JobString, &o.Job)
	}
	
	if Entities, ok := TimeoffbalancesresponseMap["entities"].([]interface{}); ok {
		EntitiesString, _ := json.Marshal(Entities)
		json.Unmarshal(EntitiesString, &o.Entities)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Timeoffbalancesresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

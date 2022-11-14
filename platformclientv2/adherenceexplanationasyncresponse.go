package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Adherenceexplanationasyncresponse
type Adherenceexplanationasyncresponse struct { 
	// Job - A reference to the job that was started by the request
	Job *Adherenceexplanationjobreference `json:"job,omitempty"`

}

func (o *Adherenceexplanationasyncresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Adherenceexplanationasyncresponse
	
	return json.Marshal(&struct { 
		Job *Adherenceexplanationjobreference `json:"job,omitempty"`
		*Alias
	}{ 
		Job: o.Job,
		Alias:    (*Alias)(o),
	})
}

func (o *Adherenceexplanationasyncresponse) UnmarshalJSON(b []byte) error {
	var AdherenceexplanationasyncresponseMap map[string]interface{}
	err := json.Unmarshal(b, &AdherenceexplanationasyncresponseMap)
	if err != nil {
		return err
	}
	
	if Job, ok := AdherenceexplanationasyncresponseMap["job"].(map[string]interface{}); ok {
		JobString, _ := json.Marshal(Job)
		json.Unmarshal(JobString, &o.Job)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Adherenceexplanationasyncresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

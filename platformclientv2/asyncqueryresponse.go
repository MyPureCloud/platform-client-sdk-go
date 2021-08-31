package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Asyncqueryresponse
type Asyncqueryresponse struct { 
	// JobId - Unique identifier for the async query execution. Can be used to check the status of the query and retrieve results.
	JobId *string `json:"jobId,omitempty"`

}

func (o *Asyncqueryresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Asyncqueryresponse
	
	return json.Marshal(&struct { 
		JobId *string `json:"jobId,omitempty"`
		*Alias
	}{ 
		JobId: o.JobId,
		Alias:    (*Alias)(o),
	})
}

func (o *Asyncqueryresponse) UnmarshalJSON(b []byte) error {
	var AsyncqueryresponseMap map[string]interface{}
	err := json.Unmarshal(b, &AsyncqueryresponseMap)
	if err != nil {
		return err
	}
	
	if JobId, ok := AsyncqueryresponseMap["jobId"].(string); ok {
		o.JobId = &JobId
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Asyncqueryresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

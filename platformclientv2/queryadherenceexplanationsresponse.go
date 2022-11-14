package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queryadherenceexplanationsresponse
type Queryadherenceexplanationsresponse struct { 
	// Job - The asynchronous job handling the query
	Job *Adherenceexplanationjobreference `json:"job,omitempty"`


	// Result - The result of the query. May come via notification
	Result *Adherenceexplanationlisting `json:"result,omitempty"`


	// DownloadUrl - The URL from which to download the result. May come via notification
	DownloadUrl *string `json:"downloadUrl,omitempty"`

}

func (o *Queryadherenceexplanationsresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queryadherenceexplanationsresponse
	
	return json.Marshal(&struct { 
		Job *Adherenceexplanationjobreference `json:"job,omitempty"`
		
		Result *Adherenceexplanationlisting `json:"result,omitempty"`
		
		DownloadUrl *string `json:"downloadUrl,omitempty"`
		*Alias
	}{ 
		Job: o.Job,
		
		Result: o.Result,
		
		DownloadUrl: o.DownloadUrl,
		Alias:    (*Alias)(o),
	})
}

func (o *Queryadherenceexplanationsresponse) UnmarshalJSON(b []byte) error {
	var QueryadherenceexplanationsresponseMap map[string]interface{}
	err := json.Unmarshal(b, &QueryadherenceexplanationsresponseMap)
	if err != nil {
		return err
	}
	
	if Job, ok := QueryadherenceexplanationsresponseMap["job"].(map[string]interface{}); ok {
		JobString, _ := json.Marshal(Job)
		json.Unmarshal(JobString, &o.Job)
	}
	
	if Result, ok := QueryadherenceexplanationsresponseMap["result"].(map[string]interface{}); ok {
		ResultString, _ := json.Marshal(Result)
		json.Unmarshal(ResultString, &o.Result)
	}
	
	if DownloadUrl, ok := QueryadherenceexplanationsresponseMap["downloadUrl"].(string); ok {
		o.DownloadUrl = &DownloadUrl
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Queryadherenceexplanationsresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

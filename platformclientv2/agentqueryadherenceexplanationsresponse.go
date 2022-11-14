package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Agentqueryadherenceexplanationsresponse
type Agentqueryadherenceexplanationsresponse struct { 
	// Job - The asynchronous job handling the query
	Job *Adherenceexplanationjobreference `json:"job,omitempty"`


	// Result - The result of the query. May come via notification
	Result *Adherenceexplanationlistingagentqueryresponse `json:"result,omitempty"`


	// DownloadUrl - The URL from which to download the result. May come via notification
	DownloadUrl *string `json:"downloadUrl,omitempty"`

}

func (o *Agentqueryadherenceexplanationsresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Agentqueryadherenceexplanationsresponse
	
	return json.Marshal(&struct { 
		Job *Adherenceexplanationjobreference `json:"job,omitempty"`
		
		Result *Adherenceexplanationlistingagentqueryresponse `json:"result,omitempty"`
		
		DownloadUrl *string `json:"downloadUrl,omitempty"`
		*Alias
	}{ 
		Job: o.Job,
		
		Result: o.Result,
		
		DownloadUrl: o.DownloadUrl,
		Alias:    (*Alias)(o),
	})
}

func (o *Agentqueryadherenceexplanationsresponse) UnmarshalJSON(b []byte) error {
	var AgentqueryadherenceexplanationsresponseMap map[string]interface{}
	err := json.Unmarshal(b, &AgentqueryadherenceexplanationsresponseMap)
	if err != nil {
		return err
	}
	
	if Job, ok := AgentqueryadherenceexplanationsresponseMap["job"].(map[string]interface{}); ok {
		JobString, _ := json.Marshal(Job)
		json.Unmarshal(JobString, &o.Job)
	}
	
	if Result, ok := AgentqueryadherenceexplanationsresponseMap["result"].(map[string]interface{}); ok {
		ResultString, _ := json.Marshal(Result)
		json.Unmarshal(ResultString, &o.Result)
	}
	
	if DownloadUrl, ok := AgentqueryadherenceexplanationsresponseMap["downloadUrl"].(string); ok {
		o.DownloadUrl = &DownloadUrl
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Agentqueryadherenceexplanationsresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmhistoricaladherencebulkresponse
type Wfmhistoricaladherencebulkresponse struct { 
	// Job - A reference to the job that was started by the request
	Job *Wfmhistoricaladherencebulkjobreference `json:"job,omitempty"`


	// DownloadUrls - The uri list to GET the results of the Historical Adherence query. This field is populated only if query state is Complete
	DownloadUrls *[]string `json:"downloadUrls,omitempty"`


	// DownloadResult - Results will always come via downloadUrls; however the schema is included for documentation
	DownloadResult *Wfmhistoricaladherencebulkresult `json:"downloadResult,omitempty"`

}

func (o *Wfmhistoricaladherencebulkresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmhistoricaladherencebulkresponse
	
	return json.Marshal(&struct { 
		Job *Wfmhistoricaladherencebulkjobreference `json:"job,omitempty"`
		
		DownloadUrls *[]string `json:"downloadUrls,omitempty"`
		
		DownloadResult *Wfmhistoricaladherencebulkresult `json:"downloadResult,omitempty"`
		*Alias
	}{ 
		Job: o.Job,
		
		DownloadUrls: o.DownloadUrls,
		
		DownloadResult: o.DownloadResult,
		Alias:    (*Alias)(o),
	})
}

func (o *Wfmhistoricaladherencebulkresponse) UnmarshalJSON(b []byte) error {
	var WfmhistoricaladherencebulkresponseMap map[string]interface{}
	err := json.Unmarshal(b, &WfmhistoricaladherencebulkresponseMap)
	if err != nil {
		return err
	}
	
	if Job, ok := WfmhistoricaladherencebulkresponseMap["job"].(map[string]interface{}); ok {
		JobString, _ := json.Marshal(Job)
		json.Unmarshal(JobString, &o.Job)
	}
	
	if DownloadUrls, ok := WfmhistoricaladherencebulkresponseMap["downloadUrls"].([]interface{}); ok {
		DownloadUrlsString, _ := json.Marshal(DownloadUrls)
		json.Unmarshal(DownloadUrlsString, &o.DownloadUrls)
	}
	
	if DownloadResult, ok := WfmhistoricaladherencebulkresponseMap["downloadResult"].(map[string]interface{}); ok {
		DownloadResultString, _ := json.Marshal(DownloadResult)
		json.Unmarshal(DownloadResultString, &o.DownloadResult)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmhistoricaladherencebulkresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

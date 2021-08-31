package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Batchdownloadjobstatusresult
type Batchdownloadjobstatusresult struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// JobId - JobId returned when job was initially submitted
	JobId *string `json:"jobId,omitempty"`


	// ExpectedResultCount - Number of results expected when job is completed
	ExpectedResultCount *int `json:"expectedResultCount,omitempty"`


	// ResultCount - Current number of results available
	ResultCount *int `json:"resultCount,omitempty"`


	// ErrorCount - Number of error results produced so far
	ErrorCount *int `json:"errorCount,omitempty"`


	// Results - Current set of results for the job
	Results *[]Batchdownloadjobresult `json:"results,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Batchdownloadjobstatusresult) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Batchdownloadjobstatusresult
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		JobId *string `json:"jobId,omitempty"`
		
		ExpectedResultCount *int `json:"expectedResultCount,omitempty"`
		
		ResultCount *int `json:"resultCount,omitempty"`
		
		ErrorCount *int `json:"errorCount,omitempty"`
		
		Results *[]Batchdownloadjobresult `json:"results,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		JobId: o.JobId,
		
		ExpectedResultCount: o.ExpectedResultCount,
		
		ResultCount: o.ResultCount,
		
		ErrorCount: o.ErrorCount,
		
		Results: o.Results,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Batchdownloadjobstatusresult) UnmarshalJSON(b []byte) error {
	var BatchdownloadjobstatusresultMap map[string]interface{}
	err := json.Unmarshal(b, &BatchdownloadjobstatusresultMap)
	if err != nil {
		return err
	}
	
	if Id, ok := BatchdownloadjobstatusresultMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if JobId, ok := BatchdownloadjobstatusresultMap["jobId"].(string); ok {
		o.JobId = &JobId
	}
	
	if ExpectedResultCount, ok := BatchdownloadjobstatusresultMap["expectedResultCount"].(float64); ok {
		ExpectedResultCountInt := int(ExpectedResultCount)
		o.ExpectedResultCount = &ExpectedResultCountInt
	}
	
	if ResultCount, ok := BatchdownloadjobstatusresultMap["resultCount"].(float64); ok {
		ResultCountInt := int(ResultCount)
		o.ResultCount = &ResultCountInt
	}
	
	if ErrorCount, ok := BatchdownloadjobstatusresultMap["errorCount"].(float64); ok {
		ErrorCountInt := int(ErrorCount)
		o.ErrorCount = &ErrorCountInt
	}
	
	if Results, ok := BatchdownloadjobstatusresultMap["results"].([]interface{}); ok {
		ResultsString, _ := json.Marshal(Results)
		json.Unmarshal(ResultsString, &o.Results)
	}
	
	if SelfUri, ok := BatchdownloadjobstatusresultMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Batchdownloadjobstatusresult) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmhistoricalshrinkageresponse
type Wfmhistoricalshrinkageresponse struct { 
	// OperationId - The operationId for which to listen
	OperationId *string `json:"operationId,omitempty"`


	// DownloadUrls - The url list to GET the results of the Historical Shrinkage query. This field is populated only if query state is Complete
	DownloadUrls *[]string `json:"downloadUrls,omitempty"`


	// DownloadResult - Result will always come via downloadUrls; however the schema is included for documentation
	DownloadResult *Historicalshrinkageresultlisting `json:"downloadResult,omitempty"`


	// State - The state of the shrinkage query
	State *string `json:"state,omitempty"`

}

func (o *Wfmhistoricalshrinkageresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmhistoricalshrinkageresponse
	
	return json.Marshal(&struct { 
		OperationId *string `json:"operationId,omitempty"`
		
		DownloadUrls *[]string `json:"downloadUrls,omitempty"`
		
		DownloadResult *Historicalshrinkageresultlisting `json:"downloadResult,omitempty"`
		
		State *string `json:"state,omitempty"`
		*Alias
	}{ 
		OperationId: o.OperationId,
		
		DownloadUrls: o.DownloadUrls,
		
		DownloadResult: o.DownloadResult,
		
		State: o.State,
		Alias:    (*Alias)(o),
	})
}

func (o *Wfmhistoricalshrinkageresponse) UnmarshalJSON(b []byte) error {
	var WfmhistoricalshrinkageresponseMap map[string]interface{}
	err := json.Unmarshal(b, &WfmhistoricalshrinkageresponseMap)
	if err != nil {
		return err
	}
	
	if OperationId, ok := WfmhistoricalshrinkageresponseMap["operationId"].(string); ok {
		o.OperationId = &OperationId
	}
    
	if DownloadUrls, ok := WfmhistoricalshrinkageresponseMap["downloadUrls"].([]interface{}); ok {
		DownloadUrlsString, _ := json.Marshal(DownloadUrls)
		json.Unmarshal(DownloadUrlsString, &o.DownloadUrls)
	}
	
	if DownloadResult, ok := WfmhistoricalshrinkageresponseMap["downloadResult"].(map[string]interface{}); ok {
		DownloadResultString, _ := json.Marshal(DownloadResult)
		json.Unmarshal(DownloadResultString, &o.DownloadResult)
	}
	
	if State, ok := WfmhistoricalshrinkageresponseMap["state"].(string); ok {
		o.State = &State
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmhistoricalshrinkageresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

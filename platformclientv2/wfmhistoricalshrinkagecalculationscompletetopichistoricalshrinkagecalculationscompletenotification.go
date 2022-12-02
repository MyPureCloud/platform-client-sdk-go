package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmhistoricalshrinkagecalculationscompletetopichistoricalshrinkagecalculationscompletenotification
type Wfmhistoricalshrinkagecalculationscompletetopichistoricalshrinkagecalculationscompletenotification struct { 
	// OperationId
	OperationId *string `json:"operationId,omitempty"`


	// DownloadUrls
	DownloadUrls *[]string `json:"downloadUrls,omitempty"`


	// State
	State *string `json:"state,omitempty"`

}

func (o *Wfmhistoricalshrinkagecalculationscompletetopichistoricalshrinkagecalculationscompletenotification) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmhistoricalshrinkagecalculationscompletetopichistoricalshrinkagecalculationscompletenotification
	
	return json.Marshal(&struct { 
		OperationId *string `json:"operationId,omitempty"`
		
		DownloadUrls *[]string `json:"downloadUrls,omitempty"`
		
		State *string `json:"state,omitempty"`
		*Alias
	}{ 
		OperationId: o.OperationId,
		
		DownloadUrls: o.DownloadUrls,
		
		State: o.State,
		Alias:    (*Alias)(o),
	})
}

func (o *Wfmhistoricalshrinkagecalculationscompletetopichistoricalshrinkagecalculationscompletenotification) UnmarshalJSON(b []byte) error {
	var WfmhistoricalshrinkagecalculationscompletetopichistoricalshrinkagecalculationscompletenotificationMap map[string]interface{}
	err := json.Unmarshal(b, &WfmhistoricalshrinkagecalculationscompletetopichistoricalshrinkagecalculationscompletenotificationMap)
	if err != nil {
		return err
	}
	
	if OperationId, ok := WfmhistoricalshrinkagecalculationscompletetopichistoricalshrinkagecalculationscompletenotificationMap["operationId"].(string); ok {
		o.OperationId = &OperationId
	}
    
	if DownloadUrls, ok := WfmhistoricalshrinkagecalculationscompletetopichistoricalshrinkagecalculationscompletenotificationMap["downloadUrls"].([]interface{}); ok {
		DownloadUrlsString, _ := json.Marshal(DownloadUrls)
		json.Unmarshal(DownloadUrlsString, &o.DownloadUrls)
	}
	
	if State, ok := WfmhistoricalshrinkagecalculationscompletetopichistoricalshrinkagecalculationscompletenotificationMap["state"].(string); ok {
		o.State = &State
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmhistoricalshrinkagecalculationscompletetopichistoricalshrinkagecalculationscompletenotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

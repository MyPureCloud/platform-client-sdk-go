package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmhistoricaladherenceresponse - Response for Historical Adherence Query, intended to tell the client what to listen for on a notification topic
type Wfmhistoricaladherenceresponse struct { 
	// Id - The query ID to listen for
	Id *string `json:"id,omitempty"`


	// DownloadUrl - Deprecated. Use downloadUrls instead.
	DownloadUrl *string `json:"downloadUrl,omitempty"`


	// DownloadResult - Result will always come via downloadUrls; however the schema is included for documentation
	DownloadResult *Wfmhistoricaladherenceresultwrapper `json:"downloadResult,omitempty"`


	// DownloadUrls - The uri list to GET the results of the Historical Adherence query. For notification purposes only
	DownloadUrls *[]string `json:"downloadUrls,omitempty"`


	// QueryState - The state of the adherence query
	QueryState *string `json:"queryState,omitempty"`

}

func (o *Wfmhistoricaladherenceresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmhistoricaladherenceresponse
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		DownloadUrl *string `json:"downloadUrl,omitempty"`
		
		DownloadResult *Wfmhistoricaladherenceresultwrapper `json:"downloadResult,omitempty"`
		
		DownloadUrls *[]string `json:"downloadUrls,omitempty"`
		
		QueryState *string `json:"queryState,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		DownloadUrl: o.DownloadUrl,
		
		DownloadResult: o.DownloadResult,
		
		DownloadUrls: o.DownloadUrls,
		
		QueryState: o.QueryState,
		Alias:    (*Alias)(o),
	})
}

func (o *Wfmhistoricaladherenceresponse) UnmarshalJSON(b []byte) error {
	var WfmhistoricaladherenceresponseMap map[string]interface{}
	err := json.Unmarshal(b, &WfmhistoricaladherenceresponseMap)
	if err != nil {
		return err
	}
	
	if Id, ok := WfmhistoricaladherenceresponseMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if DownloadUrl, ok := WfmhistoricaladherenceresponseMap["downloadUrl"].(string); ok {
		o.DownloadUrl = &DownloadUrl
	}
	
	if DownloadResult, ok := WfmhistoricaladherenceresponseMap["downloadResult"].(map[string]interface{}); ok {
		DownloadResultString, _ := json.Marshal(DownloadResult)
		json.Unmarshal(DownloadResultString, &o.DownloadResult)
	}
	
	if DownloadUrls, ok := WfmhistoricaladherenceresponseMap["downloadUrls"].([]interface{}); ok {
		DownloadUrlsString, _ := json.Marshal(DownloadUrls)
		json.Unmarshal(DownloadUrlsString, &o.DownloadUrls)
	}
	
	if QueryState, ok := WfmhistoricaladherenceresponseMap["queryState"].(string); ok {
		o.QueryState = &QueryState
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmhistoricaladherenceresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

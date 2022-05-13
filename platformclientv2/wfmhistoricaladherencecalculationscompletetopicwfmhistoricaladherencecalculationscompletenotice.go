package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmhistoricaladherencecalculationscompletetopicwfmhistoricaladherencecalculationscompletenotice
type Wfmhistoricaladherencecalculationscompletetopicwfmhistoricaladherencecalculationscompletenotice struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// DownloadUrl
	DownloadUrl *string `json:"downloadUrl,omitempty"`


	// DownloadUrls
	DownloadUrls *[]string `json:"downloadUrls,omitempty"`


	// QueryState
	QueryState *string `json:"queryState,omitempty"`

}

func (o *Wfmhistoricaladherencecalculationscompletetopicwfmhistoricaladherencecalculationscompletenotice) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmhistoricaladherencecalculationscompletetopicwfmhistoricaladherencecalculationscompletenotice
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		DownloadUrl *string `json:"downloadUrl,omitempty"`
		
		DownloadUrls *[]string `json:"downloadUrls,omitempty"`
		
		QueryState *string `json:"queryState,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		DownloadUrl: o.DownloadUrl,
		
		DownloadUrls: o.DownloadUrls,
		
		QueryState: o.QueryState,
		Alias:    (*Alias)(o),
	})
}

func (o *Wfmhistoricaladherencecalculationscompletetopicwfmhistoricaladherencecalculationscompletenotice) UnmarshalJSON(b []byte) error {
	var WfmhistoricaladherencecalculationscompletetopicwfmhistoricaladherencecalculationscompletenoticeMap map[string]interface{}
	err := json.Unmarshal(b, &WfmhistoricaladherencecalculationscompletetopicwfmhistoricaladherencecalculationscompletenoticeMap)
	if err != nil {
		return err
	}
	
	if Id, ok := WfmhistoricaladherencecalculationscompletetopicwfmhistoricaladherencecalculationscompletenoticeMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if DownloadUrl, ok := WfmhistoricaladherencecalculationscompletetopicwfmhistoricaladherencecalculationscompletenoticeMap["downloadUrl"].(string); ok {
		o.DownloadUrl = &DownloadUrl
	}
    
	if DownloadUrls, ok := WfmhistoricaladherencecalculationscompletetopicwfmhistoricaladherencecalculationscompletenoticeMap["downloadUrls"].([]interface{}); ok {
		DownloadUrlsString, _ := json.Marshal(DownloadUrls)
		json.Unmarshal(DownloadUrlsString, &o.DownloadUrls)
	}
	
	if QueryState, ok := WfmhistoricaladherencecalculationscompletetopicwfmhistoricaladherencecalculationscompletenoticeMap["queryState"].(string); ok {
		o.QueryState = &QueryState
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmhistoricaladherencecalculationscompletetopicwfmhistoricaladherencecalculationscompletenotice) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

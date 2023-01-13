package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmhistoricaladherencebulkcalculationscompletetopicwfmhistoricaladherencebulkcalculationscompletenotification
type Wfmhistoricaladherencebulkcalculationscompletetopicwfmhistoricaladherencebulkcalculationscompletenotification struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// DownloadUrls
	DownloadUrls *[]string `json:"downloadUrls,omitempty"`


	// QueryState
	QueryState *string `json:"queryState,omitempty"`

}

func (o *Wfmhistoricaladherencebulkcalculationscompletetopicwfmhistoricaladherencebulkcalculationscompletenotification) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmhistoricaladherencebulkcalculationscompletetopicwfmhistoricaladherencebulkcalculationscompletenotification
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		DownloadUrls *[]string `json:"downloadUrls,omitempty"`
		
		QueryState *string `json:"queryState,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		DownloadUrls: o.DownloadUrls,
		
		QueryState: o.QueryState,
		Alias:    (*Alias)(o),
	})
}

func (o *Wfmhistoricaladherencebulkcalculationscompletetopicwfmhistoricaladherencebulkcalculationscompletenotification) UnmarshalJSON(b []byte) error {
	var WfmhistoricaladherencebulkcalculationscompletetopicwfmhistoricaladherencebulkcalculationscompletenotificationMap map[string]interface{}
	err := json.Unmarshal(b, &WfmhistoricaladherencebulkcalculationscompletetopicwfmhistoricaladherencebulkcalculationscompletenotificationMap)
	if err != nil {
		return err
	}
	
	if Id, ok := WfmhistoricaladherencebulkcalculationscompletetopicwfmhistoricaladherencebulkcalculationscompletenotificationMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if DownloadUrls, ok := WfmhistoricaladherencebulkcalculationscompletetopicwfmhistoricaladherencebulkcalculationscompletenotificationMap["downloadUrls"].([]interface{}); ok {
		DownloadUrlsString, _ := json.Marshal(DownloadUrls)
		json.Unmarshal(DownloadUrlsString, &o.DownloadUrls)
	}
	
	if QueryState, ok := WfmhistoricaladherencebulkcalculationscompletetopicwfmhistoricaladherencebulkcalculationscompletenotificationMap["queryState"].(string); ok {
		o.QueryState = &QueryState
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmhistoricaladherencebulkcalculationscompletetopicwfmhistoricaladherencebulkcalculationscompletenotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

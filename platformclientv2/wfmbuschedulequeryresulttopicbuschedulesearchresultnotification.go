package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmbuschedulequeryresulttopicbuschedulesearchresultnotification
type Wfmbuschedulequeryresulttopicbuschedulesearchresultnotification struct { 
	// OperationId
	OperationId *string `json:"operationId,omitempty"`


	// BusinessUnitId
	BusinessUnitId *string `json:"businessUnitId,omitempty"`


	// DownloadUrl
	DownloadUrl *string `json:"downloadUrl,omitempty"`

}

func (o *Wfmbuschedulequeryresulttopicbuschedulesearchresultnotification) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmbuschedulequeryresulttopicbuschedulesearchresultnotification
	
	return json.Marshal(&struct { 
		OperationId *string `json:"operationId,omitempty"`
		
		BusinessUnitId *string `json:"businessUnitId,omitempty"`
		
		DownloadUrl *string `json:"downloadUrl,omitempty"`
		*Alias
	}{ 
		OperationId: o.OperationId,
		
		BusinessUnitId: o.BusinessUnitId,
		
		DownloadUrl: o.DownloadUrl,
		Alias:    (*Alias)(o),
	})
}

func (o *Wfmbuschedulequeryresulttopicbuschedulesearchresultnotification) UnmarshalJSON(b []byte) error {
	var WfmbuschedulequeryresulttopicbuschedulesearchresultnotificationMap map[string]interface{}
	err := json.Unmarshal(b, &WfmbuschedulequeryresulttopicbuschedulesearchresultnotificationMap)
	if err != nil {
		return err
	}
	
	if OperationId, ok := WfmbuschedulequeryresulttopicbuschedulesearchresultnotificationMap["operationId"].(string); ok {
		o.OperationId = &OperationId
	}
    
	if BusinessUnitId, ok := WfmbuschedulequeryresulttopicbuschedulesearchresultnotificationMap["businessUnitId"].(string); ok {
		o.BusinessUnitId = &BusinessUnitId
	}
    
	if DownloadUrl, ok := WfmbuschedulequeryresulttopicbuschedulesearchresultnotificationMap["downloadUrl"].(string); ok {
		o.DownloadUrl = &DownloadUrl
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmbuschedulequeryresulttopicbuschedulesearchresultnotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmbuschedulesearchresulttopicbuschedulesearchresultnotification
type Wfmbuschedulesearchresulttopicbuschedulesearchresultnotification struct { 
	// OperationId
	OperationId *string `json:"operationId,omitempty"`


	// BusinessUnitId
	BusinessUnitId *string `json:"businessUnitId,omitempty"`


	// DownloadUrl
	DownloadUrl *string `json:"downloadUrl,omitempty"`

}

func (o *Wfmbuschedulesearchresulttopicbuschedulesearchresultnotification) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmbuschedulesearchresulttopicbuschedulesearchresultnotification
	
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

func (o *Wfmbuschedulesearchresulttopicbuschedulesearchresultnotification) UnmarshalJSON(b []byte) error {
	var WfmbuschedulesearchresulttopicbuschedulesearchresultnotificationMap map[string]interface{}
	err := json.Unmarshal(b, &WfmbuschedulesearchresulttopicbuschedulesearchresultnotificationMap)
	if err != nil {
		return err
	}
	
	if OperationId, ok := WfmbuschedulesearchresulttopicbuschedulesearchresultnotificationMap["operationId"].(string); ok {
		o.OperationId = &OperationId
	}
    
	if BusinessUnitId, ok := WfmbuschedulesearchresulttopicbuschedulesearchresultnotificationMap["businessUnitId"].(string); ok {
		o.BusinessUnitId = &BusinessUnitId
	}
    
	if DownloadUrl, ok := WfmbuschedulesearchresulttopicbuschedulesearchresultnotificationMap["downloadUrl"].(string); ok {
		o.DownloadUrl = &DownloadUrl
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmbuschedulesearchresulttopicbuschedulesearchresultnotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

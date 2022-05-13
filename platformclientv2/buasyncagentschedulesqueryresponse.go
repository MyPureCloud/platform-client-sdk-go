package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Buasyncagentschedulesqueryresponse
type Buasyncagentschedulesqueryresponse struct { 
	// Status - The status of the operation
	Status *string `json:"status,omitempty"`


	// OperationId - The ID for the operation
	OperationId *string `json:"operationId,omitempty"`


	// Result - The result of the operation.  Null unless status == Complete
	Result *Buagentschedulesqueryresponse `json:"result,omitempty"`


	// Progress - Percent progress for the operation
	Progress *int `json:"progress,omitempty"`


	// DownloadUrl - The URL from which to download the result if it is too large to pass directly
	DownloadUrl *string `json:"downloadUrl,omitempty"`

}

func (o *Buasyncagentschedulesqueryresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Buasyncagentschedulesqueryresponse
	
	return json.Marshal(&struct { 
		Status *string `json:"status,omitempty"`
		
		OperationId *string `json:"operationId,omitempty"`
		
		Result *Buagentschedulesqueryresponse `json:"result,omitempty"`
		
		Progress *int `json:"progress,omitempty"`
		
		DownloadUrl *string `json:"downloadUrl,omitempty"`
		*Alias
	}{ 
		Status: o.Status,
		
		OperationId: o.OperationId,
		
		Result: o.Result,
		
		Progress: o.Progress,
		
		DownloadUrl: o.DownloadUrl,
		Alias:    (*Alias)(o),
	})
}

func (o *Buasyncagentschedulesqueryresponse) UnmarshalJSON(b []byte) error {
	var BuasyncagentschedulesqueryresponseMap map[string]interface{}
	err := json.Unmarshal(b, &BuasyncagentschedulesqueryresponseMap)
	if err != nil {
		return err
	}
	
	if Status, ok := BuasyncagentschedulesqueryresponseMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if OperationId, ok := BuasyncagentschedulesqueryresponseMap["operationId"].(string); ok {
		o.OperationId = &OperationId
	}
    
	if Result, ok := BuasyncagentschedulesqueryresponseMap["result"].(map[string]interface{}); ok {
		ResultString, _ := json.Marshal(Result)
		json.Unmarshal(ResultString, &o.Result)
	}
	
	if Progress, ok := BuasyncagentschedulesqueryresponseMap["progress"].(float64); ok {
		ProgressInt := int(Progress)
		o.Progress = &ProgressInt
	}
	
	if DownloadUrl, ok := BuasyncagentschedulesqueryresponseMap["downloadUrl"].(string); ok {
		o.DownloadUrl = &DownloadUrl
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Buasyncagentschedulesqueryresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

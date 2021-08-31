package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Buasyncagentschedulessearchresponse
type Buasyncagentschedulessearchresponse struct { 
	// Status - The status of the operation
	Status *string `json:"status,omitempty"`


	// OperationId - The ID for the operation
	OperationId *string `json:"operationId,omitempty"`


	// Result - The result of the operation.  Null unless status == Complete
	Result *Buagentschedulessearchresponse `json:"result,omitempty"`


	// Progress - Percent progress for the operation
	Progress *int `json:"progress,omitempty"`


	// DownloadUrl - The URL from which to download the result if it is too large to pass directly
	DownloadUrl *string `json:"downloadUrl,omitempty"`

}

func (o *Buasyncagentschedulessearchresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Buasyncagentschedulessearchresponse
	
	return json.Marshal(&struct { 
		Status *string `json:"status,omitempty"`
		
		OperationId *string `json:"operationId,omitempty"`
		
		Result *Buagentschedulessearchresponse `json:"result,omitempty"`
		
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

func (o *Buasyncagentschedulessearchresponse) UnmarshalJSON(b []byte) error {
	var BuasyncagentschedulessearchresponseMap map[string]interface{}
	err := json.Unmarshal(b, &BuasyncagentschedulessearchresponseMap)
	if err != nil {
		return err
	}
	
	if Status, ok := BuasyncagentschedulessearchresponseMap["status"].(string); ok {
		o.Status = &Status
	}
	
	if OperationId, ok := BuasyncagentschedulessearchresponseMap["operationId"].(string); ok {
		o.OperationId = &OperationId
	}
	
	if Result, ok := BuasyncagentschedulessearchresponseMap["result"].(map[string]interface{}); ok {
		ResultString, _ := json.Marshal(Result)
		json.Unmarshal(ResultString, &o.Result)
	}
	
	if Progress, ok := BuasyncagentschedulessearchresponseMap["progress"].(float64); ok {
		ProgressInt := int(Progress)
		o.Progress = &ProgressInt
	}
	
	if DownloadUrl, ok := BuasyncagentschedulessearchresponseMap["downloadUrl"].(string); ok {
		o.DownloadUrl = &DownloadUrl
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Buasyncagentschedulessearchresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

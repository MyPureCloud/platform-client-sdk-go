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

func (u *Buasyncagentschedulessearchresponse) MarshalJSON() ([]byte, error) {
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
		Status: u.Status,
		
		OperationId: u.OperationId,
		
		Result: u.Result,
		
		Progress: u.Progress,
		
		DownloadUrl: u.DownloadUrl,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Buasyncagentschedulessearchresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

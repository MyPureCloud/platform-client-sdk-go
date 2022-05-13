package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmbushorttermforecastgenerateprogresstopicgeneratebushorttermforecastprogressnotification
type Wfmbushorttermforecastgenerateprogresstopicgeneratebushorttermforecastprogressnotification struct { 
	// Status
	Status *string `json:"status,omitempty"`


	// Result
	Result *Wfmbushorttermforecastgenerateprogresstopicbushorttermforecast `json:"result,omitempty"`


	// OperationId
	OperationId *string `json:"operationId,omitempty"`


	// Progress
	Progress *int `json:"progress,omitempty"`

}

func (o *Wfmbushorttermforecastgenerateprogresstopicgeneratebushorttermforecastprogressnotification) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmbushorttermforecastgenerateprogresstopicgeneratebushorttermforecastprogressnotification
	
	return json.Marshal(&struct { 
		Status *string `json:"status,omitempty"`
		
		Result *Wfmbushorttermforecastgenerateprogresstopicbushorttermforecast `json:"result,omitempty"`
		
		OperationId *string `json:"operationId,omitempty"`
		
		Progress *int `json:"progress,omitempty"`
		*Alias
	}{ 
		Status: o.Status,
		
		Result: o.Result,
		
		OperationId: o.OperationId,
		
		Progress: o.Progress,
		Alias:    (*Alias)(o),
	})
}

func (o *Wfmbushorttermforecastgenerateprogresstopicgeneratebushorttermforecastprogressnotification) UnmarshalJSON(b []byte) error {
	var WfmbushorttermforecastgenerateprogresstopicgeneratebushorttermforecastprogressnotificationMap map[string]interface{}
	err := json.Unmarshal(b, &WfmbushorttermforecastgenerateprogresstopicgeneratebushorttermforecastprogressnotificationMap)
	if err != nil {
		return err
	}
	
	if Status, ok := WfmbushorttermforecastgenerateprogresstopicgeneratebushorttermforecastprogressnotificationMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if Result, ok := WfmbushorttermforecastgenerateprogresstopicgeneratebushorttermforecastprogressnotificationMap["result"].(map[string]interface{}); ok {
		ResultString, _ := json.Marshal(Result)
		json.Unmarshal(ResultString, &o.Result)
	}
	
	if OperationId, ok := WfmbushorttermforecastgenerateprogresstopicgeneratebushorttermforecastprogressnotificationMap["operationId"].(string); ok {
		o.OperationId = &OperationId
	}
    
	if Progress, ok := WfmbushorttermforecastgenerateprogresstopicgeneratebushorttermforecastprogressnotificationMap["progress"].(float64); ok {
		ProgressInt := int(Progress)
		o.Progress = &ProgressInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmbushorttermforecastgenerateprogresstopicgeneratebushorttermforecastprogressnotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

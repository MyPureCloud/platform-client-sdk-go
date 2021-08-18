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

func (u *Wfmbushorttermforecastgenerateprogresstopicgeneratebushorttermforecastprogressnotification) MarshalJSON() ([]byte, error) {
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
		Status: u.Status,
		
		Result: u.Result,
		
		OperationId: u.OperationId,
		
		Progress: u.Progress,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Wfmbushorttermforecastgenerateprogresstopicgeneratebushorttermforecastprogressnotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

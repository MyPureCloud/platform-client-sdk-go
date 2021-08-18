package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Usageexecutionresult
type Usageexecutionresult struct { 
	// ExecutionId - The id of the query execution
	ExecutionId *string `json:"executionId,omitempty"`


	// ResultsUri - URI where the query results can be retrieved
	ResultsUri *string `json:"resultsUri,omitempty"`

}

func (u *Usageexecutionresult) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Usageexecutionresult

	

	return json.Marshal(&struct { 
		ExecutionId *string `json:"executionId,omitempty"`
		
		ResultsUri *string `json:"resultsUri,omitempty"`
		*Alias
	}{ 
		ExecutionId: u.ExecutionId,
		
		ResultsUri: u.ResultsUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Usageexecutionresult) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Asyncquerystatus
type Asyncquerystatus struct { 
	// State - The current state of the asynchronous query
	State *string `json:"state,omitempty"`


	// ErrorMessage - The error associated with the current query, if the state is FAILED
	ErrorMessage *string `json:"errorMessage,omitempty"`


	// ExpirationDate - The time at which results for this query will expire. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ExpirationDate *time.Time `json:"expirationDate,omitempty"`


	// SubmissionDate - The time at which the query was submitted. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	SubmissionDate *time.Time `json:"submissionDate,omitempty"`


	// CompletionDate - The time at which the query completed. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	CompletionDate *time.Time `json:"completionDate,omitempty"`

}

func (o *Asyncquerystatus) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Asyncquerystatus
	
	ExpirationDate := new(string)
	if o.ExpirationDate != nil {
		
		*ExpirationDate = timeutil.Strftime(o.ExpirationDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ExpirationDate = nil
	}
	
	SubmissionDate := new(string)
	if o.SubmissionDate != nil {
		
		*SubmissionDate = timeutil.Strftime(o.SubmissionDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		SubmissionDate = nil
	}
	
	CompletionDate := new(string)
	if o.CompletionDate != nil {
		
		*CompletionDate = timeutil.Strftime(o.CompletionDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		CompletionDate = nil
	}
	
	return json.Marshal(&struct { 
		State *string `json:"state,omitempty"`
		
		ErrorMessage *string `json:"errorMessage,omitempty"`
		
		ExpirationDate *string `json:"expirationDate,omitempty"`
		
		SubmissionDate *string `json:"submissionDate,omitempty"`
		
		CompletionDate *string `json:"completionDate,omitempty"`
		*Alias
	}{ 
		State: o.State,
		
		ErrorMessage: o.ErrorMessage,
		
		ExpirationDate: ExpirationDate,
		
		SubmissionDate: SubmissionDate,
		
		CompletionDate: CompletionDate,
		Alias:    (*Alias)(o),
	})
}

func (o *Asyncquerystatus) UnmarshalJSON(b []byte) error {
	var AsyncquerystatusMap map[string]interface{}
	err := json.Unmarshal(b, &AsyncquerystatusMap)
	if err != nil {
		return err
	}
	
	if State, ok := AsyncquerystatusMap["state"].(string); ok {
		o.State = &State
	}
	
	if ErrorMessage, ok := AsyncquerystatusMap["errorMessage"].(string); ok {
		o.ErrorMessage = &ErrorMessage
	}
	
	if expirationDateString, ok := AsyncquerystatusMap["expirationDate"].(string); ok {
		ExpirationDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", expirationDateString)
		o.ExpirationDate = &ExpirationDate
	}
	
	if submissionDateString, ok := AsyncquerystatusMap["submissionDate"].(string); ok {
		SubmissionDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", submissionDateString)
		o.SubmissionDate = &SubmissionDate
	}
	
	if completionDateString, ok := AsyncquerystatusMap["completionDate"].(string); ok {
		CompletionDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", completionDateString)
		o.CompletionDate = &CompletionDate
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Asyncquerystatus) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

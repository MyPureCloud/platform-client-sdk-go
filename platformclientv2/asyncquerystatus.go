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

func (u *Asyncquerystatus) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Asyncquerystatus

	
	ExpirationDate := new(string)
	if u.ExpirationDate != nil {
		
		*ExpirationDate = timeutil.Strftime(u.ExpirationDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ExpirationDate = nil
	}
	
	SubmissionDate := new(string)
	if u.SubmissionDate != nil {
		
		*SubmissionDate = timeutil.Strftime(u.SubmissionDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		SubmissionDate = nil
	}
	
	CompletionDate := new(string)
	if u.CompletionDate != nil {
		
		*CompletionDate = timeutil.Strftime(u.CompletionDate, "%Y-%m-%dT%H:%M:%S.%fZ")
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
		State: u.State,
		
		ErrorMessage: u.ErrorMessage,
		
		ExpirationDate: ExpirationDate,
		
		SubmissionDate: SubmissionDate,
		
		CompletionDate: CompletionDate,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Asyncquerystatus) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

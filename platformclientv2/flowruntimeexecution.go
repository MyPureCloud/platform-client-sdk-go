package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Flowruntimeexecution - Details about the current state of a Flow execution
type Flowruntimeexecution struct { 
	// Id - The flow execution ID
	Id *string `json:"id,omitempty"`


	// Name - The flow execution name.
	Name *string `json:"name,omitempty"`


	// FlowVersion - The Version of the flow definition of the flow execution.
	FlowVersion *Flowversion `json:"flowVersion,omitempty"`


	// DateLaunched - The time the flow was launched. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateLaunched *time.Time `json:"dateLaunched,omitempty"`


	// Status - The flow's running status, which indicates whether the flow is running normally or completed, etc.
	Status *string `json:"status,omitempty"`


	// DateCompleted - The time the flow completed, if applicable. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCompleted *time.Time `json:"dateCompleted,omitempty"`


	// CompletionReason - The completion reason set at the flow completion time, if applicable.
	CompletionReason *string `json:"completionReason,omitempty"`


	// FlowErrorInfo - Additional information if the flow is in error
	FlowErrorInfo *Errorbody `json:"flowErrorInfo,omitempty"`


	// OutputData - List of the flow's output variables, if any. Output variables are only supplied for Completed flows.
	OutputData *map[string]interface{} `json:"outputData,omitempty"`


	// Conversation - The conversation to which this Flow execution is related
	Conversation *Domainentityref `json:"conversation,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (u *Flowruntimeexecution) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Flowruntimeexecution

	
	DateLaunched := new(string)
	if u.DateLaunched != nil {
		
		*DateLaunched = timeutil.Strftime(u.DateLaunched, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateLaunched = nil
	}
	
	DateCompleted := new(string)
	if u.DateCompleted != nil {
		
		*DateCompleted = timeutil.Strftime(u.DateCompleted, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCompleted = nil
	}
	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		FlowVersion *Flowversion `json:"flowVersion,omitempty"`
		
		DateLaunched *string `json:"dateLaunched,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		DateCompleted *string `json:"dateCompleted,omitempty"`
		
		CompletionReason *string `json:"completionReason,omitempty"`
		
		FlowErrorInfo *Errorbody `json:"flowErrorInfo,omitempty"`
		
		OutputData *map[string]interface{} `json:"outputData,omitempty"`
		
		Conversation *Domainentityref `json:"conversation,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		FlowVersion: u.FlowVersion,
		
		DateLaunched: DateLaunched,
		
		Status: u.Status,
		
		DateCompleted: DateCompleted,
		
		CompletionReason: u.CompletionReason,
		
		FlowErrorInfo: u.FlowErrorInfo,
		
		OutputData: u.OutputData,
		
		Conversation: u.Conversation,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Flowruntimeexecution) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

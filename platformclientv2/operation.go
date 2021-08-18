package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Operation
type Operation struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Complete
	Complete *bool `json:"complete,omitempty"`


	// User
	User *User `json:"user,omitempty"`


	// Client
	Client *Domainentityref `json:"client,omitempty"`


	// ErrorMessage
	ErrorMessage *string `json:"errorMessage,omitempty"`


	// ErrorCode
	ErrorCode *string `json:"errorCode,omitempty"`


	// ErrorDetails
	ErrorDetails *[]Detail `json:"errorDetails,omitempty"`


	// ErrorMessageParams
	ErrorMessageParams *map[string]string `json:"errorMessageParams,omitempty"`


	// ActionName - Action name
	ActionName *string `json:"actionName,omitempty"`


	// ActionStatus - Action status
	ActionStatus *string `json:"actionStatus,omitempty"`

}

func (u *Operation) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Operation

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Complete *bool `json:"complete,omitempty"`
		
		User *User `json:"user,omitempty"`
		
		Client *Domainentityref `json:"client,omitempty"`
		
		ErrorMessage *string `json:"errorMessage,omitempty"`
		
		ErrorCode *string `json:"errorCode,omitempty"`
		
		ErrorDetails *[]Detail `json:"errorDetails,omitempty"`
		
		ErrorMessageParams *map[string]string `json:"errorMessageParams,omitempty"`
		
		ActionName *string `json:"actionName,omitempty"`
		
		ActionStatus *string `json:"actionStatus,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Complete: u.Complete,
		
		User: u.User,
		
		Client: u.Client,
		
		ErrorMessage: u.ErrorMessage,
		
		ErrorCode: u.ErrorCode,
		
		ErrorDetails: u.ErrorDetails,
		
		ErrorMessageParams: u.ErrorMessageParams,
		
		ActionName: u.ActionName,
		
		ActionStatus: u.ActionStatus,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Operation) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

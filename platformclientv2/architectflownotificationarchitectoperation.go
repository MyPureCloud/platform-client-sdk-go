package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Architectflownotificationarchitectoperation
type Architectflownotificationarchitectoperation struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Complete
	Complete *bool `json:"complete,omitempty"`


	// User
	User *Architectflownotificationuser `json:"user,omitempty"`


	// Client
	Client *Architectflownotificationclient `json:"client,omitempty"`


	// ActionName
	ActionName *string `json:"actionName,omitempty"`


	// ActionStatus
	ActionStatus *string `json:"actionStatus,omitempty"`


	// ErrorMessage
	ErrorMessage *string `json:"errorMessage,omitempty"`


	// ErrorCode
	ErrorCode *string `json:"errorCode,omitempty"`


	// ErrorMessageParams
	ErrorMessageParams *Architectflownotificationerrormessageparams `json:"errorMessageParams,omitempty"`


	// ErrorDetails
	ErrorDetails *[]Architectflownotificationerrordetail `json:"errorDetails,omitempty"`

}

func (u *Architectflownotificationarchitectoperation) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Architectflownotificationarchitectoperation

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Complete *bool `json:"complete,omitempty"`
		
		User *Architectflownotificationuser `json:"user,omitempty"`
		
		Client *Architectflownotificationclient `json:"client,omitempty"`
		
		ActionName *string `json:"actionName,omitempty"`
		
		ActionStatus *string `json:"actionStatus,omitempty"`
		
		ErrorMessage *string `json:"errorMessage,omitempty"`
		
		ErrorCode *string `json:"errorCode,omitempty"`
		
		ErrorMessageParams *Architectflownotificationerrormessageparams `json:"errorMessageParams,omitempty"`
		
		ErrorDetails *[]Architectflownotificationerrordetail `json:"errorDetails,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Complete: u.Complete,
		
		User: u.User,
		
		Client: u.Client,
		
		ActionName: u.ActionName,
		
		ActionStatus: u.ActionStatus,
		
		ErrorMessage: u.ErrorMessage,
		
		ErrorCode: u.ErrorCode,
		
		ErrorMessageParams: u.ErrorMessageParams,
		
		ErrorDetails: u.ErrorDetails,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Architectflownotificationarchitectoperation) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

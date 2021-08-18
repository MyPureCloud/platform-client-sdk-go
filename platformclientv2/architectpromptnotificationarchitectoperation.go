package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Architectpromptnotificationarchitectoperation
type Architectpromptnotificationarchitectoperation struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Complete
	Complete *bool `json:"complete,omitempty"`


	// User
	User *Architectpromptnotificationuser `json:"user,omitempty"`


	// Client
	Client *Architectpromptnotificationclient `json:"client,omitempty"`


	// ActionName
	ActionName *string `json:"actionName,omitempty"`


	// ActionStatus
	ActionStatus *string `json:"actionStatus,omitempty"`


	// ErrorMessage
	ErrorMessage *string `json:"errorMessage,omitempty"`


	// ErrorCode
	ErrorCode *string `json:"errorCode,omitempty"`


	// ErrorMessageParams
	ErrorMessageParams *Architectpromptnotificationerrormessageparams `json:"errorMessageParams,omitempty"`


	// ErrorDetails
	ErrorDetails *[]Architectpromptnotificationerrordetail `json:"errorDetails,omitempty"`

}

func (u *Architectpromptnotificationarchitectoperation) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Architectpromptnotificationarchitectoperation

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Complete *bool `json:"complete,omitempty"`
		
		User *Architectpromptnotificationuser `json:"user,omitempty"`
		
		Client *Architectpromptnotificationclient `json:"client,omitempty"`
		
		ActionName *string `json:"actionName,omitempty"`
		
		ActionStatus *string `json:"actionStatus,omitempty"`
		
		ErrorMessage *string `json:"errorMessage,omitempty"`
		
		ErrorCode *string `json:"errorCode,omitempty"`
		
		ErrorMessageParams *Architectpromptnotificationerrormessageparams `json:"errorMessageParams,omitempty"`
		
		ErrorDetails *[]Architectpromptnotificationerrordetail `json:"errorDetails,omitempty"`
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
func (o *Architectpromptnotificationarchitectoperation) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

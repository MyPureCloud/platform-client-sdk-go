package platformclientv2
import (
	"encoding/json"
)

// Architectflowoutcomenotificationarchitectoperation
type Architectflowoutcomenotificationarchitectoperation struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Complete
	Complete *bool `json:"complete,omitempty"`


	// User
	User *Architectflowoutcomenotificationuser `json:"user,omitempty"`


	// Client
	Client *Architectflowoutcomenotificationclient `json:"client,omitempty"`


	// ActionName
	ActionName *string `json:"actionName,omitempty"`


	// ActionStatus
	ActionStatus *string `json:"actionStatus,omitempty"`


	// ErrorMessage
	ErrorMessage *string `json:"errorMessage,omitempty"`


	// ErrorCode
	ErrorCode *string `json:"errorCode,omitempty"`


	// ErrorMessageParams
	ErrorMessageParams *Architectflowoutcomenotificationerrormessageparams `json:"errorMessageParams,omitempty"`


	// ErrorDetails
	ErrorDetails *[]Architectflowoutcomenotificationerrordetail `json:"errorDetails,omitempty"`

}

// String returns a JSON representation of the model
func (o *Architectflowoutcomenotificationarchitectoperation) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

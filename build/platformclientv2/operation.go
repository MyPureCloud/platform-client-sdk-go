package platformclientv2
import (
	"encoding/json"
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

// String returns a JSON representation of the model
func (o *Operation) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

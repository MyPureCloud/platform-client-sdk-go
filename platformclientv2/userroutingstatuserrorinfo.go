package platformclientv2
import (
	"encoding/json"
)

// Userroutingstatuserrorinfo
type Userroutingstatuserrorinfo struct { 
	// ErrorCode
	ErrorCode *string `json:"errorCode,omitempty"`


	// Status
	Status *int32 `json:"status,omitempty"`


	// CorrelationId
	CorrelationId *string `json:"correlationId,omitempty"`


	// UserMessage
	UserMessage *string `json:"userMessage,omitempty"`


	// UserParamsMessage
	UserParamsMessage *string `json:"userParamsMessage,omitempty"`


	// UserParams
	UserParams *[]Userroutingstatususerparam `json:"userParams,omitempty"`

}

// String returns a JSON representation of the model
func (o *Userroutingstatuserrorinfo) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

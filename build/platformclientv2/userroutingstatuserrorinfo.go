package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Userroutingstatuserrorinfo
type Userroutingstatuserrorinfo struct { 
	// ErrorCode
	ErrorCode *string `json:"errorCode,omitempty"`


	// Status
	Status *int `json:"status,omitempty"`


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
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

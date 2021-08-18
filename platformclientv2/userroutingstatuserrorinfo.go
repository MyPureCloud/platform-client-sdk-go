package platformclientv2
import (
	"github.com/leekchan/timeutil"
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

func (u *Userroutingstatuserrorinfo) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Userroutingstatuserrorinfo

	

	return json.Marshal(&struct { 
		ErrorCode *string `json:"errorCode,omitempty"`
		
		Status *int `json:"status,omitempty"`
		
		CorrelationId *string `json:"correlationId,omitempty"`
		
		UserMessage *string `json:"userMessage,omitempty"`
		
		UserParamsMessage *string `json:"userParamsMessage,omitempty"`
		
		UserParams *[]Userroutingstatususerparam `json:"userParams,omitempty"`
		*Alias
	}{ 
		ErrorCode: u.ErrorCode,
		
		Status: u.Status,
		
		CorrelationId: u.CorrelationId,
		
		UserMessage: u.UserMessage,
		
		UserParamsMessage: u.UserParamsMessage,
		
		UserParams: u.UserParams,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Userroutingstatuserrorinfo) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

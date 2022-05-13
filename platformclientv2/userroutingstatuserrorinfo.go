package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Userroutingstatuserrorinfo - Error information that the Public API will receive in a response body. This allows backend services to pass an error message to consumers of the Public API.
type Userroutingstatuserrorinfo struct { 
	// ErrorCode - A code unique to this error. Typically prefixed with the service that originated the error. For example CONFIG_USER_NOT_FOUND
	ErrorCode *string `json:"errorCode,omitempty"`


	// Status - The HTTP status code for this message. If left blank the status code from the HTTP response is used.
	Status *int `json:"status,omitempty"`


	// CorrelationId - The correlation Id or context Id for this message. If left blank the Public API will look at the HTTP response header 'ININ-Correlation-Id' instead.
	CorrelationId *string `json:"correlationId,omitempty"`


	// UserMessage - A customer friendly message. This should be a complete sentence, use proper grammar and only include information useful to a customer. This is not a dev message and should not include things like Org Id
	UserMessage *string `json:"userMessage,omitempty"`


	// UserParamsMessage - This is the same as userMessage except it uses template fields for variable replacement. For instance: 'User {username} was not found'
	UserParamsMessage *string `json:"userParamsMessage,omitempty"`


	// UserParams - Used in conjunction with userParamsMessage. These are the template parameters. For instance: UserParam.key = 'username', UserParam.value = 'chuck.pulfer'
	UserParams *[]Userroutingstatususerparam `json:"userParams,omitempty"`

}

func (o *Userroutingstatuserrorinfo) MarshalJSON() ([]byte, error) {
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
		ErrorCode: o.ErrorCode,
		
		Status: o.Status,
		
		CorrelationId: o.CorrelationId,
		
		UserMessage: o.UserMessage,
		
		UserParamsMessage: o.UserParamsMessage,
		
		UserParams: o.UserParams,
		Alias:    (*Alias)(o),
	})
}

func (o *Userroutingstatuserrorinfo) UnmarshalJSON(b []byte) error {
	var UserroutingstatuserrorinfoMap map[string]interface{}
	err := json.Unmarshal(b, &UserroutingstatuserrorinfoMap)
	if err != nil {
		return err
	}
	
	if ErrorCode, ok := UserroutingstatuserrorinfoMap["errorCode"].(string); ok {
		o.ErrorCode = &ErrorCode
	}
    
	if Status, ok := UserroutingstatuserrorinfoMap["status"].(float64); ok {
		StatusInt := int(Status)
		o.Status = &StatusInt
	}
	
	if CorrelationId, ok := UserroutingstatuserrorinfoMap["correlationId"].(string); ok {
		o.CorrelationId = &CorrelationId
	}
    
	if UserMessage, ok := UserroutingstatuserrorinfoMap["userMessage"].(string); ok {
		o.UserMessage = &UserMessage
	}
    
	if UserParamsMessage, ok := UserroutingstatuserrorinfoMap["userParamsMessage"].(string); ok {
		o.UserParamsMessage = &UserParamsMessage
	}
    
	if UserParams, ok := UserroutingstatuserrorinfoMap["userParams"].([]interface{}); ok {
		UserParamsString, _ := json.Marshal(UserParams)
		json.Unmarshal(UserParamsString, &o.UserParams)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Userroutingstatuserrorinfo) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

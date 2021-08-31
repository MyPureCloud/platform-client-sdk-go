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

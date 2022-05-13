package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Policyerrormessage
type Policyerrormessage struct { 
	// StatusCode
	StatusCode *int `json:"statusCode,omitempty"`


	// UserMessage
	UserMessage *interface{} `json:"userMessage,omitempty"`


	// UserParamsMessage
	UserParamsMessage *string `json:"userParamsMessage,omitempty"`


	// ErrorCode
	ErrorCode *string `json:"errorCode,omitempty"`


	// CorrelationId
	CorrelationId *string `json:"correlationId,omitempty"`


	// UserParams
	UserParams *[]Userparam `json:"userParams,omitempty"`


	// InsertDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	InsertDate *time.Time `json:"insertDate,omitempty"`

}

func (o *Policyerrormessage) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Policyerrormessage
	
	InsertDate := new(string)
	if o.InsertDate != nil {
		
		*InsertDate = timeutil.Strftime(o.InsertDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		InsertDate = nil
	}
	
	return json.Marshal(&struct { 
		StatusCode *int `json:"statusCode,omitempty"`
		
		UserMessage *interface{} `json:"userMessage,omitempty"`
		
		UserParamsMessage *string `json:"userParamsMessage,omitempty"`
		
		ErrorCode *string `json:"errorCode,omitempty"`
		
		CorrelationId *string `json:"correlationId,omitempty"`
		
		UserParams *[]Userparam `json:"userParams,omitempty"`
		
		InsertDate *string `json:"insertDate,omitempty"`
		*Alias
	}{ 
		StatusCode: o.StatusCode,
		
		UserMessage: o.UserMessage,
		
		UserParamsMessage: o.UserParamsMessage,
		
		ErrorCode: o.ErrorCode,
		
		CorrelationId: o.CorrelationId,
		
		UserParams: o.UserParams,
		
		InsertDate: InsertDate,
		Alias:    (*Alias)(o),
	})
}

func (o *Policyerrormessage) UnmarshalJSON(b []byte) error {
	var PolicyerrormessageMap map[string]interface{}
	err := json.Unmarshal(b, &PolicyerrormessageMap)
	if err != nil {
		return err
	}
	
	if StatusCode, ok := PolicyerrormessageMap["statusCode"].(float64); ok {
		StatusCodeInt := int(StatusCode)
		o.StatusCode = &StatusCodeInt
	}
	
	if UserMessage, ok := PolicyerrormessageMap["userMessage"].(map[string]interface{}); ok {
		UserMessageString, _ := json.Marshal(UserMessage)
		json.Unmarshal(UserMessageString, &o.UserMessage)
	}
	
	if UserParamsMessage, ok := PolicyerrormessageMap["userParamsMessage"].(string); ok {
		o.UserParamsMessage = &UserParamsMessage
	}
    
	if ErrorCode, ok := PolicyerrormessageMap["errorCode"].(string); ok {
		o.ErrorCode = &ErrorCode
	}
    
	if CorrelationId, ok := PolicyerrormessageMap["correlationId"].(string); ok {
		o.CorrelationId = &CorrelationId
	}
    
	if UserParams, ok := PolicyerrormessageMap["userParams"].([]interface{}); ok {
		UserParamsString, _ := json.Marshal(UserParams)
		json.Unmarshal(UserParamsString, &o.UserParams)
	}
	
	if insertDateString, ok := PolicyerrormessageMap["insertDate"].(string); ok {
		InsertDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", insertDateString)
		o.InsertDate = &InsertDate
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Policyerrormessage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

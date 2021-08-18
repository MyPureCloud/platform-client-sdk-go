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

func (u *Policyerrormessage) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Policyerrormessage

	
	InsertDate := new(string)
	if u.InsertDate != nil {
		
		*InsertDate = timeutil.Strftime(u.InsertDate, "%Y-%m-%dT%H:%M:%S.%fZ")
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
		StatusCode: u.StatusCode,
		
		UserMessage: u.UserMessage,
		
		UserParamsMessage: u.UserParamsMessage,
		
		ErrorCode: u.ErrorCode,
		
		CorrelationId: u.CorrelationId,
		
		UserParams: u.UserParams,
		
		InsertDate: InsertDate,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Policyerrormessage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

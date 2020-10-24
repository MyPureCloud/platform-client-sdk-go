package platformclientv2
import (
	"time"
	"encoding/json"
)

// Policyerrormessage
type Policyerrormessage struct { 
	// StatusCode
	StatusCode *int32 `json:"statusCode,omitempty"`


	// UserMessage
	UserMessage *map[string]interface{} `json:"userMessage,omitempty"`


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

// String returns a JSON representation of the model
func (o *Policyerrormessage) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

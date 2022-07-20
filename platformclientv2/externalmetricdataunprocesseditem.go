package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Externalmetricdataunprocesseditem
type Externalmetricdataunprocesseditem struct { 
	// UserId - The user ID. Must provide either userId or userEmail, but not both.
	UserId *string `json:"userId,omitempty"`


	// UserEmail - The user main email used in user's GenesysCloud account. Must provide either userId or userEmail, but not both.
	UserEmail *string `json:"userEmail,omitempty"`


	// MetricId - The ID of the external metric definition
	MetricId *string `json:"metricId,omitempty"`


	// DateOccurred - The date of the metric data. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	DateOccurred *time.Time `json:"dateOccurred,omitempty"`


	// Value - The value of the metric data. When value is null, the metric data will be deleted.
	Value *float64 `json:"value,omitempty"`


	// Count - The number of data points. The default value is 1.
	Count *int `json:"count,omitempty"`


	// Message - The error message
	Message *string `json:"message,omitempty"`


	// Code - The error code
	Code *string `json:"code,omitempty"`

}

func (o *Externalmetricdataunprocesseditem) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Externalmetricdataunprocesseditem
	
	DateOccurred := new(string)
	if o.DateOccurred != nil {
		*DateOccurred = timeutil.Strftime(o.DateOccurred, "%Y-%m-%d")
	} else {
		DateOccurred = nil
	}
	
	return json.Marshal(&struct { 
		UserId *string `json:"userId,omitempty"`
		
		UserEmail *string `json:"userEmail,omitempty"`
		
		MetricId *string `json:"metricId,omitempty"`
		
		DateOccurred *string `json:"dateOccurred,omitempty"`
		
		Value *float64 `json:"value,omitempty"`
		
		Count *int `json:"count,omitempty"`
		
		Message *string `json:"message,omitempty"`
		
		Code *string `json:"code,omitempty"`
		*Alias
	}{ 
		UserId: o.UserId,
		
		UserEmail: o.UserEmail,
		
		MetricId: o.MetricId,
		
		DateOccurred: DateOccurred,
		
		Value: o.Value,
		
		Count: o.Count,
		
		Message: o.Message,
		
		Code: o.Code,
		Alias:    (*Alias)(o),
	})
}

func (o *Externalmetricdataunprocesseditem) UnmarshalJSON(b []byte) error {
	var ExternalmetricdataunprocesseditemMap map[string]interface{}
	err := json.Unmarshal(b, &ExternalmetricdataunprocesseditemMap)
	if err != nil {
		return err
	}
	
	if UserId, ok := ExternalmetricdataunprocesseditemMap["userId"].(string); ok {
		o.UserId = &UserId
	}
    
	if UserEmail, ok := ExternalmetricdataunprocesseditemMap["userEmail"].(string); ok {
		o.UserEmail = &UserEmail
	}
    
	if MetricId, ok := ExternalmetricdataunprocesseditemMap["metricId"].(string); ok {
		o.MetricId = &MetricId
	}
    
	if dateOccurredString, ok := ExternalmetricdataunprocesseditemMap["dateOccurred"].(string); ok {
		DateOccurred, _ := time.Parse("2006-01-02", dateOccurredString)
		o.DateOccurred = &DateOccurred
	}
	
	if Value, ok := ExternalmetricdataunprocesseditemMap["value"].(float64); ok {
		o.Value = &Value
	}
    
	if Count, ok := ExternalmetricdataunprocesseditemMap["count"].(float64); ok {
		CountInt := int(Count)
		o.Count = &CountInt
	}
	
	if Message, ok := ExternalmetricdataunprocesseditemMap["message"].(string); ok {
		o.Message = &Message
	}
    
	if Code, ok := ExternalmetricdataunprocesseditemMap["code"].(string); ok {
		o.Code = &Code
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Externalmetricdataunprocesseditem) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

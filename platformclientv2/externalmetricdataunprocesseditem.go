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
	Value *float32 `json:"value,omitempty"`


	// Count - The number of data points. The default value is 0 when type is Cumulative and the metric data already exists, otherwise 1. When total count reaches 0, the metric data will be deleted.
	Count *int `json:"count,omitempty"`


	// VarType - The type of the metric data. The default value is Total.
	VarType *string `json:"type,omitempty"`


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
		
		Value *float32 `json:"value,omitempty"`
		
		Count *int `json:"count,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
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
		
		VarType: o.VarType,
		
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
		ValueFloat32 := float32(Value)
		o.Value = &ValueFloat32
	}
    
	if Count, ok := ExternalmetricdataunprocesseditemMap["count"].(float64); ok {
		CountInt := int(Count)
		o.Count = &CountInt
	}
	
	if VarType, ok := ExternalmetricdataunprocesseditemMap["type"].(string); ok {
		o.VarType = &VarType
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

package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Externalmetricdataprocesseditem
type Externalmetricdataprocesseditem struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
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

	// TotalValue - The total value of the metric data.
	TotalValue *float32 `json:"totalValue,omitempty"`

	// TotalCount - The total number of data points.
	TotalCount *int `json:"totalCount,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Externalmetricdataprocesseditem) SetField(field string, fieldValue interface{}) {
	// Get Value object for field
	target := reflect.ValueOf(o)
	targetField := reflect.Indirect(target).FieldByName(field)

	// Set value
	if fieldValue != nil {
		targetField.Set(reflect.ValueOf(fieldValue))
	} else {
		// Must create a new Value (creates **type) then get its element (*type), which will be nil pointer of the appropriate type
		x := reflect.Indirect(reflect.New(targetField.Type()))
		targetField.Set(x)
	}

	// Add field to set field names list
	if o.SetFieldNames == nil {
		o.SetFieldNames = make(map[string]bool)
	}
	o.SetFieldNames[field] = true
}

func (o Externalmetricdataprocesseditem) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{  }
		localDateTimeFields := []string{  }
		dateFields := []string{ "DateOccurred", }

		// Construct object
		newObj := make(map[string]interface{})
		for fieldName := range o.SetFieldNames {
			// Get initial field value
			fieldValue := val.FieldByName(fieldName).Interface()

			// Apply value formatting overrides
			if fieldValue == nil || reflect.ValueOf(fieldValue).IsNil()  {
				// Do nothing. Just catching this case to avoid trying to custom serialize a nil value
			} else if contains(dateTimeFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%dT%H:%M:%S.%fZ")
			} else if contains(localDateTimeFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%dT%H:%M:%S.%f")
			} else if contains(dateFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%d")
			}

			// Assign value to field using JSON tag name
			newObj[getFieldName(reflect.TypeOf(&o), fieldName)] = fieldValue
		}

		// Marshal and return dynamically constructed interface
		return json.Marshal(newObj)
	}

	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Externalmetricdataprocesseditem
	
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
		
		TotalValue *float32 `json:"totalValue,omitempty"`
		
		TotalCount *int `json:"totalCount,omitempty"`
		Alias
	}{ 
		UserId: o.UserId,
		
		UserEmail: o.UserEmail,
		
		MetricId: o.MetricId,
		
		DateOccurred: DateOccurred,
		
		Value: o.Value,
		
		Count: o.Count,
		
		VarType: o.VarType,
		
		TotalValue: o.TotalValue,
		
		TotalCount: o.TotalCount,
		Alias:    (Alias)(o),
	})
}

func (o *Externalmetricdataprocesseditem) UnmarshalJSON(b []byte) error {
	var ExternalmetricdataprocesseditemMap map[string]interface{}
	err := json.Unmarshal(b, &ExternalmetricdataprocesseditemMap)
	if err != nil {
		return err
	}
	
	if UserId, ok := ExternalmetricdataprocesseditemMap["userId"].(string); ok {
		o.UserId = &UserId
	}
    
	if UserEmail, ok := ExternalmetricdataprocesseditemMap["userEmail"].(string); ok {
		o.UserEmail = &UserEmail
	}
    
	if MetricId, ok := ExternalmetricdataprocesseditemMap["metricId"].(string); ok {
		o.MetricId = &MetricId
	}
    
	if dateOccurredString, ok := ExternalmetricdataprocesseditemMap["dateOccurred"].(string); ok {
		DateOccurred, _ := time.Parse("2006-01-02", dateOccurredString)
		o.DateOccurred = &DateOccurred
	}
	
	if Value, ok := ExternalmetricdataprocesseditemMap["value"].(float64); ok {
		ValueFloat32 := float32(Value)
		o.Value = &ValueFloat32
	}
    
	if Count, ok := ExternalmetricdataprocesseditemMap["count"].(float64); ok {
		CountInt := int(Count)
		o.Count = &CountInt
	}
	
	if VarType, ok := ExternalmetricdataprocesseditemMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if TotalValue, ok := ExternalmetricdataprocesseditemMap["totalValue"].(float64); ok {
		TotalValueFloat32 := float32(TotalValue)
		o.TotalValue = &TotalValueFloat32
	}
    
	if TotalCount, ok := ExternalmetricdataprocesseditemMap["totalCount"].(float64); ok {
		TotalCountInt := int(TotalCount)
		o.TotalCount = &TotalCountInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Externalmetricdataprocesseditem) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

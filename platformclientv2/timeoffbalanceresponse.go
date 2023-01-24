package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Timeoffbalanceresponse
type Timeoffbalanceresponse struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// ActivityCodeId - The ID for activity code associated with time off balance
	ActivityCodeId *string `json:"activityCodeId,omitempty"`

	// HrisTimeOffTypeId - The ID of the time off type configured in HRIS integration
	HrisTimeOffTypeId *string `json:"hrisTimeOffTypeId,omitempty"`

	// HrisTimeOffTypeSecondaryId - The secondary ID of the time off type configured in HRIS integration
	HrisTimeOffTypeSecondaryId *string `json:"hrisTimeOffTypeSecondaryId,omitempty"`

	// StartDate - The Start date of the requested date range. The end date is determined by the size of interval list. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	StartDate *time.Time `json:"startDate,omitempty"`

	// BalanceMinutesPerDay - The list of available time off balance values in minutes for each day
	BalanceMinutesPerDay *[]int `json:"balanceMinutesPerDay,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Timeoffbalanceresponse) SetField(field string, fieldValue interface{}) {
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

func (o Timeoffbalanceresponse) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{  }
		localDateTimeFields := []string{  }
		dateFields := []string{ "StartDate", }

		// Construct object
		newObj := make(map[string]interface{})
		for fieldName := range o.SetFieldNames {
			// Get initial field value
			fieldValue := val.FieldByName(fieldName).Interface()

			// Apply value formatting overrides
			if contains(dateTimeFields, fieldName) {
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
	type Alias Timeoffbalanceresponse
	
	StartDate := new(string)
	if o.StartDate != nil {
		*StartDate = timeutil.Strftime(o.StartDate, "%Y-%m-%d")
	} else {
		StartDate = nil
	}
	
	return json.Marshal(&struct { 
		ActivityCodeId *string `json:"activityCodeId,omitempty"`
		
		HrisTimeOffTypeId *string `json:"hrisTimeOffTypeId,omitempty"`
		
		HrisTimeOffTypeSecondaryId *string `json:"hrisTimeOffTypeSecondaryId,omitempty"`
		
		StartDate *string `json:"startDate,omitempty"`
		
		BalanceMinutesPerDay *[]int `json:"balanceMinutesPerDay,omitempty"`
		Alias
	}{ 
		ActivityCodeId: o.ActivityCodeId,
		
		HrisTimeOffTypeId: o.HrisTimeOffTypeId,
		
		HrisTimeOffTypeSecondaryId: o.HrisTimeOffTypeSecondaryId,
		
		StartDate: StartDate,
		
		BalanceMinutesPerDay: o.BalanceMinutesPerDay,
		Alias:    (Alias)(o),
	})
}

func (o *Timeoffbalanceresponse) UnmarshalJSON(b []byte) error {
	var TimeoffbalanceresponseMap map[string]interface{}
	err := json.Unmarshal(b, &TimeoffbalanceresponseMap)
	if err != nil {
		return err
	}
	
	if ActivityCodeId, ok := TimeoffbalanceresponseMap["activityCodeId"].(string); ok {
		o.ActivityCodeId = &ActivityCodeId
	}
    
	if HrisTimeOffTypeId, ok := TimeoffbalanceresponseMap["hrisTimeOffTypeId"].(string); ok {
		o.HrisTimeOffTypeId = &HrisTimeOffTypeId
	}
    
	if HrisTimeOffTypeSecondaryId, ok := TimeoffbalanceresponseMap["hrisTimeOffTypeSecondaryId"].(string); ok {
		o.HrisTimeOffTypeSecondaryId = &HrisTimeOffTypeSecondaryId
	}
    
	if startDateString, ok := TimeoffbalanceresponseMap["startDate"].(string); ok {
		StartDate, _ := time.Parse("2006-01-02", startDateString)
		o.StartDate = &StartDate
	}
	
	if BalanceMinutesPerDay, ok := TimeoffbalanceresponseMap["balanceMinutesPerDay"].([]interface{}); ok {
		BalanceMinutesPerDayString, _ := json.Marshal(BalanceMinutesPerDay)
		json.Unmarshal(BalanceMinutesPerDayString, &o.BalanceMinutesPerDay)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Timeoffbalanceresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

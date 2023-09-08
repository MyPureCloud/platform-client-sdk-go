package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Estimateavailablepartialdaytimeoffresponse
type Estimateavailablepartialdaytimeoffresponse struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Date - Start date-time in ISO-8601 format for partial day request
	Date *time.Time `json:"date,omitempty"`

	// DurationMinutes - An estimation of time off request length in minutes
	DurationMinutes *int `json:"durationMinutes,omitempty"`

	// PayableMinutes - An estimation of payable part of time off request in minutes
	PayableMinutes *int `json:"payableMinutes,omitempty"`

	// Flexible - Whether there is flexibility for a user to choose different hours than the system estimated
	Flexible *bool `json:"flexible,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Estimateavailablepartialdaytimeoffresponse) SetField(field string, fieldValue interface{}) {
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

func (o Estimateavailablepartialdaytimeoffresponse) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "Date", }
		localDateTimeFields := []string{  }
		dateFields := []string{  }

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
	type Alias Estimateavailablepartialdaytimeoffresponse
	
	Date := new(string)
	if o.Date != nil {
		
		*Date = timeutil.Strftime(o.Date, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		Date = nil
	}
	
	return json.Marshal(&struct { 
		Date *string `json:"date,omitempty"`
		
		DurationMinutes *int `json:"durationMinutes,omitempty"`
		
		PayableMinutes *int `json:"payableMinutes,omitempty"`
		
		Flexible *bool `json:"flexible,omitempty"`
		Alias
	}{ 
		Date: Date,
		
		DurationMinutes: o.DurationMinutes,
		
		PayableMinutes: o.PayableMinutes,
		
		Flexible: o.Flexible,
		Alias:    (Alias)(o),
	})
}

func (o *Estimateavailablepartialdaytimeoffresponse) UnmarshalJSON(b []byte) error {
	var EstimateavailablepartialdaytimeoffresponseMap map[string]interface{}
	err := json.Unmarshal(b, &EstimateavailablepartialdaytimeoffresponseMap)
	if err != nil {
		return err
	}
	
	if dateString, ok := EstimateavailablepartialdaytimeoffresponseMap["date"].(string); ok {
		Date, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateString)
		o.Date = &Date
	}
	
	if DurationMinutes, ok := EstimateavailablepartialdaytimeoffresponseMap["durationMinutes"].(float64); ok {
		DurationMinutesInt := int(DurationMinutes)
		o.DurationMinutes = &DurationMinutesInt
	}
	
	if PayableMinutes, ok := EstimateavailablepartialdaytimeoffresponseMap["payableMinutes"].(float64); ok {
		PayableMinutesInt := int(PayableMinutes)
		o.PayableMinutes = &PayableMinutesInt
	}
	
	if Flexible, ok := EstimateavailablepartialdaytimeoffresponseMap["flexible"].(bool); ok {
		o.Flexible = &Flexible
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Estimateavailablepartialdaytimeoffresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

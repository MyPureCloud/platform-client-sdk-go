package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Locationaddressverificationdetails
type Locationaddressverificationdetails struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Status - Status of address verification process
	Status *string `json:"status,omitempty"`

	// DateFinished - Finished time of address verification process. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateFinished *time.Time `json:"dateFinished,omitempty"`

	// DateStarted - Time started of address verification process. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateStarted *time.Time `json:"dateStarted,omitempty"`

	// Service - Third party service used for address verification
	Service *string `json:"service,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Locationaddressverificationdetails) SetField(field string, fieldValue interface{}) {
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

func (o Locationaddressverificationdetails) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateFinished","DateStarted", }
		localDateTimeFields := []string{  }
		dateFields := []string{  }

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
	type Alias Locationaddressverificationdetails
	
	DateFinished := new(string)
	if o.DateFinished != nil {
		
		*DateFinished = timeutil.Strftime(o.DateFinished, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateFinished = nil
	}
	
	DateStarted := new(string)
	if o.DateStarted != nil {
		
		*DateStarted = timeutil.Strftime(o.DateStarted, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateStarted = nil
	}
	
	return json.Marshal(&struct { 
		Status *string `json:"status,omitempty"`
		
		DateFinished *string `json:"dateFinished,omitempty"`
		
		DateStarted *string `json:"dateStarted,omitempty"`
		
		Service *string `json:"service,omitempty"`
		Alias
	}{ 
		Status: o.Status,
		
		DateFinished: DateFinished,
		
		DateStarted: DateStarted,
		
		Service: o.Service,
		Alias:    (Alias)(o),
	})
}

func (o *Locationaddressverificationdetails) UnmarshalJSON(b []byte) error {
	var LocationaddressverificationdetailsMap map[string]interface{}
	err := json.Unmarshal(b, &LocationaddressverificationdetailsMap)
	if err != nil {
		return err
	}
	
	if Status, ok := LocationaddressverificationdetailsMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if dateFinishedString, ok := LocationaddressverificationdetailsMap["dateFinished"].(string); ok {
		DateFinished, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateFinishedString)
		o.DateFinished = &DateFinished
	}
	
	if dateStartedString, ok := LocationaddressverificationdetailsMap["dateStarted"].(string); ok {
		DateStarted, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateStartedString)
		o.DateStarted = &DateStarted
	}
	
	if Service, ok := LocationaddressverificationdetailsMap["service"].(string); ok {
		o.Service = &Service
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Locationaddressverificationdetails) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Gamificationstatus
type Gamificationstatus struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// IsActive - Gamification status of the organization.
	IsActive *bool `json:"isActive,omitempty"`

	// DateStart - Gamification start date. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	DateStart *time.Time `json:"dateStart,omitempty"`

	// AutomaticUserAssignment - Automatic assignment of users to the default profile
	AutomaticUserAssignment *bool `json:"automaticUserAssignment,omitempty"`

	// DateStartPersonalBest - Personal best aggregation starting date. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	DateStartPersonalBest *time.Time `json:"dateStartPersonalBest,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Gamificationstatus) SetField(field string, fieldValue interface{}) {
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

func (o Gamificationstatus) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{  }
		localDateTimeFields := []string{  }
		dateFields := []string{ "DateStart","DateStartPersonalBest", }

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
	type Alias Gamificationstatus
	
	DateStart := new(string)
	if o.DateStart != nil {
		*DateStart = timeutil.Strftime(o.DateStart, "%Y-%m-%d")
	} else {
		DateStart = nil
	}
	
	DateStartPersonalBest := new(string)
	if o.DateStartPersonalBest != nil {
		*DateStartPersonalBest = timeutil.Strftime(o.DateStartPersonalBest, "%Y-%m-%d")
	} else {
		DateStartPersonalBest = nil
	}
	
	return json.Marshal(&struct { 
		IsActive *bool `json:"isActive,omitempty"`
		
		DateStart *string `json:"dateStart,omitempty"`
		
		AutomaticUserAssignment *bool `json:"automaticUserAssignment,omitempty"`
		
		DateStartPersonalBest *string `json:"dateStartPersonalBest,omitempty"`
		Alias
	}{ 
		IsActive: o.IsActive,
		
		DateStart: DateStart,
		
		AutomaticUserAssignment: o.AutomaticUserAssignment,
		
		DateStartPersonalBest: DateStartPersonalBest,
		Alias:    (Alias)(o),
	})
}

func (o *Gamificationstatus) UnmarshalJSON(b []byte) error {
	var GamificationstatusMap map[string]interface{}
	err := json.Unmarshal(b, &GamificationstatusMap)
	if err != nil {
		return err
	}
	
	if IsActive, ok := GamificationstatusMap["isActive"].(bool); ok {
		o.IsActive = &IsActive
	}
    
	if dateStartString, ok := GamificationstatusMap["dateStart"].(string); ok {
		DateStart, _ := time.Parse("2006-01-02", dateStartString)
		o.DateStart = &DateStart
	}
	
	if AutomaticUserAssignment, ok := GamificationstatusMap["automaticUserAssignment"].(bool); ok {
		o.AutomaticUserAssignment = &AutomaticUserAssignment
	}
    
	if dateStartPersonalBestString, ok := GamificationstatusMap["dateStartPersonalBest"].(string); ok {
		DateStartPersonalBest, _ := time.Parse("2006-01-02", dateStartPersonalBestString)
		o.DateStartPersonalBest = &DateStartPersonalBest
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Gamificationstatus) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

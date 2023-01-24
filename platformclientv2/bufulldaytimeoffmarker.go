package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Bufulldaytimeoffmarker
type Bufulldaytimeoffmarker struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// BusinessUnitDate - The date of the time off marker, interpreted in the business unit's time zone. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	BusinessUnitDate *time.Time `json:"businessUnitDate,omitempty"`

	// LengthMinutes - The length of the time off marker in minutes
	LengthMinutes *int `json:"lengthMinutes,omitempty"`

	// Description - The description of the time off marker
	Description *string `json:"description,omitempty"`

	// ActivityCodeId - The ID of the activity code associated with the time off marker
	ActivityCodeId *string `json:"activityCodeId,omitempty"`

	// Paid - Whether the time off marker is paid
	Paid *bool `json:"paid,omitempty"`

	// TimeOffRequestId - The ID of the time off request
	TimeOffRequestId *string `json:"timeOffRequestId,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Bufulldaytimeoffmarker) SetField(field string, fieldValue interface{}) {
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

func (o Bufulldaytimeoffmarker) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{  }
		localDateTimeFields := []string{  }
		dateFields := []string{ "BusinessUnitDate", }

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
	type Alias Bufulldaytimeoffmarker
	
	BusinessUnitDate := new(string)
	if o.BusinessUnitDate != nil {
		*BusinessUnitDate = timeutil.Strftime(o.BusinessUnitDate, "%Y-%m-%d")
	} else {
		BusinessUnitDate = nil
	}
	
	return json.Marshal(&struct { 
		BusinessUnitDate *string `json:"businessUnitDate,omitempty"`
		
		LengthMinutes *int `json:"lengthMinutes,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		ActivityCodeId *string `json:"activityCodeId,omitempty"`
		
		Paid *bool `json:"paid,omitempty"`
		
		TimeOffRequestId *string `json:"timeOffRequestId,omitempty"`
		Alias
	}{ 
		BusinessUnitDate: BusinessUnitDate,
		
		LengthMinutes: o.LengthMinutes,
		
		Description: o.Description,
		
		ActivityCodeId: o.ActivityCodeId,
		
		Paid: o.Paid,
		
		TimeOffRequestId: o.TimeOffRequestId,
		Alias:    (Alias)(o),
	})
}

func (o *Bufulldaytimeoffmarker) UnmarshalJSON(b []byte) error {
	var BufulldaytimeoffmarkerMap map[string]interface{}
	err := json.Unmarshal(b, &BufulldaytimeoffmarkerMap)
	if err != nil {
		return err
	}
	
	if businessUnitDateString, ok := BufulldaytimeoffmarkerMap["businessUnitDate"].(string); ok {
		BusinessUnitDate, _ := time.Parse("2006-01-02", businessUnitDateString)
		o.BusinessUnitDate = &BusinessUnitDate
	}
	
	if LengthMinutes, ok := BufulldaytimeoffmarkerMap["lengthMinutes"].(float64); ok {
		LengthMinutesInt := int(LengthMinutes)
		o.LengthMinutes = &LengthMinutesInt
	}
	
	if Description, ok := BufulldaytimeoffmarkerMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if ActivityCodeId, ok := BufulldaytimeoffmarkerMap["activityCodeId"].(string); ok {
		o.ActivityCodeId = &ActivityCodeId
	}
    
	if Paid, ok := BufulldaytimeoffmarkerMap["paid"].(bool); ok {
		o.Paid = &Paid
	}
    
	if TimeOffRequestId, ok := BufulldaytimeoffmarkerMap["timeOffRequestId"].(string); ok {
		o.TimeOffRequestId = &TimeOffRequestId
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Bufulldaytimeoffmarker) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

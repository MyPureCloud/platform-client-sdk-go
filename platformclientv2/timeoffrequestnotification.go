package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Timeoffrequestnotification
type Timeoffrequestnotification struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// TimeOffRequestId - The ID of this time off request
	TimeOffRequestId *string `json:"timeOffRequestId,omitempty"`

	// User - The user associated with this time off request
	User *Userreference `json:"user,omitempty"`

	// IsFullDayRequest - Whether this is a full day request (false means partial day)
	IsFullDayRequest *bool `json:"isFullDayRequest,omitempty"`

	// Status - The status of this time off request
	Status *string `json:"status,omitempty"`

	// PartialDayStartDateTimes - A set of start date-times in ISO-8601 format for partial day requests.  Will be not empty if isFullDayRequest == false
	PartialDayStartDateTimes *[]time.Time `json:"partialDayStartDateTimes,omitempty"`

	// FullDayManagementUnitDates - A set of dates in yyyy-MM-dd format.  Should be interpreted in the management unit's configured time zone.  Will be not empty if isFullDayRequest == true
	FullDayManagementUnitDates *[]string `json:"fullDayManagementUnitDates,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Timeoffrequestnotification) SetField(field string, fieldValue interface{}) {
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

func (o Timeoffrequestnotification) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{  }
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
	type Alias Timeoffrequestnotification
	
	return json.Marshal(&struct { 
		TimeOffRequestId *string `json:"timeOffRequestId,omitempty"`
		
		User *Userreference `json:"user,omitempty"`
		
		IsFullDayRequest *bool `json:"isFullDayRequest,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		PartialDayStartDateTimes *[]time.Time `json:"partialDayStartDateTimes,omitempty"`
		
		FullDayManagementUnitDates *[]string `json:"fullDayManagementUnitDates,omitempty"`
		Alias
	}{ 
		TimeOffRequestId: o.TimeOffRequestId,
		
		User: o.User,
		
		IsFullDayRequest: o.IsFullDayRequest,
		
		Status: o.Status,
		
		PartialDayStartDateTimes: o.PartialDayStartDateTimes,
		
		FullDayManagementUnitDates: o.FullDayManagementUnitDates,
		Alias:    (Alias)(o),
	})
}

func (o *Timeoffrequestnotification) UnmarshalJSON(b []byte) error {
	var TimeoffrequestnotificationMap map[string]interface{}
	err := json.Unmarshal(b, &TimeoffrequestnotificationMap)
	if err != nil {
		return err
	}
	
	if TimeOffRequestId, ok := TimeoffrequestnotificationMap["timeOffRequestId"].(string); ok {
		o.TimeOffRequestId = &TimeOffRequestId
	}
    
	if User, ok := TimeoffrequestnotificationMap["user"].(map[string]interface{}); ok {
		UserString, _ := json.Marshal(User)
		json.Unmarshal(UserString, &o.User)
	}
	
	if IsFullDayRequest, ok := TimeoffrequestnotificationMap["isFullDayRequest"].(bool); ok {
		o.IsFullDayRequest = &IsFullDayRequest
	}
    
	if Status, ok := TimeoffrequestnotificationMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if PartialDayStartDateTimes, ok := TimeoffrequestnotificationMap["partialDayStartDateTimes"].([]interface{}); ok {
		PartialDayStartDateTimesString, _ := json.Marshal(PartialDayStartDateTimes)
		json.Unmarshal(PartialDayStartDateTimesString, &o.PartialDayStartDateTimes)
	}
	
	if FullDayManagementUnitDates, ok := TimeoffrequestnotificationMap["fullDayManagementUnitDates"].([]interface{}); ok {
		FullDayManagementUnitDatesString, _ := json.Marshal(FullDayManagementUnitDates)
		json.Unmarshal(FullDayManagementUnitDatesString, &o.FullDayManagementUnitDates)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Timeoffrequestnotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

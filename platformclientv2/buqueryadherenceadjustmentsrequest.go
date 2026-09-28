package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Buqueryadherenceadjustmentsrequest
type Buqueryadherenceadjustmentsrequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// StartDate - The start timestamp of the range to query in ISO-8601 format
	StartDate *time.Time `json:"startDate,omitempty"`

	// EndDate - The end timestamp of the range to query in ISO-8601 format
	EndDate *time.Time `json:"endDate,omitempty"`

	// ReasonCodeIds - A filter for the reason codes to include. Leave empty or omit entirely for all reason codes
	ReasonCodeIds *[]string `json:"reasonCodeIds,omitempty"`

	// Statuses - A filter for which adherence adjustment statuses to include. Leave empty or omit entirely for all statuses
	Statuses *[]string `json:"statuses,omitempty"`

	// UserIds - A filter for which users within the business unit to query. Leave empty or omit entirely for all users
	UserIds *[]string `json:"userIds,omitempty"`

	// ManagementUnitIds - A filter for which management units to query. Leave empty or omit entirely for all management units in the business unit
	ManagementUnitIds *[]string `json:"managementUnitIds,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Buqueryadherenceadjustmentsrequest) SetField(field string, fieldValue interface{}) {
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

func (o Buqueryadherenceadjustmentsrequest) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "StartDate","EndDate", }
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
	type Alias Buqueryadherenceadjustmentsrequest
	
	StartDate := new(string)
	if o.StartDate != nil {
		
		*StartDate = timeutil.Strftime(o.StartDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartDate = nil
	}
	
	EndDate := new(string)
	if o.EndDate != nil {
		
		*EndDate = timeutil.Strftime(o.EndDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		EndDate = nil
	}
	
	return json.Marshal(&struct { 
		StartDate *string `json:"startDate,omitempty"`
		
		EndDate *string `json:"endDate,omitempty"`
		
		ReasonCodeIds *[]string `json:"reasonCodeIds,omitempty"`
		
		Statuses *[]string `json:"statuses,omitempty"`
		
		UserIds *[]string `json:"userIds,omitempty"`
		
		ManagementUnitIds *[]string `json:"managementUnitIds,omitempty"`
		Alias
	}{ 
		StartDate: StartDate,
		
		EndDate: EndDate,
		
		ReasonCodeIds: o.ReasonCodeIds,
		
		Statuses: o.Statuses,
		
		UserIds: o.UserIds,
		
		ManagementUnitIds: o.ManagementUnitIds,
		Alias:    (Alias)(o),
	})
}

func (o *Buqueryadherenceadjustmentsrequest) UnmarshalJSON(b []byte) error {
	var BuqueryadherenceadjustmentsrequestMap map[string]interface{}
	err := json.Unmarshal(b, &BuqueryadherenceadjustmentsrequestMap)
	if err != nil {
		return err
	}
	
	if startDateString, ok := BuqueryadherenceadjustmentsrequestMap["startDate"].(string); ok {
		StartDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", startDateString)
		o.StartDate = &StartDate
	}
	
	if endDateString, ok := BuqueryadherenceadjustmentsrequestMap["endDate"].(string); ok {
		EndDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", endDateString)
		o.EndDate = &EndDate
	}
	
	if ReasonCodeIds, ok := BuqueryadherenceadjustmentsrequestMap["reasonCodeIds"].([]interface{}); ok {
		ReasonCodeIdsString, _ := json.Marshal(ReasonCodeIds)
		json.Unmarshal(ReasonCodeIdsString, &o.ReasonCodeIds)
	}
	
	if Statuses, ok := BuqueryadherenceadjustmentsrequestMap["statuses"].([]interface{}); ok {
		StatusesString, _ := json.Marshal(Statuses)
		json.Unmarshal(StatusesString, &o.Statuses)
	}
	
	if UserIds, ok := BuqueryadherenceadjustmentsrequestMap["userIds"].([]interface{}); ok {
		UserIdsString, _ := json.Marshal(UserIds)
		json.Unmarshal(UserIdsString, &o.UserIds)
	}
	
	if ManagementUnitIds, ok := BuqueryadherenceadjustmentsrequestMap["managementUnitIds"].([]interface{}); ok {
		ManagementUnitIdsString, _ := json.Marshal(ManagementUnitIds)
		json.Unmarshal(ManagementUnitIdsString, &o.ManagementUnitIds)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Buqueryadherenceadjustmentsrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

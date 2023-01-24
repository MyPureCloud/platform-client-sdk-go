package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Updatetimeoffplanrequest
type Updatetimeoffplanrequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Name - The name of this time off plan.
	Name *string `json:"name,omitempty"`

	// ActivityCodeIds - The set of activity code IDs to associate with this time off plan.
	ActivityCodeIds *Setwrapperstring `json:"activityCodeIds,omitempty"`

	// TimeOffLimitIds - The set of time off limit IDs to associate with this time off plan.
	TimeOffLimitIds *Setwrapperstring `json:"timeOffLimitIds,omitempty"`

	// AutoApprovalRule - Auto approval rule for the time off plan.
	AutoApprovalRule *string `json:"autoApprovalRule,omitempty"`

	// DaysBeforeStartToExpireFromWaitlist - The number of days before the time off request start date for when the request will be expired from the waitlist.
	DaysBeforeStartToExpireFromWaitlist *int `json:"daysBeforeStartToExpireFromWaitlist,omitempty"`

	// Active - Whether this time off plan should be used by agents.
	Active *bool `json:"active,omitempty"`

	// Metadata - Version metadata for the time off plan
	Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Updatetimeoffplanrequest) SetField(field string, fieldValue interface{}) {
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

func (o Updatetimeoffplanrequest) MarshalJSON() ([]byte, error) {
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
	type Alias Updatetimeoffplanrequest
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		ActivityCodeIds *Setwrapperstring `json:"activityCodeIds,omitempty"`
		
		TimeOffLimitIds *Setwrapperstring `json:"timeOffLimitIds,omitempty"`
		
		AutoApprovalRule *string `json:"autoApprovalRule,omitempty"`
		
		DaysBeforeStartToExpireFromWaitlist *int `json:"daysBeforeStartToExpireFromWaitlist,omitempty"`
		
		Active *bool `json:"active,omitempty"`
		
		Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`
		Alias
	}{ 
		Name: o.Name,
		
		ActivityCodeIds: o.ActivityCodeIds,
		
		TimeOffLimitIds: o.TimeOffLimitIds,
		
		AutoApprovalRule: o.AutoApprovalRule,
		
		DaysBeforeStartToExpireFromWaitlist: o.DaysBeforeStartToExpireFromWaitlist,
		
		Active: o.Active,
		
		Metadata: o.Metadata,
		Alias:    (Alias)(o),
	})
}

func (o *Updatetimeoffplanrequest) UnmarshalJSON(b []byte) error {
	var UpdatetimeoffplanrequestMap map[string]interface{}
	err := json.Unmarshal(b, &UpdatetimeoffplanrequestMap)
	if err != nil {
		return err
	}
	
	if Name, ok := UpdatetimeoffplanrequestMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if ActivityCodeIds, ok := UpdatetimeoffplanrequestMap["activityCodeIds"].(map[string]interface{}); ok {
		ActivityCodeIdsString, _ := json.Marshal(ActivityCodeIds)
		json.Unmarshal(ActivityCodeIdsString, &o.ActivityCodeIds)
	}
	
	if TimeOffLimitIds, ok := UpdatetimeoffplanrequestMap["timeOffLimitIds"].(map[string]interface{}); ok {
		TimeOffLimitIdsString, _ := json.Marshal(TimeOffLimitIds)
		json.Unmarshal(TimeOffLimitIdsString, &o.TimeOffLimitIds)
	}
	
	if AutoApprovalRule, ok := UpdatetimeoffplanrequestMap["autoApprovalRule"].(string); ok {
		o.AutoApprovalRule = &AutoApprovalRule
	}
    
	if DaysBeforeStartToExpireFromWaitlist, ok := UpdatetimeoffplanrequestMap["daysBeforeStartToExpireFromWaitlist"].(float64); ok {
		DaysBeforeStartToExpireFromWaitlistInt := int(DaysBeforeStartToExpireFromWaitlist)
		o.DaysBeforeStartToExpireFromWaitlist = &DaysBeforeStartToExpireFromWaitlistInt
	}
	
	if Active, ok := UpdatetimeoffplanrequestMap["active"].(bool); ok {
		o.Active = &Active
	}
    
	if Metadata, ok := UpdatetimeoffplanrequestMap["metadata"].(map[string]interface{}); ok {
		MetadataString, _ := json.Marshal(Metadata)
		json.Unmarshal(MetadataString, &o.Metadata)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Updatetimeoffplanrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

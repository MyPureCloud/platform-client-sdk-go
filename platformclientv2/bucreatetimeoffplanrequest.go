package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Bucreatetimeoffplanrequest
type Bucreatetimeoffplanrequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Name - The name of this time-off plan
	Name *string `json:"name,omitempty"`

	// ActivityCodeIds - The IDs of activity codes to associate with this time-off plan
	ActivityCodeIds *[]string `json:"activityCodeIds,omitempty"`

	// AutoApprovalRule - Auto approval rule for this time-off plan. Default is Never
	AutoApprovalRule *string `json:"autoApprovalRule,omitempty"`

	// DaysBeforeStartToExpireFromWaitlist - The number of days before the time-off request start date for when the request will be expired from the waitlist. Default is 0
	DaysBeforeStartToExpireFromWaitlist *int `json:"daysBeforeStartToExpireFromWaitlist,omitempty"`

	// HrisTimeOffType - Time-off type, if this time-off plan is associated with the integration
	HrisTimeOffType *Hristimeofftype `json:"hrisTimeOffType,omitempty"`

	// Enabled - Whether this time-off plan should be used by agents. Default is true
	Enabled *bool `json:"enabled,omitempty"`

	// CountAgainstTimeOffLimits - Whether this time-off plan should count against time-off limits. Default is false
	CountAgainstTimeOffLimits *bool `json:"countAgainstTimeOffLimits,omitempty"`

	// BusinessUnitAssociation - Business unit association, if the time-off plan belongs to a business unit. managementUnitAssociation must not be set if this is populated
	BusinessUnitAssociation *Createtimeoffplanbusinessunitassociation `json:"businessUnitAssociation,omitempty"`

	// ManagementUnitAssociation - Management unit association, if the time-off plan belongs to a management unit. businessUnitAssociation must not be set if this is populated
	ManagementUnitAssociation *Createtimeoffplanmanagementunitassociation `json:"managementUnitAssociation,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Bucreatetimeoffplanrequest) SetField(field string, fieldValue interface{}) {
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

func (o Bucreatetimeoffplanrequest) MarshalJSON() ([]byte, error) {
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
	type Alias Bucreatetimeoffplanrequest
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		ActivityCodeIds *[]string `json:"activityCodeIds,omitempty"`
		
		AutoApprovalRule *string `json:"autoApprovalRule,omitempty"`
		
		DaysBeforeStartToExpireFromWaitlist *int `json:"daysBeforeStartToExpireFromWaitlist,omitempty"`
		
		HrisTimeOffType *Hristimeofftype `json:"hrisTimeOffType,omitempty"`
		
		Enabled *bool `json:"enabled,omitempty"`
		
		CountAgainstTimeOffLimits *bool `json:"countAgainstTimeOffLimits,omitempty"`
		
		BusinessUnitAssociation *Createtimeoffplanbusinessunitassociation `json:"businessUnitAssociation,omitempty"`
		
		ManagementUnitAssociation *Createtimeoffplanmanagementunitassociation `json:"managementUnitAssociation,omitempty"`
		Alias
	}{ 
		Name: o.Name,
		
		ActivityCodeIds: o.ActivityCodeIds,
		
		AutoApprovalRule: o.AutoApprovalRule,
		
		DaysBeforeStartToExpireFromWaitlist: o.DaysBeforeStartToExpireFromWaitlist,
		
		HrisTimeOffType: o.HrisTimeOffType,
		
		Enabled: o.Enabled,
		
		CountAgainstTimeOffLimits: o.CountAgainstTimeOffLimits,
		
		BusinessUnitAssociation: o.BusinessUnitAssociation,
		
		ManagementUnitAssociation: o.ManagementUnitAssociation,
		Alias:    (Alias)(o),
	})
}

func (o *Bucreatetimeoffplanrequest) UnmarshalJSON(b []byte) error {
	var BucreatetimeoffplanrequestMap map[string]interface{}
	err := json.Unmarshal(b, &BucreatetimeoffplanrequestMap)
	if err != nil {
		return err
	}
	
	if Name, ok := BucreatetimeoffplanrequestMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if ActivityCodeIds, ok := BucreatetimeoffplanrequestMap["activityCodeIds"].([]interface{}); ok {
		ActivityCodeIdsString, _ := json.Marshal(ActivityCodeIds)
		json.Unmarshal(ActivityCodeIdsString, &o.ActivityCodeIds)
	}
	
	if AutoApprovalRule, ok := BucreatetimeoffplanrequestMap["autoApprovalRule"].(string); ok {
		o.AutoApprovalRule = &AutoApprovalRule
	}
    
	if DaysBeforeStartToExpireFromWaitlist, ok := BucreatetimeoffplanrequestMap["daysBeforeStartToExpireFromWaitlist"].(float64); ok {
		DaysBeforeStartToExpireFromWaitlistInt := int(DaysBeforeStartToExpireFromWaitlist)
		o.DaysBeforeStartToExpireFromWaitlist = &DaysBeforeStartToExpireFromWaitlistInt
	}
	
	if HrisTimeOffType, ok := BucreatetimeoffplanrequestMap["hrisTimeOffType"].(map[string]interface{}); ok {
		HrisTimeOffTypeString, _ := json.Marshal(HrisTimeOffType)
		json.Unmarshal(HrisTimeOffTypeString, &o.HrisTimeOffType)
	}
	
	if Enabled, ok := BucreatetimeoffplanrequestMap["enabled"].(bool); ok {
		o.Enabled = &Enabled
	}
    
	if CountAgainstTimeOffLimits, ok := BucreatetimeoffplanrequestMap["countAgainstTimeOffLimits"].(bool); ok {
		o.CountAgainstTimeOffLimits = &CountAgainstTimeOffLimits
	}
    
	if BusinessUnitAssociation, ok := BucreatetimeoffplanrequestMap["businessUnitAssociation"].(map[string]interface{}); ok {
		BusinessUnitAssociationString, _ := json.Marshal(BusinessUnitAssociation)
		json.Unmarshal(BusinessUnitAssociationString, &o.BusinessUnitAssociation)
	}
	
	if ManagementUnitAssociation, ok := BucreatetimeoffplanrequestMap["managementUnitAssociation"].(map[string]interface{}); ok {
		ManagementUnitAssociationString, _ := json.Marshal(ManagementUnitAssociation)
		json.Unmarshal(ManagementUnitAssociationString, &o.ManagementUnitAssociation)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Bucreatetimeoffplanrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Timeoffplan
type Timeoffplan struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name - The name of this time off plan.
	Name *string `json:"name,omitempty"`

	// ActivityCodeIds - The set of activity code IDs associated with this time off plan.
	ActivityCodeIds *[]string `json:"activityCodeIds,omitempty"`

	// TimeOffLimits - The set of time off limit IDs associated with this time off plan.
	TimeOffLimits *[]Timeofflimitreference `json:"timeOffLimits,omitempty"`

	// AutoApprovalRule - Auto approval rule for this time off plan
	AutoApprovalRule *string `json:"autoApprovalRule,omitempty"`

	// DaysBeforeStartToExpireFromWaitlist - The number of days before the time off request start date for when the request will be expired from the waitlist.
	DaysBeforeStartToExpireFromWaitlist *int `json:"daysBeforeStartToExpireFromWaitlist,omitempty"`

	// HrisTimeOffType - Time off type, if this time off plan is associated with the integration.
	HrisTimeOffType *Hristimeofftype `json:"hrisTimeOffType,omitempty"`

	// Active - Whether this time off plan is currently being used by agents.
	Active *bool `json:"active,omitempty"`

	// Metadata - Version metadata for the time off plan.
	Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Timeoffplan) SetField(field string, fieldValue interface{}) {
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

func (o Timeoffplan) MarshalJSON() ([]byte, error) {
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
	type Alias Timeoffplan
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		ActivityCodeIds *[]string `json:"activityCodeIds,omitempty"`
		
		TimeOffLimits *[]Timeofflimitreference `json:"timeOffLimits,omitempty"`
		
		AutoApprovalRule *string `json:"autoApprovalRule,omitempty"`
		
		DaysBeforeStartToExpireFromWaitlist *int `json:"daysBeforeStartToExpireFromWaitlist,omitempty"`
		
		HrisTimeOffType *Hristimeofftype `json:"hrisTimeOffType,omitempty"`
		
		Active *bool `json:"active,omitempty"`
		
		Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		ActivityCodeIds: o.ActivityCodeIds,
		
		TimeOffLimits: o.TimeOffLimits,
		
		AutoApprovalRule: o.AutoApprovalRule,
		
		DaysBeforeStartToExpireFromWaitlist: o.DaysBeforeStartToExpireFromWaitlist,
		
		HrisTimeOffType: o.HrisTimeOffType,
		
		Active: o.Active,
		
		Metadata: o.Metadata,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Timeoffplan) UnmarshalJSON(b []byte) error {
	var TimeoffplanMap map[string]interface{}
	err := json.Unmarshal(b, &TimeoffplanMap)
	if err != nil {
		return err
	}
	
	if Id, ok := TimeoffplanMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := TimeoffplanMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if ActivityCodeIds, ok := TimeoffplanMap["activityCodeIds"].([]interface{}); ok {
		ActivityCodeIdsString, _ := json.Marshal(ActivityCodeIds)
		json.Unmarshal(ActivityCodeIdsString, &o.ActivityCodeIds)
	}
	
	if TimeOffLimits, ok := TimeoffplanMap["timeOffLimits"].([]interface{}); ok {
		TimeOffLimitsString, _ := json.Marshal(TimeOffLimits)
		json.Unmarshal(TimeOffLimitsString, &o.TimeOffLimits)
	}
	
	if AutoApprovalRule, ok := TimeoffplanMap["autoApprovalRule"].(string); ok {
		o.AutoApprovalRule = &AutoApprovalRule
	}
    
	if DaysBeforeStartToExpireFromWaitlist, ok := TimeoffplanMap["daysBeforeStartToExpireFromWaitlist"].(float64); ok {
		DaysBeforeStartToExpireFromWaitlistInt := int(DaysBeforeStartToExpireFromWaitlist)
		o.DaysBeforeStartToExpireFromWaitlist = &DaysBeforeStartToExpireFromWaitlistInt
	}
	
	if HrisTimeOffType, ok := TimeoffplanMap["hrisTimeOffType"].(map[string]interface{}); ok {
		HrisTimeOffTypeString, _ := json.Marshal(HrisTimeOffType)
		json.Unmarshal(HrisTimeOffTypeString, &o.HrisTimeOffType)
	}
	
	if Active, ok := TimeoffplanMap["active"].(bool); ok {
		o.Active = &Active
	}
    
	if Metadata, ok := TimeoffplanMap["metadata"].(map[string]interface{}); ok {
		MetadataString, _ := json.Marshal(Metadata)
		json.Unmarshal(MetadataString, &o.Metadata)
	}
	
	if SelfUri, ok := TimeoffplanMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Timeoffplan) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

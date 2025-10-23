package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Buupdatetimeoffplanrequest
type Buupdatetimeoffplanrequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Name - The name of this time-off plan
	Name *string `json:"name,omitempty"`

	// ActivityCodeIds - The IDs of activity codes to associate with this time-off plan
	ActivityCodeIds *Setwrapperstring `json:"activityCodeIds,omitempty"`

	// AutoApprovalRule - Auto approval rule for this time-off plan
	AutoApprovalRule *string `json:"autoApprovalRule,omitempty"`

	// DaysBeforeStartToExpireFromWaitlist - The number of days before the time-off request start date for when the request will be expired from the waitlist
	DaysBeforeStartToExpireFromWaitlist *int `json:"daysBeforeStartToExpireFromWaitlist,omitempty"`

	// AutoPublishApprovedTimeOffRequests - Whether newly approved time-off requests with activity codes associated with this time-off plan should be automatically published to the schedule
	AutoPublishApprovedTimeOffRequests *bool `json:"autoPublishApprovedTimeOffRequests,omitempty"`

	// RestrictedActivityCodeIds - The IDs of non time-off activity codes to check for conflicts in case the auto approval rule specifies checking activity codes. If these activity codes are present in schedule and overlap with the time-off request duration, the request will not be auto approved
	RestrictedActivityCodeIds *Setwrapperstring `json:"restrictedActivityCodeIds,omitempty"`

	// HrisTimeOffType - Time-off type, if this time-off plan is associated with the integration
	HrisTimeOffType *Valuewrapperhristimeofftype `json:"hrisTimeOffType,omitempty"`

	// Enabled - Whether this time-off plan should be used by agents
	Enabled *bool `json:"enabled,omitempty"`

	// CountAgainstTimeOffLimits - Whether this time-off plan should count against time-off limits
	CountAgainstTimeOffLimits *bool `json:"countAgainstTimeOffLimits,omitempty"`

	// BusinessUnitAssociation - Business unit association, if the time-off plan belongs to a business unit. managementUnitAssociation must not be set if this is populated
	BusinessUnitAssociation *Updatetimeoffplanbusinessunitassociation `json:"businessUnitAssociation,omitempty"`

	// ManagementUnitAssociation - Management unit association, if the time-off plan belongs to a management unit. businessUnitAssociation must not be set if this is populated
	ManagementUnitAssociation *Updatetimeoffplanmanagementunitassociation `json:"managementUnitAssociation,omitempty"`

	// Metadata - Version metadata for this time-off plan
	Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Buupdatetimeoffplanrequest) SetField(field string, fieldValue interface{}) {
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

func (o Buupdatetimeoffplanrequest) MarshalJSON() ([]byte, error) {
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
	type Alias Buupdatetimeoffplanrequest
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		ActivityCodeIds *Setwrapperstring `json:"activityCodeIds,omitempty"`
		
		AutoApprovalRule *string `json:"autoApprovalRule,omitempty"`
		
		DaysBeforeStartToExpireFromWaitlist *int `json:"daysBeforeStartToExpireFromWaitlist,omitempty"`
		
		AutoPublishApprovedTimeOffRequests *bool `json:"autoPublishApprovedTimeOffRequests,omitempty"`
		
		RestrictedActivityCodeIds *Setwrapperstring `json:"restrictedActivityCodeIds,omitempty"`
		
		HrisTimeOffType *Valuewrapperhristimeofftype `json:"hrisTimeOffType,omitempty"`
		
		Enabled *bool `json:"enabled,omitempty"`
		
		CountAgainstTimeOffLimits *bool `json:"countAgainstTimeOffLimits,omitempty"`
		
		BusinessUnitAssociation *Updatetimeoffplanbusinessunitassociation `json:"businessUnitAssociation,omitempty"`
		
		ManagementUnitAssociation *Updatetimeoffplanmanagementunitassociation `json:"managementUnitAssociation,omitempty"`
		
		Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`
		Alias
	}{ 
		Name: o.Name,
		
		ActivityCodeIds: o.ActivityCodeIds,
		
		AutoApprovalRule: o.AutoApprovalRule,
		
		DaysBeforeStartToExpireFromWaitlist: o.DaysBeforeStartToExpireFromWaitlist,
		
		AutoPublishApprovedTimeOffRequests: o.AutoPublishApprovedTimeOffRequests,
		
		RestrictedActivityCodeIds: o.RestrictedActivityCodeIds,
		
		HrisTimeOffType: o.HrisTimeOffType,
		
		Enabled: o.Enabled,
		
		CountAgainstTimeOffLimits: o.CountAgainstTimeOffLimits,
		
		BusinessUnitAssociation: o.BusinessUnitAssociation,
		
		ManagementUnitAssociation: o.ManagementUnitAssociation,
		
		Metadata: o.Metadata,
		Alias:    (Alias)(o),
	})
}

func (o *Buupdatetimeoffplanrequest) UnmarshalJSON(b []byte) error {
	var BuupdatetimeoffplanrequestMap map[string]interface{}
	err := json.Unmarshal(b, &BuupdatetimeoffplanrequestMap)
	if err != nil {
		return err
	}
	
	if Name, ok := BuupdatetimeoffplanrequestMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if ActivityCodeIds, ok := BuupdatetimeoffplanrequestMap["activityCodeIds"].(map[string]interface{}); ok {
		ActivityCodeIdsString, _ := json.Marshal(ActivityCodeIds)
		json.Unmarshal(ActivityCodeIdsString, &o.ActivityCodeIds)
	}
	
	if AutoApprovalRule, ok := BuupdatetimeoffplanrequestMap["autoApprovalRule"].(string); ok {
		o.AutoApprovalRule = &AutoApprovalRule
	}
    
	if DaysBeforeStartToExpireFromWaitlist, ok := BuupdatetimeoffplanrequestMap["daysBeforeStartToExpireFromWaitlist"].(float64); ok {
		DaysBeforeStartToExpireFromWaitlistInt := int(DaysBeforeStartToExpireFromWaitlist)
		o.DaysBeforeStartToExpireFromWaitlist = &DaysBeforeStartToExpireFromWaitlistInt
	}
	
	if AutoPublishApprovedTimeOffRequests, ok := BuupdatetimeoffplanrequestMap["autoPublishApprovedTimeOffRequests"].(bool); ok {
		o.AutoPublishApprovedTimeOffRequests = &AutoPublishApprovedTimeOffRequests
	}
    
	if RestrictedActivityCodeIds, ok := BuupdatetimeoffplanrequestMap["restrictedActivityCodeIds"].(map[string]interface{}); ok {
		RestrictedActivityCodeIdsString, _ := json.Marshal(RestrictedActivityCodeIds)
		json.Unmarshal(RestrictedActivityCodeIdsString, &o.RestrictedActivityCodeIds)
	}
	
	if HrisTimeOffType, ok := BuupdatetimeoffplanrequestMap["hrisTimeOffType"].(map[string]interface{}); ok {
		HrisTimeOffTypeString, _ := json.Marshal(HrisTimeOffType)
		json.Unmarshal(HrisTimeOffTypeString, &o.HrisTimeOffType)
	}
	
	if Enabled, ok := BuupdatetimeoffplanrequestMap["enabled"].(bool); ok {
		o.Enabled = &Enabled
	}
    
	if CountAgainstTimeOffLimits, ok := BuupdatetimeoffplanrequestMap["countAgainstTimeOffLimits"].(bool); ok {
		o.CountAgainstTimeOffLimits = &CountAgainstTimeOffLimits
	}
    
	if BusinessUnitAssociation, ok := BuupdatetimeoffplanrequestMap["businessUnitAssociation"].(map[string]interface{}); ok {
		BusinessUnitAssociationString, _ := json.Marshal(BusinessUnitAssociation)
		json.Unmarshal(BusinessUnitAssociationString, &o.BusinessUnitAssociation)
	}
	
	if ManagementUnitAssociation, ok := BuupdatetimeoffplanrequestMap["managementUnitAssociation"].(map[string]interface{}); ok {
		ManagementUnitAssociationString, _ := json.Marshal(ManagementUnitAssociation)
		json.Unmarshal(ManagementUnitAssociationString, &o.ManagementUnitAssociation)
	}
	
	if Metadata, ok := BuupdatetimeoffplanrequestMap["metadata"].(map[string]interface{}); ok {
		MetadataString, _ := json.Marshal(Metadata)
		json.Unmarshal(MetadataString, &o.Metadata)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Buupdatetimeoffplanrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

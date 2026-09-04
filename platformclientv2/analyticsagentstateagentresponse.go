package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Analyticsagentstateagentresponse
type Analyticsagentstateagentresponse struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// UserId - User Id - only returned if division is covered by agentStateNames permission
	UserId *string `json:"userId,omitempty"`

	// DivisionId - Division Id
	DivisionId *string `json:"divisionId,omitempty"`

	// UserName - User name - only returned if division is covered by agentStateNames permission
	UserName *string `json:"userName,omitempty"`

	// ManagerId - The user that this user reports to
	ManagerId *string `json:"managerId,omitempty"`

	// SessionCount - The count of sessions
	SessionCount *int `json:"sessionCount,omitempty"`

	// Sessions - List of sessions
	Sessions *[]Analyticsagentstateagentsessionresult `json:"sessions,omitempty"`

	// SystemPresence - The user's system presence
	SystemPresence *string `json:"systemPresence,omitempty"`

	// OrganizationPresenceId - The identifier for the user's organization presence
	OrganizationPresenceId *string `json:"organizationPresenceId,omitempty"`

	// PresenceDate - The timestamp for when the user's presence began. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	PresenceDate *time.Time `json:"presenceDate,omitempty"`

	// RoutingStatus - The user's routing status
	RoutingStatus *string `json:"routingStatus,omitempty"`

	// RoutingStatusDate - The timestamp for when the user's routing status began. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	RoutingStatusDate *time.Time `json:"routingStatusDate,omitempty"`

	// IsOutOfOffice - Whether the user is out of office
	IsOutOfOffice *bool `json:"isOutOfOffice,omitempty"`

	// ManagementUnitId - The id of the user's management unit
	ManagementUnitId *string `json:"managementUnitId,omitempty"`

	// BusinessUnitId - The id of the user's business unit
	BusinessUnitId *string `json:"businessUnitId,omitempty"`

	// AdherenceState - The user's adherence state
	AdherenceState *string `json:"adherenceState,omitempty"`

	// AdherenceImpact - The user's adherence impact
	AdherenceImpact *string `json:"adherenceImpact,omitempty"`

	// AdherenceDate - The timestamp for when the user's adherence state began. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	AdherenceDate *time.Time `json:"adherenceDate,omitempty"`

	// ScheduledActivityCodeId - The id of the user's scheduled activity code
	ScheduledActivityCodeId *string `json:"scheduledActivityCodeId,omitempty"`

	// ScheduledActivityCategory - The user's scheduled activity category
	ScheduledActivityCategory *string `json:"scheduledActivityCategory,omitempty"`

	// ActualActivityCategory - The user's actual activity category
	ActualActivityCategory *string `json:"actualActivityCategory,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Analyticsagentstateagentresponse) SetField(field string, fieldValue interface{}) {
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

func (o Analyticsagentstateagentresponse) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "PresenceDate","RoutingStatusDate","AdherenceDate", }
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
	type Alias Analyticsagentstateagentresponse
	
	PresenceDate := new(string)
	if o.PresenceDate != nil {
		
		*PresenceDate = timeutil.Strftime(o.PresenceDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		PresenceDate = nil
	}
	
	RoutingStatusDate := new(string)
	if o.RoutingStatusDate != nil {
		
		*RoutingStatusDate = timeutil.Strftime(o.RoutingStatusDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		RoutingStatusDate = nil
	}
	
	AdherenceDate := new(string)
	if o.AdherenceDate != nil {
		
		*AdherenceDate = timeutil.Strftime(o.AdherenceDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		AdherenceDate = nil
	}
	
	return json.Marshal(&struct { 
		UserId *string `json:"userId,omitempty"`
		
		DivisionId *string `json:"divisionId,omitempty"`
		
		UserName *string `json:"userName,omitempty"`
		
		ManagerId *string `json:"managerId,omitempty"`
		
		SessionCount *int `json:"sessionCount,omitempty"`
		
		Sessions *[]Analyticsagentstateagentsessionresult `json:"sessions,omitempty"`
		
		SystemPresence *string `json:"systemPresence,omitempty"`
		
		OrganizationPresenceId *string `json:"organizationPresenceId,omitempty"`
		
		PresenceDate *string `json:"presenceDate,omitempty"`
		
		RoutingStatus *string `json:"routingStatus,omitempty"`
		
		RoutingStatusDate *string `json:"routingStatusDate,omitempty"`
		
		IsOutOfOffice *bool `json:"isOutOfOffice,omitempty"`
		
		ManagementUnitId *string `json:"managementUnitId,omitempty"`
		
		BusinessUnitId *string `json:"businessUnitId,omitempty"`
		
		AdherenceState *string `json:"adherenceState,omitempty"`
		
		AdherenceImpact *string `json:"adherenceImpact,omitempty"`
		
		AdherenceDate *string `json:"adherenceDate,omitempty"`
		
		ScheduledActivityCodeId *string `json:"scheduledActivityCodeId,omitempty"`
		
		ScheduledActivityCategory *string `json:"scheduledActivityCategory,omitempty"`
		
		ActualActivityCategory *string `json:"actualActivityCategory,omitempty"`
		Alias
	}{ 
		UserId: o.UserId,
		
		DivisionId: o.DivisionId,
		
		UserName: o.UserName,
		
		ManagerId: o.ManagerId,
		
		SessionCount: o.SessionCount,
		
		Sessions: o.Sessions,
		
		SystemPresence: o.SystemPresence,
		
		OrganizationPresenceId: o.OrganizationPresenceId,
		
		PresenceDate: PresenceDate,
		
		RoutingStatus: o.RoutingStatus,
		
		RoutingStatusDate: RoutingStatusDate,
		
		IsOutOfOffice: o.IsOutOfOffice,
		
		ManagementUnitId: o.ManagementUnitId,
		
		BusinessUnitId: o.BusinessUnitId,
		
		AdherenceState: o.AdherenceState,
		
		AdherenceImpact: o.AdherenceImpact,
		
		AdherenceDate: AdherenceDate,
		
		ScheduledActivityCodeId: o.ScheduledActivityCodeId,
		
		ScheduledActivityCategory: o.ScheduledActivityCategory,
		
		ActualActivityCategory: o.ActualActivityCategory,
		Alias:    (Alias)(o),
	})
}

func (o *Analyticsagentstateagentresponse) UnmarshalJSON(b []byte) error {
	var AnalyticsagentstateagentresponseMap map[string]interface{}
	err := json.Unmarshal(b, &AnalyticsagentstateagentresponseMap)
	if err != nil {
		return err
	}
	
	if UserId, ok := AnalyticsagentstateagentresponseMap["userId"].(string); ok {
		o.UserId = &UserId
	}
    
	if DivisionId, ok := AnalyticsagentstateagentresponseMap["divisionId"].(string); ok {
		o.DivisionId = &DivisionId
	}
    
	if UserName, ok := AnalyticsagentstateagentresponseMap["userName"].(string); ok {
		o.UserName = &UserName
	}
    
	if ManagerId, ok := AnalyticsagentstateagentresponseMap["managerId"].(string); ok {
		o.ManagerId = &ManagerId
	}
    
	if SessionCount, ok := AnalyticsagentstateagentresponseMap["sessionCount"].(float64); ok {
		SessionCountInt := int(SessionCount)
		o.SessionCount = &SessionCountInt
	}
	
	if Sessions, ok := AnalyticsagentstateagentresponseMap["sessions"].([]interface{}); ok {
		SessionsString, _ := json.Marshal(Sessions)
		json.Unmarshal(SessionsString, &o.Sessions)
	}
	
	if SystemPresence, ok := AnalyticsagentstateagentresponseMap["systemPresence"].(string); ok {
		o.SystemPresence = &SystemPresence
	}
    
	if OrganizationPresenceId, ok := AnalyticsagentstateagentresponseMap["organizationPresenceId"].(string); ok {
		o.OrganizationPresenceId = &OrganizationPresenceId
	}
    
	if presenceDateString, ok := AnalyticsagentstateagentresponseMap["presenceDate"].(string); ok {
		PresenceDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", presenceDateString)
		o.PresenceDate = &PresenceDate
	}
	
	if RoutingStatus, ok := AnalyticsagentstateagentresponseMap["routingStatus"].(string); ok {
		o.RoutingStatus = &RoutingStatus
	}
    
	if routingStatusDateString, ok := AnalyticsagentstateagentresponseMap["routingStatusDate"].(string); ok {
		RoutingStatusDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", routingStatusDateString)
		o.RoutingStatusDate = &RoutingStatusDate
	}
	
	if IsOutOfOffice, ok := AnalyticsagentstateagentresponseMap["isOutOfOffice"].(bool); ok {
		o.IsOutOfOffice = &IsOutOfOffice
	}
    
	if ManagementUnitId, ok := AnalyticsagentstateagentresponseMap["managementUnitId"].(string); ok {
		o.ManagementUnitId = &ManagementUnitId
	}
    
	if BusinessUnitId, ok := AnalyticsagentstateagentresponseMap["businessUnitId"].(string); ok {
		o.BusinessUnitId = &BusinessUnitId
	}
    
	if AdherenceState, ok := AnalyticsagentstateagentresponseMap["adherenceState"].(string); ok {
		o.AdherenceState = &AdherenceState
	}
    
	if AdherenceImpact, ok := AnalyticsagentstateagentresponseMap["adherenceImpact"].(string); ok {
		o.AdherenceImpact = &AdherenceImpact
	}
    
	if adherenceDateString, ok := AnalyticsagentstateagentresponseMap["adherenceDate"].(string); ok {
		AdherenceDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", adherenceDateString)
		o.AdherenceDate = &AdherenceDate
	}
	
	if ScheduledActivityCodeId, ok := AnalyticsagentstateagentresponseMap["scheduledActivityCodeId"].(string); ok {
		o.ScheduledActivityCodeId = &ScheduledActivityCodeId
	}
    
	if ScheduledActivityCategory, ok := AnalyticsagentstateagentresponseMap["scheduledActivityCategory"].(string); ok {
		o.ScheduledActivityCategory = &ScheduledActivityCategory
	}
    
	if ActualActivityCategory, ok := AnalyticsagentstateagentresponseMap["actualActivityCategory"].(string); ok {
		o.ActualActivityCategory = &ActualActivityCategory
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Analyticsagentstateagentresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

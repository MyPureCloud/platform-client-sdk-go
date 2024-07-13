package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Userscheduleadherence
type Userscheduleadherence struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name
	Name *string `json:"name,omitempty"`

	// User - The user for whom this status applies
	User *Userreference `json:"user,omitempty"`

	// BusinessUnit - The business unit to which this user belongs
	BusinessUnit *Businessunitreference `json:"businessUnit,omitempty"`

	// ManagementUnit - The management unit to which this user belongs
	ManagementUnit *Managementunitreference `json:"managementUnit,omitempty"`

	// Team - The team to which this user belongs
	Team *Teamreference `json:"team,omitempty"`

	// ScheduledActivityCategory - Activity for which the user is scheduled
	ScheduledActivityCategory *string `json:"scheduledActivityCategory,omitempty"`

	// ScheduledActivityCode - Activity code for which the user is currently scheduled
	ScheduledActivityCode *Activitycodesummary `json:"scheduledActivityCode,omitempty"`

	// SystemPresence - Actual underlying system presence value
	SystemPresence *string `json:"systemPresence,omitempty"`

	// OrganizationSecondaryPresenceId - Organization Secondary Presence Id.
	OrganizationSecondaryPresenceId *string `json:"organizationSecondaryPresenceId,omitempty"`

	// RoutingStatus - Actual underlying routing status, used to determine whether a user is actually in adherence when OnQueue
	RoutingStatus *string `json:"routingStatus,omitempty"`

	// ActualActivityCategory - Activity in which the user is actually engaged
	ActualActivityCategory *string `json:"actualActivityCategory,omitempty"`

	// IsOutOfOffice - Whether the user is marked OutOfOffice
	IsOutOfOffice *bool `json:"isOutOfOffice,omitempty"`

	// AdherenceState - The user's current adherence state
	AdherenceState *string `json:"adherenceState,omitempty"`

	// Impact - The impact of the user's current adherenceState
	Impact *string `json:"impact,omitempty"`

	// AdherenceExplanation - Currently applicable explanation for the adherence state
	AdherenceExplanation *Realtimeadherenceexplanation `json:"adherenceExplanation,omitempty"`

	// TimeOfAdherenceChange - Time when the user entered the current adherenceState in ISO-8601 format
	TimeOfAdherenceChange *time.Time `json:"timeOfAdherenceChange,omitempty"`

	// PresenceUpdateTime - Time when presence was last updated. Used to calculate time in current status. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	PresenceUpdateTime *time.Time `json:"presenceUpdateTime,omitempty"`

	// ActiveQueues - The list of queues to which this user is joined
	ActiveQueues *[]Queuereference `json:"activeQueues,omitempty"`

	// ActiveQueuesModifiedTime - Time when the list of active queues for this user was last updated. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ActiveQueuesModifiedTime *time.Time `json:"activeQueuesModifiedTime,omitempty"`

	// RemovedFromManagementUnit - For notification purposes. Used to indicate that a user was removed from the management unit
	RemovedFromManagementUnit *bool `json:"removedFromManagementUnit,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Userscheduleadherence) SetField(field string, fieldValue interface{}) {
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

func (o Userscheduleadherence) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "TimeOfAdherenceChange","PresenceUpdateTime","ActiveQueuesModifiedTime", }
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
	type Alias Userscheduleadherence
	
	TimeOfAdherenceChange := new(string)
	if o.TimeOfAdherenceChange != nil {
		
		*TimeOfAdherenceChange = timeutil.Strftime(o.TimeOfAdherenceChange, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		TimeOfAdherenceChange = nil
	}
	
	PresenceUpdateTime := new(string)
	if o.PresenceUpdateTime != nil {
		
		*PresenceUpdateTime = timeutil.Strftime(o.PresenceUpdateTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		PresenceUpdateTime = nil
	}
	
	ActiveQueuesModifiedTime := new(string)
	if o.ActiveQueuesModifiedTime != nil {
		
		*ActiveQueuesModifiedTime = timeutil.Strftime(o.ActiveQueuesModifiedTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ActiveQueuesModifiedTime = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		User *Userreference `json:"user,omitempty"`
		
		BusinessUnit *Businessunitreference `json:"businessUnit,omitempty"`
		
		ManagementUnit *Managementunitreference `json:"managementUnit,omitempty"`
		
		Team *Teamreference `json:"team,omitempty"`
		
		ScheduledActivityCategory *string `json:"scheduledActivityCategory,omitempty"`
		
		ScheduledActivityCode *Activitycodesummary `json:"scheduledActivityCode,omitempty"`
		
		SystemPresence *string `json:"systemPresence,omitempty"`
		
		OrganizationSecondaryPresenceId *string `json:"organizationSecondaryPresenceId,omitempty"`
		
		RoutingStatus *string `json:"routingStatus,omitempty"`
		
		ActualActivityCategory *string `json:"actualActivityCategory,omitempty"`
		
		IsOutOfOffice *bool `json:"isOutOfOffice,omitempty"`
		
		AdherenceState *string `json:"adherenceState,omitempty"`
		
		Impact *string `json:"impact,omitempty"`
		
		AdherenceExplanation *Realtimeadherenceexplanation `json:"adherenceExplanation,omitempty"`
		
		TimeOfAdherenceChange *string `json:"timeOfAdherenceChange,omitempty"`
		
		PresenceUpdateTime *string `json:"presenceUpdateTime,omitempty"`
		
		ActiveQueues *[]Queuereference `json:"activeQueues,omitempty"`
		
		ActiveQueuesModifiedTime *string `json:"activeQueuesModifiedTime,omitempty"`
		
		RemovedFromManagementUnit *bool `json:"removedFromManagementUnit,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		User: o.User,
		
		BusinessUnit: o.BusinessUnit,
		
		ManagementUnit: o.ManagementUnit,
		
		Team: o.Team,
		
		ScheduledActivityCategory: o.ScheduledActivityCategory,
		
		ScheduledActivityCode: o.ScheduledActivityCode,
		
		SystemPresence: o.SystemPresence,
		
		OrganizationSecondaryPresenceId: o.OrganizationSecondaryPresenceId,
		
		RoutingStatus: o.RoutingStatus,
		
		ActualActivityCategory: o.ActualActivityCategory,
		
		IsOutOfOffice: o.IsOutOfOffice,
		
		AdherenceState: o.AdherenceState,
		
		Impact: o.Impact,
		
		AdherenceExplanation: o.AdherenceExplanation,
		
		TimeOfAdherenceChange: TimeOfAdherenceChange,
		
		PresenceUpdateTime: PresenceUpdateTime,
		
		ActiveQueues: o.ActiveQueues,
		
		ActiveQueuesModifiedTime: ActiveQueuesModifiedTime,
		
		RemovedFromManagementUnit: o.RemovedFromManagementUnit,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Userscheduleadherence) UnmarshalJSON(b []byte) error {
	var UserscheduleadherenceMap map[string]interface{}
	err := json.Unmarshal(b, &UserscheduleadherenceMap)
	if err != nil {
		return err
	}
	
	if Id, ok := UserscheduleadherenceMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := UserscheduleadherenceMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if User, ok := UserscheduleadherenceMap["user"].(map[string]interface{}); ok {
		UserString, _ := json.Marshal(User)
		json.Unmarshal(UserString, &o.User)
	}
	
	if BusinessUnit, ok := UserscheduleadherenceMap["businessUnit"].(map[string]interface{}); ok {
		BusinessUnitString, _ := json.Marshal(BusinessUnit)
		json.Unmarshal(BusinessUnitString, &o.BusinessUnit)
	}
	
	if ManagementUnit, ok := UserscheduleadherenceMap["managementUnit"].(map[string]interface{}); ok {
		ManagementUnitString, _ := json.Marshal(ManagementUnit)
		json.Unmarshal(ManagementUnitString, &o.ManagementUnit)
	}
	
	if Team, ok := UserscheduleadherenceMap["team"].(map[string]interface{}); ok {
		TeamString, _ := json.Marshal(Team)
		json.Unmarshal(TeamString, &o.Team)
	}
	
	if ScheduledActivityCategory, ok := UserscheduleadherenceMap["scheduledActivityCategory"].(string); ok {
		o.ScheduledActivityCategory = &ScheduledActivityCategory
	}
    
	if ScheduledActivityCode, ok := UserscheduleadherenceMap["scheduledActivityCode"].(map[string]interface{}); ok {
		ScheduledActivityCodeString, _ := json.Marshal(ScheduledActivityCode)
		json.Unmarshal(ScheduledActivityCodeString, &o.ScheduledActivityCode)
	}
	
	if SystemPresence, ok := UserscheduleadherenceMap["systemPresence"].(string); ok {
		o.SystemPresence = &SystemPresence
	}
    
	if OrganizationSecondaryPresenceId, ok := UserscheduleadherenceMap["organizationSecondaryPresenceId"].(string); ok {
		o.OrganizationSecondaryPresenceId = &OrganizationSecondaryPresenceId
	}
    
	if RoutingStatus, ok := UserscheduleadherenceMap["routingStatus"].(string); ok {
		o.RoutingStatus = &RoutingStatus
	}
    
	if ActualActivityCategory, ok := UserscheduleadherenceMap["actualActivityCategory"].(string); ok {
		o.ActualActivityCategory = &ActualActivityCategory
	}
    
	if IsOutOfOffice, ok := UserscheduleadherenceMap["isOutOfOffice"].(bool); ok {
		o.IsOutOfOffice = &IsOutOfOffice
	}
    
	if AdherenceState, ok := UserscheduleadherenceMap["adherenceState"].(string); ok {
		o.AdherenceState = &AdherenceState
	}
    
	if Impact, ok := UserscheduleadherenceMap["impact"].(string); ok {
		o.Impact = &Impact
	}
    
	if AdherenceExplanation, ok := UserscheduleadherenceMap["adherenceExplanation"].(map[string]interface{}); ok {
		AdherenceExplanationString, _ := json.Marshal(AdherenceExplanation)
		json.Unmarshal(AdherenceExplanationString, &o.AdherenceExplanation)
	}
	
	if timeOfAdherenceChangeString, ok := UserscheduleadherenceMap["timeOfAdherenceChange"].(string); ok {
		TimeOfAdherenceChange, _ := time.Parse("2006-01-02T15:04:05.999999Z", timeOfAdherenceChangeString)
		o.TimeOfAdherenceChange = &TimeOfAdherenceChange
	}
	
	if presenceUpdateTimeString, ok := UserscheduleadherenceMap["presenceUpdateTime"].(string); ok {
		PresenceUpdateTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", presenceUpdateTimeString)
		o.PresenceUpdateTime = &PresenceUpdateTime
	}
	
	if ActiveQueues, ok := UserscheduleadherenceMap["activeQueues"].([]interface{}); ok {
		ActiveQueuesString, _ := json.Marshal(ActiveQueues)
		json.Unmarshal(ActiveQueuesString, &o.ActiveQueues)
	}
	
	if activeQueuesModifiedTimeString, ok := UserscheduleadherenceMap["activeQueuesModifiedTime"].(string); ok {
		ActiveQueuesModifiedTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", activeQueuesModifiedTimeString)
		o.ActiveQueuesModifiedTime = &ActiveQueuesModifiedTime
	}
	
	if RemovedFromManagementUnit, ok := UserscheduleadherenceMap["removedFromManagementUnit"].(bool); ok {
		o.RemovedFromManagementUnit = &RemovedFromManagementUnit
	}
    
	if SelfUri, ok := UserscheduleadherenceMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Userscheduleadherence) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

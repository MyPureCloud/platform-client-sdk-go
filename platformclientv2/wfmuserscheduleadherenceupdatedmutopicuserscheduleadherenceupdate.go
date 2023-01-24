package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmuserscheduleadherenceupdatedmutopicuserscheduleadherenceupdate
type Wfmuserscheduleadherenceupdatedmutopicuserscheduleadherenceupdate struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// User
	User *Wfmuserscheduleadherenceupdatedmutopicuserreference `json:"user,omitempty"`

	// ManagementUnitId
	ManagementUnitId *string `json:"managementUnitId,omitempty"`

	// Team
	Team *Wfmuserscheduleadherenceupdatedmutopicurireference `json:"team,omitempty"`

	// ScheduledActivityCategory
	ScheduledActivityCategory *string `json:"scheduledActivityCategory,omitempty"`

	// ScheduledActivityCode
	ScheduledActivityCode *Wfmuserscheduleadherenceupdatedmutopicactivitycodereference `json:"scheduledActivityCode,omitempty"`

	// SystemPresence
	SystemPresence *string `json:"systemPresence,omitempty"`

	// OrganizationSecondaryPresenceId
	OrganizationSecondaryPresenceId *string `json:"organizationSecondaryPresenceId,omitempty"`

	// RoutingStatus
	RoutingStatus *string `json:"routingStatus,omitempty"`

	// ActualActivityCategory
	ActualActivityCategory *string `json:"actualActivityCategory,omitempty"`

	// IsOutOfOffice
	IsOutOfOffice *bool `json:"isOutOfOffice,omitempty"`

	// AdherenceState
	AdherenceState *string `json:"adherenceState,omitempty"`

	// Impact
	Impact *string `json:"impact,omitempty"`

	// AdherenceExplanation
	AdherenceExplanation *Wfmuserscheduleadherenceupdatedmutopicrealtimeadherenceexplanation `json:"adherenceExplanation,omitempty"`

	// AdherenceChangeTime
	AdherenceChangeTime *time.Time `json:"adherenceChangeTime,omitempty"`

	// PresenceUpdateTime
	PresenceUpdateTime *time.Time `json:"presenceUpdateTime,omitempty"`

	// ActiveQueues
	ActiveQueues *[]Wfmuserscheduleadherenceupdatedmutopicqueuereference `json:"activeQueues,omitempty"`

	// ActiveQueuesModifiedTime
	ActiveQueuesModifiedTime *time.Time `json:"activeQueuesModifiedTime,omitempty"`

	// RemovedFromManagementUnit
	RemovedFromManagementUnit *bool `json:"removedFromManagementUnit,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Wfmuserscheduleadherenceupdatedmutopicuserscheduleadherenceupdate) SetField(field string, fieldValue interface{}) {
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

func (o Wfmuserscheduleadherenceupdatedmutopicuserscheduleadherenceupdate) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "AdherenceChangeTime","PresenceUpdateTime","ActiveQueuesModifiedTime", }
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
	type Alias Wfmuserscheduleadherenceupdatedmutopicuserscheduleadherenceupdate
	
	AdherenceChangeTime := new(string)
	if o.AdherenceChangeTime != nil {
		
		*AdherenceChangeTime = timeutil.Strftime(o.AdherenceChangeTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		AdherenceChangeTime = nil
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
		User *Wfmuserscheduleadherenceupdatedmutopicuserreference `json:"user,omitempty"`
		
		ManagementUnitId *string `json:"managementUnitId,omitempty"`
		
		Team *Wfmuserscheduleadherenceupdatedmutopicurireference `json:"team,omitempty"`
		
		ScheduledActivityCategory *string `json:"scheduledActivityCategory,omitempty"`
		
		ScheduledActivityCode *Wfmuserscheduleadherenceupdatedmutopicactivitycodereference `json:"scheduledActivityCode,omitempty"`
		
		SystemPresence *string `json:"systemPresence,omitempty"`
		
		OrganizationSecondaryPresenceId *string `json:"organizationSecondaryPresenceId,omitempty"`
		
		RoutingStatus *string `json:"routingStatus,omitempty"`
		
		ActualActivityCategory *string `json:"actualActivityCategory,omitempty"`
		
		IsOutOfOffice *bool `json:"isOutOfOffice,omitempty"`
		
		AdherenceState *string `json:"adherenceState,omitempty"`
		
		Impact *string `json:"impact,omitempty"`
		
		AdherenceExplanation *Wfmuserscheduleadherenceupdatedmutopicrealtimeadherenceexplanation `json:"adherenceExplanation,omitempty"`
		
		AdherenceChangeTime *string `json:"adherenceChangeTime,omitempty"`
		
		PresenceUpdateTime *string `json:"presenceUpdateTime,omitempty"`
		
		ActiveQueues *[]Wfmuserscheduleadherenceupdatedmutopicqueuereference `json:"activeQueues,omitempty"`
		
		ActiveQueuesModifiedTime *string `json:"activeQueuesModifiedTime,omitempty"`
		
		RemovedFromManagementUnit *bool `json:"removedFromManagementUnit,omitempty"`
		Alias
	}{ 
		User: o.User,
		
		ManagementUnitId: o.ManagementUnitId,
		
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
		
		AdherenceChangeTime: AdherenceChangeTime,
		
		PresenceUpdateTime: PresenceUpdateTime,
		
		ActiveQueues: o.ActiveQueues,
		
		ActiveQueuesModifiedTime: ActiveQueuesModifiedTime,
		
		RemovedFromManagementUnit: o.RemovedFromManagementUnit,
		Alias:    (Alias)(o),
	})
}

func (o *Wfmuserscheduleadherenceupdatedmutopicuserscheduleadherenceupdate) UnmarshalJSON(b []byte) error {
	var WfmuserscheduleadherenceupdatedmutopicuserscheduleadherenceupdateMap map[string]interface{}
	err := json.Unmarshal(b, &WfmuserscheduleadherenceupdatedmutopicuserscheduleadherenceupdateMap)
	if err != nil {
		return err
	}
	
	if User, ok := WfmuserscheduleadherenceupdatedmutopicuserscheduleadherenceupdateMap["user"].(map[string]interface{}); ok {
		UserString, _ := json.Marshal(User)
		json.Unmarshal(UserString, &o.User)
	}
	
	if ManagementUnitId, ok := WfmuserscheduleadherenceupdatedmutopicuserscheduleadherenceupdateMap["managementUnitId"].(string); ok {
		o.ManagementUnitId = &ManagementUnitId
	}
    
	if Team, ok := WfmuserscheduleadherenceupdatedmutopicuserscheduleadherenceupdateMap["team"].(map[string]interface{}); ok {
		TeamString, _ := json.Marshal(Team)
		json.Unmarshal(TeamString, &o.Team)
	}
	
	if ScheduledActivityCategory, ok := WfmuserscheduleadherenceupdatedmutopicuserscheduleadherenceupdateMap["scheduledActivityCategory"].(string); ok {
		o.ScheduledActivityCategory = &ScheduledActivityCategory
	}
    
	if ScheduledActivityCode, ok := WfmuserscheduleadherenceupdatedmutopicuserscheduleadherenceupdateMap["scheduledActivityCode"].(map[string]interface{}); ok {
		ScheduledActivityCodeString, _ := json.Marshal(ScheduledActivityCode)
		json.Unmarshal(ScheduledActivityCodeString, &o.ScheduledActivityCode)
	}
	
	if SystemPresence, ok := WfmuserscheduleadherenceupdatedmutopicuserscheduleadherenceupdateMap["systemPresence"].(string); ok {
		o.SystemPresence = &SystemPresence
	}
    
	if OrganizationSecondaryPresenceId, ok := WfmuserscheduleadherenceupdatedmutopicuserscheduleadherenceupdateMap["organizationSecondaryPresenceId"].(string); ok {
		o.OrganizationSecondaryPresenceId = &OrganizationSecondaryPresenceId
	}
    
	if RoutingStatus, ok := WfmuserscheduleadherenceupdatedmutopicuserscheduleadherenceupdateMap["routingStatus"].(string); ok {
		o.RoutingStatus = &RoutingStatus
	}
    
	if ActualActivityCategory, ok := WfmuserscheduleadherenceupdatedmutopicuserscheduleadherenceupdateMap["actualActivityCategory"].(string); ok {
		o.ActualActivityCategory = &ActualActivityCategory
	}
    
	if IsOutOfOffice, ok := WfmuserscheduleadherenceupdatedmutopicuserscheduleadherenceupdateMap["isOutOfOffice"].(bool); ok {
		o.IsOutOfOffice = &IsOutOfOffice
	}
    
	if AdherenceState, ok := WfmuserscheduleadherenceupdatedmutopicuserscheduleadherenceupdateMap["adherenceState"].(string); ok {
		o.AdherenceState = &AdherenceState
	}
    
	if Impact, ok := WfmuserscheduleadherenceupdatedmutopicuserscheduleadherenceupdateMap["impact"].(string); ok {
		o.Impact = &Impact
	}
    
	if AdherenceExplanation, ok := WfmuserscheduleadherenceupdatedmutopicuserscheduleadherenceupdateMap["adherenceExplanation"].(map[string]interface{}); ok {
		AdherenceExplanationString, _ := json.Marshal(AdherenceExplanation)
		json.Unmarshal(AdherenceExplanationString, &o.AdherenceExplanation)
	}
	
	if adherenceChangeTimeString, ok := WfmuserscheduleadherenceupdatedmutopicuserscheduleadherenceupdateMap["adherenceChangeTime"].(string); ok {
		AdherenceChangeTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", adherenceChangeTimeString)
		o.AdherenceChangeTime = &AdherenceChangeTime
	}
	
	if presenceUpdateTimeString, ok := WfmuserscheduleadherenceupdatedmutopicuserscheduleadherenceupdateMap["presenceUpdateTime"].(string); ok {
		PresenceUpdateTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", presenceUpdateTimeString)
		o.PresenceUpdateTime = &PresenceUpdateTime
	}
	
	if ActiveQueues, ok := WfmuserscheduleadherenceupdatedmutopicuserscheduleadherenceupdateMap["activeQueues"].([]interface{}); ok {
		ActiveQueuesString, _ := json.Marshal(ActiveQueues)
		json.Unmarshal(ActiveQueuesString, &o.ActiveQueues)
	}
	
	if activeQueuesModifiedTimeString, ok := WfmuserscheduleadherenceupdatedmutopicuserscheduleadherenceupdateMap["activeQueuesModifiedTime"].(string); ok {
		ActiveQueuesModifiedTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", activeQueuesModifiedTimeString)
		o.ActiveQueuesModifiedTime = &ActiveQueuesModifiedTime
	}
	
	if RemovedFromManagementUnit, ok := WfmuserscheduleadherenceupdatedmutopicuserscheduleadherenceupdateMap["removedFromManagementUnit"].(bool); ok {
		o.RemovedFromManagementUnit = &RemovedFromManagementUnit
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmuserscheduleadherenceupdatedmutopicuserscheduleadherenceupdate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

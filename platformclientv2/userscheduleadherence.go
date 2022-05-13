package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Userscheduleadherence
type Userscheduleadherence struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// User - The user for whom this status applies
	User *Userreference `json:"user,omitempty"`


	// BusinessUnit - The business unit to which this user belongs
	BusinessUnit *Businessunit `json:"businessUnit,omitempty"`


	// ManagementUnit - The management unit to which this user belongs
	ManagementUnit *Managementunit `json:"managementUnit,omitempty"`


	// Team - The team to which this user belongs
	Team *Team `json:"team,omitempty"`


	// ScheduledActivityCategory - Activity for which the user is scheduled
	ScheduledActivityCategory *string `json:"scheduledActivityCategory,omitempty"`


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

func (o *Userscheduleadherence) MarshalJSON() ([]byte, error) {
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
		
		BusinessUnit *Businessunit `json:"businessUnit,omitempty"`
		
		ManagementUnit *Managementunit `json:"managementUnit,omitempty"`
		
		Team *Team `json:"team,omitempty"`
		
		ScheduledActivityCategory *string `json:"scheduledActivityCategory,omitempty"`
		
		SystemPresence *string `json:"systemPresence,omitempty"`
		
		OrganizationSecondaryPresenceId *string `json:"organizationSecondaryPresenceId,omitempty"`
		
		RoutingStatus *string `json:"routingStatus,omitempty"`
		
		ActualActivityCategory *string `json:"actualActivityCategory,omitempty"`
		
		IsOutOfOffice *bool `json:"isOutOfOffice,omitempty"`
		
		AdherenceState *string `json:"adherenceState,omitempty"`
		
		Impact *string `json:"impact,omitempty"`
		
		TimeOfAdherenceChange *string `json:"timeOfAdherenceChange,omitempty"`
		
		PresenceUpdateTime *string `json:"presenceUpdateTime,omitempty"`
		
		ActiveQueues *[]Queuereference `json:"activeQueues,omitempty"`
		
		ActiveQueuesModifiedTime *string `json:"activeQueuesModifiedTime,omitempty"`
		
		RemovedFromManagementUnit *bool `json:"removedFromManagementUnit,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		User: o.User,
		
		BusinessUnit: o.BusinessUnit,
		
		ManagementUnit: o.ManagementUnit,
		
		Team: o.Team,
		
		ScheduledActivityCategory: o.ScheduledActivityCategory,
		
		SystemPresence: o.SystemPresence,
		
		OrganizationSecondaryPresenceId: o.OrganizationSecondaryPresenceId,
		
		RoutingStatus: o.RoutingStatus,
		
		ActualActivityCategory: o.ActualActivityCategory,
		
		IsOutOfOffice: o.IsOutOfOffice,
		
		AdherenceState: o.AdherenceState,
		
		Impact: o.Impact,
		
		TimeOfAdherenceChange: TimeOfAdherenceChange,
		
		PresenceUpdateTime: PresenceUpdateTime,
		
		ActiveQueues: o.ActiveQueues,
		
		ActiveQueuesModifiedTime: ActiveQueuesModifiedTime,
		
		RemovedFromManagementUnit: o.RemovedFromManagementUnit,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
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

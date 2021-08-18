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


	// PresenceUpdateTime - Time when presence was last updated.  Used to calculate time in current status. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
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

func (u *Userscheduleadherence) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Userscheduleadherence

	
	TimeOfAdherenceChange := new(string)
	if u.TimeOfAdherenceChange != nil {
		
		*TimeOfAdherenceChange = timeutil.Strftime(u.TimeOfAdherenceChange, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		TimeOfAdherenceChange = nil
	}
	
	PresenceUpdateTime := new(string)
	if u.PresenceUpdateTime != nil {
		
		*PresenceUpdateTime = timeutil.Strftime(u.PresenceUpdateTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		PresenceUpdateTime = nil
	}
	
	ActiveQueuesModifiedTime := new(string)
	if u.ActiveQueuesModifiedTime != nil {
		
		*ActiveQueuesModifiedTime = timeutil.Strftime(u.ActiveQueuesModifiedTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ActiveQueuesModifiedTime = nil
	}
	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		User *Userreference `json:"user,omitempty"`
		
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
		Id: u.Id,
		
		Name: u.Name,
		
		User: u.User,
		
		ManagementUnit: u.ManagementUnit,
		
		Team: u.Team,
		
		ScheduledActivityCategory: u.ScheduledActivityCategory,
		
		SystemPresence: u.SystemPresence,
		
		OrganizationSecondaryPresenceId: u.OrganizationSecondaryPresenceId,
		
		RoutingStatus: u.RoutingStatus,
		
		ActualActivityCategory: u.ActualActivityCategory,
		
		IsOutOfOffice: u.IsOutOfOffice,
		
		AdherenceState: u.AdherenceState,
		
		Impact: u.Impact,
		
		TimeOfAdherenceChange: TimeOfAdherenceChange,
		
		PresenceUpdateTime: PresenceUpdateTime,
		
		ActiveQueues: u.ActiveQueues,
		
		ActiveQueuesModifiedTime: ActiveQueuesModifiedTime,
		
		RemovedFromManagementUnit: u.RemovedFromManagementUnit,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Userscheduleadherence) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

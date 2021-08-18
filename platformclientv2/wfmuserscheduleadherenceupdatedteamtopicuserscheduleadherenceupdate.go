package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmuserscheduleadherenceupdatedteamtopicuserscheduleadherenceupdate
type Wfmuserscheduleadherenceupdatedteamtopicuserscheduleadherenceupdate struct { 
	// User
	User *Wfmuserscheduleadherenceupdatedteamtopicuserreference `json:"user,omitempty"`


	// ManagementUnitId
	ManagementUnitId *string `json:"managementUnitId,omitempty"`


	// Team
	Team *Wfmuserscheduleadherenceupdatedteamtopicurireference `json:"team,omitempty"`


	// ScheduledActivityCategory
	ScheduledActivityCategory *string `json:"scheduledActivityCategory,omitempty"`


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


	// AdherenceChangeTime
	AdherenceChangeTime *time.Time `json:"adherenceChangeTime,omitempty"`


	// PresenceUpdateTime
	PresenceUpdateTime *time.Time `json:"presenceUpdateTime,omitempty"`


	// ActiveQueues
	ActiveQueues *[]Wfmuserscheduleadherenceupdatedteamtopicqueuereference `json:"activeQueues,omitempty"`


	// ActiveQueuesModifiedTime
	ActiveQueuesModifiedTime *time.Time `json:"activeQueuesModifiedTime,omitempty"`


	// RemovedFromManagementUnit
	RemovedFromManagementUnit *bool `json:"removedFromManagementUnit,omitempty"`

}

func (u *Wfmuserscheduleadherenceupdatedteamtopicuserscheduleadherenceupdate) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmuserscheduleadherenceupdatedteamtopicuserscheduleadherenceupdate

	
	AdherenceChangeTime := new(string)
	if u.AdherenceChangeTime != nil {
		
		*AdherenceChangeTime = timeutil.Strftime(u.AdherenceChangeTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		AdherenceChangeTime = nil
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
		User *Wfmuserscheduleadherenceupdatedteamtopicuserreference `json:"user,omitempty"`
		
		ManagementUnitId *string `json:"managementUnitId,omitempty"`
		
		Team *Wfmuserscheduleadherenceupdatedteamtopicurireference `json:"team,omitempty"`
		
		ScheduledActivityCategory *string `json:"scheduledActivityCategory,omitempty"`
		
		SystemPresence *string `json:"systemPresence,omitempty"`
		
		OrganizationSecondaryPresenceId *string `json:"organizationSecondaryPresenceId,omitempty"`
		
		RoutingStatus *string `json:"routingStatus,omitempty"`
		
		ActualActivityCategory *string `json:"actualActivityCategory,omitempty"`
		
		IsOutOfOffice *bool `json:"isOutOfOffice,omitempty"`
		
		AdherenceState *string `json:"adherenceState,omitempty"`
		
		Impact *string `json:"impact,omitempty"`
		
		AdherenceChangeTime *string `json:"adherenceChangeTime,omitempty"`
		
		PresenceUpdateTime *string `json:"presenceUpdateTime,omitempty"`
		
		ActiveQueues *[]Wfmuserscheduleadherenceupdatedteamtopicqueuereference `json:"activeQueues,omitempty"`
		
		ActiveQueuesModifiedTime *string `json:"activeQueuesModifiedTime,omitempty"`
		
		RemovedFromManagementUnit *bool `json:"removedFromManagementUnit,omitempty"`
		*Alias
	}{ 
		User: u.User,
		
		ManagementUnitId: u.ManagementUnitId,
		
		Team: u.Team,
		
		ScheduledActivityCategory: u.ScheduledActivityCategory,
		
		SystemPresence: u.SystemPresence,
		
		OrganizationSecondaryPresenceId: u.OrganizationSecondaryPresenceId,
		
		RoutingStatus: u.RoutingStatus,
		
		ActualActivityCategory: u.ActualActivityCategory,
		
		IsOutOfOffice: u.IsOutOfOffice,
		
		AdherenceState: u.AdherenceState,
		
		Impact: u.Impact,
		
		AdherenceChangeTime: AdherenceChangeTime,
		
		PresenceUpdateTime: PresenceUpdateTime,
		
		ActiveQueues: u.ActiveQueues,
		
		ActiveQueuesModifiedTime: ActiveQueuesModifiedTime,
		
		RemovedFromManagementUnit: u.RemovedFromManagementUnit,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Wfmuserscheduleadherenceupdatedteamtopicuserscheduleadherenceupdate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

package platformclientv2
import (
	"time"
	"encoding/json"
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


	// PresenceUpdateTime - Time when presence was last updated.  Used to calculate time in current status. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	PresenceUpdateTime *time.Time `json:"presenceUpdateTime,omitempty"`


	// ActiveQueues - The list of queues to which this user is joined
	ActiveQueues *[]Queuereference `json:"activeQueues,omitempty"`


	// ActiveQueuesModifiedTime - Time when the list of active queues for this user was last updated. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	ActiveQueuesModifiedTime *time.Time `json:"activeQueuesModifiedTime,omitempty"`


	// RemovedFromManagementUnit - For notification purposes. Used to indicate that a user was removed from the management unit
	RemovedFromManagementUnit *bool `json:"removedFromManagementUnit,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Userscheduleadherence) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

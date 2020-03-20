package platformclientv2
import (
	"encoding/json"
)

// Queueutilizationdiagnostic
type Queueutilizationdiagnostic struct { 
	// Queue - Identifier of the queue
	Queue *Domainentityref `json:"queue,omitempty"`


	// UsersInQueue - The number of users joined to the queue
	UsersInQueue *int32 `json:"usersInQueue,omitempty"`


	// ActiveUsersInQueue - The number of users active on the queue
	ActiveUsersInQueue *int32 `json:"activeUsersInQueue,omitempty"`


	// UsersOnQueue - The number of users with a status of on-queue
	UsersOnQueue *int32 `json:"usersOnQueue,omitempty"`


	// UsersNotUtilized - The number of users in the queue currently not engaged
	UsersNotUtilized *int32 `json:"usersNotUtilized,omitempty"`


	// UsersOnQueueWithStation - The number of users in the queue with a station
	UsersOnQueueWithStation *int32 `json:"usersOnQueueWithStation,omitempty"`


	// UsersOnACampaignCall - The number of users currently engaged in a campaign call
	UsersOnACampaignCall *int32 `json:"usersOnACampaignCall,omitempty"`


	// UsersOnDifferentEdgeGroup - The number of users whose station is homed to an edge different from the campaign
	UsersOnDifferentEdgeGroup *int32 `json:"usersOnDifferentEdgeGroup,omitempty"`


	// UsersOnANonCampaignCall - The number of users currently engaged in a communication that is not part of the campaign
	UsersOnANonCampaignCall *int32 `json:"usersOnANonCampaignCall,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queueutilizationdiagnostic) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Queueutilizationdiagnostic
type Queueutilizationdiagnostic struct { 
	// Queue - Identifier of the queue
	Queue *Domainentityref `json:"queue,omitempty"`


	// UsersInQueue - The number of users joined to the queue
	UsersInQueue *int `json:"usersInQueue,omitempty"`


	// ActiveUsersInQueue - The number of users active on the queue
	ActiveUsersInQueue *int `json:"activeUsersInQueue,omitempty"`


	// UsersOnQueue - The number of users with a status of on-queue
	UsersOnQueue *int `json:"usersOnQueue,omitempty"`


	// UsersNotUtilized - The number of users in the queue currently not engaged
	UsersNotUtilized *int `json:"usersNotUtilized,omitempty"`


	// UsersOnQueueWithStation - The number of users in the queue with a station
	UsersOnQueueWithStation *int `json:"usersOnQueueWithStation,omitempty"`


	// UsersOnACampaignCall - The number of users currently engaged in a campaign call
	UsersOnACampaignCall *int `json:"usersOnACampaignCall,omitempty"`


	// UsersOnDifferentEdgeGroup - The number of users whose station is homed to an edge different from the campaign
	UsersOnDifferentEdgeGroup *int `json:"usersOnDifferentEdgeGroup,omitempty"`


	// UsersOnANonCampaignCall - The number of users currently engaged in a communication that is not part of the campaign
	UsersOnANonCampaignCall *int `json:"usersOnANonCampaignCall,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queueutilizationdiagnostic) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

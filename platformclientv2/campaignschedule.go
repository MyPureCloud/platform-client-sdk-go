package platformclientv2
import (
	"time"
	"encoding/json"
)

// Campaignschedule
type Campaignschedule struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// DateCreated - Creation time of the entity. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateModified - Last modified time of the entity. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	DateModified *time.Time `json:"dateModified,omitempty"`


	// Version - Required for updates, must match the version number of the most recent update
	Version *int32 `json:"version,omitempty"`


	// Intervals - A list of intervals during which to run the associated Campaign.
	Intervals *[]Scheduleinterval `json:"intervals,omitempty"`


	// TimeZone - The time zone for this CampaignSchedule. For example, Africa/Abidjan.
	TimeZone *string `json:"timeZone,omitempty"`


	// Campaign - The Campaign that this CampaignSchedule is for.
	Campaign *Domainentityref `json:"campaign,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Campaignschedule) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

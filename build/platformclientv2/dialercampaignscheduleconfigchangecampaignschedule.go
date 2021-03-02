package platformclientv2
import (
	"time"
	"encoding/json"
)

// Dialercampaignscheduleconfigchangecampaignschedule
type Dialercampaignscheduleconfigchangecampaignschedule struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// DateCreated
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateModified
	DateModified *time.Time `json:"dateModified,omitempty"`


	// Version
	Version *int `json:"version,omitempty"`


	// Intervals
	Intervals *[]Dialercampaignscheduleconfigchangescheduleinterval `json:"intervals,omitempty"`


	// TimeZone
	TimeZone *string `json:"timeZone,omitempty"`


	// Campaign
	Campaign *Dialercampaignscheduleconfigchangeurireference `json:"campaign,omitempty"`


	// AdditionalProperties
	AdditionalProperties *map[string]interface{} `json:"additionalProperties,omitempty"`

}

// String returns a JSON representation of the model
func (o *Dialercampaignscheduleconfigchangecampaignschedule) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

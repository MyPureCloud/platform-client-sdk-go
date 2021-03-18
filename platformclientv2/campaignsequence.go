package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Campaignsequence
type Campaignsequence struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// DateCreated - Creation time of the entity. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateModified - Last modified time of the entity. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`


	// Version - Required for updates, must match the version number of the most recent update
	Version *int `json:"version,omitempty"`


	// Campaigns - The ordered list of Campaigns that this CampaignSequence will run.
	Campaigns *[]Domainentityref `json:"campaigns,omitempty"`


	// CurrentCampaign - A zero-based index indicating which Campaign this CampaignSequence is currently on.
	CurrentCampaign *int `json:"currentCampaign,omitempty"`


	// Status - The current status of the CampaignSequence. A CampaignSequence can be turned 'on' or 'off'.
	Status *string `json:"status,omitempty"`


	// StopMessage - A message indicating if and why a CampaignSequence has stopped unexpectedly.
	StopMessage *string `json:"stopMessage,omitempty"`


	// Repeat - Indicates if a sequence should repeat from the beginning after the last campaign completes. Default is false.
	Repeat *bool `json:"repeat,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Campaignsequence) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

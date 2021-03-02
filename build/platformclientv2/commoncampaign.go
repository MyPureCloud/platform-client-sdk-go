package platformclientv2
import (
	"encoding/json"
)

// Commoncampaign
type Commoncampaign struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name - The name of the Campaign.
	Name *string `json:"name,omitempty"`


	// Division - The division to which this entity belongs.
	Division *Division `json:"division,omitempty"`


	// MediaType - The media type used for this campaign.
	MediaType *string `json:"mediaType,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Commoncampaign) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

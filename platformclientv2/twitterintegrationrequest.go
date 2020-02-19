package platformclientv2
import (
	"encoding/json"
)

// Twitterintegrationrequest
type Twitterintegrationrequest struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name - The name of the Twitter Integration
	Name *string `json:"name,omitempty"`


	// AccessTokenKey - The Access Token Key from Twitter messenger
	AccessTokenKey *string `json:"accessTokenKey,omitempty"`


	// AccessTokenSecret - The Access Token Secret from Twitter messenger
	AccessTokenSecret *string `json:"accessTokenSecret,omitempty"`


	// ConsumerKey - The Consumer Key from Twitter messenger
	ConsumerKey *string `json:"consumerKey,omitempty"`


	// ConsumerSecret - The Consumer Secret from Twitter messenger
	ConsumerSecret *string `json:"consumerSecret,omitempty"`


	// Tier - The type of twitter account to be used for the integration
	Tier *string `json:"tier,omitempty"`


	// EnvName - The Twitter environment name, e.g.: env-beta (required for premium tier)
	EnvName *string `json:"envName,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Twitterintegrationrequest) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

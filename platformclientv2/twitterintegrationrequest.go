package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
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

func (u *Twitterintegrationrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Twitterintegrationrequest

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		AccessTokenKey *string `json:"accessTokenKey,omitempty"`
		
		AccessTokenSecret *string `json:"accessTokenSecret,omitempty"`
		
		ConsumerKey *string `json:"consumerKey,omitempty"`
		
		ConsumerSecret *string `json:"consumerSecret,omitempty"`
		
		Tier *string `json:"tier,omitempty"`
		
		EnvName *string `json:"envName,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		AccessTokenKey: u.AccessTokenKey,
		
		AccessTokenSecret: u.AccessTokenSecret,
		
		ConsumerKey: u.ConsumerKey,
		
		ConsumerSecret: u.ConsumerSecret,
		
		Tier: u.Tier,
		
		EnvName: u.EnvName,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Twitterintegrationrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

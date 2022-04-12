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


	// SupportedContent - Defines the SupportedContent profile configured for an integration
	SupportedContent *Supportedcontentreference `json:"supportedContent,omitempty"`


	// MessagingSetting
	MessagingSetting *Messagingsettingreference `json:"messagingSetting,omitempty"`


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

func (o *Twitterintegrationrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Twitterintegrationrequest
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		SupportedContent *Supportedcontentreference `json:"supportedContent,omitempty"`
		
		MessagingSetting *Messagingsettingreference `json:"messagingSetting,omitempty"`
		
		AccessTokenKey *string `json:"accessTokenKey,omitempty"`
		
		AccessTokenSecret *string `json:"accessTokenSecret,omitempty"`
		
		ConsumerKey *string `json:"consumerKey,omitempty"`
		
		ConsumerSecret *string `json:"consumerSecret,omitempty"`
		
		Tier *string `json:"tier,omitempty"`
		
		EnvName *string `json:"envName,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		SupportedContent: o.SupportedContent,
		
		MessagingSetting: o.MessagingSetting,
		
		AccessTokenKey: o.AccessTokenKey,
		
		AccessTokenSecret: o.AccessTokenSecret,
		
		ConsumerKey: o.ConsumerKey,
		
		ConsumerSecret: o.ConsumerSecret,
		
		Tier: o.Tier,
		
		EnvName: o.EnvName,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Twitterintegrationrequest) UnmarshalJSON(b []byte) error {
	var TwitterintegrationrequestMap map[string]interface{}
	err := json.Unmarshal(b, &TwitterintegrationrequestMap)
	if err != nil {
		return err
	}
	
	if Id, ok := TwitterintegrationrequestMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := TwitterintegrationrequestMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if SupportedContent, ok := TwitterintegrationrequestMap["supportedContent"].(map[string]interface{}); ok {
		SupportedContentString, _ := json.Marshal(SupportedContent)
		json.Unmarshal(SupportedContentString, &o.SupportedContent)
	}
	
	if MessagingSetting, ok := TwitterintegrationrequestMap["messagingSetting"].(map[string]interface{}); ok {
		MessagingSettingString, _ := json.Marshal(MessagingSetting)
		json.Unmarshal(MessagingSettingString, &o.MessagingSetting)
	}
	
	if AccessTokenKey, ok := TwitterintegrationrequestMap["accessTokenKey"].(string); ok {
		o.AccessTokenKey = &AccessTokenKey
	}
	
	if AccessTokenSecret, ok := TwitterintegrationrequestMap["accessTokenSecret"].(string); ok {
		o.AccessTokenSecret = &AccessTokenSecret
	}
	
	if ConsumerKey, ok := TwitterintegrationrequestMap["consumerKey"].(string); ok {
		o.ConsumerKey = &ConsumerKey
	}
	
	if ConsumerSecret, ok := TwitterintegrationrequestMap["consumerSecret"].(string); ok {
		o.ConsumerSecret = &ConsumerSecret
	}
	
	if Tier, ok := TwitterintegrationrequestMap["tier"].(string); ok {
		o.Tier = &Tier
	}
	
	if EnvName, ok := TwitterintegrationrequestMap["envName"].(string); ok {
		o.EnvName = &EnvName
	}
	
	if SelfUri, ok := TwitterintegrationrequestMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Twitterintegrationrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

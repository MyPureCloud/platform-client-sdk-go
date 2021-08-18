package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Webdeploymentconfigurationversion - Details about the configuration version of a Web Deployment
type Webdeploymentconfigurationversion struct { 
	// Id - The configuration version ID
	Id *string `json:"id,omitempty"`


	// Name - The configuration version name
	Name *string `json:"name,omitempty"`


	// Version - The version of the configuration
	Version *string `json:"version,omitempty"`


	// Description - The description of the configuration
	Description *string `json:"description,omitempty"`


	// Languages - A list of languages supported on the configuration
	Languages *[]string `json:"languages,omitempty"`


	// DefaultLanguage - The default language to use for the configuration
	DefaultLanguage *string `json:"defaultLanguage,omitempty"`


	// Messenger - The settings for messenger
	Messenger *Messengersettings `json:"messenger,omitempty"`


	// JourneyEvents - The settings for journey events
	JourneyEvents *Journeyeventssettings `json:"journeyEvents,omitempty"`


	// AuthenticationSettings - The settings for authenticated deployments
	AuthenticationSettings *Authenticationsettings `json:"authenticationSettings,omitempty"`


	// DateCreated - The date the configuration version was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateModified - The date the configuration version was most recently modified. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`


	// DatePublished - The date the configuration version was most recently published. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DatePublished *time.Time `json:"datePublished,omitempty"`


	// LastModifiedUser - A reference to the user who most recently modified the configuration version
	LastModifiedUser *Addressableentityref `json:"lastModifiedUser,omitempty"`


	// CreatedUser - A reference to the user who created the configuration version
	CreatedUser *Addressableentityref `json:"createdUser,omitempty"`


	// PublishedUser - A reference to the user who published the configuration version
	PublishedUser *Addressableentityref `json:"publishedUser,omitempty"`


	// Status - The current status of the configuration version
	Status *string `json:"status,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (u *Webdeploymentconfigurationversion) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Webdeploymentconfigurationversion

	
	DateCreated := new(string)
	if u.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(u.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	DateModified := new(string)
	if u.DateModified != nil {
		
		*DateModified = timeutil.Strftime(u.DateModified, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateModified = nil
	}
	
	DatePublished := new(string)
	if u.DatePublished != nil {
		
		*DatePublished = timeutil.Strftime(u.DatePublished, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DatePublished = nil
	}
	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Version *string `json:"version,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		Languages *[]string `json:"languages,omitempty"`
		
		DefaultLanguage *string `json:"defaultLanguage,omitempty"`
		
		Messenger *Messengersettings `json:"messenger,omitempty"`
		
		JourneyEvents *Journeyeventssettings `json:"journeyEvents,omitempty"`
		
		AuthenticationSettings *Authenticationsettings `json:"authenticationSettings,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		DatePublished *string `json:"datePublished,omitempty"`
		
		LastModifiedUser *Addressableentityref `json:"lastModifiedUser,omitempty"`
		
		CreatedUser *Addressableentityref `json:"createdUser,omitempty"`
		
		PublishedUser *Addressableentityref `json:"publishedUser,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		Version: u.Version,
		
		Description: u.Description,
		
		Languages: u.Languages,
		
		DefaultLanguage: u.DefaultLanguage,
		
		Messenger: u.Messenger,
		
		JourneyEvents: u.JourneyEvents,
		
		AuthenticationSettings: u.AuthenticationSettings,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		DatePublished: DatePublished,
		
		LastModifiedUser: u.LastModifiedUser,
		
		CreatedUser: u.CreatedUser,
		
		PublishedUser: u.PublishedUser,
		
		Status: u.Status,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Webdeploymentconfigurationversion) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

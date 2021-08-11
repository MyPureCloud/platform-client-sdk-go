package platformclientv2
import (
	"time"
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

// String returns a JSON representation of the model
func (o *Webdeploymentconfigurationversion) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

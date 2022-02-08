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


	// Position - The settings for position
	Position *Positionsettings `json:"position,omitempty"`


	// SupportCenter - The settings for support center
	SupportCenter *Supportcentersettings `json:"supportCenter,omitempty"`


	// Cobrowse - The settings for cobrowse
	Cobrowse *Cobrowsesettings `json:"cobrowse,omitempty"`


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

func (o *Webdeploymentconfigurationversion) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Webdeploymentconfigurationversion
	
	DateCreated := new(string)
	if o.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(o.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	DateModified := new(string)
	if o.DateModified != nil {
		
		*DateModified = timeutil.Strftime(o.DateModified, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateModified = nil
	}
	
	DatePublished := new(string)
	if o.DatePublished != nil {
		
		*DatePublished = timeutil.Strftime(o.DatePublished, "%Y-%m-%dT%H:%M:%S.%fZ")
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
		
		Position *Positionsettings `json:"position,omitempty"`
		
		SupportCenter *Supportcentersettings `json:"supportCenter,omitempty"`
		
		Cobrowse *Cobrowsesettings `json:"cobrowse,omitempty"`
		
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
		Id: o.Id,
		
		Name: o.Name,
		
		Version: o.Version,
		
		Description: o.Description,
		
		Languages: o.Languages,
		
		DefaultLanguage: o.DefaultLanguage,
		
		Messenger: o.Messenger,
		
		Position: o.Position,
		
		SupportCenter: o.SupportCenter,
		
		Cobrowse: o.Cobrowse,
		
		JourneyEvents: o.JourneyEvents,
		
		AuthenticationSettings: o.AuthenticationSettings,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		DatePublished: DatePublished,
		
		LastModifiedUser: o.LastModifiedUser,
		
		CreatedUser: o.CreatedUser,
		
		PublishedUser: o.PublishedUser,
		
		Status: o.Status,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Webdeploymentconfigurationversion) UnmarshalJSON(b []byte) error {
	var WebdeploymentconfigurationversionMap map[string]interface{}
	err := json.Unmarshal(b, &WebdeploymentconfigurationversionMap)
	if err != nil {
		return err
	}
	
	if Id, ok := WebdeploymentconfigurationversionMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := WebdeploymentconfigurationversionMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if Version, ok := WebdeploymentconfigurationversionMap["version"].(string); ok {
		o.Version = &Version
	}
	
	if Description, ok := WebdeploymentconfigurationversionMap["description"].(string); ok {
		o.Description = &Description
	}
	
	if Languages, ok := WebdeploymentconfigurationversionMap["languages"].([]interface{}); ok {
		LanguagesString, _ := json.Marshal(Languages)
		json.Unmarshal(LanguagesString, &o.Languages)
	}
	
	if DefaultLanguage, ok := WebdeploymentconfigurationversionMap["defaultLanguage"].(string); ok {
		o.DefaultLanguage = &DefaultLanguage
	}
	
	if Messenger, ok := WebdeploymentconfigurationversionMap["messenger"].(map[string]interface{}); ok {
		MessengerString, _ := json.Marshal(Messenger)
		json.Unmarshal(MessengerString, &o.Messenger)
	}
	
	if Position, ok := WebdeploymentconfigurationversionMap["position"].(map[string]interface{}); ok {
		PositionString, _ := json.Marshal(Position)
		json.Unmarshal(PositionString, &o.Position)
	}
	
	if SupportCenter, ok := WebdeploymentconfigurationversionMap["supportCenter"].(map[string]interface{}); ok {
		SupportCenterString, _ := json.Marshal(SupportCenter)
		json.Unmarshal(SupportCenterString, &o.SupportCenter)
	}
	
	if Cobrowse, ok := WebdeploymentconfigurationversionMap["cobrowse"].(map[string]interface{}); ok {
		CobrowseString, _ := json.Marshal(Cobrowse)
		json.Unmarshal(CobrowseString, &o.Cobrowse)
	}
	
	if JourneyEvents, ok := WebdeploymentconfigurationversionMap["journeyEvents"].(map[string]interface{}); ok {
		JourneyEventsString, _ := json.Marshal(JourneyEvents)
		json.Unmarshal(JourneyEventsString, &o.JourneyEvents)
	}
	
	if AuthenticationSettings, ok := WebdeploymentconfigurationversionMap["authenticationSettings"].(map[string]interface{}); ok {
		AuthenticationSettingsString, _ := json.Marshal(AuthenticationSettings)
		json.Unmarshal(AuthenticationSettingsString, &o.AuthenticationSettings)
	}
	
	if dateCreatedString, ok := WebdeploymentconfigurationversionMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := WebdeploymentconfigurationversionMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if datePublishedString, ok := WebdeploymentconfigurationversionMap["datePublished"].(string); ok {
		DatePublished, _ := time.Parse("2006-01-02T15:04:05.999999Z", datePublishedString)
		o.DatePublished = &DatePublished
	}
	
	if LastModifiedUser, ok := WebdeploymentconfigurationversionMap["lastModifiedUser"].(map[string]interface{}); ok {
		LastModifiedUserString, _ := json.Marshal(LastModifiedUser)
		json.Unmarshal(LastModifiedUserString, &o.LastModifiedUser)
	}
	
	if CreatedUser, ok := WebdeploymentconfigurationversionMap["createdUser"].(map[string]interface{}); ok {
		CreatedUserString, _ := json.Marshal(CreatedUser)
		json.Unmarshal(CreatedUserString, &o.CreatedUser)
	}
	
	if PublishedUser, ok := WebdeploymentconfigurationversionMap["publishedUser"].(map[string]interface{}); ok {
		PublishedUserString, _ := json.Marshal(PublishedUser)
		json.Unmarshal(PublishedUserString, &o.PublishedUser)
	}
	
	if Status, ok := WebdeploymentconfigurationversionMap["status"].(string); ok {
		o.Status = &Status
	}
	
	if SelfUri, ok := WebdeploymentconfigurationversionMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Webdeploymentconfigurationversion) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

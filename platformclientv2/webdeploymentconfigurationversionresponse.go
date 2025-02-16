package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Webdeploymentconfigurationversionresponse - Details about the configuration version of a Web Deployment
type Webdeploymentconfigurationversionresponse struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The configuration version ID
	Id *string `json:"id,omitempty"`

	// Name - The configuration version name
	Name *string `json:"name,omitempty"`

	// Version - The version of the configuration
	Version *string `json:"version,omitempty"`

	// HeadlessMode - Headless Mode Support which Controls UI components. When enabled, native UI components will be disabled and allows for custom-built UI.
	HeadlessMode *Webdeploymentheadlessmode `json:"headlessMode,omitempty"`

	// Description - The description of the configuration
	Description *string `json:"description,omitempty"`

	// Languages - A list of languages supported on the configuration required if the messenger is enabled
	Languages *[]string `json:"languages,omitempty"`

	// DefaultLanguage - The default language to use for the configuration required if the messenger is enabled
	DefaultLanguage *string `json:"defaultLanguage,omitempty"`

	// CustomI18nLabels - The localization settings for homescreen app
	CustomI18nLabels *[]Customi18nlabels `json:"customI18nLabels,omitempty"`

	// Messenger - The settings for messenger
	Messenger *Messengersettings `json:"messenger,omitempty"`

	// Position - The settings for position
	Position *Positionsettings `json:"position,omitempty"`

	// SupportCenter - The settings for knowledge portal (previously support center)
	SupportCenter *Supportcentersettings `json:"supportCenter,omitempty"`

	// Cobrowse - The settings for cobrowse
	Cobrowse *Cobrowsesettings `json:"cobrowse,omitempty"`

	// JourneyEvents - The settings for journey events
	JourneyEvents *Journeyeventssettings `json:"journeyEvents,omitempty"`

	// AuthenticationSettings - The settings for authenticated deployments
	AuthenticationSettings *Authenticationsettings `json:"authenticationSettings,omitempty"`

	// Video - The settings for video
	Video *Videosettings `json:"video,omitempty"`

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

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Webdeploymentconfigurationversionresponse) SetField(field string, fieldValue interface{}) {
	// Get Value object for field
	target := reflect.ValueOf(o)
	targetField := reflect.Indirect(target).FieldByName(field)

	// Set value
	if fieldValue != nil {
		targetField.Set(reflect.ValueOf(fieldValue))
	} else {
		// Must create a new Value (creates **type) then get its element (*type), which will be nil pointer of the appropriate type
		x := reflect.Indirect(reflect.New(targetField.Type()))
		targetField.Set(x)
	}

	// Add field to set field names list
	if o.SetFieldNames == nil {
		o.SetFieldNames = make(map[string]bool)
	}
	o.SetFieldNames[field] = true
}

func (o Webdeploymentconfigurationversionresponse) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateCreated","DateModified","DatePublished", }
		localDateTimeFields := []string{  }
		dateFields := []string{  }

		// Construct object
		newObj := make(map[string]interface{})
		for fieldName := range o.SetFieldNames {
			// Get initial field value
			fieldValue := val.FieldByName(fieldName).Interface()

			// Apply value formatting overrides
			if fieldValue == nil || reflect.ValueOf(fieldValue).IsNil()  {
				// Do nothing. Just catching this case to avoid trying to custom serialize a nil value
			} else if contains(dateTimeFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%dT%H:%M:%S.%fZ")
			} else if contains(localDateTimeFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%dT%H:%M:%S.%f")
			} else if contains(dateFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%d")
			}

			// Assign value to field using JSON tag name
			newObj[getFieldName(reflect.TypeOf(&o), fieldName)] = fieldValue
		}

		// Marshal and return dynamically constructed interface
		return json.Marshal(newObj)
	}

	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Webdeploymentconfigurationversionresponse
	
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
		
		HeadlessMode *Webdeploymentheadlessmode `json:"headlessMode,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		Languages *[]string `json:"languages,omitempty"`
		
		DefaultLanguage *string `json:"defaultLanguage,omitempty"`
		
		CustomI18nLabels *[]Customi18nlabels `json:"customI18nLabels,omitempty"`
		
		Messenger *Messengersettings `json:"messenger,omitempty"`
		
		Position *Positionsettings `json:"position,omitempty"`
		
		SupportCenter *Supportcentersettings `json:"supportCenter,omitempty"`
		
		Cobrowse *Cobrowsesettings `json:"cobrowse,omitempty"`
		
		JourneyEvents *Journeyeventssettings `json:"journeyEvents,omitempty"`
		
		AuthenticationSettings *Authenticationsettings `json:"authenticationSettings,omitempty"`
		
		Video *Videosettings `json:"video,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		DatePublished *string `json:"datePublished,omitempty"`
		
		LastModifiedUser *Addressableentityref `json:"lastModifiedUser,omitempty"`
		
		CreatedUser *Addressableentityref `json:"createdUser,omitempty"`
		
		PublishedUser *Addressableentityref `json:"publishedUser,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Version: o.Version,
		
		HeadlessMode: o.HeadlessMode,
		
		Description: o.Description,
		
		Languages: o.Languages,
		
		DefaultLanguage: o.DefaultLanguage,
		
		CustomI18nLabels: o.CustomI18nLabels,
		
		Messenger: o.Messenger,
		
		Position: o.Position,
		
		SupportCenter: o.SupportCenter,
		
		Cobrowse: o.Cobrowse,
		
		JourneyEvents: o.JourneyEvents,
		
		AuthenticationSettings: o.AuthenticationSettings,
		
		Video: o.Video,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		DatePublished: DatePublished,
		
		LastModifiedUser: o.LastModifiedUser,
		
		CreatedUser: o.CreatedUser,
		
		PublishedUser: o.PublishedUser,
		
		Status: o.Status,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Webdeploymentconfigurationversionresponse) UnmarshalJSON(b []byte) error {
	var WebdeploymentconfigurationversionresponseMap map[string]interface{}
	err := json.Unmarshal(b, &WebdeploymentconfigurationversionresponseMap)
	if err != nil {
		return err
	}
	
	if Id, ok := WebdeploymentconfigurationversionresponseMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := WebdeploymentconfigurationversionresponseMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Version, ok := WebdeploymentconfigurationversionresponseMap["version"].(string); ok {
		o.Version = &Version
	}
    
	if HeadlessMode, ok := WebdeploymentconfigurationversionresponseMap["headlessMode"].(map[string]interface{}); ok {
		HeadlessModeString, _ := json.Marshal(HeadlessMode)
		json.Unmarshal(HeadlessModeString, &o.HeadlessMode)
	}
	
	if Description, ok := WebdeploymentconfigurationversionresponseMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if Languages, ok := WebdeploymentconfigurationversionresponseMap["languages"].([]interface{}); ok {
		LanguagesString, _ := json.Marshal(Languages)
		json.Unmarshal(LanguagesString, &o.Languages)
	}
	
	if DefaultLanguage, ok := WebdeploymentconfigurationversionresponseMap["defaultLanguage"].(string); ok {
		o.DefaultLanguage = &DefaultLanguage
	}
    
	if CustomI18nLabels, ok := WebdeploymentconfigurationversionresponseMap["customI18nLabels"].([]interface{}); ok {
		CustomI18nLabelsString, _ := json.Marshal(CustomI18nLabels)
		json.Unmarshal(CustomI18nLabelsString, &o.CustomI18nLabels)
	}
	
	if Messenger, ok := WebdeploymentconfigurationversionresponseMap["messenger"].(map[string]interface{}); ok {
		MessengerString, _ := json.Marshal(Messenger)
		json.Unmarshal(MessengerString, &o.Messenger)
	}
	
	if Position, ok := WebdeploymentconfigurationversionresponseMap["position"].(map[string]interface{}); ok {
		PositionString, _ := json.Marshal(Position)
		json.Unmarshal(PositionString, &o.Position)
	}
	
	if SupportCenter, ok := WebdeploymentconfigurationversionresponseMap["supportCenter"].(map[string]interface{}); ok {
		SupportCenterString, _ := json.Marshal(SupportCenter)
		json.Unmarshal(SupportCenterString, &o.SupportCenter)
	}
	
	if Cobrowse, ok := WebdeploymentconfigurationversionresponseMap["cobrowse"].(map[string]interface{}); ok {
		CobrowseString, _ := json.Marshal(Cobrowse)
		json.Unmarshal(CobrowseString, &o.Cobrowse)
	}
	
	if JourneyEvents, ok := WebdeploymentconfigurationversionresponseMap["journeyEvents"].(map[string]interface{}); ok {
		JourneyEventsString, _ := json.Marshal(JourneyEvents)
		json.Unmarshal(JourneyEventsString, &o.JourneyEvents)
	}
	
	if AuthenticationSettings, ok := WebdeploymentconfigurationversionresponseMap["authenticationSettings"].(map[string]interface{}); ok {
		AuthenticationSettingsString, _ := json.Marshal(AuthenticationSettings)
		json.Unmarshal(AuthenticationSettingsString, &o.AuthenticationSettings)
	}
	
	if Video, ok := WebdeploymentconfigurationversionresponseMap["video"].(map[string]interface{}); ok {
		VideoString, _ := json.Marshal(Video)
		json.Unmarshal(VideoString, &o.Video)
	}
	
	if dateCreatedString, ok := WebdeploymentconfigurationversionresponseMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := WebdeploymentconfigurationversionresponseMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if datePublishedString, ok := WebdeploymentconfigurationversionresponseMap["datePublished"].(string); ok {
		DatePublished, _ := time.Parse("2006-01-02T15:04:05.999999Z", datePublishedString)
		o.DatePublished = &DatePublished
	}
	
	if LastModifiedUser, ok := WebdeploymentconfigurationversionresponseMap["lastModifiedUser"].(map[string]interface{}); ok {
		LastModifiedUserString, _ := json.Marshal(LastModifiedUser)
		json.Unmarshal(LastModifiedUserString, &o.LastModifiedUser)
	}
	
	if CreatedUser, ok := WebdeploymentconfigurationversionresponseMap["createdUser"].(map[string]interface{}); ok {
		CreatedUserString, _ := json.Marshal(CreatedUser)
		json.Unmarshal(CreatedUserString, &o.CreatedUser)
	}
	
	if PublishedUser, ok := WebdeploymentconfigurationversionresponseMap["publishedUser"].(map[string]interface{}); ok {
		PublishedUserString, _ := json.Marshal(PublishedUser)
		json.Unmarshal(PublishedUserString, &o.PublishedUser)
	}
	
	if Status, ok := WebdeploymentconfigurationversionresponseMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if SelfUri, ok := WebdeploymentconfigurationversionresponseMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Webdeploymentconfigurationversionresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

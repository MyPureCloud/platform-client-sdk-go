package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Userme
type Userme struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name
	Name *string `json:"name,omitempty"`

	// Division - The division to which this entity belongs.
	Division *Division `json:"division,omitempty"`

	// Chat
	Chat *Chat `json:"chat,omitempty"`

	// Department
	Department *string `json:"department,omitempty"`

	// Email
	Email *string `json:"email,omitempty"`

	// PrimaryContactInfo - Auto populated from addresses.
	PrimaryContactInfo *[]Contact `json:"primaryContactInfo,omitempty"`

	// Addresses - Email addresses and phone numbers for this user
	Addresses *[]Contact `json:"addresses,omitempty"`

	// State - The current state for this user.
	State *string `json:"state,omitempty"`

	// Title
	Title *string `json:"title,omitempty"`

	// Username
	Username *string `json:"username,omitempty"`

	// Manager
	Manager *User `json:"manager,omitempty"`

	// Images
	Images *[]Userimage `json:"images,omitempty"`

	// Version - Required when updating a user, this value should be the current version of the user.  The current version can be obtained with a GET on the user before doing a PATCH.
	Version *int `json:"version,omitempty"`

	// Certifications
	Certifications *[]string `json:"certifications,omitempty"`

	// Biography
	Biography *Biography `json:"biography,omitempty"`

	// EmployerInfo
	EmployerInfo *Employerinfo `json:"employerInfo,omitempty"`

	// RoutingStatus - ACD routing status
	RoutingStatus *Routingstatus `json:"routingStatus,omitempty"`

	// Presence - Active presence
	Presence *Userpresence `json:"presence,omitempty"`

	// IntegrationPresence - Integration presence
	IntegrationPresence *Userpresence `json:"integrationPresence,omitempty"`

	// ConversationSummary - Summary of conversion statistics for conversation types.
	ConversationSummary *Userconversationsummary `json:"conversationSummary,omitempty"`

	// OutOfOffice - Determine if out of office is enabled
	OutOfOffice *Outofoffice `json:"outOfOffice,omitempty"`

	// Geolocation - Current geolocation position
	Geolocation *Geolocation `json:"geolocation,omitempty"`

	// Station - Effective, default, and last station information
	Station *Userstations `json:"station,omitempty"`

	// Authorization - Roles and permissions assigned to the user
	Authorization *Userauthorization `json:"authorization,omitempty"`

	// ProfileSkills - Profile skills possessed by the user
	ProfileSkills *[]string `json:"profileSkills,omitempty"`

	// Locations - The user placement at each site location.
	Locations *[]Location `json:"locations,omitempty"`

	// Groups - The groups the user is a member of
	Groups *[]Group `json:"groups,omitempty"`

	// Team - The team the user is a member of
	Team *Team `json:"team,omitempty"`

	// Skills - Routing (ACD) skills possessed by the user
	Skills *[]Userroutingskill `json:"skills,omitempty"`

	// Languages - Routing (ACD) languages possessed by the user
	Languages *[]Userroutinglanguage `json:"languages,omitempty"`

	// AcdAutoAnswer - acd auto answer
	AcdAutoAnswer *bool `json:"acdAutoAnswer,omitempty"`

	// LanguagePreference - preferred language by the user
	LanguagePreference *string `json:"languagePreference,omitempty"`

	// LastTokenIssued
	LastTokenIssued *Oauthlasttokenissued `json:"lastTokenIssued,omitempty"`

	// DateLastLogin - The last time the user logged in using username and password. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateLastLogin *time.Time `json:"dateLastLogin,omitempty"`

	// Date - The PureCloud system date time.
	Date *Serverdate `json:"date,omitempty"`

	// GeolocationSettings - Geolocation settings for user's organization.
	GeolocationSettings *Geolocationsettings `json:"geolocationSettings,omitempty"`

	// Organization - Organization details for this user.
	Organization *Organization `json:"organization,omitempty"`

	// PresenceDefinitions - The first 100 presence definitions for user's organization.
	PresenceDefinitions *[]Organizationpresence `json:"presenceDefinitions,omitempty"`

	// LocationDefinitions - The first 100 site locations for user's organization
	LocationDefinitions *[]Locationdefinition `json:"locationDefinitions,omitempty"`

	// OrgAuthorization - The first 100 organization roles, with applicable permission policies, for user's organization.
	OrgAuthorization *[]Domainorganizationrole `json:"orgAuthorization,omitempty"`

	// Favorites - The first 50 favorited users.
	Favorites *[]User `json:"favorites,omitempty"`

	// Superiors - The first 50 superiors of this user.
	Superiors *[]User `json:"superiors,omitempty"`

	// DirectReports - The first 50 direct reports to this user.
	DirectReports *[]User `json:"directReports,omitempty"`

	// Adjacents - The first 50 superiors, direct reports, and siblings of this user. Mutually exclusive with superiors and direct reports expands.
	Adjacents *Adjacents `json:"adjacents,omitempty"`

	// RoutingSkills - The first 50 routing skills for user's organizations
	RoutingSkills *[]Routingskill `json:"routingSkills,omitempty"`

	// FieldConfigs - The field config for all entities types of user's organization
	FieldConfigs *Fieldconfigs `json:"fieldConfigs,omitempty"`

	// Token - Information about the current token
	Token *Tokeninfo `json:"token,omitempty"`

	// Trustors - Organizations having this user as a trustee
	Trustors *[]Trustor `json:"trustors,omitempty"`

	// OrgProducts - Products enabled in this organization
	OrgProducts *[]Domainorganizationproduct `json:"orgProducts,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Userme) SetField(field string, fieldValue interface{}) {
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

func (o Userme) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateLastLogin", }
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
	type Alias Userme
	
	DateLastLogin := new(string)
	if o.DateLastLogin != nil {
		
		*DateLastLogin = timeutil.Strftime(o.DateLastLogin, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateLastLogin = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Division *Division `json:"division,omitempty"`
		
		Chat *Chat `json:"chat,omitempty"`
		
		Department *string `json:"department,omitempty"`
		
		Email *string `json:"email,omitempty"`
		
		PrimaryContactInfo *[]Contact `json:"primaryContactInfo,omitempty"`
		
		Addresses *[]Contact `json:"addresses,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		Title *string `json:"title,omitempty"`
		
		Username *string `json:"username,omitempty"`
		
		Manager *User `json:"manager,omitempty"`
		
		Images *[]Userimage `json:"images,omitempty"`
		
		Version *int `json:"version,omitempty"`
		
		Certifications *[]string `json:"certifications,omitempty"`
		
		Biography *Biography `json:"biography,omitempty"`
		
		EmployerInfo *Employerinfo `json:"employerInfo,omitempty"`
		
		RoutingStatus *Routingstatus `json:"routingStatus,omitempty"`
		
		Presence *Userpresence `json:"presence,omitempty"`
		
		IntegrationPresence *Userpresence `json:"integrationPresence,omitempty"`
		
		ConversationSummary *Userconversationsummary `json:"conversationSummary,omitempty"`
		
		OutOfOffice *Outofoffice `json:"outOfOffice,omitempty"`
		
		Geolocation *Geolocation `json:"geolocation,omitempty"`
		
		Station *Userstations `json:"station,omitempty"`
		
		Authorization *Userauthorization `json:"authorization,omitempty"`
		
		ProfileSkills *[]string `json:"profileSkills,omitempty"`
		
		Locations *[]Location `json:"locations,omitempty"`
		
		Groups *[]Group `json:"groups,omitempty"`
		
		Team *Team `json:"team,omitempty"`
		
		Skills *[]Userroutingskill `json:"skills,omitempty"`
		
		Languages *[]Userroutinglanguage `json:"languages,omitempty"`
		
		AcdAutoAnswer *bool `json:"acdAutoAnswer,omitempty"`
		
		LanguagePreference *string `json:"languagePreference,omitempty"`
		
		LastTokenIssued *Oauthlasttokenissued `json:"lastTokenIssued,omitempty"`
		
		DateLastLogin *string `json:"dateLastLogin,omitempty"`
		
		Date *Serverdate `json:"date,omitempty"`
		
		GeolocationSettings *Geolocationsettings `json:"geolocationSettings,omitempty"`
		
		Organization *Organization `json:"organization,omitempty"`
		
		PresenceDefinitions *[]Organizationpresence `json:"presenceDefinitions,omitempty"`
		
		LocationDefinitions *[]Locationdefinition `json:"locationDefinitions,omitempty"`
		
		OrgAuthorization *[]Domainorganizationrole `json:"orgAuthorization,omitempty"`
		
		Favorites *[]User `json:"favorites,omitempty"`
		
		Superiors *[]User `json:"superiors,omitempty"`
		
		DirectReports *[]User `json:"directReports,omitempty"`
		
		Adjacents *Adjacents `json:"adjacents,omitempty"`
		
		RoutingSkills *[]Routingskill `json:"routingSkills,omitempty"`
		
		FieldConfigs *Fieldconfigs `json:"fieldConfigs,omitempty"`
		
		Token *Tokeninfo `json:"token,omitempty"`
		
		Trustors *[]Trustor `json:"trustors,omitempty"`
		
		OrgProducts *[]Domainorganizationproduct `json:"orgProducts,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Division: o.Division,
		
		Chat: o.Chat,
		
		Department: o.Department,
		
		Email: o.Email,
		
		PrimaryContactInfo: o.PrimaryContactInfo,
		
		Addresses: o.Addresses,
		
		State: o.State,
		
		Title: o.Title,
		
		Username: o.Username,
		
		Manager: o.Manager,
		
		Images: o.Images,
		
		Version: o.Version,
		
		Certifications: o.Certifications,
		
		Biography: o.Biography,
		
		EmployerInfo: o.EmployerInfo,
		
		RoutingStatus: o.RoutingStatus,
		
		Presence: o.Presence,
		
		IntegrationPresence: o.IntegrationPresence,
		
		ConversationSummary: o.ConversationSummary,
		
		OutOfOffice: o.OutOfOffice,
		
		Geolocation: o.Geolocation,
		
		Station: o.Station,
		
		Authorization: o.Authorization,
		
		ProfileSkills: o.ProfileSkills,
		
		Locations: o.Locations,
		
		Groups: o.Groups,
		
		Team: o.Team,
		
		Skills: o.Skills,
		
		Languages: o.Languages,
		
		AcdAutoAnswer: o.AcdAutoAnswer,
		
		LanguagePreference: o.LanguagePreference,
		
		LastTokenIssued: o.LastTokenIssued,
		
		DateLastLogin: DateLastLogin,
		
		Date: o.Date,
		
		GeolocationSettings: o.GeolocationSettings,
		
		Organization: o.Organization,
		
		PresenceDefinitions: o.PresenceDefinitions,
		
		LocationDefinitions: o.LocationDefinitions,
		
		OrgAuthorization: o.OrgAuthorization,
		
		Favorites: o.Favorites,
		
		Superiors: o.Superiors,
		
		DirectReports: o.DirectReports,
		
		Adjacents: o.Adjacents,
		
		RoutingSkills: o.RoutingSkills,
		
		FieldConfigs: o.FieldConfigs,
		
		Token: o.Token,
		
		Trustors: o.Trustors,
		
		OrgProducts: o.OrgProducts,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Userme) UnmarshalJSON(b []byte) error {
	var UsermeMap map[string]interface{}
	err := json.Unmarshal(b, &UsermeMap)
	if err != nil {
		return err
	}
	
	if Id, ok := UsermeMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := UsermeMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Division, ok := UsermeMap["division"].(map[string]interface{}); ok {
		DivisionString, _ := json.Marshal(Division)
		json.Unmarshal(DivisionString, &o.Division)
	}
	
	if Chat, ok := UsermeMap["chat"].(map[string]interface{}); ok {
		ChatString, _ := json.Marshal(Chat)
		json.Unmarshal(ChatString, &o.Chat)
	}
	
	if Department, ok := UsermeMap["department"].(string); ok {
		o.Department = &Department
	}
    
	if Email, ok := UsermeMap["email"].(string); ok {
		o.Email = &Email
	}
    
	if PrimaryContactInfo, ok := UsermeMap["primaryContactInfo"].([]interface{}); ok {
		PrimaryContactInfoString, _ := json.Marshal(PrimaryContactInfo)
		json.Unmarshal(PrimaryContactInfoString, &o.PrimaryContactInfo)
	}
	
	if Addresses, ok := UsermeMap["addresses"].([]interface{}); ok {
		AddressesString, _ := json.Marshal(Addresses)
		json.Unmarshal(AddressesString, &o.Addresses)
	}
	
	if State, ok := UsermeMap["state"].(string); ok {
		o.State = &State
	}
    
	if Title, ok := UsermeMap["title"].(string); ok {
		o.Title = &Title
	}
    
	if Username, ok := UsermeMap["username"].(string); ok {
		o.Username = &Username
	}
    
	if Manager, ok := UsermeMap["manager"].(map[string]interface{}); ok {
		ManagerString, _ := json.Marshal(Manager)
		json.Unmarshal(ManagerString, &o.Manager)
	}
	
	if Images, ok := UsermeMap["images"].([]interface{}); ok {
		ImagesString, _ := json.Marshal(Images)
		json.Unmarshal(ImagesString, &o.Images)
	}
	
	if Version, ok := UsermeMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	
	if Certifications, ok := UsermeMap["certifications"].([]interface{}); ok {
		CertificationsString, _ := json.Marshal(Certifications)
		json.Unmarshal(CertificationsString, &o.Certifications)
	}
	
	if Biography, ok := UsermeMap["biography"].(map[string]interface{}); ok {
		BiographyString, _ := json.Marshal(Biography)
		json.Unmarshal(BiographyString, &o.Biography)
	}
	
	if EmployerInfo, ok := UsermeMap["employerInfo"].(map[string]interface{}); ok {
		EmployerInfoString, _ := json.Marshal(EmployerInfo)
		json.Unmarshal(EmployerInfoString, &o.EmployerInfo)
	}
	
	if RoutingStatus, ok := UsermeMap["routingStatus"].(map[string]interface{}); ok {
		RoutingStatusString, _ := json.Marshal(RoutingStatus)
		json.Unmarshal(RoutingStatusString, &o.RoutingStatus)
	}
	
	if Presence, ok := UsermeMap["presence"].(map[string]interface{}); ok {
		PresenceString, _ := json.Marshal(Presence)
		json.Unmarshal(PresenceString, &o.Presence)
	}
	
	if IntegrationPresence, ok := UsermeMap["integrationPresence"].(map[string]interface{}); ok {
		IntegrationPresenceString, _ := json.Marshal(IntegrationPresence)
		json.Unmarshal(IntegrationPresenceString, &o.IntegrationPresence)
	}
	
	if ConversationSummary, ok := UsermeMap["conversationSummary"].(map[string]interface{}); ok {
		ConversationSummaryString, _ := json.Marshal(ConversationSummary)
		json.Unmarshal(ConversationSummaryString, &o.ConversationSummary)
	}
	
	if OutOfOffice, ok := UsermeMap["outOfOffice"].(map[string]interface{}); ok {
		OutOfOfficeString, _ := json.Marshal(OutOfOffice)
		json.Unmarshal(OutOfOfficeString, &o.OutOfOffice)
	}
	
	if Geolocation, ok := UsermeMap["geolocation"].(map[string]interface{}); ok {
		GeolocationString, _ := json.Marshal(Geolocation)
		json.Unmarshal(GeolocationString, &o.Geolocation)
	}
	
	if Station, ok := UsermeMap["station"].(map[string]interface{}); ok {
		StationString, _ := json.Marshal(Station)
		json.Unmarshal(StationString, &o.Station)
	}
	
	if Authorization, ok := UsermeMap["authorization"].(map[string]interface{}); ok {
		AuthorizationString, _ := json.Marshal(Authorization)
		json.Unmarshal(AuthorizationString, &o.Authorization)
	}
	
	if ProfileSkills, ok := UsermeMap["profileSkills"].([]interface{}); ok {
		ProfileSkillsString, _ := json.Marshal(ProfileSkills)
		json.Unmarshal(ProfileSkillsString, &o.ProfileSkills)
	}
	
	if Locations, ok := UsermeMap["locations"].([]interface{}); ok {
		LocationsString, _ := json.Marshal(Locations)
		json.Unmarshal(LocationsString, &o.Locations)
	}
	
	if Groups, ok := UsermeMap["groups"].([]interface{}); ok {
		GroupsString, _ := json.Marshal(Groups)
		json.Unmarshal(GroupsString, &o.Groups)
	}
	
	if Team, ok := UsermeMap["team"].(map[string]interface{}); ok {
		TeamString, _ := json.Marshal(Team)
		json.Unmarshal(TeamString, &o.Team)
	}
	
	if Skills, ok := UsermeMap["skills"].([]interface{}); ok {
		SkillsString, _ := json.Marshal(Skills)
		json.Unmarshal(SkillsString, &o.Skills)
	}
	
	if Languages, ok := UsermeMap["languages"].([]interface{}); ok {
		LanguagesString, _ := json.Marshal(Languages)
		json.Unmarshal(LanguagesString, &o.Languages)
	}
	
	if AcdAutoAnswer, ok := UsermeMap["acdAutoAnswer"].(bool); ok {
		o.AcdAutoAnswer = &AcdAutoAnswer
	}
    
	if LanguagePreference, ok := UsermeMap["languagePreference"].(string); ok {
		o.LanguagePreference = &LanguagePreference
	}
    
	if LastTokenIssued, ok := UsermeMap["lastTokenIssued"].(map[string]interface{}); ok {
		LastTokenIssuedString, _ := json.Marshal(LastTokenIssued)
		json.Unmarshal(LastTokenIssuedString, &o.LastTokenIssued)
	}
	
	if dateLastLoginString, ok := UsermeMap["dateLastLogin"].(string); ok {
		DateLastLogin, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateLastLoginString)
		o.DateLastLogin = &DateLastLogin
	}
	
	if Date, ok := UsermeMap["date"].(map[string]interface{}); ok {
		DateString, _ := json.Marshal(Date)
		json.Unmarshal(DateString, &o.Date)
	}
	
	if GeolocationSettings, ok := UsermeMap["geolocationSettings"].(map[string]interface{}); ok {
		GeolocationSettingsString, _ := json.Marshal(GeolocationSettings)
		json.Unmarshal(GeolocationSettingsString, &o.GeolocationSettings)
	}
	
	if Organization, ok := UsermeMap["organization"].(map[string]interface{}); ok {
		OrganizationString, _ := json.Marshal(Organization)
		json.Unmarshal(OrganizationString, &o.Organization)
	}
	
	if PresenceDefinitions, ok := UsermeMap["presenceDefinitions"].([]interface{}); ok {
		PresenceDefinitionsString, _ := json.Marshal(PresenceDefinitions)
		json.Unmarshal(PresenceDefinitionsString, &o.PresenceDefinitions)
	}
	
	if LocationDefinitions, ok := UsermeMap["locationDefinitions"].([]interface{}); ok {
		LocationDefinitionsString, _ := json.Marshal(LocationDefinitions)
		json.Unmarshal(LocationDefinitionsString, &o.LocationDefinitions)
	}
	
	if OrgAuthorization, ok := UsermeMap["orgAuthorization"].([]interface{}); ok {
		OrgAuthorizationString, _ := json.Marshal(OrgAuthorization)
		json.Unmarshal(OrgAuthorizationString, &o.OrgAuthorization)
	}
	
	if Favorites, ok := UsermeMap["favorites"].([]interface{}); ok {
		FavoritesString, _ := json.Marshal(Favorites)
		json.Unmarshal(FavoritesString, &o.Favorites)
	}
	
	if Superiors, ok := UsermeMap["superiors"].([]interface{}); ok {
		SuperiorsString, _ := json.Marshal(Superiors)
		json.Unmarshal(SuperiorsString, &o.Superiors)
	}
	
	if DirectReports, ok := UsermeMap["directReports"].([]interface{}); ok {
		DirectReportsString, _ := json.Marshal(DirectReports)
		json.Unmarshal(DirectReportsString, &o.DirectReports)
	}
	
	if Adjacents, ok := UsermeMap["adjacents"].(map[string]interface{}); ok {
		AdjacentsString, _ := json.Marshal(Adjacents)
		json.Unmarshal(AdjacentsString, &o.Adjacents)
	}
	
	if RoutingSkills, ok := UsermeMap["routingSkills"].([]interface{}); ok {
		RoutingSkillsString, _ := json.Marshal(RoutingSkills)
		json.Unmarshal(RoutingSkillsString, &o.RoutingSkills)
	}
	
	if FieldConfigs, ok := UsermeMap["fieldConfigs"].(map[string]interface{}); ok {
		FieldConfigsString, _ := json.Marshal(FieldConfigs)
		json.Unmarshal(FieldConfigsString, &o.FieldConfigs)
	}
	
	if Token, ok := UsermeMap["token"].(map[string]interface{}); ok {
		TokenString, _ := json.Marshal(Token)
		json.Unmarshal(TokenString, &o.Token)
	}
	
	if Trustors, ok := UsermeMap["trustors"].([]interface{}); ok {
		TrustorsString, _ := json.Marshal(Trustors)
		json.Unmarshal(TrustorsString, &o.Trustors)
	}
	
	if OrgProducts, ok := UsermeMap["orgProducts"].([]interface{}); ok {
		OrgProductsString, _ := json.Marshal(OrgProducts)
		json.Unmarshal(OrgProductsString, &o.OrgProducts)
	}
	
	if SelfUri, ok := UsermeMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Userme) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

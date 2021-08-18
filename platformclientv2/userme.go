package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Userme
type Userme struct { 
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

func (u *Userme) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Userme

	

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
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		Division: u.Division,
		
		Chat: u.Chat,
		
		Department: u.Department,
		
		Email: u.Email,
		
		PrimaryContactInfo: u.PrimaryContactInfo,
		
		Addresses: u.Addresses,
		
		State: u.State,
		
		Title: u.Title,
		
		Username: u.Username,
		
		Manager: u.Manager,
		
		Images: u.Images,
		
		Version: u.Version,
		
		Certifications: u.Certifications,
		
		Biography: u.Biography,
		
		EmployerInfo: u.EmployerInfo,
		
		RoutingStatus: u.RoutingStatus,
		
		Presence: u.Presence,
		
		ConversationSummary: u.ConversationSummary,
		
		OutOfOffice: u.OutOfOffice,
		
		Geolocation: u.Geolocation,
		
		Station: u.Station,
		
		Authorization: u.Authorization,
		
		ProfileSkills: u.ProfileSkills,
		
		Locations: u.Locations,
		
		Groups: u.Groups,
		
		Team: u.Team,
		
		Skills: u.Skills,
		
		Languages: u.Languages,
		
		AcdAutoAnswer: u.AcdAutoAnswer,
		
		LanguagePreference: u.LanguagePreference,
		
		LastTokenIssued: u.LastTokenIssued,
		
		Date: u.Date,
		
		GeolocationSettings: u.GeolocationSettings,
		
		Organization: u.Organization,
		
		PresenceDefinitions: u.PresenceDefinitions,
		
		LocationDefinitions: u.LocationDefinitions,
		
		OrgAuthorization: u.OrgAuthorization,
		
		Favorites: u.Favorites,
		
		Superiors: u.Superiors,
		
		DirectReports: u.DirectReports,
		
		Adjacents: u.Adjacents,
		
		RoutingSkills: u.RoutingSkills,
		
		FieldConfigs: u.FieldConfigs,
		
		Token: u.Token,
		
		Trustors: u.Trustors,
		
		OrgProducts: u.OrgProducts,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Userme) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

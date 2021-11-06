package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Orguser
type Orguser struct { 
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
	Manager **User `json:"manager,omitempty"`


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


	// DateLastLogin - The last time the user logged in using username and password. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateLastLogin *time.Time `json:"dateLastLogin,omitempty"`


	// Organization
	Organization *Organization `json:"organization,omitempty"`

}

func (o *Orguser) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Orguser
	
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
		
		Manager **User `json:"manager,omitempty"`
		
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
		
		DateLastLogin *string `json:"dateLastLogin,omitempty"`
		
		Organization *Organization `json:"organization,omitempty"`
		*Alias
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
		
		Organization: o.Organization,
		Alias:    (*Alias)(o),
	})
}

func (o *Orguser) UnmarshalJSON(b []byte) error {
	var OrguserMap map[string]interface{}
	err := json.Unmarshal(b, &OrguserMap)
	if err != nil {
		return err
	}
	
	if Id, ok := OrguserMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := OrguserMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if Division, ok := OrguserMap["division"].(map[string]interface{}); ok {
		DivisionString, _ := json.Marshal(Division)
		json.Unmarshal(DivisionString, &o.Division)
	}
	
	if Chat, ok := OrguserMap["chat"].(map[string]interface{}); ok {
		ChatString, _ := json.Marshal(Chat)
		json.Unmarshal(ChatString, &o.Chat)
	}
	
	if Department, ok := OrguserMap["department"].(string); ok {
		o.Department = &Department
	}
	
	if Email, ok := OrguserMap["email"].(string); ok {
		o.Email = &Email
	}
	
	if PrimaryContactInfo, ok := OrguserMap["primaryContactInfo"].([]interface{}); ok {
		PrimaryContactInfoString, _ := json.Marshal(PrimaryContactInfo)
		json.Unmarshal(PrimaryContactInfoString, &o.PrimaryContactInfo)
	}
	
	if Addresses, ok := OrguserMap["addresses"].([]interface{}); ok {
		AddressesString, _ := json.Marshal(Addresses)
		json.Unmarshal(AddressesString, &o.Addresses)
	}
	
	if State, ok := OrguserMap["state"].(string); ok {
		o.State = &State
	}
	
	if Title, ok := OrguserMap["title"].(string); ok {
		o.Title = &Title
	}
	
	if Username, ok := OrguserMap["username"].(string); ok {
		o.Username = &Username
	}
	
	if Manager, ok := OrguserMap["manager"].(map[string]interface{}); ok {
		ManagerString, _ := json.Marshal(Manager)
		json.Unmarshal(ManagerString, &o.Manager)
	}
	
	if Images, ok := OrguserMap["images"].([]interface{}); ok {
		ImagesString, _ := json.Marshal(Images)
		json.Unmarshal(ImagesString, &o.Images)
	}
	
	if Version, ok := OrguserMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	
	if Certifications, ok := OrguserMap["certifications"].([]interface{}); ok {
		CertificationsString, _ := json.Marshal(Certifications)
		json.Unmarshal(CertificationsString, &o.Certifications)
	}
	
	if Biography, ok := OrguserMap["biography"].(map[string]interface{}); ok {
		BiographyString, _ := json.Marshal(Biography)
		json.Unmarshal(BiographyString, &o.Biography)
	}
	
	if EmployerInfo, ok := OrguserMap["employerInfo"].(map[string]interface{}); ok {
		EmployerInfoString, _ := json.Marshal(EmployerInfo)
		json.Unmarshal(EmployerInfoString, &o.EmployerInfo)
	}
	
	if RoutingStatus, ok := OrguserMap["routingStatus"].(map[string]interface{}); ok {
		RoutingStatusString, _ := json.Marshal(RoutingStatus)
		json.Unmarshal(RoutingStatusString, &o.RoutingStatus)
	}
	
	if Presence, ok := OrguserMap["presence"].(map[string]interface{}); ok {
		PresenceString, _ := json.Marshal(Presence)
		json.Unmarshal(PresenceString, &o.Presence)
	}
	
	if ConversationSummary, ok := OrguserMap["conversationSummary"].(map[string]interface{}); ok {
		ConversationSummaryString, _ := json.Marshal(ConversationSummary)
		json.Unmarshal(ConversationSummaryString, &o.ConversationSummary)
	}
	
	if OutOfOffice, ok := OrguserMap["outOfOffice"].(map[string]interface{}); ok {
		OutOfOfficeString, _ := json.Marshal(OutOfOffice)
		json.Unmarshal(OutOfOfficeString, &o.OutOfOffice)
	}
	
	if Geolocation, ok := OrguserMap["geolocation"].(map[string]interface{}); ok {
		GeolocationString, _ := json.Marshal(Geolocation)
		json.Unmarshal(GeolocationString, &o.Geolocation)
	}
	
	if Station, ok := OrguserMap["station"].(map[string]interface{}); ok {
		StationString, _ := json.Marshal(Station)
		json.Unmarshal(StationString, &o.Station)
	}
	
	if Authorization, ok := OrguserMap["authorization"].(map[string]interface{}); ok {
		AuthorizationString, _ := json.Marshal(Authorization)
		json.Unmarshal(AuthorizationString, &o.Authorization)
	}
	
	if ProfileSkills, ok := OrguserMap["profileSkills"].([]interface{}); ok {
		ProfileSkillsString, _ := json.Marshal(ProfileSkills)
		json.Unmarshal(ProfileSkillsString, &o.ProfileSkills)
	}
	
	if Locations, ok := OrguserMap["locations"].([]interface{}); ok {
		LocationsString, _ := json.Marshal(Locations)
		json.Unmarshal(LocationsString, &o.Locations)
	}
	
	if Groups, ok := OrguserMap["groups"].([]interface{}); ok {
		GroupsString, _ := json.Marshal(Groups)
		json.Unmarshal(GroupsString, &o.Groups)
	}
	
	if Team, ok := OrguserMap["team"].(map[string]interface{}); ok {
		TeamString, _ := json.Marshal(Team)
		json.Unmarshal(TeamString, &o.Team)
	}
	
	if Skills, ok := OrguserMap["skills"].([]interface{}); ok {
		SkillsString, _ := json.Marshal(Skills)
		json.Unmarshal(SkillsString, &o.Skills)
	}
	
	if Languages, ok := OrguserMap["languages"].([]interface{}); ok {
		LanguagesString, _ := json.Marshal(Languages)
		json.Unmarshal(LanguagesString, &o.Languages)
	}
	
	if AcdAutoAnswer, ok := OrguserMap["acdAutoAnswer"].(bool); ok {
		o.AcdAutoAnswer = &AcdAutoAnswer
	}
	
	if LanguagePreference, ok := OrguserMap["languagePreference"].(string); ok {
		o.LanguagePreference = &LanguagePreference
	}
	
	if LastTokenIssued, ok := OrguserMap["lastTokenIssued"].(map[string]interface{}); ok {
		LastTokenIssuedString, _ := json.Marshal(LastTokenIssued)
		json.Unmarshal(LastTokenIssuedString, &o.LastTokenIssued)
	}
	
	if dateLastLoginString, ok := OrguserMap["dateLastLogin"].(string); ok {
		DateLastLogin, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateLastLoginString)
		o.DateLastLogin = &DateLastLogin
	}
	
	if Organization, ok := OrguserMap["organization"].(map[string]interface{}); ok {
		OrganizationString, _ := json.Marshal(Organization)
		json.Unmarshal(OrganizationString, &o.Organization)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Orguser) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

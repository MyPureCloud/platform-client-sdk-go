package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// User
type User struct { 
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
	Manager **User `json:"manager,omitempty"`

	// Images
	Images *[]Image `json:"images,omitempty"`

	// Version - Required when updating a user, this value should be the current version of the user.  The current version can be obtained with a GET on the user before doing a PATCH.
	Version *int `json:"version,omitempty"`

	// Certifications
	Certifications *[]string `json:"certifications,omitempty"`

	// Biography
	Biography *Biography `json:"biography,omitempty"`

	// EmployerInfo
	EmployerInfo *Employerinfo `json:"employerInfo,omitempty"`

	// PreferredName - Preferred full name of the agent
	PreferredName *string `json:"preferredName,omitempty"`

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
	Station **Userstations `json:"station,omitempty"`

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

	// WorkPlanBidRanks - The WFM work plan bid rank settings for the user
	WorkPlanBidRanks *Workplanbidranks `json:"workPlanBidRanks,omitempty"`

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

	// DateWelcomeSent - The date & time the user was sent their welcome email. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateWelcomeSent *time.Time `json:"dateWelcomeSent,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *User) SetField(field string, fieldValue interface{}) {
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

func (o User) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateLastLogin","DateWelcomeSent", }
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
	type Alias User
	
	DateLastLogin := new(string)
	if o.DateLastLogin != nil {
		
		*DateLastLogin = timeutil.Strftime(o.DateLastLogin, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateLastLogin = nil
	}
	
	DateWelcomeSent := new(string)
	if o.DateWelcomeSent != nil {
		
		*DateWelcomeSent = timeutil.Strftime(o.DateWelcomeSent, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateWelcomeSent = nil
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
		
		Images *[]Image `json:"images,omitempty"`
		
		Version *int `json:"version,omitempty"`
		
		Certifications *[]string `json:"certifications,omitempty"`
		
		Biography *Biography `json:"biography,omitempty"`
		
		EmployerInfo *Employerinfo `json:"employerInfo,omitempty"`
		
		PreferredName *string `json:"preferredName,omitempty"`
		
		RoutingStatus *Routingstatus `json:"routingStatus,omitempty"`
		
		Presence *Userpresence `json:"presence,omitempty"`
		
		IntegrationPresence *Userpresence `json:"integrationPresence,omitempty"`
		
		ConversationSummary *Userconversationsummary `json:"conversationSummary,omitempty"`
		
		OutOfOffice *Outofoffice `json:"outOfOffice,omitempty"`
		
		Geolocation *Geolocation `json:"geolocation,omitempty"`
		
		Station **Userstations `json:"station,omitempty"`
		
		Authorization *Userauthorization `json:"authorization,omitempty"`
		
		ProfileSkills *[]string `json:"profileSkills,omitempty"`
		
		Locations *[]Location `json:"locations,omitempty"`
		
		Groups *[]Group `json:"groups,omitempty"`
		
		Team *Team `json:"team,omitempty"`
		
		WorkPlanBidRanks *Workplanbidranks `json:"workPlanBidRanks,omitempty"`
		
		Skills *[]Userroutingskill `json:"skills,omitempty"`
		
		Languages *[]Userroutinglanguage `json:"languages,omitempty"`
		
		AcdAutoAnswer *bool `json:"acdAutoAnswer,omitempty"`
		
		LanguagePreference *string `json:"languagePreference,omitempty"`
		
		LastTokenIssued *Oauthlasttokenissued `json:"lastTokenIssued,omitempty"`
		
		DateLastLogin *string `json:"dateLastLogin,omitempty"`
		
		DateWelcomeSent *string `json:"dateWelcomeSent,omitempty"`
		
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
		
		PreferredName: o.PreferredName,
		
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
		
		WorkPlanBidRanks: o.WorkPlanBidRanks,
		
		Skills: o.Skills,
		
		Languages: o.Languages,
		
		AcdAutoAnswer: o.AcdAutoAnswer,
		
		LanguagePreference: o.LanguagePreference,
		
		LastTokenIssued: o.LastTokenIssued,
		
		DateLastLogin: DateLastLogin,
		
		DateWelcomeSent: DateWelcomeSent,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *User) UnmarshalJSON(b []byte) error {
	var UserMap map[string]interface{}
	err := json.Unmarshal(b, &UserMap)
	if err != nil {
		return err
	}
	
	if Id, ok := UserMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := UserMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Division, ok := UserMap["division"].(map[string]interface{}); ok {
		DivisionString, _ := json.Marshal(Division)
		json.Unmarshal(DivisionString, &o.Division)
	}
	
	if Chat, ok := UserMap["chat"].(map[string]interface{}); ok {
		ChatString, _ := json.Marshal(Chat)
		json.Unmarshal(ChatString, &o.Chat)
	}
	
	if Department, ok := UserMap["department"].(string); ok {
		o.Department = &Department
	}
    
	if Email, ok := UserMap["email"].(string); ok {
		o.Email = &Email
	}
    
	if PrimaryContactInfo, ok := UserMap["primaryContactInfo"].([]interface{}); ok {
		PrimaryContactInfoString, _ := json.Marshal(PrimaryContactInfo)
		json.Unmarshal(PrimaryContactInfoString, &o.PrimaryContactInfo)
	}
	
	if Addresses, ok := UserMap["addresses"].([]interface{}); ok {
		AddressesString, _ := json.Marshal(Addresses)
		json.Unmarshal(AddressesString, &o.Addresses)
	}
	
	if State, ok := UserMap["state"].(string); ok {
		o.State = &State
	}
    
	if Title, ok := UserMap["title"].(string); ok {
		o.Title = &Title
	}
    
	if Username, ok := UserMap["username"].(string); ok {
		o.Username = &Username
	}
    
	if Manager, ok := UserMap["manager"].(map[string]interface{}); ok {
		ManagerString, _ := json.Marshal(Manager)
		json.Unmarshal(ManagerString, &o.Manager)
	}
	
	if Images, ok := UserMap["images"].([]interface{}); ok {
		ImagesString, _ := json.Marshal(Images)
		json.Unmarshal(ImagesString, &o.Images)
	}
	
	if Version, ok := UserMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	
	if Certifications, ok := UserMap["certifications"].([]interface{}); ok {
		CertificationsString, _ := json.Marshal(Certifications)
		json.Unmarshal(CertificationsString, &o.Certifications)
	}
	
	if Biography, ok := UserMap["biography"].(map[string]interface{}); ok {
		BiographyString, _ := json.Marshal(Biography)
		json.Unmarshal(BiographyString, &o.Biography)
	}
	
	if EmployerInfo, ok := UserMap["employerInfo"].(map[string]interface{}); ok {
		EmployerInfoString, _ := json.Marshal(EmployerInfo)
		json.Unmarshal(EmployerInfoString, &o.EmployerInfo)
	}
	
	if PreferredName, ok := UserMap["preferredName"].(string); ok {
		o.PreferredName = &PreferredName
	}
    
	if RoutingStatus, ok := UserMap["routingStatus"].(map[string]interface{}); ok {
		RoutingStatusString, _ := json.Marshal(RoutingStatus)
		json.Unmarshal(RoutingStatusString, &o.RoutingStatus)
	}
	
	if Presence, ok := UserMap["presence"].(map[string]interface{}); ok {
		PresenceString, _ := json.Marshal(Presence)
		json.Unmarshal(PresenceString, &o.Presence)
	}
	
	if IntegrationPresence, ok := UserMap["integrationPresence"].(map[string]interface{}); ok {
		IntegrationPresenceString, _ := json.Marshal(IntegrationPresence)
		json.Unmarshal(IntegrationPresenceString, &o.IntegrationPresence)
	}
	
	if ConversationSummary, ok := UserMap["conversationSummary"].(map[string]interface{}); ok {
		ConversationSummaryString, _ := json.Marshal(ConversationSummary)
		json.Unmarshal(ConversationSummaryString, &o.ConversationSummary)
	}
	
	if OutOfOffice, ok := UserMap["outOfOffice"].(map[string]interface{}); ok {
		OutOfOfficeString, _ := json.Marshal(OutOfOffice)
		json.Unmarshal(OutOfOfficeString, &o.OutOfOffice)
	}
	
	if Geolocation, ok := UserMap["geolocation"].(map[string]interface{}); ok {
		GeolocationString, _ := json.Marshal(Geolocation)
		json.Unmarshal(GeolocationString, &o.Geolocation)
	}
	
	if Station, ok := UserMap["station"].(map[string]interface{}); ok {
		StationString, _ := json.Marshal(Station)
		json.Unmarshal(StationString, &o.Station)
	}
	
	if Authorization, ok := UserMap["authorization"].(map[string]interface{}); ok {
		AuthorizationString, _ := json.Marshal(Authorization)
		json.Unmarshal(AuthorizationString, &o.Authorization)
	}
	
	if ProfileSkills, ok := UserMap["profileSkills"].([]interface{}); ok {
		ProfileSkillsString, _ := json.Marshal(ProfileSkills)
		json.Unmarshal(ProfileSkillsString, &o.ProfileSkills)
	}
	
	if Locations, ok := UserMap["locations"].([]interface{}); ok {
		LocationsString, _ := json.Marshal(Locations)
		json.Unmarshal(LocationsString, &o.Locations)
	}
	
	if Groups, ok := UserMap["groups"].([]interface{}); ok {
		GroupsString, _ := json.Marshal(Groups)
		json.Unmarshal(GroupsString, &o.Groups)
	}
	
	if Team, ok := UserMap["team"].(map[string]interface{}); ok {
		TeamString, _ := json.Marshal(Team)
		json.Unmarshal(TeamString, &o.Team)
	}
	
	if WorkPlanBidRanks, ok := UserMap["workPlanBidRanks"].(map[string]interface{}); ok {
		WorkPlanBidRanksString, _ := json.Marshal(WorkPlanBidRanks)
		json.Unmarshal(WorkPlanBidRanksString, &o.WorkPlanBidRanks)
	}
	
	if Skills, ok := UserMap["skills"].([]interface{}); ok {
		SkillsString, _ := json.Marshal(Skills)
		json.Unmarshal(SkillsString, &o.Skills)
	}
	
	if Languages, ok := UserMap["languages"].([]interface{}); ok {
		LanguagesString, _ := json.Marshal(Languages)
		json.Unmarshal(LanguagesString, &o.Languages)
	}
	
	if AcdAutoAnswer, ok := UserMap["acdAutoAnswer"].(bool); ok {
		o.AcdAutoAnswer = &AcdAutoAnswer
	}
    
	if LanguagePreference, ok := UserMap["languagePreference"].(string); ok {
		o.LanguagePreference = &LanguagePreference
	}
    
	if LastTokenIssued, ok := UserMap["lastTokenIssued"].(map[string]interface{}); ok {
		LastTokenIssuedString, _ := json.Marshal(LastTokenIssued)
		json.Unmarshal(LastTokenIssuedString, &o.LastTokenIssued)
	}
	
	if dateLastLoginString, ok := UserMap["dateLastLogin"].(string); ok {
		DateLastLogin, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateLastLoginString)
		o.DateLastLogin = &DateLastLogin
	}
	
	if dateWelcomeSentString, ok := UserMap["dateWelcomeSent"].(string); ok {
		DateWelcomeSent, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateWelcomeSentString)
		o.DateWelcomeSent = &DateWelcomeSent
	}
	
	if SelfUri, ok := UserMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *User) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

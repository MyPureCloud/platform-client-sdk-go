package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Trustuser
type Trustuser struct { 
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


	// TrustUserDetails
	TrustUserDetails *Trustuserdetails `json:"trustUserDetails,omitempty"`

}

// String returns a JSON representation of the model
func (o *Trustuser) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Updateuser
type Updateuser struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Chat
	Chat *Chat `json:"chat,omitempty"`


	// Department
	Department *string `json:"department,omitempty"`


	// Email
	Email *string `json:"email,omitempty"`


	// PrimaryContactInfo - The address(s) used for primary contact. Updates to the corresponding address in the addresses list will be reflected here.
	PrimaryContactInfo *[]Contact `json:"primaryContactInfo,omitempty"`


	// Addresses - Email address, phone number, and/or extension for this user. One entry is allowed per media type
	Addresses *[]Contact `json:"addresses,omitempty"`


	// Title
	Title *string `json:"title,omitempty"`


	// Username
	Username *string `json:"username,omitempty"`


	// Manager
	Manager *string `json:"manager,omitempty"`


	// Images
	Images *[]Userimage `json:"images,omitempty"`


	// Version - This value should be the current version of the user. The current version can be obtained with a GET on the user before doing a PATCH.
	Version *int `json:"version,omitempty"`


	// ProfileSkills - Profile skills possessed by the user
	ProfileSkills *[]string `json:"profileSkills,omitempty"`


	// Locations - The user placement at each site location.
	Locations *[]Location `json:"locations,omitempty"`


	// Groups - The groups the user is a member of
	Groups *[]Group `json:"groups,omitempty"`


	// State - The state of the user. This property can be used to restore a deleted user or transition between active and inactive. If specified, it is the only modifiable field.
	State *string `json:"state,omitempty"`


	// AcdAutoAnswer - The value that denotes if acdAutoAnswer is set on the user
	AcdAutoAnswer *bool `json:"acdAutoAnswer,omitempty"`


	// Certifications
	Certifications *[]string `json:"certifications,omitempty"`


	// Biography
	Biography *Biography `json:"biography,omitempty"`


	// EmployerInfo
	EmployerInfo *Employerinfo `json:"employerInfo,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (u *Updateuser) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Updateuser

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Chat *Chat `json:"chat,omitempty"`
		
		Department *string `json:"department,omitempty"`
		
		Email *string `json:"email,omitempty"`
		
		PrimaryContactInfo *[]Contact `json:"primaryContactInfo,omitempty"`
		
		Addresses *[]Contact `json:"addresses,omitempty"`
		
		Title *string `json:"title,omitempty"`
		
		Username *string `json:"username,omitempty"`
		
		Manager *string `json:"manager,omitempty"`
		
		Images *[]Userimage `json:"images,omitempty"`
		
		Version *int `json:"version,omitempty"`
		
		ProfileSkills *[]string `json:"profileSkills,omitempty"`
		
		Locations *[]Location `json:"locations,omitempty"`
		
		Groups *[]Group `json:"groups,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		AcdAutoAnswer *bool `json:"acdAutoAnswer,omitempty"`
		
		Certifications *[]string `json:"certifications,omitempty"`
		
		Biography *Biography `json:"biography,omitempty"`
		
		EmployerInfo *Employerinfo `json:"employerInfo,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		Chat: u.Chat,
		
		Department: u.Department,
		
		Email: u.Email,
		
		PrimaryContactInfo: u.PrimaryContactInfo,
		
		Addresses: u.Addresses,
		
		Title: u.Title,
		
		Username: u.Username,
		
		Manager: u.Manager,
		
		Images: u.Images,
		
		Version: u.Version,
		
		ProfileSkills: u.ProfileSkills,
		
		Locations: u.Locations,
		
		Groups: u.Groups,
		
		State: u.State,
		
		AcdAutoAnswer: u.AcdAutoAnswer,
		
		Certifications: u.Certifications,
		
		Biography: u.Biography,
		
		EmployerInfo: u.EmployerInfo,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Updateuser) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

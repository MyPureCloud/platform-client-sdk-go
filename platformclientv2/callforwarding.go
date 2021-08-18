package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Callforwarding
type Callforwarding struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// User
	User *User `json:"user,omitempty"`


	// Enabled - Whether or not CallForwarding is enabled
	Enabled *bool `json:"enabled,omitempty"`


	// PhoneNumber - This property is deprecated. Please use the calls property
	PhoneNumber *string `json:"phoneNumber,omitempty"`


	// Calls - An ordered list of CallRoutes to be executed when CallForwarding is enabled
	Calls *[]Callroute `json:"calls,omitempty"`


	// Voicemail - The type of voicemail to use with the callForwarding configuration
	Voicemail *string `json:"voicemail,omitempty"`


	// ModifiedDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ModifiedDate *time.Time `json:"modifiedDate,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (u *Callforwarding) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Callforwarding

	
	ModifiedDate := new(string)
	if u.ModifiedDate != nil {
		
		*ModifiedDate = timeutil.Strftime(u.ModifiedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ModifiedDate = nil
	}
	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		User *User `json:"user,omitempty"`
		
		Enabled *bool `json:"enabled,omitempty"`
		
		PhoneNumber *string `json:"phoneNumber,omitempty"`
		
		Calls *[]Callroute `json:"calls,omitempty"`
		
		Voicemail *string `json:"voicemail,omitempty"`
		
		ModifiedDate *string `json:"modifiedDate,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		User: u.User,
		
		Enabled: u.Enabled,
		
		PhoneNumber: u.PhoneNumber,
		
		Calls: u.Calls,
		
		Voicemail: u.Voicemail,
		
		ModifiedDate: ModifiedDate,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Callforwarding) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Userstation
type Userstation struct { 
	// Id - A globally unique identifier for this station
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// VarType
	VarType *string `json:"type,omitempty"`


	// AssociatedUser
	AssociatedUser *User `json:"associatedUser,omitempty"`


	// AssociatedDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	AssociatedDate *time.Time `json:"associatedDate,omitempty"`


	// DefaultUser
	DefaultUser *User `json:"defaultUser,omitempty"`


	// ProviderInfo - Provider-specific info for this station, e.g. { \"edgeGroupId\": \"ffe7b15c-a9cc-4f4c-88f5-781327819a49\" }
	ProviderInfo *map[string]string `json:"providerInfo,omitempty"`

}

func (u *Userstation) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Userstation

	
	AssociatedDate := new(string)
	if u.AssociatedDate != nil {
		
		*AssociatedDate = timeutil.Strftime(u.AssociatedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		AssociatedDate = nil
	}
	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		AssociatedUser *User `json:"associatedUser,omitempty"`
		
		AssociatedDate *string `json:"associatedDate,omitempty"`
		
		DefaultUser *User `json:"defaultUser,omitempty"`
		
		ProviderInfo *map[string]string `json:"providerInfo,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		VarType: u.VarType,
		
		AssociatedUser: u.AssociatedUser,
		
		AssociatedDate: AssociatedDate,
		
		DefaultUser: u.DefaultUser,
		
		ProviderInfo: u.ProviderInfo,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Userstation) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

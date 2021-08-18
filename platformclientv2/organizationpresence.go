package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Organizationpresence
type Organizationpresence struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// LanguageLabels - The label used for the system presence in each specified language
	LanguageLabels *map[string]string `json:"languageLabels,omitempty"`


	// SystemPresence
	SystemPresence *string `json:"systemPresence,omitempty"`


	// Deactivated
	Deactivated *bool `json:"deactivated,omitempty"`


	// Primary
	Primary *bool `json:"primary,omitempty"`


	// CreatedBy
	CreatedBy *User `json:"createdBy,omitempty"`


	// CreatedDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	CreatedDate *time.Time `json:"createdDate,omitempty"`


	// ModifiedBy
	ModifiedBy *User `json:"modifiedBy,omitempty"`


	// ModifiedDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ModifiedDate *time.Time `json:"modifiedDate,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (u *Organizationpresence) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Organizationpresence

	
	CreatedDate := new(string)
	if u.CreatedDate != nil {
		
		*CreatedDate = timeutil.Strftime(u.CreatedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		CreatedDate = nil
	}
	
	ModifiedDate := new(string)
	if u.ModifiedDate != nil {
		
		*ModifiedDate = timeutil.Strftime(u.ModifiedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ModifiedDate = nil
	}
	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		LanguageLabels *map[string]string `json:"languageLabels,omitempty"`
		
		SystemPresence *string `json:"systemPresence,omitempty"`
		
		Deactivated *bool `json:"deactivated,omitempty"`
		
		Primary *bool `json:"primary,omitempty"`
		
		CreatedBy *User `json:"createdBy,omitempty"`
		
		CreatedDate *string `json:"createdDate,omitempty"`
		
		ModifiedBy *User `json:"modifiedBy,omitempty"`
		
		ModifiedDate *string `json:"modifiedDate,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		LanguageLabels: u.LanguageLabels,
		
		SystemPresence: u.SystemPresence,
		
		Deactivated: u.Deactivated,
		
		Primary: u.Primary,
		
		CreatedBy: u.CreatedBy,
		
		CreatedDate: CreatedDate,
		
		ModifiedBy: u.ModifiedBy,
		
		ModifiedDate: ModifiedDate,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Organizationpresence) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

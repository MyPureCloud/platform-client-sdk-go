package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Trustee
type Trustee struct { 
	// Id - Organization Id for this trust.
	Id *string `json:"id,omitempty"`


	// Enabled - If disabled no trustee user will have access, even if they were previously added.
	Enabled *bool `json:"enabled,omitempty"`


	// UsesDefaultRole - Denotes if trustee uses admin role by default.
	UsesDefaultRole *bool `json:"usesDefaultRole,omitempty"`


	// DateCreated - Date Trust was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateExpired - The expiration date of the trust. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateExpired *time.Time `json:"dateExpired,omitempty"`


	// CreatedBy - User that created trust.
	CreatedBy *Orguser `json:"createdBy,omitempty"`


	// Organization - Organization associated with this trust.
	Organization *Organization `json:"organization,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (u *Trustee) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Trustee

	
	DateCreated := new(string)
	if u.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(u.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	DateExpired := new(string)
	if u.DateExpired != nil {
		
		*DateExpired = timeutil.Strftime(u.DateExpired, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateExpired = nil
	}
	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Enabled *bool `json:"enabled,omitempty"`
		
		UsesDefaultRole *bool `json:"usesDefaultRole,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateExpired *string `json:"dateExpired,omitempty"`
		
		CreatedBy *Orguser `json:"createdBy,omitempty"`
		
		Organization *Organization `json:"organization,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Enabled: u.Enabled,
		
		UsesDefaultRole: u.UsesDefaultRole,
		
		DateCreated: DateCreated,
		
		DateExpired: DateExpired,
		
		CreatedBy: u.CreatedBy,
		
		Organization: u.Organization,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Trustee) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

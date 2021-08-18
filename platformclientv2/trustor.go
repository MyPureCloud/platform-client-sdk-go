package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Trustor
type Trustor struct { 
	// Id - Organization Id for this trust.
	Id *string `json:"id,omitempty"`


	// Enabled - If disabled no trustee user will have access, even if they were previously added.
	Enabled *bool `json:"enabled,omitempty"`


	// DateCreated - Date Trust was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// CreatedBy - User that created trust.
	CreatedBy *Orguser `json:"createdBy,omitempty"`


	// Organization - Organization associated with this trust.
	Organization *Organization `json:"organization,omitempty"`


	// Authorization - Authorization for the trustee user has in this trustor organization
	Authorization *Trusteeauthorization `json:"authorization,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (u *Trustor) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Trustor

	
	DateCreated := new(string)
	if u.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(u.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Enabled *bool `json:"enabled,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		CreatedBy *Orguser `json:"createdBy,omitempty"`
		
		Organization *Organization `json:"organization,omitempty"`
		
		Authorization *Trusteeauthorization `json:"authorization,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Enabled: u.Enabled,
		
		DateCreated: DateCreated,
		
		CreatedBy: u.CreatedBy,
		
		Organization: u.Organization,
		
		Authorization: u.Authorization,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Trustor) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Externalorganizationtrustorlink
type Externalorganizationtrustorlink struct { 
	// ExternalOrganizationId - The id of a PureCloud External Organization entity in the External Contacts system that will be used to represent the trustor org
	ExternalOrganizationId *string `json:"externalOrganizationId,omitempty"`


	// TrustorOrgId - The id of a PureCloud organization that has granted trust to this PureCloud organization
	TrustorOrgId *string `json:"trustorOrgId,omitempty"`


	// DateCreated - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// ExternalOrganizationUri - The URI for the External Organization that is linked to the trustor org
	ExternalOrganizationUri *string `json:"externalOrganizationUri,omitempty"`

}

func (u *Externalorganizationtrustorlink) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Externalorganizationtrustorlink

	
	DateCreated := new(string)
	if u.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(u.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	

	return json.Marshal(&struct { 
		ExternalOrganizationId *string `json:"externalOrganizationId,omitempty"`
		
		TrustorOrgId *string `json:"trustorOrgId,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		ExternalOrganizationUri *string `json:"externalOrganizationUri,omitempty"`
		*Alias
	}{ 
		ExternalOrganizationId: u.ExternalOrganizationId,
		
		TrustorOrgId: u.TrustorOrgId,
		
		DateCreated: DateCreated,
		
		ExternalOrganizationUri: u.ExternalOrganizationUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Externalorganizationtrustorlink) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

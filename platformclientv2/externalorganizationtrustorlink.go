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

func (o *Externalorganizationtrustorlink) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Externalorganizationtrustorlink
	
	DateCreated := new(string)
	if o.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(o.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
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
		ExternalOrganizationId: o.ExternalOrganizationId,
		
		TrustorOrgId: o.TrustorOrgId,
		
		DateCreated: DateCreated,
		
		ExternalOrganizationUri: o.ExternalOrganizationUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Externalorganizationtrustorlink) UnmarshalJSON(b []byte) error {
	var ExternalorganizationtrustorlinkMap map[string]interface{}
	err := json.Unmarshal(b, &ExternalorganizationtrustorlinkMap)
	if err != nil {
		return err
	}
	
	if ExternalOrganizationId, ok := ExternalorganizationtrustorlinkMap["externalOrganizationId"].(string); ok {
		o.ExternalOrganizationId = &ExternalOrganizationId
	}
    
	if TrustorOrgId, ok := ExternalorganizationtrustorlinkMap["trustorOrgId"].(string); ok {
		o.TrustorOrgId = &TrustorOrgId
	}
    
	if dateCreatedString, ok := ExternalorganizationtrustorlinkMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if ExternalOrganizationUri, ok := ExternalorganizationtrustorlinkMap["externalOrganizationUri"].(string); ok {
		o.ExternalOrganizationUri = &ExternalOrganizationUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Externalorganizationtrustorlink) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

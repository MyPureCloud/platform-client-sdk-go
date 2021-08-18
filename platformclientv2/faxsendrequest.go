package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Faxsendrequest
type Faxsendrequest struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Addresses - A list of outbound fax dialing addresses. E.g. +13175555555 or 3175555555
	Addresses *[]string `json:"addresses,omitempty"`


	// DocumentId - DocumentId of Content Management artifact. If Content Management document is not used for faxing, documentId should be null
	DocumentId *string `json:"documentId,omitempty"`


	// ContentType - The content type that is going to be uploaded. If Content Management document is used for faxing, contentType will be ignored
	ContentType *string `json:"contentType,omitempty"`


	// Workspace - Workspace in which the document should be stored. If Content Management document is used for faxing, workspace will be ignored
	Workspace *Workspace `json:"workspace,omitempty"`


	// CoverSheet - Data for coversheet generation.
	CoverSheet *Coversheet `json:"coverSheet,omitempty"`


	// TimeZoneOffsetMinutes - Time zone offset minutes from GMT
	TimeZoneOffsetMinutes *int `json:"timeZoneOffsetMinutes,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (u *Faxsendrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Faxsendrequest

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Addresses *[]string `json:"addresses,omitempty"`
		
		DocumentId *string `json:"documentId,omitempty"`
		
		ContentType *string `json:"contentType,omitempty"`
		
		Workspace *Workspace `json:"workspace,omitempty"`
		
		CoverSheet *Coversheet `json:"coverSheet,omitempty"`
		
		TimeZoneOffsetMinutes *int `json:"timeZoneOffsetMinutes,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		Addresses: u.Addresses,
		
		DocumentId: u.DocumentId,
		
		ContentType: u.ContentType,
		
		Workspace: u.Workspace,
		
		CoverSheet: u.CoverSheet,
		
		TimeZoneOffsetMinutes: u.TimeZoneOffsetMinutes,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Faxsendrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

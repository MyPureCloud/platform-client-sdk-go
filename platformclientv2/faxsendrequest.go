package platformclientv2
import (
	"encoding/json"
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
	TimeZoneOffsetMinutes *int32 `json:"timeZoneOffsetMinutes,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Faxsendrequest) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

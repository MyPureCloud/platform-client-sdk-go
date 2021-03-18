package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Externalcontact
type Externalcontact struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// FirstName - The first name of the contact.
	FirstName *string `json:"firstName,omitempty"`


	// MiddleName
	MiddleName *string `json:"middleName,omitempty"`


	// LastName - The last name of the contact.
	LastName *string `json:"lastName,omitempty"`


	// Salutation
	Salutation *string `json:"salutation,omitempty"`


	// Title
	Title *string `json:"title,omitempty"`


	// WorkPhone
	WorkPhone *Phonenumber `json:"workPhone,omitempty"`


	// CellPhone
	CellPhone *Phonenumber `json:"cellPhone,omitempty"`


	// HomePhone
	HomePhone *Phonenumber `json:"homePhone,omitempty"`


	// OtherPhone
	OtherPhone *Phonenumber `json:"otherPhone,omitempty"`


	// WorkEmail
	WorkEmail *string `json:"workEmail,omitempty"`


	// PersonalEmail
	PersonalEmail *string `json:"personalEmail,omitempty"`


	// OtherEmail
	OtherEmail *string `json:"otherEmail,omitempty"`


	// Address
	Address *Contactaddress `json:"address,omitempty"`


	// TwitterId
	TwitterId *Twitterid `json:"twitterId,omitempty"`


	// LineId
	LineId *Lineid `json:"lineId,omitempty"`


	// WhatsAppId
	WhatsAppId *Whatsappid `json:"whatsAppId,omitempty"`


	// FacebookId
	FacebookId *Facebookid `json:"facebookId,omitempty"`


	// ModifyDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ModifyDate *time.Time `json:"modifyDate,omitempty"`


	// CreateDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	CreateDate *time.Time `json:"createDate,omitempty"`


	// ExternalOrganization
	ExternalOrganization *Externalorganization `json:"externalOrganization,omitempty"`


	// SurveyOptOut
	SurveyOptOut *bool `json:"surveyOptOut,omitempty"`


	// ExternalSystemUrl - A string that identifies an external system-of-record resource that may have more detailed information on the contact. It should be a valid URL (including the http/https protocol, port, and path [if any]). The value is automatically trimmed of any leading and trailing whitespace.
	ExternalSystemUrl *string `json:"externalSystemUrl,omitempty"`


	// Schema - The schema defining custom fields for this contact
	Schema *Dataschema `json:"schema,omitempty"`


	// CustomFields - Custom fields defined in the schema referenced by schemaId and schemaVersion.
	CustomFields *map[string]interface{} `json:"customFields,omitempty"`


	// ExternalDataSources - Links to the sources of data (e.g. one source might be a CRM) that contributed data to this record.  Read-only, and only populated when requested via expand param.
	ExternalDataSources *[]Externaldatasource `json:"externalDataSources,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Externalcontact) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

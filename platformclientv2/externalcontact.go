package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
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

func (u *Externalcontact) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Externalcontact

	
	ModifyDate := new(string)
	if u.ModifyDate != nil {
		
		*ModifyDate = timeutil.Strftime(u.ModifyDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ModifyDate = nil
	}
	
	CreateDate := new(string)
	if u.CreateDate != nil {
		
		*CreateDate = timeutil.Strftime(u.CreateDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		CreateDate = nil
	}
	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		FirstName *string `json:"firstName,omitempty"`
		
		MiddleName *string `json:"middleName,omitempty"`
		
		LastName *string `json:"lastName,omitempty"`
		
		Salutation *string `json:"salutation,omitempty"`
		
		Title *string `json:"title,omitempty"`
		
		WorkPhone *Phonenumber `json:"workPhone,omitempty"`
		
		CellPhone *Phonenumber `json:"cellPhone,omitempty"`
		
		HomePhone *Phonenumber `json:"homePhone,omitempty"`
		
		OtherPhone *Phonenumber `json:"otherPhone,omitempty"`
		
		WorkEmail *string `json:"workEmail,omitempty"`
		
		PersonalEmail *string `json:"personalEmail,omitempty"`
		
		OtherEmail *string `json:"otherEmail,omitempty"`
		
		Address *Contactaddress `json:"address,omitempty"`
		
		TwitterId *Twitterid `json:"twitterId,omitempty"`
		
		LineId *Lineid `json:"lineId,omitempty"`
		
		WhatsAppId *Whatsappid `json:"whatsAppId,omitempty"`
		
		FacebookId *Facebookid `json:"facebookId,omitempty"`
		
		ModifyDate *string `json:"modifyDate,omitempty"`
		
		CreateDate *string `json:"createDate,omitempty"`
		
		ExternalOrganization *Externalorganization `json:"externalOrganization,omitempty"`
		
		SurveyOptOut *bool `json:"surveyOptOut,omitempty"`
		
		ExternalSystemUrl *string `json:"externalSystemUrl,omitempty"`
		
		Schema *Dataschema `json:"schema,omitempty"`
		
		CustomFields *map[string]interface{} `json:"customFields,omitempty"`
		
		ExternalDataSources *[]Externaldatasource `json:"externalDataSources,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		FirstName: u.FirstName,
		
		MiddleName: u.MiddleName,
		
		LastName: u.LastName,
		
		Salutation: u.Salutation,
		
		Title: u.Title,
		
		WorkPhone: u.WorkPhone,
		
		CellPhone: u.CellPhone,
		
		HomePhone: u.HomePhone,
		
		OtherPhone: u.OtherPhone,
		
		WorkEmail: u.WorkEmail,
		
		PersonalEmail: u.PersonalEmail,
		
		OtherEmail: u.OtherEmail,
		
		Address: u.Address,
		
		TwitterId: u.TwitterId,
		
		LineId: u.LineId,
		
		WhatsAppId: u.WhatsAppId,
		
		FacebookId: u.FacebookId,
		
		ModifyDate: ModifyDate,
		
		CreateDate: CreateDate,
		
		ExternalOrganization: u.ExternalOrganization,
		
		SurveyOptOut: u.SurveyOptOut,
		
		ExternalSystemUrl: u.ExternalSystemUrl,
		
		Schema: u.Schema,
		
		CustomFields: u.CustomFields,
		
		ExternalDataSources: u.ExternalDataSources,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Externalcontact) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

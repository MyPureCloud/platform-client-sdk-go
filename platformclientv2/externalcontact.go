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

func (o *Externalcontact) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Externalcontact
	
	ModifyDate := new(string)
	if o.ModifyDate != nil {
		
		*ModifyDate = timeutil.Strftime(o.ModifyDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ModifyDate = nil
	}
	
	CreateDate := new(string)
	if o.CreateDate != nil {
		
		*CreateDate = timeutil.Strftime(o.CreateDate, "%Y-%m-%dT%H:%M:%S.%fZ")
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
		Id: o.Id,
		
		FirstName: o.FirstName,
		
		MiddleName: o.MiddleName,
		
		LastName: o.LastName,
		
		Salutation: o.Salutation,
		
		Title: o.Title,
		
		WorkPhone: o.WorkPhone,
		
		CellPhone: o.CellPhone,
		
		HomePhone: o.HomePhone,
		
		OtherPhone: o.OtherPhone,
		
		WorkEmail: o.WorkEmail,
		
		PersonalEmail: o.PersonalEmail,
		
		OtherEmail: o.OtherEmail,
		
		Address: o.Address,
		
		TwitterId: o.TwitterId,
		
		LineId: o.LineId,
		
		WhatsAppId: o.WhatsAppId,
		
		FacebookId: o.FacebookId,
		
		ModifyDate: ModifyDate,
		
		CreateDate: CreateDate,
		
		ExternalOrganization: o.ExternalOrganization,
		
		SurveyOptOut: o.SurveyOptOut,
		
		ExternalSystemUrl: o.ExternalSystemUrl,
		
		Schema: o.Schema,
		
		CustomFields: o.CustomFields,
		
		ExternalDataSources: o.ExternalDataSources,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Externalcontact) UnmarshalJSON(b []byte) error {
	var ExternalcontactMap map[string]interface{}
	err := json.Unmarshal(b, &ExternalcontactMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ExternalcontactMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if FirstName, ok := ExternalcontactMap["firstName"].(string); ok {
		o.FirstName = &FirstName
	}
	
	if MiddleName, ok := ExternalcontactMap["middleName"].(string); ok {
		o.MiddleName = &MiddleName
	}
	
	if LastName, ok := ExternalcontactMap["lastName"].(string); ok {
		o.LastName = &LastName
	}
	
	if Salutation, ok := ExternalcontactMap["salutation"].(string); ok {
		o.Salutation = &Salutation
	}
	
	if Title, ok := ExternalcontactMap["title"].(string); ok {
		o.Title = &Title
	}
	
	if WorkPhone, ok := ExternalcontactMap["workPhone"].(map[string]interface{}); ok {
		WorkPhoneString, _ := json.Marshal(WorkPhone)
		json.Unmarshal(WorkPhoneString, &o.WorkPhone)
	}
	
	if CellPhone, ok := ExternalcontactMap["cellPhone"].(map[string]interface{}); ok {
		CellPhoneString, _ := json.Marshal(CellPhone)
		json.Unmarshal(CellPhoneString, &o.CellPhone)
	}
	
	if HomePhone, ok := ExternalcontactMap["homePhone"].(map[string]interface{}); ok {
		HomePhoneString, _ := json.Marshal(HomePhone)
		json.Unmarshal(HomePhoneString, &o.HomePhone)
	}
	
	if OtherPhone, ok := ExternalcontactMap["otherPhone"].(map[string]interface{}); ok {
		OtherPhoneString, _ := json.Marshal(OtherPhone)
		json.Unmarshal(OtherPhoneString, &o.OtherPhone)
	}
	
	if WorkEmail, ok := ExternalcontactMap["workEmail"].(string); ok {
		o.WorkEmail = &WorkEmail
	}
	
	if PersonalEmail, ok := ExternalcontactMap["personalEmail"].(string); ok {
		o.PersonalEmail = &PersonalEmail
	}
	
	if OtherEmail, ok := ExternalcontactMap["otherEmail"].(string); ok {
		o.OtherEmail = &OtherEmail
	}
	
	if Address, ok := ExternalcontactMap["address"].(map[string]interface{}); ok {
		AddressString, _ := json.Marshal(Address)
		json.Unmarshal(AddressString, &o.Address)
	}
	
	if TwitterId, ok := ExternalcontactMap["twitterId"].(map[string]interface{}); ok {
		TwitterIdString, _ := json.Marshal(TwitterId)
		json.Unmarshal(TwitterIdString, &o.TwitterId)
	}
	
	if LineId, ok := ExternalcontactMap["lineId"].(map[string]interface{}); ok {
		LineIdString, _ := json.Marshal(LineId)
		json.Unmarshal(LineIdString, &o.LineId)
	}
	
	if WhatsAppId, ok := ExternalcontactMap["whatsAppId"].(map[string]interface{}); ok {
		WhatsAppIdString, _ := json.Marshal(WhatsAppId)
		json.Unmarshal(WhatsAppIdString, &o.WhatsAppId)
	}
	
	if FacebookId, ok := ExternalcontactMap["facebookId"].(map[string]interface{}); ok {
		FacebookIdString, _ := json.Marshal(FacebookId)
		json.Unmarshal(FacebookIdString, &o.FacebookId)
	}
	
	if modifyDateString, ok := ExternalcontactMap["modifyDate"].(string); ok {
		ModifyDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", modifyDateString)
		o.ModifyDate = &ModifyDate
	}
	
	if createDateString, ok := ExternalcontactMap["createDate"].(string); ok {
		CreateDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", createDateString)
		o.CreateDate = &CreateDate
	}
	
	if ExternalOrganization, ok := ExternalcontactMap["externalOrganization"].(map[string]interface{}); ok {
		ExternalOrganizationString, _ := json.Marshal(ExternalOrganization)
		json.Unmarshal(ExternalOrganizationString, &o.ExternalOrganization)
	}
	
	if SurveyOptOut, ok := ExternalcontactMap["surveyOptOut"].(bool); ok {
		o.SurveyOptOut = &SurveyOptOut
	}
	
	if ExternalSystemUrl, ok := ExternalcontactMap["externalSystemUrl"].(string); ok {
		o.ExternalSystemUrl = &ExternalSystemUrl
	}
	
	if Schema, ok := ExternalcontactMap["schema"].(map[string]interface{}); ok {
		SchemaString, _ := json.Marshal(Schema)
		json.Unmarshal(SchemaString, &o.Schema)
	}
	
	if CustomFields, ok := ExternalcontactMap["customFields"].(map[string]interface{}); ok {
		CustomFieldsString, _ := json.Marshal(CustomFields)
		json.Unmarshal(CustomFieldsString, &o.CustomFields)
	}
	
	if ExternalDataSources, ok := ExternalcontactMap["externalDataSources"].([]interface{}); ok {
		ExternalDataSourcesString, _ := json.Marshal(ExternalDataSources)
		json.Unmarshal(ExternalDataSourcesString, &o.ExternalDataSources)
	}
	
	if SelfUri, ok := ExternalcontactMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Externalcontact) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

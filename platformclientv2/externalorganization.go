package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Externalorganization
type Externalorganization struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name - The name of the company.
	Name *string `json:"name,omitempty"`


	// CompanyType
	CompanyType *string `json:"companyType,omitempty"`


	// Industry
	Industry *string `json:"industry,omitempty"`


	// PrimaryContactId
	PrimaryContactId *string `json:"primaryContactId,omitempty"`


	// Address
	Address *Contactaddress `json:"address,omitempty"`


	// PhoneNumber
	PhoneNumber *Phonenumber `json:"phoneNumber,omitempty"`


	// FaxNumber
	FaxNumber *Phonenumber `json:"faxNumber,omitempty"`


	// EmployeeCount
	EmployeeCount *int `json:"employeeCount,omitempty"`


	// Revenue
	Revenue *int `json:"revenue,omitempty"`


	// Tags
	Tags *[]string `json:"tags,omitempty"`


	// Websites
	Websites *[]string `json:"websites,omitempty"`


	// Tickers
	Tickers *[]Ticker `json:"tickers,omitempty"`


	// TwitterId
	TwitterId *Twitterid `json:"twitterId,omitempty"`


	// ExternalSystemUrl - A string that identifies an external system-of-record resource that may have more detailed information on the organization. It should be a valid URL (including the http/https protocol, port, and path [if any]). The value is automatically trimmed of any leading and trailing whitespace.
	ExternalSystemUrl *string `json:"externalSystemUrl,omitempty"`


	// ModifyDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ModifyDate *time.Time `json:"modifyDate,omitempty"`


	// CreateDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	CreateDate *time.Time `json:"createDate,omitempty"`


	// Trustor
	Trustor *Trustor `json:"trustor,omitempty"`


	// Schema - The schema defining custom fields for this contact
	Schema *Dataschema `json:"schema,omitempty"`


	// CustomFields - Custom fields defined in the schema referenced by schemaId and schemaVersion.
	CustomFields *map[string]interface{} `json:"customFields,omitempty"`


	// ExternalDataSources - Links to the sources of data (e.g. one source might be a CRM) that contributed data to this record.  Read-only, and only populated when requested via expand param.
	ExternalDataSources *[]Externaldatasource `json:"externalDataSources,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (u *Externalorganization) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Externalorganization

	
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
		
		Name *string `json:"name,omitempty"`
		
		CompanyType *string `json:"companyType,omitempty"`
		
		Industry *string `json:"industry,omitempty"`
		
		PrimaryContactId *string `json:"primaryContactId,omitempty"`
		
		Address *Contactaddress `json:"address,omitempty"`
		
		PhoneNumber *Phonenumber `json:"phoneNumber,omitempty"`
		
		FaxNumber *Phonenumber `json:"faxNumber,omitempty"`
		
		EmployeeCount *int `json:"employeeCount,omitempty"`
		
		Revenue *int `json:"revenue,omitempty"`
		
		Tags *[]string `json:"tags,omitempty"`
		
		Websites *[]string `json:"websites,omitempty"`
		
		Tickers *[]Ticker `json:"tickers,omitempty"`
		
		TwitterId *Twitterid `json:"twitterId,omitempty"`
		
		ExternalSystemUrl *string `json:"externalSystemUrl,omitempty"`
		
		ModifyDate *string `json:"modifyDate,omitempty"`
		
		CreateDate *string `json:"createDate,omitempty"`
		
		Trustor *Trustor `json:"trustor,omitempty"`
		
		Schema *Dataschema `json:"schema,omitempty"`
		
		CustomFields *map[string]interface{} `json:"customFields,omitempty"`
		
		ExternalDataSources *[]Externaldatasource `json:"externalDataSources,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		CompanyType: u.CompanyType,
		
		Industry: u.Industry,
		
		PrimaryContactId: u.PrimaryContactId,
		
		Address: u.Address,
		
		PhoneNumber: u.PhoneNumber,
		
		FaxNumber: u.FaxNumber,
		
		EmployeeCount: u.EmployeeCount,
		
		Revenue: u.Revenue,
		
		Tags: u.Tags,
		
		Websites: u.Websites,
		
		Tickers: u.Tickers,
		
		TwitterId: u.TwitterId,
		
		ExternalSystemUrl: u.ExternalSystemUrl,
		
		ModifyDate: ModifyDate,
		
		CreateDate: CreateDate,
		
		Trustor: u.Trustor,
		
		Schema: u.Schema,
		
		CustomFields: u.CustomFields,
		
		ExternalDataSources: u.ExternalDataSources,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Externalorganization) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Externalorganization
type Externalorganization struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name - The name of the company. Max: 1000 characters. Leading and trailing whitespace stripped.
	Name *string `json:"name,omitempty"`

	// Division - The division to which this entity belongs.
	Division *Writablestarrabledivision `json:"division,omitempty"`

	// CompanyType
	CompanyType *string `json:"companyType,omitempty"`

	// Industry
	Industry *string `json:"industry,omitempty"`

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

	// Identifiers - Identifiers claimed by this external org
	Identifiers *[]Externalorganizationidentifier `json:"identifiers,omitempty"`

	// ExternalIds - A list of external identifiers that identify this External Organization in an external system. Max 10 items.
	ExternalIds *[]Externalid `json:"externalIds,omitempty"`

	// ExternalDataSources - Links to the sources of data (e.g. one source might be a CRM) that contributed data to this record.  Read-only, and only populated when requested via expand param.
	ExternalDataSources *[]Externaldatasource `json:"externalDataSources,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Externalorganization) SetField(field string, fieldValue interface{}) {
	// Get Value object for field
	target := reflect.ValueOf(o)
	targetField := reflect.Indirect(target).FieldByName(field)

	// Set value
	if fieldValue != nil {
		targetField.Set(reflect.ValueOf(fieldValue))
	} else {
		// Must create a new Value (creates **type) then get its element (*type), which will be nil pointer of the appropriate type
		x := reflect.Indirect(reflect.New(targetField.Type()))
		targetField.Set(x)
	}

	// Add field to set field names list
	if o.SetFieldNames == nil {
		o.SetFieldNames = make(map[string]bool)
	}
	o.SetFieldNames[field] = true
}

func (o Externalorganization) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "ModifyDate","CreateDate", }
		localDateTimeFields := []string{  }
		dateFields := []string{  }

		// Construct object
		newObj := make(map[string]interface{})
		for fieldName := range o.SetFieldNames {
			// Get initial field value
			fieldValue := val.FieldByName(fieldName).Interface()

			// Apply value formatting overrides
			if fieldValue == nil || reflect.ValueOf(fieldValue).IsNil()  {
				// Do nothing. Just catching this case to avoid trying to custom serialize a nil value
			} else if contains(dateTimeFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%dT%H:%M:%S.%fZ")
			} else if contains(localDateTimeFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%dT%H:%M:%S.%f")
			} else if contains(dateFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%d")
			}

			// Assign value to field using JSON tag name
			newObj[getFieldName(reflect.TypeOf(&o), fieldName)] = fieldValue
		}

		// Marshal and return dynamically constructed interface
		return json.Marshal(newObj)
	}

	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Externalorganization
	
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
		
		Name *string `json:"name,omitempty"`
		
		Division *Writablestarrabledivision `json:"division,omitempty"`
		
		CompanyType *string `json:"companyType,omitempty"`
		
		Industry *string `json:"industry,omitempty"`
		
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
		
		Identifiers *[]Externalorganizationidentifier `json:"identifiers,omitempty"`
		
		ExternalIds *[]Externalid `json:"externalIds,omitempty"`
		
		ExternalDataSources *[]Externaldatasource `json:"externalDataSources,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Division: o.Division,
		
		CompanyType: o.CompanyType,
		
		Industry: o.Industry,
		
		Address: o.Address,
		
		PhoneNumber: o.PhoneNumber,
		
		FaxNumber: o.FaxNumber,
		
		EmployeeCount: o.EmployeeCount,
		
		Revenue: o.Revenue,
		
		Tags: o.Tags,
		
		Websites: o.Websites,
		
		Tickers: o.Tickers,
		
		TwitterId: o.TwitterId,
		
		ExternalSystemUrl: o.ExternalSystemUrl,
		
		ModifyDate: ModifyDate,
		
		CreateDate: CreateDate,
		
		Trustor: o.Trustor,
		
		Schema: o.Schema,
		
		CustomFields: o.CustomFields,
		
		Identifiers: o.Identifiers,
		
		ExternalIds: o.ExternalIds,
		
		ExternalDataSources: o.ExternalDataSources,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Externalorganization) UnmarshalJSON(b []byte) error {
	var ExternalorganizationMap map[string]interface{}
	err := json.Unmarshal(b, &ExternalorganizationMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ExternalorganizationMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := ExternalorganizationMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Division, ok := ExternalorganizationMap["division"].(map[string]interface{}); ok {
		DivisionString, _ := json.Marshal(Division)
		json.Unmarshal(DivisionString, &o.Division)
	}
	
	if CompanyType, ok := ExternalorganizationMap["companyType"].(string); ok {
		o.CompanyType = &CompanyType
	}
    
	if Industry, ok := ExternalorganizationMap["industry"].(string); ok {
		o.Industry = &Industry
	}
    
	if Address, ok := ExternalorganizationMap["address"].(map[string]interface{}); ok {
		AddressString, _ := json.Marshal(Address)
		json.Unmarshal(AddressString, &o.Address)
	}
	
	if PhoneNumber, ok := ExternalorganizationMap["phoneNumber"].(map[string]interface{}); ok {
		PhoneNumberString, _ := json.Marshal(PhoneNumber)
		json.Unmarshal(PhoneNumberString, &o.PhoneNumber)
	}
	
	if FaxNumber, ok := ExternalorganizationMap["faxNumber"].(map[string]interface{}); ok {
		FaxNumberString, _ := json.Marshal(FaxNumber)
		json.Unmarshal(FaxNumberString, &o.FaxNumber)
	}
	
	if EmployeeCount, ok := ExternalorganizationMap["employeeCount"].(float64); ok {
		EmployeeCountInt := int(EmployeeCount)
		o.EmployeeCount = &EmployeeCountInt
	}
	
	if Revenue, ok := ExternalorganizationMap["revenue"].(float64); ok {
		RevenueInt := int(Revenue)
		o.Revenue = &RevenueInt
	}
	
	if Tags, ok := ExternalorganizationMap["tags"].([]interface{}); ok {
		TagsString, _ := json.Marshal(Tags)
		json.Unmarshal(TagsString, &o.Tags)
	}
	
	if Websites, ok := ExternalorganizationMap["websites"].([]interface{}); ok {
		WebsitesString, _ := json.Marshal(Websites)
		json.Unmarshal(WebsitesString, &o.Websites)
	}
	
	if Tickers, ok := ExternalorganizationMap["tickers"].([]interface{}); ok {
		TickersString, _ := json.Marshal(Tickers)
		json.Unmarshal(TickersString, &o.Tickers)
	}
	
	if TwitterId, ok := ExternalorganizationMap["twitterId"].(map[string]interface{}); ok {
		TwitterIdString, _ := json.Marshal(TwitterId)
		json.Unmarshal(TwitterIdString, &o.TwitterId)
	}
	
	if ExternalSystemUrl, ok := ExternalorganizationMap["externalSystemUrl"].(string); ok {
		o.ExternalSystemUrl = &ExternalSystemUrl
	}
    
	if modifyDateString, ok := ExternalorganizationMap["modifyDate"].(string); ok {
		ModifyDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", modifyDateString)
		o.ModifyDate = &ModifyDate
	}
	
	if createDateString, ok := ExternalorganizationMap["createDate"].(string); ok {
		CreateDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", createDateString)
		o.CreateDate = &CreateDate
	}
	
	if Trustor, ok := ExternalorganizationMap["trustor"].(map[string]interface{}); ok {
		TrustorString, _ := json.Marshal(Trustor)
		json.Unmarshal(TrustorString, &o.Trustor)
	}
	
	if Schema, ok := ExternalorganizationMap["schema"].(map[string]interface{}); ok {
		SchemaString, _ := json.Marshal(Schema)
		json.Unmarshal(SchemaString, &o.Schema)
	}
	
	if CustomFields, ok := ExternalorganizationMap["customFields"].(map[string]interface{}); ok {
		CustomFieldsString, _ := json.Marshal(CustomFields)
		json.Unmarshal(CustomFieldsString, &o.CustomFields)
	}
	
	if Identifiers, ok := ExternalorganizationMap["identifiers"].([]interface{}); ok {
		IdentifiersString, _ := json.Marshal(Identifiers)
		json.Unmarshal(IdentifiersString, &o.Identifiers)
	}
	
	if ExternalIds, ok := ExternalorganizationMap["externalIds"].([]interface{}); ok {
		ExternalIdsString, _ := json.Marshal(ExternalIds)
		json.Unmarshal(ExternalIdsString, &o.ExternalIds)
	}
	
	if ExternalDataSources, ok := ExternalorganizationMap["externalDataSources"].([]interface{}); ok {
		ExternalDataSourcesString, _ := json.Marshal(ExternalDataSources)
		json.Unmarshal(ExternalDataSourcesString, &o.ExternalDataSources)
	}
	
	if SelfUri, ok := ExternalorganizationMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Externalorganization) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

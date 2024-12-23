package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Externalcontactsrelationshipchangedtopicexternalorganization
type Externalcontactsrelationshipchangedtopicexternalorganization struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id
	Id *string `json:"id,omitempty"`

	// Division
	Division *Externalcontactsrelationshipchangedtopicdivision `json:"division,omitempty"`

	// Name
	Name *string `json:"name,omitempty"`

	// CompanyType
	CompanyType *string `json:"companyType,omitempty"`

	// Industry
	Industry *string `json:"industry,omitempty"`

	// PrimaryContactId
	PrimaryContactId *string `json:"primaryContactId,omitempty"`

	// Address
	Address *Externalcontactsrelationshipchangedtopiccontactaddress `json:"address,omitempty"`

	// PhoneNumber
	PhoneNumber *Externalcontactsrelationshipchangedtopicphonenumber `json:"phoneNumber,omitempty"`

	// FaxNumber
	FaxNumber *Externalcontactsrelationshipchangedtopicphonenumber `json:"faxNumber,omitempty"`

	// EmployeeCount
	EmployeeCount *int `json:"employeeCount,omitempty"`

	// Revenue
	Revenue *int `json:"revenue,omitempty"`

	// Tags
	Tags *[]string `json:"tags,omitempty"`

	// Websites
	Websites *[]string `json:"websites,omitempty"`

	// Tickers
	Tickers *[]Externalcontactsrelationshipchangedtopicticker `json:"tickers,omitempty"`

	// TwitterId
	TwitterId *Externalcontactsrelationshipchangedtopictwitterid `json:"twitterId,omitempty"`

	// ExternalSystemUrl
	ExternalSystemUrl *string `json:"externalSystemUrl,omitempty"`

	// CustomFields
	CustomFields *map[string]interface{} `json:"customFields,omitempty"`

	// CreateDate
	CreateDate *time.Time `json:"createDate,omitempty"`

	// ModifyDate
	ModifyDate *time.Time `json:"modifyDate,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Externalcontactsrelationshipchangedtopicexternalorganization) SetField(field string, fieldValue interface{}) {
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

func (o Externalcontactsrelationshipchangedtopicexternalorganization) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "CreateDate","ModifyDate", }
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
	type Alias Externalcontactsrelationshipchangedtopicexternalorganization
	
	CreateDate := new(string)
	if o.CreateDate != nil {
		
		*CreateDate = timeutil.Strftime(o.CreateDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		CreateDate = nil
	}
	
	ModifyDate := new(string)
	if o.ModifyDate != nil {
		
		*ModifyDate = timeutil.Strftime(o.ModifyDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ModifyDate = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Division *Externalcontactsrelationshipchangedtopicdivision `json:"division,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		CompanyType *string `json:"companyType,omitempty"`
		
		Industry *string `json:"industry,omitempty"`
		
		PrimaryContactId *string `json:"primaryContactId,omitempty"`
		
		Address *Externalcontactsrelationshipchangedtopiccontactaddress `json:"address,omitempty"`
		
		PhoneNumber *Externalcontactsrelationshipchangedtopicphonenumber `json:"phoneNumber,omitempty"`
		
		FaxNumber *Externalcontactsrelationshipchangedtopicphonenumber `json:"faxNumber,omitempty"`
		
		EmployeeCount *int `json:"employeeCount,omitempty"`
		
		Revenue *int `json:"revenue,omitempty"`
		
		Tags *[]string `json:"tags,omitempty"`
		
		Websites *[]string `json:"websites,omitempty"`
		
		Tickers *[]Externalcontactsrelationshipchangedtopicticker `json:"tickers,omitempty"`
		
		TwitterId *Externalcontactsrelationshipchangedtopictwitterid `json:"twitterId,omitempty"`
		
		ExternalSystemUrl *string `json:"externalSystemUrl,omitempty"`
		
		CustomFields *map[string]interface{} `json:"customFields,omitempty"`
		
		CreateDate *string `json:"createDate,omitempty"`
		
		ModifyDate *string `json:"modifyDate,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Division: o.Division,
		
		Name: o.Name,
		
		CompanyType: o.CompanyType,
		
		Industry: o.Industry,
		
		PrimaryContactId: o.PrimaryContactId,
		
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
		
		CustomFields: o.CustomFields,
		
		CreateDate: CreateDate,
		
		ModifyDate: ModifyDate,
		Alias:    (Alias)(o),
	})
}

func (o *Externalcontactsrelationshipchangedtopicexternalorganization) UnmarshalJSON(b []byte) error {
	var ExternalcontactsrelationshipchangedtopicexternalorganizationMap map[string]interface{}
	err := json.Unmarshal(b, &ExternalcontactsrelationshipchangedtopicexternalorganizationMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ExternalcontactsrelationshipchangedtopicexternalorganizationMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Division, ok := ExternalcontactsrelationshipchangedtopicexternalorganizationMap["division"].(map[string]interface{}); ok {
		DivisionString, _ := json.Marshal(Division)
		json.Unmarshal(DivisionString, &o.Division)
	}
	
	if Name, ok := ExternalcontactsrelationshipchangedtopicexternalorganizationMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if CompanyType, ok := ExternalcontactsrelationshipchangedtopicexternalorganizationMap["companyType"].(string); ok {
		o.CompanyType = &CompanyType
	}
    
	if Industry, ok := ExternalcontactsrelationshipchangedtopicexternalorganizationMap["industry"].(string); ok {
		o.Industry = &Industry
	}
    
	if PrimaryContactId, ok := ExternalcontactsrelationshipchangedtopicexternalorganizationMap["primaryContactId"].(string); ok {
		o.PrimaryContactId = &PrimaryContactId
	}
    
	if Address, ok := ExternalcontactsrelationshipchangedtopicexternalorganizationMap["address"].(map[string]interface{}); ok {
		AddressString, _ := json.Marshal(Address)
		json.Unmarshal(AddressString, &o.Address)
	}
	
	if PhoneNumber, ok := ExternalcontactsrelationshipchangedtopicexternalorganizationMap["phoneNumber"].(map[string]interface{}); ok {
		PhoneNumberString, _ := json.Marshal(PhoneNumber)
		json.Unmarshal(PhoneNumberString, &o.PhoneNumber)
	}
	
	if FaxNumber, ok := ExternalcontactsrelationshipchangedtopicexternalorganizationMap["faxNumber"].(map[string]interface{}); ok {
		FaxNumberString, _ := json.Marshal(FaxNumber)
		json.Unmarshal(FaxNumberString, &o.FaxNumber)
	}
	
	if EmployeeCount, ok := ExternalcontactsrelationshipchangedtopicexternalorganizationMap["employeeCount"].(float64); ok {
		EmployeeCountInt := int(EmployeeCount)
		o.EmployeeCount = &EmployeeCountInt
	}
	
	if Revenue, ok := ExternalcontactsrelationshipchangedtopicexternalorganizationMap["revenue"].(float64); ok {
		RevenueInt := int(Revenue)
		o.Revenue = &RevenueInt
	}
	
	if Tags, ok := ExternalcontactsrelationshipchangedtopicexternalorganizationMap["tags"].([]interface{}); ok {
		TagsString, _ := json.Marshal(Tags)
		json.Unmarshal(TagsString, &o.Tags)
	}
	
	if Websites, ok := ExternalcontactsrelationshipchangedtopicexternalorganizationMap["websites"].([]interface{}); ok {
		WebsitesString, _ := json.Marshal(Websites)
		json.Unmarshal(WebsitesString, &o.Websites)
	}
	
	if Tickers, ok := ExternalcontactsrelationshipchangedtopicexternalorganizationMap["tickers"].([]interface{}); ok {
		TickersString, _ := json.Marshal(Tickers)
		json.Unmarshal(TickersString, &o.Tickers)
	}
	
	if TwitterId, ok := ExternalcontactsrelationshipchangedtopicexternalorganizationMap["twitterId"].(map[string]interface{}); ok {
		TwitterIdString, _ := json.Marshal(TwitterId)
		json.Unmarshal(TwitterIdString, &o.TwitterId)
	}
	
	if ExternalSystemUrl, ok := ExternalcontactsrelationshipchangedtopicexternalorganizationMap["externalSystemUrl"].(string); ok {
		o.ExternalSystemUrl = &ExternalSystemUrl
	}
    
	if CustomFields, ok := ExternalcontactsrelationshipchangedtopicexternalorganizationMap["customFields"].(map[string]interface{}); ok {
		CustomFieldsString, _ := json.Marshal(CustomFields)
		json.Unmarshal(CustomFieldsString, &o.CustomFields)
	}
	
	if createDateString, ok := ExternalcontactsrelationshipchangedtopicexternalorganizationMap["createDate"].(string); ok {
		CreateDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", createDateString)
		o.CreateDate = &CreateDate
	}
	
	if modifyDateString, ok := ExternalcontactsrelationshipchangedtopicexternalorganizationMap["modifyDate"].(string); ok {
		ModifyDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", modifyDateString)
		o.ModifyDate = &ModifyDate
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Externalcontactsrelationshipchangedtopicexternalorganization) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

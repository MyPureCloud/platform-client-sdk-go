package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Dnclist
type Dnclist struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name - The name of the DncList.
	Name *string `json:"name,omitempty"`

	// DateCreated - Creation time of the entity. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`

	// DateModified - Last modified time of the entity. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`

	// Version - Required for updates, must match the version number of the most recent update
	Version *int `json:"version,omitempty"`

	// ImportStatus - The status of the import process
	ImportStatus *Importstatus `json:"importStatus,omitempty"`

	// Size - The total number of phone numbers in the DncList.
	Size *int `json:"size,omitempty"`

	// DncSourceType - The type of the DncList.
	DncSourceType *string `json:"dncSourceType,omitempty"`

	// ContactMethod - The contact method. Required if dncSourceType is rds.
	ContactMethod *string `json:"contactMethod,omitempty"`

	// LoginId - A dnc.com loginId. Required if the dncSourceType is dnc.com.
	LoginId *string `json:"loginId,omitempty"`

	// CampaignId - A dnc.com campaignId. Optional if the dncSourceType is dnc.com.
	CampaignId *string `json:"campaignId,omitempty"`

	// DncCodes - The list of dnc.com codes to be treated as DNC. Required if the dncSourceType is dnc.com.
	DncCodes *[]string `json:"dncCodes,omitempty"`

	// LicenseId - A gryphon license number. Required if the dncSourceType is gryphon.
	LicenseId *string `json:"licenseId,omitempty"`

	// Division - The division this DncList belongs to.
	Division *Domainentityref `json:"division,omitempty"`

	// CustomExclusionColumn - The column to evaluate exclusion against. Required if the dncSourceType is rds_custom.
	CustomExclusionColumn *string `json:"customExclusionColumn,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Dnclist) SetField(field string, fieldValue interface{}) {
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

func (o Dnclist) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateCreated","DateModified", }
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
	type Alias Dnclist
	
	DateCreated := new(string)
	if o.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(o.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	DateModified := new(string)
	if o.DateModified != nil {
		
		*DateModified = timeutil.Strftime(o.DateModified, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateModified = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		Version *int `json:"version,omitempty"`
		
		ImportStatus *Importstatus `json:"importStatus,omitempty"`
		
		Size *int `json:"size,omitempty"`
		
		DncSourceType *string `json:"dncSourceType,omitempty"`
		
		ContactMethod *string `json:"contactMethod,omitempty"`
		
		LoginId *string `json:"loginId,omitempty"`
		
		CampaignId *string `json:"campaignId,omitempty"`
		
		DncCodes *[]string `json:"dncCodes,omitempty"`
		
		LicenseId *string `json:"licenseId,omitempty"`
		
		Division *Domainentityref `json:"division,omitempty"`
		
		CustomExclusionColumn *string `json:"customExclusionColumn,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		Version: o.Version,
		
		ImportStatus: o.ImportStatus,
		
		Size: o.Size,
		
		DncSourceType: o.DncSourceType,
		
		ContactMethod: o.ContactMethod,
		
		LoginId: o.LoginId,
		
		CampaignId: o.CampaignId,
		
		DncCodes: o.DncCodes,
		
		LicenseId: o.LicenseId,
		
		Division: o.Division,
		
		CustomExclusionColumn: o.CustomExclusionColumn,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Dnclist) UnmarshalJSON(b []byte) error {
	var DnclistMap map[string]interface{}
	err := json.Unmarshal(b, &DnclistMap)
	if err != nil {
		return err
	}
	
	if Id, ok := DnclistMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := DnclistMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if dateCreatedString, ok := DnclistMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := DnclistMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if Version, ok := DnclistMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	
	if ImportStatus, ok := DnclistMap["importStatus"].(map[string]interface{}); ok {
		ImportStatusString, _ := json.Marshal(ImportStatus)
		json.Unmarshal(ImportStatusString, &o.ImportStatus)
	}
	
	if Size, ok := DnclistMap["size"].(float64); ok {
		SizeInt := int(Size)
		o.Size = &SizeInt
	}
	
	if DncSourceType, ok := DnclistMap["dncSourceType"].(string); ok {
		o.DncSourceType = &DncSourceType
	}
    
	if ContactMethod, ok := DnclistMap["contactMethod"].(string); ok {
		o.ContactMethod = &ContactMethod
	}
    
	if LoginId, ok := DnclistMap["loginId"].(string); ok {
		o.LoginId = &LoginId
	}
    
	if CampaignId, ok := DnclistMap["campaignId"].(string); ok {
		o.CampaignId = &CampaignId
	}
    
	if DncCodes, ok := DnclistMap["dncCodes"].([]interface{}); ok {
		DncCodesString, _ := json.Marshal(DncCodes)
		json.Unmarshal(DncCodesString, &o.DncCodes)
	}
	
	if LicenseId, ok := DnclistMap["licenseId"].(string); ok {
		o.LicenseId = &LicenseId
	}
    
	if Division, ok := DnclistMap["division"].(map[string]interface{}); ok {
		DivisionString, _ := json.Marshal(Division)
		json.Unmarshal(DivisionString, &o.Division)
	}
	
	if CustomExclusionColumn, ok := DnclistMap["customExclusionColumn"].(string); ok {
		o.CustomExclusionColumn = &CustomExclusionColumn
	}
    
	if SelfUri, ok := DnclistMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Dnclist) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Contactlisttemplate
type Contactlisttemplate struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name
	Name *string `json:"name,omitempty"`

	// DateCreated - Creation time of the entity. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`

	// DateModified - Last modified time of the entity. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`

	// Version - Required for updates, must match the version number of the most recent update
	Version *int `json:"version,omitempty"`

	// ColumnNames - The names of the contact data columns.
	ColumnNames *[]string `json:"columnNames,omitempty"`

	// PhoneColumns - Indicates which columns are phone numbers.
	PhoneColumns *[]Contactphonenumbercolumn `json:"phoneColumns,omitempty"`

	// EmailColumns - Indicates which columns are email addresses
	EmailColumns *[]Emailcolumn `json:"emailColumns,omitempty"`

	// PreviewModeColumnName - A column to check if a contact should always be dialed in preview mode.
	PreviewModeColumnName *string `json:"previewModeColumnName,omitempty"`

	// PreviewModeAcceptedValues - The values in the previewModeColumnName column that indicate a contact should always be dialed in preview mode.
	PreviewModeAcceptedValues *[]string `json:"previewModeAcceptedValues,omitempty"`

	// AttemptLimits - AttemptLimits for this ContactListTemplate.
	AttemptLimits *Domainentityref `json:"attemptLimits,omitempty"`

	// AutomaticTimeZoneMapping - Indicates if automatic time zone mapping is to be used for this ContactListTemplate.
	AutomaticTimeZoneMapping *bool `json:"automaticTimeZoneMapping,omitempty"`

	// ZipCodeColumnName - The name of ContactListTemplate column containing the zip code for use with automatic time zone mapping. Only allowed if 'automaticTimeZoneMapping' is set to true.
	ZipCodeColumnName *string `json:"zipCodeColumnName,omitempty"`

	// ColumnDataTypeSpecifications - The settings of the columns selected for dynamic queueing
	ColumnDataTypeSpecifications *[]Columndatatypespecification `json:"columnDataTypeSpecifications,omitempty"`

	// TrimWhitespace - Whether to trim white space when importing a ContactListTemplate csv file, default value = true
	TrimWhitespace *bool `json:"trimWhitespace,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Contactlisttemplate) SetField(field string, fieldValue interface{}) {
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

func (o Contactlisttemplate) MarshalJSON() ([]byte, error) {
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
	type Alias Contactlisttemplate
	
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
		
		ColumnNames *[]string `json:"columnNames,omitempty"`
		
		PhoneColumns *[]Contactphonenumbercolumn `json:"phoneColumns,omitempty"`
		
		EmailColumns *[]Emailcolumn `json:"emailColumns,omitempty"`
		
		PreviewModeColumnName *string `json:"previewModeColumnName,omitempty"`
		
		PreviewModeAcceptedValues *[]string `json:"previewModeAcceptedValues,omitempty"`
		
		AttemptLimits *Domainentityref `json:"attemptLimits,omitempty"`
		
		AutomaticTimeZoneMapping *bool `json:"automaticTimeZoneMapping,omitempty"`
		
		ZipCodeColumnName *string `json:"zipCodeColumnName,omitempty"`
		
		ColumnDataTypeSpecifications *[]Columndatatypespecification `json:"columnDataTypeSpecifications,omitempty"`
		
		TrimWhitespace *bool `json:"trimWhitespace,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		Version: o.Version,
		
		ColumnNames: o.ColumnNames,
		
		PhoneColumns: o.PhoneColumns,
		
		EmailColumns: o.EmailColumns,
		
		PreviewModeColumnName: o.PreviewModeColumnName,
		
		PreviewModeAcceptedValues: o.PreviewModeAcceptedValues,
		
		AttemptLimits: o.AttemptLimits,
		
		AutomaticTimeZoneMapping: o.AutomaticTimeZoneMapping,
		
		ZipCodeColumnName: o.ZipCodeColumnName,
		
		ColumnDataTypeSpecifications: o.ColumnDataTypeSpecifications,
		
		TrimWhitespace: o.TrimWhitespace,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Contactlisttemplate) UnmarshalJSON(b []byte) error {
	var ContactlisttemplateMap map[string]interface{}
	err := json.Unmarshal(b, &ContactlisttemplateMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ContactlisttemplateMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := ContactlisttemplateMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if dateCreatedString, ok := ContactlisttemplateMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := ContactlisttemplateMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if Version, ok := ContactlisttemplateMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	
	if ColumnNames, ok := ContactlisttemplateMap["columnNames"].([]interface{}); ok {
		ColumnNamesString, _ := json.Marshal(ColumnNames)
		json.Unmarshal(ColumnNamesString, &o.ColumnNames)
	}
	
	if PhoneColumns, ok := ContactlisttemplateMap["phoneColumns"].([]interface{}); ok {
		PhoneColumnsString, _ := json.Marshal(PhoneColumns)
		json.Unmarshal(PhoneColumnsString, &o.PhoneColumns)
	}
	
	if EmailColumns, ok := ContactlisttemplateMap["emailColumns"].([]interface{}); ok {
		EmailColumnsString, _ := json.Marshal(EmailColumns)
		json.Unmarshal(EmailColumnsString, &o.EmailColumns)
	}
	
	if PreviewModeColumnName, ok := ContactlisttemplateMap["previewModeColumnName"].(string); ok {
		o.PreviewModeColumnName = &PreviewModeColumnName
	}
    
	if PreviewModeAcceptedValues, ok := ContactlisttemplateMap["previewModeAcceptedValues"].([]interface{}); ok {
		PreviewModeAcceptedValuesString, _ := json.Marshal(PreviewModeAcceptedValues)
		json.Unmarshal(PreviewModeAcceptedValuesString, &o.PreviewModeAcceptedValues)
	}
	
	if AttemptLimits, ok := ContactlisttemplateMap["attemptLimits"].(map[string]interface{}); ok {
		AttemptLimitsString, _ := json.Marshal(AttemptLimits)
		json.Unmarshal(AttemptLimitsString, &o.AttemptLimits)
	}
	
	if AutomaticTimeZoneMapping, ok := ContactlisttemplateMap["automaticTimeZoneMapping"].(bool); ok {
		o.AutomaticTimeZoneMapping = &AutomaticTimeZoneMapping
	}
    
	if ZipCodeColumnName, ok := ContactlisttemplateMap["zipCodeColumnName"].(string); ok {
		o.ZipCodeColumnName = &ZipCodeColumnName
	}
    
	if ColumnDataTypeSpecifications, ok := ContactlisttemplateMap["columnDataTypeSpecifications"].([]interface{}); ok {
		ColumnDataTypeSpecificationsString, _ := json.Marshal(ColumnDataTypeSpecifications)
		json.Unmarshal(ColumnDataTypeSpecificationsString, &o.ColumnDataTypeSpecifications)
	}
	
	if TrimWhitespace, ok := ContactlisttemplateMap["trimWhitespace"].(bool); ok {
		o.TrimWhitespace = &TrimWhitespace
	}
    
	if SelfUri, ok := ContactlisttemplateMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Contactlisttemplate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

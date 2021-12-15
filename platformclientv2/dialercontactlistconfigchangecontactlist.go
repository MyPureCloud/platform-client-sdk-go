package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Dialercontactlistconfigchangecontactlist
type Dialercontactlistconfigchangecontactlist struct { 
	// ColumnNames - the contact column names
	ColumnNames *[]string `json:"columnNames,omitempty"`


	// PhoneColumns - the columns containing phone numbers
	PhoneColumns *[]Dialercontactlistconfigchangecontactphonenumbercolumn `json:"phoneColumns,omitempty"`


	// EmailColumns - the columns containing email addresses
	EmailColumns *[]Dialercontactlistconfigchangeemailcolumn `json:"emailColumns,omitempty"`


	// ImportStatus
	ImportStatus *Dialercontactlistconfigchangeimportstatus `json:"importStatus,omitempty"`


	// PreviewModeColumnName - the name of the column that holds the indicators for contacts that are to be dialed in preview mode only
	PreviewModeColumnName *string `json:"previewModeColumnName,omitempty"`


	// PreviewModeAcceptedValues - list of user-defined values indicating the contact is to be dialed in preview mode only
	PreviewModeAcceptedValues *[]string `json:"previewModeAcceptedValues,omitempty"`


	// Size - the number of contacts in the contact list
	Size *int `json:"size,omitempty"`


	// AttemptLimits
	AttemptLimits *Dialercontactlistconfigchangeurireference `json:"attemptLimits,omitempty"`


	// AutomaticTimeZoneMapping - whether or not automatic time zone mapping is enabled on the list
	AutomaticTimeZoneMapping *bool `json:"automaticTimeZoneMapping,omitempty"`


	// ZipCodeColumnName - zip code column from the contact list to be used optionally with automatic time zone mapping
	ZipCodeColumnName *string `json:"zipCodeColumnName,omitempty"`


	// Division - A UriReference for a resource
	Division *Dialercontactlistconfigchangeurireference `json:"division,omitempty"`


	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name - The UI-visible name of the object
	Name *string `json:"name,omitempty"`


	// DateCreated - Creation time of the entity
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateModified - Last modified time of the entity
	DateModified *time.Time `json:"dateModified,omitempty"`


	// Version - Required for updates, must match the version number of the most recent update
	Version *int `json:"version,omitempty"`

}

func (o *Dialercontactlistconfigchangecontactlist) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dialercontactlistconfigchangecontactlist
	
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
		ColumnNames *[]string `json:"columnNames,omitempty"`
		
		PhoneColumns *[]Dialercontactlistconfigchangecontactphonenumbercolumn `json:"phoneColumns,omitempty"`
		
		EmailColumns *[]Dialercontactlistconfigchangeemailcolumn `json:"emailColumns,omitempty"`
		
		ImportStatus *Dialercontactlistconfigchangeimportstatus `json:"importStatus,omitempty"`
		
		PreviewModeColumnName *string `json:"previewModeColumnName,omitempty"`
		
		PreviewModeAcceptedValues *[]string `json:"previewModeAcceptedValues,omitempty"`
		
		Size *int `json:"size,omitempty"`
		
		AttemptLimits *Dialercontactlistconfigchangeurireference `json:"attemptLimits,omitempty"`
		
		AutomaticTimeZoneMapping *bool `json:"automaticTimeZoneMapping,omitempty"`
		
		ZipCodeColumnName *string `json:"zipCodeColumnName,omitempty"`
		
		Division *Dialercontactlistconfigchangeurireference `json:"division,omitempty"`
		
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		Version *int `json:"version,omitempty"`
		*Alias
	}{ 
		ColumnNames: o.ColumnNames,
		
		PhoneColumns: o.PhoneColumns,
		
		EmailColumns: o.EmailColumns,
		
		ImportStatus: o.ImportStatus,
		
		PreviewModeColumnName: o.PreviewModeColumnName,
		
		PreviewModeAcceptedValues: o.PreviewModeAcceptedValues,
		
		Size: o.Size,
		
		AttemptLimits: o.AttemptLimits,
		
		AutomaticTimeZoneMapping: o.AutomaticTimeZoneMapping,
		
		ZipCodeColumnName: o.ZipCodeColumnName,
		
		Division: o.Division,
		
		Id: o.Id,
		
		Name: o.Name,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		Version: o.Version,
		Alias:    (*Alias)(o),
	})
}

func (o *Dialercontactlistconfigchangecontactlist) UnmarshalJSON(b []byte) error {
	var DialercontactlistconfigchangecontactlistMap map[string]interface{}
	err := json.Unmarshal(b, &DialercontactlistconfigchangecontactlistMap)
	if err != nil {
		return err
	}
	
	if ColumnNames, ok := DialercontactlistconfigchangecontactlistMap["columnNames"].([]interface{}); ok {
		ColumnNamesString, _ := json.Marshal(ColumnNames)
		json.Unmarshal(ColumnNamesString, &o.ColumnNames)
	}
	
	if PhoneColumns, ok := DialercontactlistconfigchangecontactlistMap["phoneColumns"].([]interface{}); ok {
		PhoneColumnsString, _ := json.Marshal(PhoneColumns)
		json.Unmarshal(PhoneColumnsString, &o.PhoneColumns)
	}
	
	if EmailColumns, ok := DialercontactlistconfigchangecontactlistMap["emailColumns"].([]interface{}); ok {
		EmailColumnsString, _ := json.Marshal(EmailColumns)
		json.Unmarshal(EmailColumnsString, &o.EmailColumns)
	}
	
	if ImportStatus, ok := DialercontactlistconfigchangecontactlistMap["importStatus"].(map[string]interface{}); ok {
		ImportStatusString, _ := json.Marshal(ImportStatus)
		json.Unmarshal(ImportStatusString, &o.ImportStatus)
	}
	
	if PreviewModeColumnName, ok := DialercontactlistconfigchangecontactlistMap["previewModeColumnName"].(string); ok {
		o.PreviewModeColumnName = &PreviewModeColumnName
	}
	
	if PreviewModeAcceptedValues, ok := DialercontactlistconfigchangecontactlistMap["previewModeAcceptedValues"].([]interface{}); ok {
		PreviewModeAcceptedValuesString, _ := json.Marshal(PreviewModeAcceptedValues)
		json.Unmarshal(PreviewModeAcceptedValuesString, &o.PreviewModeAcceptedValues)
	}
	
	if Size, ok := DialercontactlistconfigchangecontactlistMap["size"].(float64); ok {
		SizeInt := int(Size)
		o.Size = &SizeInt
	}
	
	if AttemptLimits, ok := DialercontactlistconfigchangecontactlistMap["attemptLimits"].(map[string]interface{}); ok {
		AttemptLimitsString, _ := json.Marshal(AttemptLimits)
		json.Unmarshal(AttemptLimitsString, &o.AttemptLimits)
	}
	
	if AutomaticTimeZoneMapping, ok := DialercontactlistconfigchangecontactlistMap["automaticTimeZoneMapping"].(bool); ok {
		o.AutomaticTimeZoneMapping = &AutomaticTimeZoneMapping
	}
	
	if ZipCodeColumnName, ok := DialercontactlistconfigchangecontactlistMap["zipCodeColumnName"].(string); ok {
		o.ZipCodeColumnName = &ZipCodeColumnName
	}
	
	if Division, ok := DialercontactlistconfigchangecontactlistMap["division"].(map[string]interface{}); ok {
		DivisionString, _ := json.Marshal(Division)
		json.Unmarshal(DivisionString, &o.Division)
	}
	
	if Id, ok := DialercontactlistconfigchangecontactlistMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := DialercontactlistconfigchangecontactlistMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if dateCreatedString, ok := DialercontactlistconfigchangecontactlistMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := DialercontactlistconfigchangecontactlistMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if Version, ok := DialercontactlistconfigchangecontactlistMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Dialercontactlistconfigchangecontactlist) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

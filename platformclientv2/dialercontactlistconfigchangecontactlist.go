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
	// Id
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// DateCreated
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateModified
	DateModified *time.Time `json:"dateModified,omitempty"`


	// Version
	Version *int `json:"version,omitempty"`


	// ColumnNames
	ColumnNames *[]string `json:"columnNames,omitempty"`


	// PhoneColumns
	PhoneColumns *[]Dialercontactlistconfigchangecontactphonenumbercolumn `json:"phoneColumns,omitempty"`


	// EmailColumns
	EmailColumns *[]Dialercontactlistconfigchangeemailcolumn `json:"emailColumns,omitempty"`


	// ImportStatus
	ImportStatus *Dialercontactlistconfigchangeimportstatus `json:"importStatus,omitempty"`


	// PreviewModeColumnName
	PreviewModeColumnName *string `json:"previewModeColumnName,omitempty"`


	// PreviewModeAcceptedValues
	PreviewModeAcceptedValues *[]string `json:"previewModeAcceptedValues,omitempty"`


	// Size
	Size *int `json:"size,omitempty"`


	// AttemptLimits
	AttemptLimits *Dialercontactlistconfigchangeurireference `json:"attemptLimits,omitempty"`


	// AutomaticTimeZoneMapping
	AutomaticTimeZoneMapping *bool `json:"automaticTimeZoneMapping,omitempty"`


	// ZipCodeColumnName
	ZipCodeColumnName *string `json:"zipCodeColumnName,omitempty"`


	// Division
	Division *Dialercontactlistconfigchangeurireference `json:"division,omitempty"`


	// AdditionalProperties
	AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`

}

func (u *Dialercontactlistconfigchangecontactlist) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dialercontactlistconfigchangecontactlist

	
	DateCreated := new(string)
	if u.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(u.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	DateModified := new(string)
	if u.DateModified != nil {
		
		*DateModified = timeutil.Strftime(u.DateModified, "%Y-%m-%dT%H:%M:%S.%fZ")
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
		
		AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		Version: u.Version,
		
		ColumnNames: u.ColumnNames,
		
		PhoneColumns: u.PhoneColumns,
		
		EmailColumns: u.EmailColumns,
		
		ImportStatus: u.ImportStatus,
		
		PreviewModeColumnName: u.PreviewModeColumnName,
		
		PreviewModeAcceptedValues: u.PreviewModeAcceptedValues,
		
		Size: u.Size,
		
		AttemptLimits: u.AttemptLimits,
		
		AutomaticTimeZoneMapping: u.AutomaticTimeZoneMapping,
		
		ZipCodeColumnName: u.ZipCodeColumnName,
		
		Division: u.Division,
		
		AdditionalProperties: u.AdditionalProperties,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Dialercontactlistconfigchangecontactlist) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

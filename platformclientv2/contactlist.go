package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Contactlist
type Contactlist struct { 
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


	// Division - The division this entity belongs to.
	Division *Domainentityref `json:"division,omitempty"`


	// ColumnNames - The names of the contact data columns.
	ColumnNames *[]string `json:"columnNames,omitempty"`


	// PhoneColumns - Indicates which columns are phone numbers.
	PhoneColumns *[]Contactphonenumbercolumn `json:"phoneColumns,omitempty"`


	// ImportStatus - The status of the import process.
	ImportStatus *Importstatus `json:"importStatus,omitempty"`


	// PreviewModeColumnName - A column to check if a contact should always be dialed in preview mode.
	PreviewModeColumnName *string `json:"previewModeColumnName,omitempty"`


	// PreviewModeAcceptedValues - The values in the previewModeColumnName column that indicate a contact should always be dialed in preview mode.
	PreviewModeAcceptedValues *[]string `json:"previewModeAcceptedValues,omitempty"`


	// Size - The number of contacts in the ContactList.
	Size *int `json:"size,omitempty"`


	// AttemptLimits - AttemptLimits for this ContactList.
	AttemptLimits *Domainentityref `json:"attemptLimits,omitempty"`


	// AutomaticTimeZoneMapping - Indicates if automatic time zone mapping is to be used for this ContactList.
	AutomaticTimeZoneMapping *bool `json:"automaticTimeZoneMapping,omitempty"`


	// ZipCodeColumnName - The name of contact list column containing the zip code for use with automatic time zone mapping. Only allowed if 'automaticTimeZoneMapping' is set to true.
	ZipCodeColumnName *string `json:"zipCodeColumnName,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (u *Contactlist) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Contactlist

	
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
		
		Division *Domainentityref `json:"division,omitempty"`
		
		ColumnNames *[]string `json:"columnNames,omitempty"`
		
		PhoneColumns *[]Contactphonenumbercolumn `json:"phoneColumns,omitempty"`
		
		ImportStatus *Importstatus `json:"importStatus,omitempty"`
		
		PreviewModeColumnName *string `json:"previewModeColumnName,omitempty"`
		
		PreviewModeAcceptedValues *[]string `json:"previewModeAcceptedValues,omitempty"`
		
		Size *int `json:"size,omitempty"`
		
		AttemptLimits *Domainentityref `json:"attemptLimits,omitempty"`
		
		AutomaticTimeZoneMapping *bool `json:"automaticTimeZoneMapping,omitempty"`
		
		ZipCodeColumnName *string `json:"zipCodeColumnName,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		Version: u.Version,
		
		Division: u.Division,
		
		ColumnNames: u.ColumnNames,
		
		PhoneColumns: u.PhoneColumns,
		
		ImportStatus: u.ImportStatus,
		
		PreviewModeColumnName: u.PreviewModeColumnName,
		
		PreviewModeAcceptedValues: u.PreviewModeAcceptedValues,
		
		Size: u.Size,
		
		AttemptLimits: u.AttemptLimits,
		
		AutomaticTimeZoneMapping: u.AutomaticTimeZoneMapping,
		
		ZipCodeColumnName: u.ZipCodeColumnName,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Contactlist) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

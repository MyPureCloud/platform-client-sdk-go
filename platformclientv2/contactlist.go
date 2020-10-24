package platformclientv2
import (
	"time"
	"encoding/json"
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
	Version *int32 `json:"version,omitempty"`


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
	Size *int64 `json:"size,omitempty"`


	// AttemptLimits - AttemptLimits for this ContactList.
	AttemptLimits *Domainentityref `json:"attemptLimits,omitempty"`


	// AutomaticTimeZoneMapping - Indicates if automatic time zone mapping is to be used for this ContactList.
	AutomaticTimeZoneMapping *bool `json:"automaticTimeZoneMapping,omitempty"`


	// ZipCodeColumnName - The name of contact list column containing the zip code for use with automatic time zone mapping. Only allowed if 'automaticTimeZoneMapping' is set to true.
	ZipCodeColumnName *string `json:"zipCodeColumnName,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Contactlist) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

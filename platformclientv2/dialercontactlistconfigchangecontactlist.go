package platformclientv2
import (
	"time"
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

// String returns a JSON representation of the model
func (o *Dialercontactlistconfigchangecontactlist) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

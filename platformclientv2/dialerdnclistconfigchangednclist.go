package platformclientv2
import (
	"time"
	"encoding/json"
)

// Dialerdnclistconfigchangednclist
type Dialerdnclistconfigchangednclist struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// DateCreated
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateModified
	DateModified *time.Time `json:"dateModified,omitempty"`


	// Version
	Version *int32 `json:"version,omitempty"`


	// ImportStatus
	ImportStatus *Dialerdnclistconfigchangeimportstatus `json:"importStatus,omitempty"`


	// Size
	Size *int32 `json:"size,omitempty"`


	// DncSourceType
	DncSourceType *string `json:"dncSourceType,omitempty"`


	// LoginId
	LoginId *string `json:"loginId,omitempty"`


	// DncCodes
	DncCodes *[]string `json:"dncCodes,omitempty"`


	// LicenseId
	LicenseId *string `json:"licenseId,omitempty"`


	// Division
	Division *Dialerdnclistconfigchangeurireference `json:"division,omitempty"`


	// AdditionalProperties
	AdditionalProperties *map[string]interface{} `json:"additionalProperties,omitempty"`

}

// String returns a JSON representation of the model
func (o *Dialerdnclistconfigchangednclist) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

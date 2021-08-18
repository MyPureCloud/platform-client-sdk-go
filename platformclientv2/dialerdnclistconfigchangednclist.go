package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
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
	Version *int `json:"version,omitempty"`


	// ImportStatus
	ImportStatus *Dialerdnclistconfigchangeimportstatus `json:"importStatus,omitempty"`


	// Size
	Size *int `json:"size,omitempty"`


	// DncSourceType
	DncSourceType *string `json:"dncSourceType,omitempty"`


	// LoginId
	LoginId *string `json:"loginId,omitempty"`


	// DncCodes
	DncCodes *[]string `json:"dncCodes,omitempty"`


	// LicenseId
	LicenseId *string `json:"licenseId,omitempty"`


	// ContactMethod
	ContactMethod *string `json:"contactMethod,omitempty"`


	// Division
	Division *Dialerdnclistconfigchangeurireference `json:"division,omitempty"`


	// AdditionalProperties
	AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`

}

func (u *Dialerdnclistconfigchangednclist) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dialerdnclistconfigchangednclist

	
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
		
		ImportStatus *Dialerdnclistconfigchangeimportstatus `json:"importStatus,omitempty"`
		
		Size *int `json:"size,omitempty"`
		
		DncSourceType *string `json:"dncSourceType,omitempty"`
		
		LoginId *string `json:"loginId,omitempty"`
		
		DncCodes *[]string `json:"dncCodes,omitempty"`
		
		LicenseId *string `json:"licenseId,omitempty"`
		
		ContactMethod *string `json:"contactMethod,omitempty"`
		
		Division *Dialerdnclistconfigchangeurireference `json:"division,omitempty"`
		
		AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		Version: u.Version,
		
		ImportStatus: u.ImportStatus,
		
		Size: u.Size,
		
		DncSourceType: u.DncSourceType,
		
		LoginId: u.LoginId,
		
		DncCodes: u.DncCodes,
		
		LicenseId: u.LicenseId,
		
		ContactMethod: u.ContactMethod,
		
		Division: u.Division,
		
		AdditionalProperties: u.AdditionalProperties,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Dialerdnclistconfigchangednclist) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

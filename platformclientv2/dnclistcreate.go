package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Dnclistcreate
type Dnclistcreate struct { 
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


	// LoginId - A dnc.com loginId. Required if the dncSourceType is dnc.com.
	LoginId *string `json:"loginId,omitempty"`


	// DncCodes - The list of dnc.com codes to be treated as DNC. Required if the dncSourceType is dnc.com.
	DncCodes *[]string `json:"dncCodes,omitempty"`


	// LicenseId - A gryphon license number. Required if the dncSourceType is gryphon.
	LicenseId *string `json:"licenseId,omitempty"`


	// Division - The division this DncList belongs to.
	Division *Domainentityref `json:"division,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (u *Dnclistcreate) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dnclistcreate

	
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
		
		ImportStatus *Importstatus `json:"importStatus,omitempty"`
		
		Size *int `json:"size,omitempty"`
		
		DncSourceType *string `json:"dncSourceType,omitempty"`
		
		LoginId *string `json:"loginId,omitempty"`
		
		DncCodes *[]string `json:"dncCodes,omitempty"`
		
		LicenseId *string `json:"licenseId,omitempty"`
		
		Division *Domainentityref `json:"division,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
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
		
		Division: u.Division,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Dnclistcreate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

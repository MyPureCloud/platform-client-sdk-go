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

func (o *Dnclistcreate) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dnclistcreate
	
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
		
		LoginId *string `json:"loginId,omitempty"`
		
		DncCodes *[]string `json:"dncCodes,omitempty"`
		
		LicenseId *string `json:"licenseId,omitempty"`
		
		Division *Domainentityref `json:"division,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		Version: o.Version,
		
		ImportStatus: o.ImportStatus,
		
		Size: o.Size,
		
		DncSourceType: o.DncSourceType,
		
		LoginId: o.LoginId,
		
		DncCodes: o.DncCodes,
		
		LicenseId: o.LicenseId,
		
		Division: o.Division,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Dnclistcreate) UnmarshalJSON(b []byte) error {
	var DnclistcreateMap map[string]interface{}
	err := json.Unmarshal(b, &DnclistcreateMap)
	if err != nil {
		return err
	}
	
	if Id, ok := DnclistcreateMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := DnclistcreateMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if dateCreatedString, ok := DnclistcreateMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := DnclistcreateMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if Version, ok := DnclistcreateMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	
	if ImportStatus, ok := DnclistcreateMap["importStatus"].(map[string]interface{}); ok {
		ImportStatusString, _ := json.Marshal(ImportStatus)
		json.Unmarshal(ImportStatusString, &o.ImportStatus)
	}
	
	if Size, ok := DnclistcreateMap["size"].(float64); ok {
		SizeInt := int(Size)
		o.Size = &SizeInt
	}
	
	if DncSourceType, ok := DnclistcreateMap["dncSourceType"].(string); ok {
		o.DncSourceType = &DncSourceType
	}
	
	if LoginId, ok := DnclistcreateMap["loginId"].(string); ok {
		o.LoginId = &LoginId
	}
	
	if DncCodes, ok := DnclistcreateMap["dncCodes"].([]interface{}); ok {
		DncCodesString, _ := json.Marshal(DncCodes)
		json.Unmarshal(DncCodesString, &o.DncCodes)
	}
	
	if LicenseId, ok := DnclistcreateMap["licenseId"].(string); ok {
		o.LicenseId = &LicenseId
	}
	
	if Division, ok := DnclistcreateMap["division"].(map[string]interface{}); ok {
		DivisionString, _ := json.Marshal(Division)
		json.Unmarshal(DivisionString, &o.Division)
	}
	
	if SelfUri, ok := DnclistcreateMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Dnclistcreate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

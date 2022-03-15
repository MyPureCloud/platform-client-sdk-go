package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Dnclistdivisionview
type Dnclistdivisionview struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Division - The division to which this entity belongs.
	Division *Division `json:"division,omitempty"`


	// ImportStatus - The status of the import process.
	ImportStatus *Importstatus `json:"importStatus,omitempty"`


	// Size - The number of contacts in the DncList.
	Size *int `json:"size,omitempty"`


	// DncSourceType - The type of the DncList.
	DncSourceType *string `json:"dncSourceType,omitempty"`


	// ContactMethod - The contact method. Required if dncSourceType is rds.
	ContactMethod *string `json:"contactMethod,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Dnclistdivisionview) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dnclistdivisionview
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Division *Division `json:"division,omitempty"`
		
		ImportStatus *Importstatus `json:"importStatus,omitempty"`
		
		Size *int `json:"size,omitempty"`
		
		DncSourceType *string `json:"dncSourceType,omitempty"`
		
		ContactMethod *string `json:"contactMethod,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Division: o.Division,
		
		ImportStatus: o.ImportStatus,
		
		Size: o.Size,
		
		DncSourceType: o.DncSourceType,
		
		ContactMethod: o.ContactMethod,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Dnclistdivisionview) UnmarshalJSON(b []byte) error {
	var DnclistdivisionviewMap map[string]interface{}
	err := json.Unmarshal(b, &DnclistdivisionviewMap)
	if err != nil {
		return err
	}
	
	if Id, ok := DnclistdivisionviewMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := DnclistdivisionviewMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if Division, ok := DnclistdivisionviewMap["division"].(map[string]interface{}); ok {
		DivisionString, _ := json.Marshal(Division)
		json.Unmarshal(DivisionString, &o.Division)
	}
	
	if ImportStatus, ok := DnclistdivisionviewMap["importStatus"].(map[string]interface{}); ok {
		ImportStatusString, _ := json.Marshal(ImportStatus)
		json.Unmarshal(ImportStatusString, &o.ImportStatus)
	}
	
	if Size, ok := DnclistdivisionviewMap["size"].(float64); ok {
		SizeInt := int(Size)
		o.Size = &SizeInt
	}
	
	if DncSourceType, ok := DnclistdivisionviewMap["dncSourceType"].(string); ok {
		o.DncSourceType = &DncSourceType
	}
	
	if ContactMethod, ok := DnclistdivisionviewMap["contactMethod"].(string); ok {
		o.ContactMethod = &ContactMethod
	}
	
	if SelfUri, ok := DnclistdivisionviewMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Dnclistdivisionview) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Adherenceexplanationjobreference
type Adherenceexplanationjobreference struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// VarType - The type of the adherence explanation job
	VarType *string `json:"type,omitempty"`


	// Status - The status of the adherence explanation job
	Status *string `json:"status,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Adherenceexplanationjobreference) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Adherenceexplanationjobreference
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		VarType: o.VarType,
		
		Status: o.Status,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Adherenceexplanationjobreference) UnmarshalJSON(b []byte) error {
	var AdherenceexplanationjobreferenceMap map[string]interface{}
	err := json.Unmarshal(b, &AdherenceexplanationjobreferenceMap)
	if err != nil {
		return err
	}
	
	if Id, ok := AdherenceexplanationjobreferenceMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if VarType, ok := AdherenceexplanationjobreferenceMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Status, ok := AdherenceexplanationjobreferenceMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if SelfUri, ok := AdherenceexplanationjobreferenceMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Adherenceexplanationjobreference) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

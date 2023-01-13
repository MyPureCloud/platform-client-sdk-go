package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmhistoricaladherencebulkjobreference
type Wfmhistoricaladherencebulkjobreference struct { 
	// Id - The ID of the historical adherence bulk job to listen for via notification or query using the jobs route
	Id *string `json:"id,omitempty"`


	// Status - The status of the historical adherence bulk job
	Status *string `json:"status,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Wfmhistoricaladherencebulkjobreference) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmhistoricaladherencebulkjobreference
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Status: o.Status,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Wfmhistoricaladherencebulkjobreference) UnmarshalJSON(b []byte) error {
	var WfmhistoricaladherencebulkjobreferenceMap map[string]interface{}
	err := json.Unmarshal(b, &WfmhistoricaladherencebulkjobreferenceMap)
	if err != nil {
		return err
	}
	
	if Id, ok := WfmhistoricaladherencebulkjobreferenceMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Status, ok := WfmhistoricaladherencebulkjobreferenceMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if SelfUri, ok := WfmhistoricaladherencebulkjobreferenceMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmhistoricaladherencebulkjobreference) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

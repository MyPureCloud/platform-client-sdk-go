package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Documentupdate
type Documentupdate struct { 
	// ChangeNumber
	ChangeNumber *int `json:"changeNumber,omitempty"`


	// Name - The name of the document
	Name *string `json:"name,omitempty"`


	// Read
	Read *bool `json:"read,omitempty"`


	// AddTags
	AddTags *[]string `json:"addTags,omitempty"`


	// RemoveTags
	RemoveTags *[]string `json:"removeTags,omitempty"`


	// AddTagIds
	AddTagIds *[]string `json:"addTagIds,omitempty"`


	// RemoveTagIds
	RemoveTagIds *[]string `json:"removeTagIds,omitempty"`


	// UpdateAttributes
	UpdateAttributes *[]Documentattribute `json:"updateAttributes,omitempty"`


	// RemoveAttributes
	RemoveAttributes *[]string `json:"removeAttributes,omitempty"`

}

func (o *Documentupdate) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Documentupdate
	
	return json.Marshal(&struct { 
		ChangeNumber *int `json:"changeNumber,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Read *bool `json:"read,omitempty"`
		
		AddTags *[]string `json:"addTags,omitempty"`
		
		RemoveTags *[]string `json:"removeTags,omitempty"`
		
		AddTagIds *[]string `json:"addTagIds,omitempty"`
		
		RemoveTagIds *[]string `json:"removeTagIds,omitempty"`
		
		UpdateAttributes *[]Documentattribute `json:"updateAttributes,omitempty"`
		
		RemoveAttributes *[]string `json:"removeAttributes,omitempty"`
		*Alias
	}{ 
		ChangeNumber: o.ChangeNumber,
		
		Name: o.Name,
		
		Read: o.Read,
		
		AddTags: o.AddTags,
		
		RemoveTags: o.RemoveTags,
		
		AddTagIds: o.AddTagIds,
		
		RemoveTagIds: o.RemoveTagIds,
		
		UpdateAttributes: o.UpdateAttributes,
		
		RemoveAttributes: o.RemoveAttributes,
		Alias:    (*Alias)(o),
	})
}

func (o *Documentupdate) UnmarshalJSON(b []byte) error {
	var DocumentupdateMap map[string]interface{}
	err := json.Unmarshal(b, &DocumentupdateMap)
	if err != nil {
		return err
	}
	
	if ChangeNumber, ok := DocumentupdateMap["changeNumber"].(float64); ok {
		ChangeNumberInt := int(ChangeNumber)
		o.ChangeNumber = &ChangeNumberInt
	}
	
	if Name, ok := DocumentupdateMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Read, ok := DocumentupdateMap["read"].(bool); ok {
		o.Read = &Read
	}
    
	if AddTags, ok := DocumentupdateMap["addTags"].([]interface{}); ok {
		AddTagsString, _ := json.Marshal(AddTags)
		json.Unmarshal(AddTagsString, &o.AddTags)
	}
	
	if RemoveTags, ok := DocumentupdateMap["removeTags"].([]interface{}); ok {
		RemoveTagsString, _ := json.Marshal(RemoveTags)
		json.Unmarshal(RemoveTagsString, &o.RemoveTags)
	}
	
	if AddTagIds, ok := DocumentupdateMap["addTagIds"].([]interface{}); ok {
		AddTagIdsString, _ := json.Marshal(AddTagIds)
		json.Unmarshal(AddTagIdsString, &o.AddTagIds)
	}
	
	if RemoveTagIds, ok := DocumentupdateMap["removeTagIds"].([]interface{}); ok {
		RemoveTagIdsString, _ := json.Marshal(RemoveTagIds)
		json.Unmarshal(RemoveTagIdsString, &o.RemoveTagIds)
	}
	
	if UpdateAttributes, ok := DocumentupdateMap["updateAttributes"].([]interface{}); ok {
		UpdateAttributesString, _ := json.Marshal(UpdateAttributes)
		json.Unmarshal(UpdateAttributesString, &o.UpdateAttributes)
	}
	
	if RemoveAttributes, ok := DocumentupdateMap["removeAttributes"].([]interface{}); ok {
		RemoveAttributesString, _ := json.Marshal(RemoveAttributes)
		json.Unmarshal(RemoveAttributesString, &o.RemoveAttributes)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Documentupdate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

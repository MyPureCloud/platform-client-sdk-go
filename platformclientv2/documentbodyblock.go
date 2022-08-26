package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Documentbodyblock
type Documentbodyblock struct { 
	// VarType - The type of the block for the body. This determines which body block object (paragraph, list, video or image) would have a value.
	VarType *string `json:"type,omitempty"`


	// Paragraph - Paragraph. It must contain a value if the type of the block is Paragraph.
	Paragraph *Documentbodyparagraph `json:"paragraph,omitempty"`


	// Image - Image. It must contain a value if the type of the block is Image.
	Image *Documentbodyimage `json:"image,omitempty"`


	// Video - Video. It must contain a value if the type of the block is Video.
	Video *Documentbodyvideo `json:"video,omitempty"`


	// List - List. It must contain a value if the type of the block is UnorderedList or OrderedList.
	List *Documentbodylist `json:"list,omitempty"`

}

func (o *Documentbodyblock) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Documentbodyblock
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		Paragraph *Documentbodyparagraph `json:"paragraph,omitempty"`
		
		Image *Documentbodyimage `json:"image,omitempty"`
		
		Video *Documentbodyvideo `json:"video,omitempty"`
		
		List *Documentbodylist `json:"list,omitempty"`
		*Alias
	}{ 
		VarType: o.VarType,
		
		Paragraph: o.Paragraph,
		
		Image: o.Image,
		
		Video: o.Video,
		
		List: o.List,
		Alias:    (*Alias)(o),
	})
}

func (o *Documentbodyblock) UnmarshalJSON(b []byte) error {
	var DocumentbodyblockMap map[string]interface{}
	err := json.Unmarshal(b, &DocumentbodyblockMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := DocumentbodyblockMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Paragraph, ok := DocumentbodyblockMap["paragraph"].(map[string]interface{}); ok {
		ParagraphString, _ := json.Marshal(Paragraph)
		json.Unmarshal(ParagraphString, &o.Paragraph)
	}
	
	if Image, ok := DocumentbodyblockMap["image"].(map[string]interface{}); ok {
		ImageString, _ := json.Marshal(Image)
		json.Unmarshal(ImageString, &o.Image)
	}
	
	if Video, ok := DocumentbodyblockMap["video"].(map[string]interface{}); ok {
		VideoString, _ := json.Marshal(Video)
		json.Unmarshal(VideoString, &o.Video)
	}
	
	if List, ok := DocumentbodyblockMap["list"].(map[string]interface{}); ok {
		ListString, _ := json.Marshal(List)
		json.Unmarshal(ListString, &o.List)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Documentbodyblock) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

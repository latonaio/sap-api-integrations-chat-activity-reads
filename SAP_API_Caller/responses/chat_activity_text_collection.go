package responses

type ChatActivityTextCollection struct {
	D struct {
		Results []struct {
			Metadata struct {
				URI  string `json:"uri"`
				Type string `json:"type"`
				Etag string `json:"etag"`
			} `json:"__metadata"`
			ObjectID            string `json:"ObjectID"`
			ParentObjectID      string `json:"ParentObjectID"`
			ETag                string `json:"ETag"`
			Text                string `json:"Text"`
			LanguageCode        string `json:"LanguageCode"`
			LanguageCodeText    string `json:"LanguageCodeText"`
			TypeCode            string `json:"TypeCode"`
			TypeCodeText        string `json:"TypeCodeText"`
			AuthorUUID          string `json:"AuthorUUID"`
			AuthorName          string `json:"AuthorName"`
			CreatedOn           string `json:"CreatedOn"`
			CreatedBy           string `json:"CreatedBy"`
			UpdatedOn           string `json:"UpdatedOn"`
			LastUpdatedBy       string `json:"LastUpdatedBy"`
			EntityLastChangedOn string `json:"EntityLastChangedOn"`
		} `json:"results"`
	} `json:"d"`
}

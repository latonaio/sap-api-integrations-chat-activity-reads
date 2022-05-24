package responses

type ToChatActivityParties struct {
	D struct {
		Results []struct {
			Metadata struct {
				URI  string `json:"uri"`
				Type string `json:"type"`
				Etag string `json:"etag"`
			} `json:"__metadata"`
			ObjectID             string `json:"ObjectID"`
			ParentObjectID       string `json:"ParentObjectID"`
			ETag                 string `json:"ETag"`
			ChatActivityID       string `json:"ChatActivityID"`
			PartyUUID            string `json:"PartyUUID"`
			PartyID              string `json:"PartyID"`
			PartyTypeCode        string `json:"PartyTypeCode"`
			PartyTypeCodeText    string `json:"PartyTypeCodeText"`
			RoleCode             string `json:"RoleCode"`
			RoleCodeText         string `json:"RoleCodeText"`
			RoleCategoryCode     string `json:"RoleCategoryCode"`
			RoleCategoryCodeText string `json:"RoleCategoryCodeText"`
			PartyEmailURI        string `json:"PartyEmailURI"`
			MainIndicator        bool   `json:"MainIndicator"`
			PartyName            string `json:"PartyName"`
			ToParentObjectID     string `json:"ToParent_ObjectID"`
		} `json:"results"`
	} `json:"d"`
}

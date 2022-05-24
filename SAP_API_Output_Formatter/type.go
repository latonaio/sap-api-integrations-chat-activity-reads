package sap_api_output_formatter

type ChatActivity struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	APISchema     string `json:"api_schema"`
	Deleted       bool   `json:"deleted"`
}

type ChatActivityCollection struct {
	ObjectID                                             string `json:"ObjectID"`
	ETag                                                 string `json:"ETag"`
	TypeCode                                             string `json:"TypeCode"`
	TypeCodeText                                         string `json:"TypeCodeText"`
	LifeCycleStatusCode                                  string `json:"LifeCycleStatusCode"`
	LifeCycleStatusCodeText                              string `json:"LifeCycleStatusCodeText"`
	LastChangeDate                                       string `json:"LastChangeDate"`
	ActivityWorklistItemUUID                             string `json:"ActivityWorklistItemUUID"`
	ActualDuration                                       string `json:"ActualDuration"`
	ActualEndDateTime                                    string `json:"ActualEndDateTime"`
	ActualStartDateTime                                  string `json:"ActualStartDateTime"`
	ActivitySentimentTypeCode                            string `json:"ActivitySentimentTypeCode"`
	ActivitySentimentTypeCodeText                        string `json:"ActivitySentimentTypeCodeText"`
	AdditionalLocationName                               string `json:"AdditionalLocationName"`
	AvailabilityCode                                     string `json:"AvailabilityCode"`
	AvailabilityCodeText                                 string `json:"AvailabilityCodeText"`
	BouncedIndicator                                     bool   `json:"BouncedIndicator"`
	CompletionDateTime                                   string `json:"CompletionDateTime"`
	CompletionPercent                                    string `json:"CompletionPercent"`
	CorrespondenceTransmissionStatusCode                 string `json:"CorrespondenceTransmissionStatusCode"`
	CorrespondenceTransmissionStatusCodeText             string `json:"CorrespondenceTransmissionStatusCodeText"`
	CreationDate                                         string `json:"CreationDate"`
	DataOriginTypeCode                                   string `json:"DataOriginTypeCode"`
	DataOriginTypeCodeText                               string `json:"DataOriginTypeCodeText"`
	DistributionChannelCode                              string `json:"DistributionChannelCode"`
	DistributionChannelCodeText                          string `json:"DistributionChannelCodeText"`
	DistributionChannelDeterminationMethodCode           string `json:"DistributionChannelDeterminationMethodCode"`
	DistributionChannelDeterminationMethodCodeText       string `json:"DistributionChannelDeterminationMethodCodeText"`
	DivisionCode                                         string `json:"DivisionCode"`
	DivisionCodeText                                     string `json:"DivisionCodeText"`
	DivisionDeterminationMethodCode                      string `json:"DivisionDeterminationMethodCode"`
	DivisionDeterminationMethodCodeText                  string `json:"DivisionDeterminationMethodCodeText"`
	ExternalName                                         string `json:"ExternalName"`
	FullDayIndicator                                     bool   `json:"FullDayIndicator"`
	GroupwareItemID                                      string `json:"GroupwareItemID"`
	GroupCode                                            string `json:"GroupCode"`
	GroupCodeText                                        string `json:"GroupCodeText"`
	GroupwareSynchronizationNonRelevanceIndicator        bool   `json:"GroupwareSynchronizationNonRelevanceIndicator"`
	ID                                                   string `json:"ID"`
	InformationSensitivityCode                           string `json:"InformationSensitivityCode"`
	InformationSensitivityCodeText                       string `json:"InformationSensitivityCodeText"`
	InitiatingActivityUUID                               string `json:"InitiatingActivityUUID"`
	InitiatorCode                                        string `json:"InitiatorCode"`
	InitiatorCodeText                                    string `json:"InitiatorCodeText"`
	LocationName                                         string `json:"LocationName"`
	MarkedForDeletionDate                                string `json:"MarkedForDeletionDate"`
	MarkedForDeletion                                    bool   `json:"MarkedForDeletion"`
	MigratedDataAdaptationTypeCode                       string `json:"MigratedDataAdaptationTypeCode"`
	MigratedDataAdaptationTypeCodeText                   string `json:"MigratedDataAdaptationTypeCodeText"`
	PerfectStoreIndicator                                bool   `json:"PerfectStoreIndicator"`
	PlannedDuration                                      string `json:"PlannedDuration"`
	PredecessorActivityUUID                              string `json:"PredecessorActivityUUID"`
	PriorityCode                                         string `json:"PriorityCode"`
	PriorityCodeText                                     string `json:"PriorityCodeText"`
	ProcessingTypeCode                                   string `json:"ProcessingTypeCode"`
	ProcessingTypeCodeText                               string `json:"ProcessingTypeCodeText"`
	ReportedDate                                         string `json:"ReportedDate"`
	ReportedDateTime                                     string `json:"ReportedDateTime"`
	SalesOrganisationID                                  string `json:"SalesOrganisationID"`
	SalesOrganisationUUID                                string `json:"SalesOrganisationUUID"`
	SalesOrganizationDeterminationMethodCode             string `json:"SalesOrganizationDeterminationMethodCode"`
	SalesOrganizationDeterminationMethodCodeText         string `json:"SalesOrganizationDeterminationMethodCodeText"`
	SalesTerritoryDeterminationMethodCode                string `json:"SalesTerritoryDeterminationMethodCode"`
	SalesTerritoryDeterminationMethodCodeText            string `json:"SalesTerritoryDeterminationMethodCodeText"`
	SalesTerritoryID                                     string `json:"SalesTerritoryID"`
	SalesTerritoryUUID                                   string `json:"SalesTerritoryUUID"`
	ScheduledEndDate                                     string `json:"ScheduledEndDate"`
	ScheduledStartDate                                   string `json:"ScheduledStartDate"`
	SilentDataMigrationID                                string `json:"SilentDataMigrationID"`
	SocialMediaActivityProviderUUID                      string `json:"SocialMediaActivityProviderUUID"`
	ActivityFollowUpServiceRequestBlockingReasonCode     string `json:"ActivityFollowUpServiceRequestBlockingReasonCode"`
	ActivityFollowUpServiceRequestBlockingReasonCodeText string `json:"ActivityFollowUpServiceRequestBlockingReasonCodeText"`
	StatusSchemaName                                     string `json:"StatusSchemaName"`
	SubjectName                                          string `json:"SubjectName"`
	TransmittedDate                                      string `json:"TransmittedDate"`
	VisitWorkItemGenerationCode                          string `json:"VisitWorkItemGenerationCode"`
	VisitWorkItemGenerationCodeText                      string `json:"VisitWorkItemGenerationCodeText"`
	VisitTypeCode                                        string `json:"VisitTypeCode"`
	VisitTypeCodeText                                    string `json:"VisitTypeCodeText"`
	VisitTourPlanUUID                                    string `json:"VisitTourPlanUUID"`
	ChatMainPartyUUID                                    string `json:"ChatMainPartyUUID"`
	ChatMainPartyID                                      string `json:"ChatMainPartyID"`
	ChatMainEmailURI                                     string `json:"ChatMainEmailURI"`
	EntityLastChangedOn                                  string `json:"EntityLastChangedOn"`
	ToChatActivityParties                                string `json:"ToChatActivityParties"`
}

type ChatActivityTextCollection struct {
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
}

type ChatActivityParties struct {
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
}

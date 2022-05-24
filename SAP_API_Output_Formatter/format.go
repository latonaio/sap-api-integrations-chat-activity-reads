package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-chat-activity-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToChatActivityCollection(raw []byte, l *logger.Logger) ([]ChatActivityCollection, error) {
	pm := &responses.ChatActivityCollection{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ChatActivityCollection. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	chatActivityCollection := make([]ChatActivityCollection, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		chatActivityCollection = append(chatActivityCollection, ChatActivityCollection{
			ObjectID:                                   data.ObjectID,
			ETag:                                       data.ETag,
			TypeCode:                                   data.TypeCode,
			TypeCodeText:                               data.TypeCodeText,
			LifeCycleStatusCode:                        data.LifeCycleStatusCode,
			LifeCycleStatusCodeText:                    data.LifeCycleStatusCodeText,
			LastChangeDate:                             data.LastChangeDate,
			ActivityWorklistItemUUID:                   data.ActivityWorklistItemUUID,
			ActualDuration:                             data.ActualDuration,
			ActualEndDateTime:                          data.ActualEndDateTime,
			ActualStartDateTime:                        data.ActualStartDateTime,
			ActivitySentimentTypeCode:                  data.ActivitySentimentTypeCode,
			ActivitySentimentTypeCodeText:              data.ActivitySentimentTypeCodeText,
			AdditionalLocationName:                     data.AdditionalLocationName,
			AvailabilityCode:                           data.AvailabilityCode,
			AvailabilityCodeText:                       data.AvailabilityCodeText,
			BouncedIndicator:                           data.BouncedIndicator,
			CompletionDateTime:                         data.CompletionDateTime,
			CompletionPercent:                          data.CompletionPercent,
			CorrespondenceTransmissionStatusCode:       data.CorrespondenceTransmissionStatusCode,
			CorrespondenceTransmissionStatusCodeText:   data.CorrespondenceTransmissionStatusCodeText,
			CreationDate:                               data.CreationDate,
			DataOriginTypeCode:                         data.DataOriginTypeCode,
			DataOriginTypeCodeText:                     data.DataOriginTypeCodeText,
			DistributionChannelCode:                    data.DistributionChannelCode,
			DistributionChannelCodeText:                data.DistributionChannelCodeText,
			DistributionChannelDeterminationMethodCode: data.DistributionChannelDeterminationMethodCode,
			DistributionChannelDeterminationMethodCodeText: data.DistributionChannelDeterminationMethodCodeText,
			DivisionCode:                        data.DivisionCode,
			DivisionCodeText:                    data.DivisionCodeText,
			DivisionDeterminationMethodCode:     data.DivisionDeterminationMethodCode,
			DivisionDeterminationMethodCodeText: data.DivisionDeterminationMethodCodeText,
			ExternalName:                        data.ExternalName,
			FullDayIndicator:                    data.FullDayIndicator,
			GroupwareItemID:                     data.GroupwareItemID,
			GroupCode:                           data.GroupCode,
			GroupCodeText:                       data.GroupCodeText,
			GroupwareSynchronizationNonRelevanceIndicator: data.GroupwareSynchronizationNonRelevanceIndicator,
			ID:                                               data.ID,
			InformationSensitivityCode:                       data.InformationSensitivityCode,
			InformationSensitivityCodeText:                   data.InformationSensitivityCodeText,
			InitiatingActivityUUID:                           data.InitiatingActivityUUID,
			InitiatorCode:                                    data.InitiatorCode,
			InitiatorCodeText:                                data.InitiatorCodeText,
			LocationName:                                     data.LocationName,
			MarkedForDeletionDate:                            data.MarkedForDeletionDate,
			MarkedForDeletion:                                data.MarkedForDeletion,
			MigratedDataAdaptationTypeCode:                   data.MigratedDataAdaptationTypeCode,
			MigratedDataAdaptationTypeCodeText:               data.MigratedDataAdaptationTypeCodeText,
			PerfectStoreIndicator:                            data.PerfectStoreIndicator,
			PlannedDuration:                                  data.PlannedDuration,
			PredecessorActivityUUID:                          data.PredecessorActivityUUID,
			PriorityCode:                                     data.PriorityCode,
			PriorityCodeText:                                 data.PriorityCodeText,
			ProcessingTypeCode:                               data.ProcessingTypeCode,
			ProcessingTypeCodeText:                           data.ProcessingTypeCodeText,
			ReportedDate:                                     data.ReportedDate,
			ReportedDateTime:                                 data.ReportedDateTime,
			SalesOrganisationID:                              data.SalesOrganisationID,
			SalesOrganisationUUID:                            data.SalesOrganisationUUID,
			SalesOrganizationDeterminationMethodCode:         data.SalesOrganizationDeterminationMethodCode,
			SalesOrganizationDeterminationMethodCodeText:     data.SalesOrganizationDeterminationMethodCodeText,
			SalesTerritoryDeterminationMethodCode:            data.SalesTerritoryDeterminationMethodCode,
			SalesTerritoryDeterminationMethodCodeText:        data.SalesTerritoryDeterminationMethodCodeText,
			SalesTerritoryID:                                 data.SalesTerritoryID,
			SalesTerritoryUUID:                               data.SalesTerritoryUUID,
			ScheduledEndDate:                                 data.ScheduledEndDate,
			ScheduledStartDate:                               data.ScheduledStartDate,
			SilentDataMigrationID:                            data.SilentDataMigrationID,
			SocialMediaActivityProviderUUID:                  data.SocialMediaActivityProviderUUID,
			ActivityFollowUpServiceRequestBlockingReasonCode: data.ActivityFollowUpServiceRequestBlockingReasonCode,
			ActivityFollowUpServiceRequestBlockingReasonCodeText: data.ActivityFollowUpServiceRequestBlockingReasonCodeText,
			StatusSchemaName:                data.StatusSchemaName,
			SubjectName:                     data.SubjectName,
			TransmittedDate:                 data.TransmittedDate,
			VisitWorkItemGenerationCode:     data.VisitWorkItemGenerationCode,
			VisitWorkItemGenerationCodeText: data.VisitWorkItemGenerationCodeText,
			VisitTypeCode:                   data.VisitTypeCode,
			VisitTypeCodeText:               data.VisitTypeCodeText,
			VisitTourPlanUUID:               data.VisitTourPlanUUID,
			ChatMainPartyUUID:               data.ChatMainPartyUUID,
			ChatMainPartyID:                 data.ChatMainPartyID,
			ChatMainEmailURI:                data.ChatMainEmailURI,
			EntityLastChangedOn:             data.EntityLastChangedOn,
			ToChatActivityParties:           data.ChatActivityParties.Deferred.URI,
		})
	}

	return chatActivityCollection, nil
}

func ConvertToChatActivityTextCollection(raw []byte, l *logger.Logger) ([]ChatActivityTextCollection, error) {
	pm := &responses.ChatActivityTextCollection{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ChatActivityTextCollection. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	chatActivityTextCollection := make([]ChatActivityTextCollection, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		chatActivityTextCollection = append(chatActivityTextCollection, ChatActivityTextCollection{
			ObjectID:            data.ObjectID,
			ParentObjectID:      data.ParentObjectID,
			ETag:                data.ETag,
			Text:                data.Text,
			LanguageCode:        data.LanguageCode,
			LanguageCodeText:    data.LanguageCodeText,
			TypeCode:            data.TypeCode,
			TypeCodeText:        data.TypeCodeText,
			AuthorUUID:          data.AuthorUUID,
			AuthorName:          data.AuthorName,
			CreatedOn:           data.CreatedOn,
			CreatedBy:           data.CreatedBy,
			UpdatedOn:           data.UpdatedOn,
			LastUpdatedBy:       data.LastUpdatedBy,
			EntityLastChangedOn: data.EntityLastChangedOn,
		})
	}

	return chatActivityTextCollection, nil
}

func ConvertToChatActivityParties(raw []byte, l *logger.Logger) ([]ChatActivityParties, error) {
	pm := &responses.ToChatActivityParties{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ChatActivityParties. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	chatActivityParties := make([]ChatActivityParties, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		chatActivityParties = append(chatActivityParties, ChatActivityParties{
			ObjectID:             data.ObjectID,
			ParentObjectID:       data.ParentObjectID,
			ETag:                 data.ETag,
			ChatActivityID:       data.ChatActivityID,
			PartyUUID:            data.PartyUUID,
			PartyID:              data.PartyID,
			PartyTypeCode:        data.PartyTypeCode,
			PartyTypeCodeText:    data.PartyTypeCodeText,
			RoleCode:             data.RoleCode,
			RoleCodeText:         data.RoleCodeText,
			RoleCategoryCode:     data.RoleCategoryCode,
			RoleCategoryCodeText: data.RoleCategoryCodeText,
			PartyEmailURI:        data.PartyEmailURI,
			MainIndicator:        data.MainIndicator,
			PartyName:            data.PartyName,
			ToParentObjectID:     data.ToParentObjectID,
		})
	}

	return chatActivityParties, nil
}

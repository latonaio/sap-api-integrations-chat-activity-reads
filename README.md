# sap-api-integrations-chat-activity-reads  
sap-api-integrations-chat-activity-reads は、外部システム(特にエッジコンピューティング環境)をSAPと統合することを目的に、SAP API チャットアクティビティデータを取得するマイクロサービスです。  
sap-api-integrations-chat-activity-reads には、サンプルのAPI Json フォーマットが含まれています。  
sap-api-integrations-chat-activity-reads は、オンプレミス版である（＝クラウド版ではない）SAPC4HANA API の利用を前提としています。クラウド版APIを利用する場合は、ご注意ください。  
https://api.sap.com/api/chatactivity/overview

## 動作環境
sap-api-integrations-chat-activity-reads は、主にエッジコンピューティング環境における動作にフォーカスしています。   
使用する際は、事前に下記の通り エッジコンピューティングの動作環境（推奨/必須）を用意してください。   
・ エッジ Kubernetes （推奨）    
・ AION のリソース （推奨)    
・ OS: LinuxOS （必須）    
・ CPU: ARM/AMD/Intel（いずれか必須） 

## クラウド環境での利用  
sap-api-integrations-chat-activity-reads は、外部システムがクラウド環境である場合にSAPと統合するときにおいても、利用可能なように設計されています。  

## 本レポジトリ が 対応する API サービス
sap-api-integrations-chat-activity-reads が対応する APIサービス は、次のものです。

* APIサービス概要説明 URL: https://api.sap.com/api/chatactivity/overview
* APIサービス名(=baseURL): c4codataapi

## 本レポジトリ に 含まれる API名
sap-api-integrations-chat-activity-reads には、次の API をコールするためのリソースが含まれています。  

* ChatActivityCollection（チャットアクティビティ - チャットアクティビティ）※会話履歴の詳細データを取得するために、ToChatActivityParties、と合わせて利用されます。
* ToChatActivityParties（チャットアクティビティ -  チャット情報※To）
* ChatActivityTextCollection（チャットアクティビティ - テキスト）

## API への 値入力条件 の 初期値
sap-api-integrations-chat-activity-reads において、API への値入力条件の初期値は、入力ファイルレイアウトの種別毎に、次の通りとなっています。  

### SDC レイアウト

* inoutSDC.ChatActivityCollection.ID（ID）
* inoutSDC.ChatActivityCollection.ChatActivityParties（チャット情報）  
* ChatActivityTextCollection.Text（テキスト）

## SAP API Bussiness Hub の API の選択的コール

Latona および AION の SAP 関連リソースでは、Inputs フォルダ下の sample.json の accepter に取得したいデータの種別（＝APIの種別）を入力し、指定することができます。  
なお、同 accepter にAll(もしくは空白)の値を入力することで、全データ（＝全APIの種別）をまとめて取得することができます。  

* sample.jsonの記載例(1)  

accepter において 下記の例のように、データの種別（＝APIの種別）を指定します。  
ここでは、"ChatActivityCollection" が指定されています。    
  
```
    "api_schema":  "ChatActivityCollection",
	"accepter": ["ChatActivityCollection"],
	"chat_activity_code": "3",
	"deleted": false
```
  
* 全データを取得する際のsample.jsonの記載例(2)  

全データを取得する場合、sample.json は以下のように記載します。  

```
    "api_schema":  "ChatActivityCollection",
	"accepter": ["All"],
	"chat_activity_code": "3",
	"deleted": false
```

## 指定されたデータ種別のコール

accepter における データ種別 の指定に基づいて SAP_API_Caller 内の caller.go で API がコールされます。  
caller.go の func() 毎 の 以下の箇所が、指定された API をコールするソースコードです。  

```
func (c *SAPAPICaller) AsyncGetChatActivityCollection(iD, text string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "ChatActivityCollection":
			func() {
				c.ChatActivityCollection(iD)
				wg.Done()
			}()
		case "ChatActivityTextCollection":
			func() {
				c.ChatActivityTextCollection(text)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}
```

## Output  
本マイクロサービスでは、[golang-logging-library-for-sap](https://github.com/latonaio/golang-logging-library-for-sap) により、以下のようなデータがJSON形式で出力されます。  
以下の sample.json の例は、SAP キャンペーン  の キャンペーンデータ が取得された結果の JSON の例です。  
以下の項目のうち、"ObjectID" ～ "ToChatActivityParties" は、/SAP_API_Output_Formatter/type.go 内 の Type ChatActivityCollection {} による出力結果です。"cursor" ～ "time"は、golang-logging-library-for-sap による 定型フォーマットの出力結果です。  

```
{
	"cursor": "/Users/latona2/bitbucket/sap-api-integrations-chat-activity-reads/SAP_API_Caller/caller.go#L58",
	"function": "sap-api-integrations-chat-activity-reads/SAP_API_Caller.(*SAPAPICaller).ChatActivityCollection",
	"level": "INFO",
	"message": [
		{
			"ObjectID": "00163E063FDC1ED48AA2AECE7D596AE4",
			"ETag": "2014-08-21T18:28:27+09:00",
			"TypeCode": "1976",
			"TypeCodeText": "Chat",
			"LifeCycleStatusCode": "3",
			"LifeCycleStatusCodeText": "Completed",
			"LastChangeDate": "2014-08-21T09:00:00+09:00",
			"ActivityWorklistItemUUID": "",
			"ActualDuration": "",
			"ActualEndDateTime": "",
			"ActualStartDateTime": "",
			"ActivitySentimentTypeCode": "",
			"ActivitySentimentTypeCodeText": "",
			"AdditionalLocationName": "",
			"AvailabilityCode": "",
			"AvailabilityCodeText": "",
			"BouncedIndicator": false,
			"CompletionDateTime": "2014-08-21T18:28:26+09:00",
			"CompletionPercent": "0.00",
			"CorrespondenceTransmissionStatusCode": "",
			"CorrespondenceTransmissionStatusCodeText": "",
			"CreationDate": "2014-08-21T09:00:00+09:00",
			"DataOriginTypeCode": "1",
			"DataOriginTypeCodeText": "Manual Data Entry",
			"DistributionChannelCode": "",
			"DistributionChannelCodeText": "",
			"DistributionChannelDeterminationMethodCode": "",
			"DistributionChannelDeterminationMethodCodeText": "",
			"DivisionCode": "",
			"DivisionCodeText": "",
			"DivisionDeterminationMethodCode": "",
			"DivisionDeterminationMethodCodeText": "",
			"ExternalName": "",
			"FullDayIndicator": false,
			"GroupwareItemID": "",
			"GroupCode": "",
			"GroupCodeText": "",
			"GroupwareSynchronizationNonRelevanceIndicator": false,
			"ID": "3",
			"InformationSensitivityCode": "1",
			"InformationSensitivityCodeText": "Normal",
			"InitiatingActivityUUID": "",
			"InitiatorCode": "2",
			"InitiatorCodeText": "Inbound",
			"LocationName": "",
			"MarkedForDeletionDate": "",
			"MarkedForDeletion": false,
			"MigratedDataAdaptationTypeCode": "",
			"MigratedDataAdaptationTypeCodeText": "",
			"PerfectStoreIndicator": false,
			"PlannedDuration": "",
			"PredecessorActivityUUID": "",
			"PriorityCode": "3",
			"PriorityCodeText": "Normal",
			"ProcessingTypeCode": "0007",
			"ProcessingTypeCodeText": "Chat Activity",
			"ReportedDate": "2014-08-21T09:00:00+09:00",
			"ReportedDateTime": "2014-08-21T18:28:09+09:00",
			"SalesOrganisationID": "",
			"SalesOrganisationUUID": "",
			"SalesOrganizationDeterminationMethodCode": "",
			"SalesOrganizationDeterminationMethodCodeText": "",
			"SalesTerritoryDeterminationMethodCode": "",
			"SalesTerritoryDeterminationMethodCodeText": "",
			"SalesTerritoryID": "",
			"SalesTerritoryUUID": "",
			"ScheduledEndDate": "",
			"ScheduledStartDate": "",
			"SilentDataMigrationID": "M",
			"SocialMediaActivityProviderUUID": "",
			"ActivityFollowUpServiceRequestBlockingReasonCode": "",
			"ActivityFollowUpServiceRequestBlockingReasonCodeText": "",
			"StatusSchemaName": "ACTIVITY_INTERACTION_STANDARD",
			"SubjectName": "Chat From:pavithra.n@sap.com",
			"TransmittedDate": "",
			"VisitWorkItemGenerationCode": "",
			"VisitWorkItemGenerationCodeText": "",
			"VisitTypeCode": "",
			"VisitTypeCodeText": "",
			"VisitTourPlanUUID": "",
			"ChatMainPartyUUID": "00163E06-3FDC-1ED4-8AA2-670F0CC829EE",
			"ChatMainPartyID": "1001461",
			"ChatMainEmailURI": "pavithra.n@sap.com",
			"EntityLastChangedOn": "2014-08-21T18:28:27+09:00",
			"ToChatActivityParties": "https://sandbox.api.sap.com/sap/c4c/odata/v1/c4codataapi/ChatActivityCollection('00163E063FDC1ED48AA2AECE7D596AE4')/ChatActivityParties"
		}
	],
	"time": "2022-05-21T18:13:04+09:00"
}

```
package dsl

import (
	"testing"
)

const json string = `{
	"id": "bf6178d0a3ec5be07d316eb34ff5c848",
	"user": {
	  "id": "6a56acf793c3d2eb192c5f8df37133ed298dd80a"
	},
	"site": {
	  "id": 1,
	  "content": {
		"url": "www.iqiyi.com",
		"len": 1200
	  }
	},
	"device": {
	  "ua": "Mozilla/5.0 (iPhone; CPU iPhone OS 15_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 QIYIVideo/14.2.0",
	  "ip": "112.32.50.136",
	  "geo": {
		"country": 86,
		"metro": 8634,
		"city": 863401
	  },
	  "connectionType": 2,
	  "platformId": 32,
	  "model": "iPhone14,5",
	  "os": "ios",
	  "osVersion": "15.2",
	  "appVersion": "14.2.0",
	  "screenHeight": 2532,
	  "screenWidth": 1170,
	  "carrierName": "中国移动",
	  "idfv": "80802AAF-6915-48DE-9226-84599F5066B4",
	  "localTimezone": "Asia/Shanghai",
	  "osUpdateTime": "1644848755117922",
	  "startupTime": "1675959974855187",
	  "cpuNum": 6,
	  "diskTotal": "255866781696",
	  "memTotal": "3866853376",
	  "authStatus": 2
	},
	"imp": [{
	  "id": "0",
	  "banner": {
		"adZoneId": "30074",
		"adType": 0
	  },
	  "bidfloor": 1200,
	  "campaignId": 81001108,
	  "isPmp": false,
	  "extendedAdsPosition": 0,
	  "impressionDate": 20230302,
	  "maxSkippableRollAds": 0,
	  "skippableRollBidfloor": 0
	}]
  }`

const json2 = `{
    "requestTime": 1677686450423,
    "id": "bf6178d0a3ec5be07d316eb34ff5c848",
    "exId": 5,
    "impList": [
        {
            "impId": "0",
            "templates": [
                {
                    "specs": [],
                    "title": null,
                    "desc": null,
                    "content": null,
                    "icon": null,
                    "dimension": ""
                }
            ],
            "adspaceId": "30074",
            "bidFloor": 12000,
            "cur": "RMB",
            "deals": null,
            "secure": false,
            "pmp": false
        }
    ],
    "test": 0,
    "site": {
        "name": null,
        "page": null,
        "domain": null,
        "referrer": null,
        "categories": null,
        "userAgent": null
    },
    "app": null,
    "device": {
        "deviceId": "80802AAF-6915-48DE-9226-84599F5066B4",
        "userId": "7443acec95767fd88d9a19b58aa7bd48",
        "idType": "IDFV",
        "idEncoding": "Raw",
        "userAgent": "Mozilla/5.0 (iPhone; CPU iPhone OS 15_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 QIYIVideo/14.2.0",
        "ip": "112.32.50.136",
        "region": "340100",
        "deviceType": "PHONE",
        "make": "",
        "model": "iPhone14,5",
        "os": "IOS",
        "network": "WIFI",
        "carrier": "CARRIER_UNKNOWN",
        "osVersion": "15.2",
        "screenWidth": 0,
        "screenHeight": 0,
        "dpi": 0,
        "geo": null,
        "userTags": null,
        "mac": "",
        "macmd5": "",
        "macroImei": "",
        "macroMac": "",
        "macroMac1": ""
    },
    "width": 0,
    "height": 0,
    "bcat": null,
    "badv": null,
    "isPing": false,
    "isDebug": false
}`

func TestBasic(t *testing.T) {

	je := JE{}
	je.parse("", json)
}

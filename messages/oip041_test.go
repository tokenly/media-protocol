package messages

import (
	"reflect"
	"testing"
)

func TestDecodeOIP041(t *testing.T) {
	// ToDo: a better test/errors
	oip041, err := DecodeOIP041(oip041_music_example)
	if err != nil {
		t.Error("error")
	}
	if !reflect.DeepEqual(oip041, oip041_music_example_obj) {
		t.Error("not equal")
	}
}

var oip041_music_example = `{
  "oip-041": {
    "artifact": {
      "publisher": "F97Tp8LYnw94CpXmAhqACXWTT36jyvLCWx",
      "timestamp": 1470269387,
      "type": "music",
      "storage":{
        "network": "IPFS",
        "location": "QmPukCZKeJD4KZFtstpvrguLaq94rsWfBxLU1QoZxvgRxA"
      },
      "files": [
	    {
	      "dname": "Skipping Stones",
	      "fname": "1 - Skipping Stones.mp3",
	      "fsize": 6515667,
	      "type": "album track",
	      "duration": 1533.603293,
	      "sugPlay": 100,
	      "minPlay": null,
	      "sugBuy": 750,
	      "minBuy": 500,
	      "promo": 10,
	      "retail": 15,
	      "ptpFT": 10,
	      "ptpDT": 20,
	      "ptpDA": 50
	    },
	    {
	      "dname": "Lessons",
	      "fname": "2 - Lessons with intro.mp3",
	      "fsize": 6515667,
	      "type": "album track",
	      "duration": 1231.155243,
	      "disallowPlay": 1,
	      "sugBuy": 750,
	      "minBuy": 500,
	      "promo": 10,
	      "retail": 15,
	      "ptpFT": 10,
	      "ptpDT": 20,
	      "ptpDA": 50
	    },
	    {
	      "dname": "Born to Roam",
	      "fname": "3 - Born to Roam.mp3",
	      "fsize": 6515667,
	      "type": "album track",
	      "duration": 2374.550714,
	      "sugPlay": 100,
	      "minPlay": 50,
	      "disallowBuy": 1,
	      "promo": 10,
	      "retail": 15,
	      "ptpFT": 10,
	      "ptpDT": 20,
	      "ptpDA": 50
	    },
	    {
	      "dname": "Cover Art",
	      "fname": "birthdayepFINAL.jpg",
	      "type": "coverArt",
	      "disallowBuy": 1
	    }
	  ],
      "info": {
        "title": "Happy Birthday EP",
        "description": "this is the second organically grown, gluten free album released by Adam B. Levine - contact adam@tokenly.com with questions or comments or discuss collaborations.",
        "year": 2016,
        "extra-info": {
          "artist": "Adam B. Levine",
          "company": "",
          "composers": [
            "Adam B. Levine"
          ],
          "copyright": "",
          "tokenly_ID": "",
          "usageProhibitions": "",
          "usageRights": "",
          "tags": []
        }
      },
      "payment": {
        "fiat": "USD",
        "scale": "1000:1",
        "sug_tip": [
          5,
          50,
          100
        ],
        "tokens": {
          "MTMCOLLECTOR": "",
          "MTMPRODUCER": "",
          "HAPPYBDAYEP": "",
          "EARLY": "",
          "LTBCOIN": "",
          "BTC": "1GMMg2J5iUKnDf5PbRr9TcKV3R6KfUiB55"
        }
      }
    },
    "signature": "H27r7UxUb8BozjEvV0v++nCyRI7S6yyroeKCJQpgU5NO3CP6FpXWs5kCxy8vhmMhbtpj/FMj+8s3+updw7g+bmE="
  }
}`

var oip041_music_example_obj Oip041 = Oip041{
	Artifact: Oip041Artifact{
		Publisher: "F97Tp8LYnw94CpXmAhqACXWTT36jyvLCWx",
		Timestamp: 1470269387,
		Type:      "music",
		Info: Oip041Info{
			Title:       "Happy Birthday EP",
			Description: "this is the second organically grown, gluten free album released by Adam B. Levine - contact adam@tokenly.com with questions or comments or discuss collaborations.",
			Year:        2016,
			ExtraInfo: Oip041MusicExtraInfo{
				Artist:  "Adam B. Levine",
				Company: "",
				Composers: []string{
					"Adam B. Levine",
				},
				Copyright:         "",
				UsageProhibitions: "",
				UsageRights:       "",
				Tags:              []string{},
			},
			ExtraInfoString: "",
		},
		Storage: Oip041Storage{
			Network:  "IPFS",
			Location: "QmPukCZKeJD4KZFtstpvrguLaq94rsWfBxLU1QoZxvgRxA"},
		Files: []Oip041Files{
			{
				DisallowBuy:  0,
				Dname:        "Skipping Stones",
				Duration:     1533.603293,
				Fname:        "1 - Skipping Stones.mp3",
				Fsize:        6515667,
				MinPlay:      0,
				SugPlay:      100,
				Promo:        10,
				Retail:       15,
				PtpFT:        10,
				PtpDT:        20,
				PtpDA:        50,
				Type:         "album track",
				TokenlyID:    "",
				DisallowPlay: 0,
				MinBuy:       500,
				SugBuy:       750,
				Storage:      Oip041Storage{},
			},
			{
				DisallowBuy:  0,
				Dname:        "Lessons",
				Duration:     1231.155243,
				Fname:        "2 - Lessons with intro.mp3",
				Fsize:        6515667,
				MinPlay:      0,
				SugPlay:      0,
				Promo:        10,
				Retail:       15,
				PtpFT:        10,
				PtpDT:        20,
				PtpDA:        50,
				Type:         "album track",
				TokenlyID:    "",
				DisallowPlay: 1,
				MinBuy:       500,
				SugBuy:       750,
				Storage:      Oip041Storage{},
			},
			{
				DisallowBuy:  1,
				Dname:        "Born to Roam",
				Duration:     2374.550714,
				Fname:        "3 - Born to Roam.mp3",
				Fsize:        6515667,
				MinPlay:      50,
				SugPlay:      100,
				Promo:        10,
				Retail:       15,
				PtpFT:        10,
				PtpDT:        20,
				PtpDA:        50,
				Type:         "album track",
				TokenlyID:    "",
				DisallowPlay: 0,
				MinBuy:       0,
				SugBuy:       0,
				Storage:      Oip041Storage{},
			},
			{
				DisallowBuy:  1,
				Dname:        "Cover Art",
				Duration:     0,
				Fname:        "birthdayepFINAL.jpg",
				Fsize:        0,
				MinPlay:      0,
				SugPlay:      0,
				Promo:        0,
				Retail:       0,
				PtpFT:        0,
				PtpDT:        0,
				PtpDA:        0,
				Type:         "coverArt",
				TokenlyID:    "",
				DisallowPlay: 0,
				MinBuy:       0,
				SugBuy:       0,
				Storage:      Oip041Storage{},
			},
		},
		Payment: Oip041Payment{
			Fiat:  "USD",
			Scale: "1000:1",
			SugTip: []int{
				5,
				50,
				100,
			},
			Tokens: map[string]string{
				"BTC":          "1GMMg2J5iUKnDf5PbRr9TcKV3R6KfUiB55",
				"MTMCOLLECTOR": "",
				"MTMPRODUCER":  "",
				"HAPPYBDAYEP":  "",
				"EARLY":        "",
				"LTBCOIN":      ""},
		},
	},
	Signature: "H27r7UxUb8BozjEvV0v++nCyRI7S6yyroeKCJQpgU5NO3CP6FpXWs5kCxy8vhmMhbtpj/FMj+8s3+updw7g+bmE=",
}
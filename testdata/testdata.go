/**
 * Copyright 2024 Cloudwego Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package testdata

import "os"

func GetData(path string) []byte {
	js, err := os.ReadFile(path)
	if err != nil {
		panic("read " + path + " failed: " + err.Error())
	}
	return js
}

// this json block is poorly formed on purpose.
var basicJSON = `  {"age":100, "name":{"here":"B\\\"R"},
	"noop":{"what is a wren?":"a bird"},
	"happy":true,"immortal":false,
	"items":[1,2,3,{"tags":[1,2,3],"points":[[1,2],[3,4]]},4,5,6,7],
	"arr":["1",2,"3",{"hello":"world"},"4",5],
	"vals":[1,2,3,{"sadf":   "asdf"}],"name":{"first":"tom","last":null},
	"created":"2014-05-16T08:28:06.989Z",
	"loggy":{
		"programmers": [
    	    {
    	        "firstName": "Brett",
    	        "lastName": "McLaughlin",
    	        "email": "aaaa",
				"tag": "good"
    	    },
    	    {
    	        "firstName": "Jason",
    	        "lastName": "Hunter",
    	        "email": "bbbb",
				"tag": "bad"
    	    },
    	    {
    	        "firstName": "Elliotte",
    	        "lastName": "Harold",
    	        "email": "cccc",
				"tag":  "good"
    	    },
			{
				"firstName": 1002.3,
				"age": 101
			}
    	]
	},
	"lastly":{"end...ing":"soon","yay":"final"}
}`

var complicatedJSON = `
{
	"tagged": "OK",
	"Tagged": "KO",
	"NotTagged": true,
	"unsettable": 101,
	"Nested": {
		"Yellow": "Green",
		"yellow": "yellow"
	},
	"nestedTagged": {
		"Green": "Green",
		"Map": {
			"this": "that",
			"and": "the other thing"
		},
		"Ints": {
			"Uint": 99,
			"Uint16": 16,
			"Uint32": 32,
			"Uint64": 65
		},
		"Uints": {
			"int": -99,
			"Int": -98,
			"Int16": -16,
			"Int32": -32,
			"int64": -64,
			"Int64": -65
		},
		"Uints": {
			"Float32": 32.32,
			"Float64": 64.64
		},
		"Byte": 254,
		"Bool": true
	},
	"LeftOut": "you shouldn't be here",
	"SelfPtr": {"tagged":"OK","nestedTagged":{"Ints":{"Uint32":32}}},
	"SelfSlice": [{"tagged":"OK","nestedTagged":{"Ints":{"Uint32":32}}}],
	"SelfSlicePtr": [{"tagged":"OK","nestedTagged":{"Ints":{"Uint32":32}}}],
	"SelfPtrSlice": [{"tagged":"OK","nestedTagged":{"Ints":{"Uint32":32}}}],
	"interface": "Tile38 Rocks!",
	"Interface": "Please Download",
	"Array": [0,2,3,4,5],
	"time": "2017-05-07T13:24:43-07:00",
	"Binary": "R0lGODlhPQBEAPeo",
	"NonBinary": [9,3,100,115]
}
`

var twitterLarge = string(GetData("./twitterescaped.json"))

var TwitterJsonMedium = `{
  "statuses": [
    {
      "coordinates": null,
      "favorited": false,
      "truncated": false,
      "created_at": "Mon Sep 24 03:35:21 +0000 2012",
      "id_str": "250075927172759552",
      "entities": {
        "urls": [
 
        ],
        "hashtags": [
          {
            "text": "freebandnames",
            "indices": [
              20,
              34
            ]
          }
        ],
        "user_mentions": [
 
        ]
      },
      "in_reply_to_user_id_str": null,
      "contributors": null,
      "text": "Aggressive Ponytail #freebandnames",
      "metadata": {
        "iso_language_code": "en",
        "result_type": "recent"
      },
      "retweet_count": 0,
      "in_reply_to_status_id_str": null,
      "id": 250075927172759552,
      "geo": null,
      "retweeted": false,
      "in_reply_to_user_id": null,
      "place": null,
      "user": {
        "profile_sidebar_fill_color": "DDEEF6",
        "profile_sidebar_border_color": "C0DEED",
        "profile_background_tile": false,
        "name": "Sean Cummings",
        "profile_image_url": "https://a0.twimg.com/profile_images/2359746665/1v6zfgqo8g0d3mk7ii5s_normal.jpeg",
        "created_at": "Mon Apr 26 06:01:55 +0000 2010",
        "location": "LA, CA",
        "follow_request_sent": null,
        "profile_link_color": "0084B4",
        "is_translator": false,
        "id_str": "137238150",
        "entities": {
          "url": {
            "urls": [
              {
                "expanded_url": null,
                "url": "",
                "indices": [
                  0,
                  0
                ]
              }
            ]
          },
          "description": {
            "urls": [
 
            ]
          }
        },
        "default_profile": true,
        "contributors_enabled": false,
        "favourites_count": 0,
        "url": null,
        "profile_image_url_https": "https://si0.twimg.com/profile_images/2359746665/1v6zfgqo8g0d3mk7ii5s_normal.jpeg",
        "utc_offset": -28800,
        "id": 137238150,
        "profile_use_background_image": true,
        "listed_count": 2,
        "profile_text_color": "333333",
        "lang": "en",
        "followers_count": 70,
        "protected": false,
        "notifications": null,
        "profile_background_image_url_https": "https://si0.twimg.com/images/themes/theme1/bg.png",
        "profile_background_color": "C0DEED",
        "verified": false,
        "geo_enabled": true,
        "time_zone": "Pacific Time (US & Canada)",
        "description": "Born 330 Live 310",
        "default_profile_image": false,
        "profile_background_image_url": "https://a0.twimg.com/images/themes/theme1/bg.png",
        "statuses_count": 579,
        "friends_count": 110,
        "following": null,
        "show_all_inline_media": false,
        "screen_name": "sean_cummings"
      },
      "in_reply_to_screen_name": null,
      "source": "<a href=\"//itunes.apple.com/us/app/twitter/id409789998?mt=12%5C%22\" rel=\"\\\"nofollow\\\"\">Twitter for Mac</a>",
      "in_reply_to_status_id": null
    },
    {
      "coordinates": null,
      "favorited": false,
      "truncated": false,
      "created_at": "Fri Sep 21 23:40:54 +0000 2012",
      "id_str": "249292149810667520",
      "entities": {
        "urls": [
 
        ],
        "hashtags": [
          {
            "text": "FreeBandNames",
            "indices": [
              20,
              34
            ]
          }
        ],
        "user_mentions": [
 
        ]
      },
      "in_reply_to_user_id_str": null,
      "contributors": null,
      "text": "Thee Namaste Nerdz. #FreeBandNames",
      "metadata": {
        "iso_language_code": "pl",
        "result_type": "recent"
      },
      "retweet_count": 0,
      "in_reply_to_status_id_str": null,
      "id": 249292149810667520,
      "geo": null,
      "retweeted": false,
      "in_reply_to_user_id": null,
      "place": null,
      "user": {
        "profile_sidebar_fill_color": "DDFFCC",
        "profile_sidebar_border_color": "BDDCAD",
        "profile_background_tile": true,
        "name": "Chaz Martenstein",
        "profile_image_url": "https://a0.twimg.com/profile_images/447958234/Lichtenstein_normal.jpg",
        "created_at": "Tue Apr 07 19:05:07 +0000 2009",
        "location": "Durham, NC",
        "follow_request_sent": null,
        "profile_link_color": "0084B4",
        "is_translator": false,
        "id_str": "29516238",
        "entities": {
          "url": {
            "urls": [
              {
                "expanded_url": null,
                "url": "https://bullcityrecords.com/wnng/",
                "indices": [
                  0,
                  32
                ]
              }
            ]
          },
          "description": {
            "urls": [
 
            ]
          }
        },
        "default_profile": false,
        "contributors_enabled": false,
        "favourites_count": 8,
        "url": "https://bullcityrecords.com/wnng/",
        "profile_image_url_https": "https://si0.twimg.com/profile_images/447958234/Lichtenstein_normal.jpg",
        "utc_offset": -18000,
        "id": 29516238,
        "profile_use_background_image": true,
        "listed_count": 118,
        "profile_text_color": "333333",
        "lang": "en",
        "followers_count": 2052,
        "protected": false,
        "notifications": null,
        "profile_background_image_url_https": "https://si0.twimg.com/profile_background_images/9423277/background_tile.bmp",
        "profile_background_color": "9AE4E8",
        "verified": false,
        "geo_enabled": false,
        "time_zone": "Eastern Time (US & Canada)",
        "description": "You will come to Durham, North Carolina. I will sell you some records then, here in Durham, North Carolina. Fun will happen.",
        "default_profile_image": false,
        "profile_background_image_url": "https://a0.twimg.com/profile_background_images/9423277/background_tile.bmp",
        "statuses_count": 7579,
        "friends_count": 348,
        "following": null,
        "show_all_inline_media": true,
        "screen_name": "bullcityrecords"
      },
      "in_reply_to_screen_name": null,
      "source": "web",
      "in_reply_to_status_id": null
    },
    {
      "coordinates": null,
      "favorited": false,
      "truncated": false,
      "created_at": "Fri Sep 21 23:30:20 +0000 2012",
      "id_str": "249289491129438208",
      "entities": {
        "urls": [
 
        ],
        "hashtags": [
          {
            "text": "freebandnames",
            "indices": [
              29,
              43
            ]
          }
        ],
        "user_mentions": [
 
        ]
      },
      "in_reply_to_user_id_str": null,
      "contributors": null,
      "text": "Mexican Heaven, Mexican Hell #freebandnames",
      "metadata": {
        "iso_language_code": "en",
        "result_type": "recent"
      },
      "retweet_count": 0,
      "in_reply_to_status_id_str": null,
      "id": 249289491129438208,
      "geo": null,
      "retweeted": false,
      "in_reply_to_user_id": null,
      "place": null,
      "user": {
        "profile_sidebar_fill_color": "99CC33",
        "profile_sidebar_border_color": "829D5E",
        "profile_background_tile": false,
        "name": "Thomas John Wakeman",
        "profile_image_url": "https://a0.twimg.com/profile_images/2219333930/Froggystyle_normal.png",
        "created_at": "Tue Sep 01 21:21:35 +0000 2009",
        "location": "Kingston New York",
        "follow_request_sent": null,
        "profile_link_color": "D02B55",
        "is_translator": false,
        "id_str": "70789458",
        "entities": {
          "url": {
            "urls": [
              {
                "expanded_url": null,
                "url": "",
                "indices": [
                  0,
                  0
                ]
              }
            ]
          },
          "description": {
            "urls": [
 
            ]
          }
        },
        "default_profile": false,
        "contributors_enabled": false,
        "favourites_count": 19,
        "url": null,
        "profile_image_url_https": "https://si0.twimg.com/profile_images/2219333930/Froggystyle_normal.png",
        "utc_offset": -18000,
        "id": 70789458,
        "profile_use_background_image": true,
        "listed_count": 1,
        "profile_text_color": "3E4415",
        "lang": "en",
        "followers_count": 63,
        "protected": false,
        "notifications": null,
        "profile_background_image_url_https": "https://si0.twimg.com/images/themes/theme5/bg.gif",
        "profile_background_color": "352726",
        "verified": false,
        "geo_enabled": false,
        "time_zone": "Eastern Time (US & Canada)",
        "description": "Science Fiction Writer, sort of. Likes Superheroes, Mole People, Alt. Timelines.",
        "default_profile_image": false,
        "profile_background_image_url": "https://a0.twimg.com/images/themes/theme5/bg.gif",
        "statuses_count": 1048,
        "friends_count": 63,
        "following": null,
        "show_all_inline_media": false,
        "screen_name": "MonkiesFist"
      },
      "in_reply_to_screen_name": null,
      "source": "web",
      "in_reply_to_status_id": null
    },
    {
      "coordinates": null,
      "favorited": false,
      "truncated": false,
      "created_at": "Fri Sep 21 22:51:18 +0000 2012",
      "id_str": "249279667666817024",
      "entities": {
        "urls": [
 
        ],
        "hashtags": [
          {
            "text": "freebandnames",
            "indices": [
              20,
              34
            ]
          }
        ],
        "user_mentions": [
 
        ]
      },
      "in_reply_to_user_id_str": null,
      "contributors": null,
      "text": "The Foolish Mortals #freebandnames",
      "metadata": {
        "iso_language_code": "en",
        "result_type": "recent"
      },
      "retweet_count": 0,
      "in_reply_to_status_id_str": null,
      "id": 249279667666817024,
      "geo": null,
      "retweeted": false,
      "in_reply_to_user_id": null,
      "place": null,
      "user": {
        "profile_sidebar_fill_color": "BFAC83",
        "profile_sidebar_border_color": "615A44",
        "profile_background_tile": true,
        "name": "Marty Elmer",
        "profile_image_url": "https://a0.twimg.com/profile_images/1629790393/shrinker_2000_trans_normal.png",
        "created_at": "Mon May 04 00:05:00 +0000 2009",
        "location": "Wisconsin, USA",
        "follow_request_sent": null,
        "profile_link_color": "3B2A26",
        "is_translator": false,
        "id_str": "37539828",
        "entities": {
          "url": {
            "urls": [
              {
                "expanded_url": null,
                "url": "https://www.omnitarian.me",
                "indices": [
                  0,
                  24
                ]
              }
            ]
          },
          "description": {
            "urls": [
 
            ]
          }
        },
        "default_profile": false,
        "contributors_enabled": false,
        "favourites_count": 647,
        "url": "https://www.omnitarian.me",
        "profile_image_url_https": "https://si0.twimg.com/profile_images/1629790393/shrinker_2000_trans_normal.png",
        "utc_offset": -21600,
        "id": 37539828,
        "profile_use_background_image": true,
        "listed_count": 52,
        "profile_text_color": "000000",
        "lang": "en",
        "followers_count": 608,
        "protected": false,
        "notifications": null,
        "profile_background_image_url_https": "https://si0.twimg.com/profile_background_images/106455659/rect6056-9.png",
        "profile_background_color": "EEE3C4",
        "verified": false,
        "geo_enabled": false,
        "time_zone": "Central Time (US & Canada)",
        "description": "Cartoonist, Illustrator, and T-Shirt connoisseur",
        "default_profile_image": false,
        "profile_background_image_url": "https://a0.twimg.com/profile_background_images/106455659/rect6056-9.png",
        "statuses_count": 3575,
        "friends_count": 249,
        "following": null,
        "show_all_inline_media": true,
        "screen_name": "Omnitarian"
      },
      "in_reply_to_screen_name": null,
      "source": "<a href=\"//twitter.com/download/iphone%5C%22\" rel=\"\\\"nofollow\\\"\">Twitter for iPhone</a>",
      "in_reply_to_status_id": null
    }
  ],
  "search_metadata": {
    "max_id": 250126199840518145,
    "since_id": 24012619984051000,
    "refresh_url": "?since_id=250126199840518145&q=%23freebandnames&result_type=mixed&include_entities=1",
    "next_results": "?max_id=249279667666817023&q=%23freebandnames&count=4&include_entities=1&result_type=mixed",
    "count": 4,
    "completed_in": 0.035,
    "since_id_str": "24012619984051000",
    "query": "%23freebandnames",
    "max_id_str": "250126199840518145"
  }
}`

package entity

import "time"

type AccuWeatherGeoResponse struct {
	Version           int    `json:"Version"`
	Key               string `json:"Key"`
	Type              string `json:"Type"`
	Rank              int    `json:"Rank"`
	LocalizedName     string `json:"LocalizedName"`
	EnglishName       string `json:"EnglishName"`
	PrimaryPostalCode string `json:"PrimaryPostalCode"`
	Region            struct {
		ID            string `json:"ID"`
		LocalizedName string `json:"LocalizedName"`
		EnglishName   string `json:"EnglishName"`
	} `json:"Region"`
	Country struct {
		ID            string `json:"ID"`
		LocalizedName string `json:"LocalizedName"`
		EnglishName   string `json:"EnglishName"`
	} `json:"Country"`
	AdministrativeArea struct {
		ID            string `json:"ID"`
		LocalizedName string `json:"LocalizedName"`
		EnglishName   string `json:"EnglishName"`
		Level         int    `json:"Level"`
		LocalizedType string `json:"LocalizedType"`
		EnglishType   string `json:"EnglishType"`
		CountryID     string `json:"CountryID"`
	} `json:"AdministrativeArea"`
}

type Measure struct {
	Value    float32 `json:"Value"`
	Unit     string  `json:"Unit"`
	UnitType int     `json:"UnitType"`
}

type AccuWeatherResponse struct {
	LocalObservationDateTime time.Time   `json:"LocalObservationDateTime"`
	EpochTime                int         `json:"EpochTime"`
	WeatherText              string      `json:"WeatherText"`
	WeatherIcon              int         `json:"WeatherIcon"`
	HasPrecipitation         bool        `json:"HasPrecipitation"`
	PrecipitationType        interface{} `json:"PrecipitationType"`
	IsDayTime                bool        `json:"IsDayTime"`
	Temperature              struct {
		Metric   Measure `json:"Metric"`
		Imperial Measure `json:"Imperial"`
	} `json:"Temperature"`
	RealFeelTemperature struct {
		Metric   Measure `json:"Metric"`
		Imperial Measure `json:"Imperial"`
	} `json:"RealFeelTemperature"`
	RealFeelTemperatureShade struct {
		Metric   Measure `json:"Metric"`
		Imperial Measure `json:"Imperial"`
	} `json:"RealFeelTemperatureShade"`
	RelativeHumidity int `json:"RelativeHumidity"`
	DewPoint         struct {
		Metric   Measure `json:"Metric"`
		Imperial Measure `json:"Imperial"`
	} `json:"DewPoint"`
	Wind struct {
		Direction struct {
			Degrees   int    `json:"Degrees"`
			Localized string `json:"Localized"`
			English   string `json:"English"`
		} `json:"Direction"`
		Speed struct {
			Metric   Measure `json:"Metric"`
			Imperial Measure `json:"Imperial"`
		} `json:"Speed"`
	} `json:"Wind"`
	WindGust struct {
		Speed struct {
			Metric   Measure `json:"Metric"`
			Imperial Measure `json:"Imperial"`
		} `json:"Speed"`
	} `json:"WindGust"`
	UVIndex     int    `json:"UVIndex"`
	UVIndexText string `json:"UVIndexText"`
	Visibility  struct {
		Metric   Measure `json:"Metric"`
		Imperial Measure `json:"Imperial"`
	} `json:"Visibility"`
	ObstructionsToVisibility string `json:"ObstructionsToVisibility"`
	CloudCover               int    `json:"CloudCover"`
	Ceiling                  struct {
		Metric   Measure `json:"Metric"`
		Imperial Measure `json:"Imperial"`
	} `json:"Ceiling"`
	Pressure struct {
		Metric   Measure `json:"Metric"`
		Imperial Measure `json:"Imperial"`
	} `json:"Pressure"`
	PressureTendency struct {
		LocalizedText string `json:"LocalizedText"`
		Code          string `json:"Code"`
	} `json:"PressureTendency"`
	Past24HourTemperatureDeparture struct {
		Metric   Measure `json:"Metric"`
		Imperial Measure `json:"Imperial"`
	} `json:"Past24HourTemperatureDeparture"`
	ApparentTemperature struct {
		Metric   Measure `json:"Metric"`
		Imperial Measure `json:"Imperial"`
	} `json:"ApparentTemperature"`
	WindChillTemperature struct {
		Metric   Measure `json:"Metric"`
		Imperial Measure `json:"Imperial"`
	} `json:"WindChillTemperature"`
	WetBulbTemperature struct {
		Metric   Measure `json:"Metric"`
		Imperial Measure `json:"Imperial"`
	} `json:"WetBulbTemperature"`
	Precip1Hr struct {
		Metric   Measure `json:"Metric"`
		Imperial Measure `json:"Imperial"`
	} `json:"Precip1hr"`
	PrecipitationSummary struct {
		Precipitation struct {
			Metric   Measure `json:"Metric"`
			Imperial Measure `json:"Imperial"`
		} `json:"Precipitation"`
		PastHour struct {
			Metric   Measure `json:"Metric"`
			Imperial Measure `json:"Imperial"`
		} `json:"PastHour"`
		Past3Hours struct {
			Metric   Measure `json:"Metric"`
			Imperial Measure `json:"Imperial"`
		} `json:"Past3Hours"`
		Past6Hours struct {
			Metric   Measure `json:"Metric"`
			Imperial Measure `json:"Imperial"`
		} `json:"Past6Hours"`
		Past9Hours struct {
			Metric   Measure `json:"Metric"`
			Imperial Measure `json:"Imperial"`
		} `json:"Past9Hours"`
		Past12Hours struct {
			Metric   Measure `json:"Metric"`
			Imperial Measure `json:"Imperial"`
		} `json:"Past12Hours"`
		Past18Hours struct {
			Metric   Measure `json:"Metric"`
			Imperial Measure `json:"Imperial"`
		} `json:"Past18Hours"`
		Past24Hours struct {
			Metric   Measure `json:"Metric"`
			Imperial Measure `json:"Imperial"`
		} `json:"Past24Hours"`
	} `json:"PrecipitationSummary"`
	TemperatureSummary struct {
		Past6HourRange struct {
			Minimum struct {
				Metric   Measure `json:"Metric"`
				Imperial Measure `json:"Imperial"`
			} `json:"Minimum"`
			Maximum struct {
				Metric   Measure `json:"Metric"`
				Imperial Measure `json:"Imperial"`
			} `json:"Maximum"`
		} `json:"Past6HourRange"`
		Past12HourRange struct {
			Minimum struct {
				Metric   Measure `json:"Metric"`
				Imperial Measure `json:"Imperial"`
			} `json:"Minimum"`
			Maximum struct {
				Metric   Measure `json:"Metric"`
				Imperial Measure `json:"Imperial"`
			} `json:"Maximum"`
		} `json:"Past12HourRange"`
		Past24HourRange struct {
			Minimum struct {
				Metric   Measure `json:"Metric"`
				Imperial Measure `json:"Imperial"`
			} `json:"Minimum"`
			Maximum struct {
				Metric   Measure `json:"Metric"`
				Imperial Measure `json:"Imperial"`
			} `json:"Maximum"`
		} `json:"Past24HourRange"`
	} `json:"TemperatureSummary"`
	MobileLink string `json:"MobileLink"`
	Link       string `json:"Link"`
}

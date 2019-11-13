package main

import "time"

type Observation struct {
	Id              int       `json:"id"`
	Latitude        float64   `json:"latitude"`
	Longitude       float64   `json:"longitude"`
	MeasurementUnit string    `json:"measurement_unit"`
	Timestamp       time.Time `json:"timestamp"`
	Value           float64   `json:"value"`
	SensorId        int       `json:"sensor_id"`
	Measuring       string    `json:"measuring"`
}
type Data struct {
	UserId        int     `json:"userId"`
	DeviceId      int     `json:"deviceId"`
	Id            int     `json:"id"`
	Timestamp     string  `json:"date"`
	Latitude      string  `json:"latitude"`
	Longitude     string  `json:"longitude"`
	Lpg           float64 `json:"lpg"`
	Co            float64 `json:"co"`
	Smoke         float64 `json:"smoke"`
	Co2           float64 `json:"co2"`
	BackTemp      float64 `json:"backTemp"`
	Humidity      float64 `json:"humidity"`
	Dust          float64 `json:"dust"`
	Pressure      float64 `json:"pressure"`
	FrontTemp     float64 `json:"frontTemp"`
	Vis           float64 `json:"vis"`
	Ir            float64 `json:"ir"`
	Uv            float64 `json:"uv"`
	FrontTempDht  float64 `json:"frontTempDht"`
	FrontHumidity float64 `json:"frontHumidity"`
}
type User struct {
	Id       int       `json:"id"`
	DeviceId int       `json:"deviceID"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	Role     string    `json:"role"`
	Email    string    `json:"email"`
	Sex      string    `json:"sex"`
	Birthday time.Time `json:"birthday"`
}

type Answer struct {
	Id         int       `json:"id"`
	Latitude   float64   `json:"latitude"`
	Longitude  float64   `json:"longitude"`
	Timestamp  time.Time `json:"date"`
	Value      int       `json:"answer"`
	QuestionId int       `json:"questionId"`
	UserId     int       `json:"userId"`
}
type Question struct {
	Id   int    `json:"id"`
	Text string `json:"text"`
	Type string `json:"type"`
}
type Alert struct {
	Id          int       `json:"id"`
	Address     string    `json:"address"`
	Description string    `json:"description"`
	Type        string    `json:"type"`
	Timestamp   time.Time `json:"timestamp"`
	Latitude    float64   `json:"latitude"`
	Longitude   float64   `json:"longitude"`
}

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

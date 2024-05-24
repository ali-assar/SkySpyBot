package handler

var weatherDescriptions = map[int]string{
	800: "Clear sky :sunny:",
	801: "Few clouds :partly_sunny:",
	802: "Scattered clouds :cloud:",
	803: "Broken clouds :cloud:",
	804: "Overcast clouds :cloud:",
	500: "Light rain :cloud_with_rain:",
	501: "Moderate rain :cloud_with_rain:",
	502: "Heavy rain :cloud_with_rain:",
	503: "Very heavy rain :cloud_with_rain:",
	504: "Extreme rain :cloud_with_rain:",
	600: "Light snow :cloud_with_snow:",
	601: "Snow :cloud_with_snow:",
	602: "Heavy snow :cloud_with_snow:",
	611: "Sleet :cloud_with_snow:",
	615: "Light rain and snow :cloud_with_snow:",
	616: "Rain and snow :cloud_with_snow:",
	// Add more conditions as needed
}


How it Works:
1)   The app listens for HTTP requests at http://localhost:8080/weather.
2)   When you provide a city name using the city query parameter (e.g., http://localhost:8080/weather?city=London), it sends a request to the OpenWeatherMap API to fetch weather details for that city.
3)   The weather data, including the temperature, humidity, and description, is then returned as a JSON response.



How to Use:
1)   To run this application, you need to get your own API key from OpenWeatherMap. You can sign up for a free account and get your API key
2)   Change the api key 

How to run:
1)   first copy this code 
2)   now type command in terminal ( nano (yourfilename).go
3)   now paste this code in it and press ctrl + o , then press enter , then press ctrl + x
4)   open terminal and write command ( go run (yourfilename).go)
5)   you will see (Weather app is running on http://localhost:8080)
6)   now paste this url in chrome (http://localhost:8080/weather?city=(yourcity name))

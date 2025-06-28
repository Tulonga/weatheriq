<!-- WeatherIQ README.md -->

# ⛅️ WeatherIQ


> **Your weather, reimagined. Fast. Simple. Beautiful.**

---

## 🌍 What is WeatherIQ?

**WeatherIQ** is a modern, blazing-fast weather API built with Go.  
It fetches real-time weather data for any city on Earth and delivers it to you in a clean, developer-friendly JSON format.

- **Live weather** for any city
- **Lightning-fast** thanks to Go’s performance
- **Secure**: API keys are never exposed
- **Ready for your frontend**: Just plug and play!

---

## 🚀 Features

- **🌦 Real-Time Weather:** Get current temperature, conditions, and more for any city.
- **🔒 Secure API Key Handling:** Your secrets stay secret.
- **⚡️ Ultra-Fast API:** Go-powered backend ensures rapid responses.
- **🖥 Simple REST Endpoint:** Easy to integrate with any frontend or tool.
- **📱 Mobile & Web Ready:** Build beautiful UIs on top of WeatherIQ.
- **🧩 Modular Design:** Easy to extend and customize.

---

## 🛠 How It Works

1. **Frontend** sends a request like  
   `GET /weather?city=Windhoek`
2. **WeatherIQ** backend fetches live data from WeatherAPI.com
3. **WeatherIQ** returns a clean JSON response.
4. **You** display it however you want!

---

## 🏁 Quick Start

1. **Clone this repo**  

https://www.github.com/tulonga/weatheriq

cd weatheriq


2. **Set your WeatherAPI key**  

export WEATHERAPI_KEY= --insert api key here--


3. **Run the server**  

go run main.go


4. **Test it!**  
Open your browser or Postman:  
[http://localhost:8080/weather?city=Manchester](http://localhost:8080/weather?city=Manchester)

---

## ✨ ASCII Weather Vibes

  \   /      _           _   _      _       
   .-.     / \   _ __ __| | | | ___| |_ ___ 
― (   ) ― / _ \ | '__/ _` | | |/ _ \ __/ _ \
   `-’   / ___ \| | | (_| | | |  __/ ||  __/
  /   \ /_/   \_\_|  \__,_| |_|\___|\__\___|

## 📬 Contact

Made with vibes by [Tulonga](https://github.com/tulonga)

---

# 🌤️ WeatherStation CLI
A simple Go command-line application to manage weather station telemetry. This tool efficiently processes incoming sensor data updates, maintains state, and provides full status reporting with minimal dependencies.

## 📘 Overview
This application simulates how a weather station data processor might function in a critical infrastructure system. It accepts incoming data from a weather station via standard input and tracks the latest known state of nine key meteorological sensors.

To optimize payload sizes, the station only sends updates when values change, and a full snapshot is sent every 10 minutes. This CLI tool is designed to interpret those partial updates and reconstruct the most complete state possible.

## ✅ Features
Handles incoming telemetry as ID,Value data (partial updates)

Maintains the latest known values for 9 meteorological sensors

### Responds to:

    get — to print the full sensor state

    clear — to reset all values

    exit — to terminate the session

    Ignores deprecated or invalid sensor IDs

    Starts with all values as NULL (unknown)

### 🗃️ Sensor ID Mapping
    The following table defines the expected sensor IDs and their internal keys:

ID	Sensor Key
1	airTemp
2	airPressure
7	precipitation
11	windSpeed
12	windDirection
13	humidity
14	dewPoint
15	soilMoisture
22	cloudCover

Any IDs not listed above will be ignored.

### 🔄 Example Commands
bash
Copy
Edit
# Update values
11,15.5
13,32.3

# Set a field to NULL (missing data)
12,NULL

# Query full state
get

# Clear/reset all values
clear

# Exit the program
exit
🖨️ Example Output
```bash
--- Weather Station ---
airTemp:NULL
airPressure:NULL
precipitation:NULL
windSpeed:15.5
windDirection:NULL
humidity:32.3
dewPoint:NULL
soilMoisture:NULL
cloudCover:NULL
```
## 🏗️ Project Structure
arduino
```bash
weatherStation/
├── main.go          # Entry point, handles command loop and stdin
The Go module name should match the repository name: weatherStation

You are free to expand the implementation into multiple files or packages, but main.go must remain in the root.

📌 Implementation Details
The program uses raw Go code without external libraries to:

Store current weather sensor states in memory

Update values based on partial or full snapshots

Output formatted state to standard output

Process and validate user input via bufio.Scanner
```

    All values are printed in ascending ID order, using NULL for any sensor not yet updated.

## 🚀 How to Run
```bash

go run main.go
``` 
Then start entering sensor updates, get, clear, or exit commands line by line.

# 🧠 Notes
Each line of input triggers exactly one action — no menus or extra prompts.

No need to handle EOF or error conditions beyond standard input validation.

Ideal for testing low-level logic for stream-based sensor integration.
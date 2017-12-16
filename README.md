# StatPi
StatPi is a simple project running on the RaspberryPi with the RaspberryPi SenseHat. It is part of a student research paper and at the moment only a prove-of-concept. It consists of 3 main Parts:

### StatPiCollector (golang)
Temperature and humidity collected from the SenseHat. Additionally values are collected from an online weatherAPI(e.g. openweathermap).
The values are safed into .csv-files.

### StatPiConverter (R)
To plot .svg-files from the given data, R-scripts are used. For each case, another script is written, that they can be executed by a cron-deamon.

### StatPiProvider (golang)
The StatPiProvider is webservice written in go. At this point in time, it simply returns .html-file, which display the .svg-files

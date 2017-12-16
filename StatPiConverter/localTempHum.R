library(readr)
mydata <- read_csv("local.csv")
names(mydata)
head(mydata)
svg(filename="localTempHum.svg", 
    width=10, 
    height=8, 
    pointsize=12)
plot(humidity ~ temperature ,type="p", main='Darstellung der lokalen Temperatur und Luftfeuchte',
     data = mydata, col= "blue", xlab="Temperatur in °C", ylab="relative Luftfeuchtigkeit in %")

dev.off()


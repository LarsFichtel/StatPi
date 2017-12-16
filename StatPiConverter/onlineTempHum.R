library(readr)
mydata <- read_csv("online.csv")
names(mydata)
head(mydata)
svg(filename="onlineTempHum.svg", 
    width=10, 
    height=8, 
    pointsize=12)
plot(humidity ~ temperature ,type="p", main='Darstellung der online Temperatur und Luftfeuchte',
     data = mydata, col= "blue", xlab="Temperatur in °C", ylab="relative Luftfeuchtigkeit in %")

dev.off()


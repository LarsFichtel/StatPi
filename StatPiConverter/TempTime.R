library(readr)
library(anytime)
library(dplyr)
local <- read_csv("local.csv")
online <- read_csv("online.csv")

svg(filename="temptime.svg", 
    width=10, 
    height=8, 
    pointsize=12)
print(as.numeric(Sys.time())-(3600*24), digits=15)



locallastday <- filter(local, local$time > as.numeric(Sys.time())-(3600*24))
onlinelastday <- filter(online, online$time > as.numeric(Sys.time())-(3600*24))

plot(local$temperature ~ format(anytime(locallastday$time), "%H") ,type="p", main='Darstellung der Temperaturen der letzten 24 Stunden',
     data = local, col= "blue", ylab="Temperatur in °C", xlab="Uhrzeit")
points(online$temperature ~ format(anytime(online$time), "%H") , col ="red", pch=20)
legend(20.275, 31.5, legend=c("online", "local"),
       col=c("red", "blue"), lty=1:2, cex=0.9,
       box.lty=1)

dev.off()


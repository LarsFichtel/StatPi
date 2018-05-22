library(readr)
library(anytime)
library(dplyr)
local <- read_csv("C:/Users/Lars Fichtel/Documents/StatPi/local.csv")

attach(local)

Y <- cbind(local$temperature)
X <- cbind(local$humidity)
png(filename="C:/Users/Lars Fichtel/Documents/StatPi/name.png")


plot(Y ~ X, data = local)

abline(lm(Y ~ X))

dev.off()


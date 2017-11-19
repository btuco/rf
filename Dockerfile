FROM alexellis2/go-armhf:1.8.4
RUN apk add --no-cache git && go get -d -u gobot.io/x/gobot/...
RUN go get github.com/gorilla/mux
#FROM hypriot/rpi-gpio
ENTRYPOINT /bin/sh

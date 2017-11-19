sudo docker run --rm --network host -it --privileged --cap-add SYS_RAWIO --device /dev/mem --device /dev/gpiomem -v src:/go/src/github.com -v /home/pirate/rf/:/go/src/github.com/btuco/rest-api rf


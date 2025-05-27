FROM ubuntu:latest
LABEL authors="mihni"

ENTRYPOINT ["top", "-b"]
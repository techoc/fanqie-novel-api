FROM ubuntu:latest
LABEL authors="techoc"


ENTRYPOINT ["top", "-b"]
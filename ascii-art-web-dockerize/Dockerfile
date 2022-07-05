# Dockerfile References: https://docs.docker.com/engine/reference/builder/
FROM ubuntu:20.04
LABEL maintainer="M.a_k & Araya"
ADD cmd/main /
ADD ./ui/html/*.html /ui/html/
ADD ./ui/static/*.css /ui/static/
ADD ./ui/static/*.png /ui/static/
ADD shadow.txt /
ADD standard.txt /
ADD thinkertoy.txt /
ADD README.md /
# ENV API_BOLTDB_PATH='bolt.db'
# ENV API_HOSTNAME=':8080'

CMD ["/main"]

EXPOSE 8080

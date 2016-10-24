FROM busybox

MAINTAINER inge4pres

RUN mkdir LEGODIR

COPY certificates/4pr.es.key /var/lego/
COPY certificates/4pr.es.crt /var/lego/
COPY shorty /shorty
COPY templates /templates

EXPOSE 1337 

ENTRYPOINT [ "/shorty", "-cert", "/var/lego/4pr.es.crt", "-key", "/var/lego/4pr.es.key" ]



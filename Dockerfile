FROM alpine:latest
MAINTAINER contact@leblanc-simon.eu

RUN apk update && \
    apk --no-cache add util-linux && \
    apk --no-cache add libreoffice-common && \
    apk --no-cache add libreoffice-writer && \
    apk --no-cache add ttf-droid-nonlatin ttf-droid ttf-dejavu ttf-freefont ttf-liberation && \
    apk --no-cache add msttcorefonts-installer fontconfig && \
    update-ms-fonts && \
    fc-cache -f && \
    rm -fr /var/cache/apk/* && \
    ln -s /usr/bin/python3 /usr/bin/python && \
    mkdir -p ~/.config/libreoffice && chmod -R 777 ~/.config/libreoffice

COPY files/. /

VOLUME ["/tmp"]

RUN chmod 755 /usr/local/sbin/unoconv /usr/local/sbin/convert.sh

ENTRYPOINT ["/usr/local/sbin/convert.sh"]


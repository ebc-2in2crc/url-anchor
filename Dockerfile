FROM alpine:3.11

LABEL com.github.ebc-2in2crc.url-anchor.image=https://github.com/ebc-2in2crc/url-anchor.git

RUN apk update && \
	apk --no-cache add curl && \
	curl --location --remote-name https://github.com/ebc-2in2crc/url-anchor/releases/download/v1.1.0/url-anchor_linux_amd64.zip && \
	apk del curl && \
	unzip url-anchor_linux_amd64.zip url-anchor_linux_amd64/url-anchor && \
	cp ./url-anchor_linux_amd64/url-anchor /usr/local/bin && \
	rm -rf url-anchor_linux_amd64.zip url-anchor_linux_amd64/url-anchor

COPY docker-entrypoint.sh /usr/local/bin

ENTRYPOINT ["docker-entrypoint.sh"]

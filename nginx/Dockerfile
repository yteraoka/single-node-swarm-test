FROM nginx:1.15.5-alpine
COPY default.conf /etc/nginx/conf.d/default.conf
RUN apk --no-cache add curl
HEALTHCHECK CMD curl -f http://127.0.0.1/healthcheck || exit 1

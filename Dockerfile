FROM alpine:edge

COPY dist /app
WORKDIR /app

RUN apk --no-cache add nodejs npm

RUN cd /app \
    && npm install pm2 -g \
    && npm install --production \
    && npm cache clean --force \
    && apk del npm

CMD ["pm2-runtime", "app.js", "-i", "max"]

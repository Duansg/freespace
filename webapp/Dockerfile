FROM node:18-alpine

ARG TIMEZONE
ENV TIMEZONE=${TIMEZONE:-"Asia/Shanghai"}

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /app

COPY package*.json ./

RUN npm install

COPY . .

RUN npm run build

RUN adduser -D -s /bin/sh astro && \
    chown -R astro:astro /app

USER astro

EXPOSE 8241

# Start the SSR server
CMD ["npm", "start"]
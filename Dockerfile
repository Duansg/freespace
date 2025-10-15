FROM node:20-bullseye

WORKDIR /app

COPY frontend ./frontend
COPY backend ./backend
COPY run.sh .

RUN chmod +x run.sh \
    && cd backend && go build -o /app/backend/server main.go \
    && cd /app/frontend && npm ci && npm run build

EXPOSE 8240 4321

CMD ["./run.sh"]

#!/bin/bash
/app/backend/server &

cd /app/frontend && node ./dist/server/entry.mjs &

wait -n

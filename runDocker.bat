@ECHO off
docker run -v SLL_LOCATON:/app/ssl -v LOG_LOCATION:/app/logs --rm --name radio-t-bot -p 505:80 -d radio-t-bot
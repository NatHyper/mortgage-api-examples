@echo off
echo Registering new API key...
curl -X POST https://cleargage.co.uk/api/register -H "Content-Type: application/json" -d "{\"name\":\"<YOUR-KEY-NAME>\",\"email\":\"<YOUR-EMAIL-ADDRESS>\"}"
echo.
pause
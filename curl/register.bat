@echo off
echo Registering new API key...
curl -X POST https://cleargage.co.uk/api/register -H "Content-Type: application/json" -d "{\"name\":\"Test User\",\"email\":\"test@example.com\"}"
echo.
pause
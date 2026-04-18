@echo off
set /p api_key="Enter your API key: "
curl -X POST http://localhost:8080/api/summary -H "X-API-Key: %api_key%" -H "Content-Type: application/json" -d "{\"principal\":200000,\"rate\":4.5,\"months\":300,\"start_date\":\"2026-03-01\"}"
echo.
pause
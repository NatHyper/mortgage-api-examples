import requests
import json

# Step 1: Register and get an API key
register_url = "http://localhost:8080/api/register"
register_data = {
    "name": "Test User",
    "email": "test@example.com"
}

response = requests.post(register_url, json=register_data)
api_key = response.json()["api_key"]
print(f"Your API key: {api_key}")

# Step 2: Make a calculation
summary_url = "http://localhost:8080/api/summary"
headers = {
    "X-API-Key": api_key,
    "Content-Type": "application/json"
}
summary_data = {
    "principal": 200000,
    "rate": 4.5,
    "months": 300,
    "start_date": "2026-03-01"
}

response = requests.post(summary_url, headers=headers, json=summary_data)
print(json.dumps(response.json(), indent=2))
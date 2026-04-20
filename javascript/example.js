const https = require('https');

const API_BASE = 'https://cleargage.co.uk/api';

// Helper function for HTTP requests
function request(url, options, data = null) {
    return new Promise((resolve, reject) => {
        const req = require('https').request(url, options, (res) => {
            let body = '';
            res.on('data', (chunk) => body += chunk);
            res.on('end', () => {
                try {
                    resolve(JSON.parse(body));
                } catch {
                    resolve(body);
                }
            });
        });
        req.on('error', reject);
        if (data) req.write(JSON.stringify(data));
        req.end();
    });
}

async function main() {
    // Step 1: Register to get an API key
    const registerOptions = {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' }
    };
    const registerData = {
        name: 'Test User',
        email: 'test@example.com'
    };
    
    const registerResult = await request(`${API_BASE}/register`, registerOptions, registerData);
    const apiKey = registerResult.api_key;
    console.log(`Your API key: ${apiKey}`);
    
    // Step 2: Make a calculation using the API key
    const summaryOptions = {
        method: 'POST',
        headers: {
            'X-API-Key': apiKey,
            'Content-Type': 'application/json'
        }
    };
    const summaryData = {
        principal: 200000,
        rate: 4.5,
        months: 300,
        start_date: '2026-03-01'
    };
    
    const summaryResult = await request(`${API_BASE}/summary`, summaryOptions, summaryData);
    console.log(JSON.stringify(summaryResult, null, 2));
}

main().catch(console.error);
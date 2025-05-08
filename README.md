# Go Challenge - Flight Price

## Requirements

### RESTful Endpoints & Authentication

- GET /flights/search?origin=XXX&destination=YYY&date=YYYY-MM-DD → Returns flight price
comparisons.
- Implement JWT authentication to restrict API access.

### Parallel API Fetching & Data Processing
- Make concurrent requests to at least three APIs.
- Aggregate and compare flight data from multiple providers.
- Identify and sort the cheapest and fastest flights.

### Security
- Keep credentials and tokens secure
- Ensure only authenticated users can access the endpoints.
- Document how HTTPS/TLS should be configured for production.

### Testing
- Unit and integration tests for core aspect
- Validate the accuracy of price comparisons.

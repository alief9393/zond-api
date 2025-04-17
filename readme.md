Zond API
Overview
Zond API is a blockchain explorer backend for the Zond network, a quantum-safe blockchain using a proof-of-stake (PoS) consensus mechanism with a beacon chain via Qrysm. This API provides endpoints to query blockchain data such as blocks, transactions, addresses, forks, reorgs, chain metadata, and validators. It is built with Go, using the Gin framework for routing, PostgreSQL for data storage, and JWT for authentication.
Features

Retrieve latest blocks and specific blocks by number
Fetch latest transactions and specific transactions by hash
Query address balances and transactions
Access fork events and chain reorganizations (reorgs)
Get chain metadata (e.g., chain ID, latest block)
Monitor validator data from the beacon chain
Secure endpoints with JWT authentication
Role-based access (e.g., admin-only endpoints)

Prerequisites

Go 1.20 or higher
PostgreSQL 15 or higher
Qrysm beacon chain client (for validator data)

Installation

Clone the repository:
git clone https://github.com/aliefchandrawijaya/zond-api.git
cd zond-api


Install dependencies:
go mod tidy


Build the project (optional):
go build -o zond-api ./cmd



Configuration
The API uses environment variables for configuration. Create a .env file in the project root with the following variables:
POSTGRES_CONN=postgresql://user:password@localhost:5432/zond_indexer_db
JWT_SECRET=1234567890
PORT=8080


POSTGRES_CONN: Connection string for PostgreSQL.
JWT_SECRET: Secret key for JWT authentication.
PORT: Port where the API will run (default: 8080).

Note: If running from the cmd directory, ensure the .env file is in the project root, or adjust the path in cmd/main.go to load it correctly.
Running the API

Start the PostgreSQL database and ensure the zond_indexer_db exists (see Database Setup below).
Start the Qrysm beacon chain client (for validator data) at http://localhost:3500.
Run the API:go run cmd/main.go

Or, if you built the binary:./zond-api



The API will be available at http://localhost:8080.
API Endpoints
All endpoints except /api/login require JWT authentication. Use the /api/login endpoint to obtain a token.



Method
Endpoint
Description
Role Required



POST
/api/login
Authenticate and get a JWT token
None


GET
/api/blocks/latest
Get the latest blocks
Any


GET
/api/blocks/:block_number
Get a block by number
Any


GET
/api/transactions/latest
Get the latest transactions
Any


GET
/api/transactions/:tx_hash
Get a transaction by hash
Any


GET
/api/addresses/:address/balance
Get an address’s balance
Admin


GET
/api/addresses/:address/transactions
Get transactions for an address
Any


GET
/api/forks
Get fork events
Any


GET
/api/chain
Get chain metadata
Any


GET
/api/reorgs
Get chain reorg events
Any


GET
/api/validators
Get validator data
Any


Example Requests

Login:
curl -X POST http://localhost:8080/api/login -H "Content-Type: application/json" -d '{"username":"admin","password":"password123"}'

Response:
{"token":"your_jwt_token"}


Get Latest Blocks:
curl http://localhost:8080/api/blocks/latest -H "Authorization: Bearer your_jwt_token"



Authentication
The API uses JWT for authentication. To access protected endpoints:

Use the /api/login endpoint to get a token (default credentials: username: admin, password: password123).
Include the token in the Authorization header as Bearer your_jwt_token.

Admin Role: The /api/addresses/:address/balance endpoint requires the user to have an admin role, which is set in the JWT claims during login.
Database Setup
The API stores data in a PostgreSQL database named zond_indexer_db. Set up the database as follows:

Indexer Integration
The API relies on an indexer to sync data from the Zond network into the database. The indexer should:

Sync block, transaction, and address data from the Zond execution layer.
Fetch fork and reorg events from the beacon chain.
Retrieve validator data from the Qrysm beacon chain client (/eth/v1/beacon/states/head/validators endpoint).

Example: To sync validator data, the indexer can periodically call the Qrysm API and insert/update the validators table. Ensure the Qrysm client is running at http://localhost:3500 (default port)
Contributing
Contributions are welcome! To contribute:

Fork the repository.
Create a new branch (git checkout -b feature/your-feature).
Make your changes and commit (git commit -m "Add your feature").
Push to your branch (git push origin feature/your-feature).
Open a Pull Request.

License
This project is licensed under the MIT License. See the LICENSE file for details.

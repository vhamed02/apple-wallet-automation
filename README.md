# Apple Wallet Automation

Records Apple Pay / iPhone Wallet transactions sent via iOS Shortcuts and automatically categorizes them by merchant name.

## How It Works

An iOS Shortcut fires a POST request after each payment. This server receives the request, validates the API key, categorizes the transaction by merchant name, and stores it in a per-user JSON file.

## Setup

### 1. Create your credentials file

Copy the example and fill in your real values:

```bash
cp credentials.example.yml credentials.yml
```

Edit `credentials.yml`:

```yaml
users:
  - username: your_username
    api_key: your_secret_api_key
```

`credentials.yml` is gitignored and never committed. `credentials.example.yml` is the committed template.

### 2. Run

```bash
go run main.go
```

Override file paths via env vars if needed:

```bash
CONFIG_PATH=config.yml CREDENTIALS_PATH=credentials.yml go run main.go
```

Or build and run:

```bash
go build -o wallet-automation .
./wallet-automation
```

### Environment Variables

| Variable            | Default              | Description                      |
|---------------------|----------------------|----------------------------------|
| `PORT`              | `3000`               | HTTP port to listen on           |
| `CONFIG_PATH`       | `config.yml`         | Path to main config file         |
| `CREDENTIALS_PATH`  | `credentials.yml`    | Path to credentials file         |

## API

### POST `/record/`

**Headers:**

| Header        | Required | Description        |
|---------------|----------|--------------------|
| `X-Api-Key`   | Yes      | User's API key     |
| `Content-Type`| Yes      | `application/json` |

**Body:**

```json
{
  "amount": "֏26 307,00",
  "transaction": "Yerevan City Komitas",
  "card": "Visa Classic",
  "name": "Yerevan City Komitas",
  "merchant": "Yerevan City Komitas"
}
```

**Response `201`:**

```json
{
  "id": "uuid",
  "category": "Groceries",
  "recorded_at": "2026-08-16T10:00:00Z"
}
```

**Response `401`:** Invalid or missing API key — transaction is not recorded.

## Data Storage

Transactions are stored in `./data/<username>.json`:

```json
{
  "transactions": [
    {
      "id": "uuid",
      "amount": "֏26 307,00",
      "transaction": "Yerevan City Komitas",
      "card": "Visa Classic",
      "name": "Yerevan City Komitas",
      "merchant": "Yerevan City Komitas",
      "category": "Groceries",
      "recorded_at": "2026-08-16T10:00:00Z"
    }
  ]
}
```

## Categories

Matching is case-insensitive substring matching on the merchant name. If no keyword matches, the transaction is placed in `Other`.

Built-in categories: `Groceries`, `Restaurant`, `Transport`, `Shopping`, `Health`, `Entertainment`, `Utilities`, `Travel`, `Other`.

All keywords are configurable in `config.yml`.

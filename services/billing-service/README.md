# Billing Service

Billing and invoicing management.

## Features

- Invoice generation
- Payment tracking
- Subscription billing
- Usage-based billing
- Tax calculation
- Dunning management
- Payment reminders
- Financial reporting

## Endpoints

- `GET /invoices` - List invoices
- `POST /invoices` - Create invoice
- `GET /invoices/:id` - Get invoice
- `POST /invoices/:id/pay` - Pay invoice
- `GET /billing/usage` - Get usage
- `GET /billing/summary` - Get billing summary
- `POST /billing/credits` - Add credits

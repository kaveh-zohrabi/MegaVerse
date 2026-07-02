# Payment Service

Handles payment processing, subscriptions, and billing.

## Features

- Stripe integration
- Subscription management
- Invoice generation
- Payment history
- Refund processing
- Multi-currency support
- Webhook handling
- Tax calculation

## Endpoints

- `POST /payments/create` - Create payment
- `GET /payments/:id` - Get payment
- `POST /payments/:id/refund` - Refund payment
- `GET /payments/history` - Payment history
- `POST /subscriptions` - Create subscription
- `GET /subscriptions/:id` - Get subscription
- `PUT /subscriptions/:id` - Update subscription
- `DELETE /subscriptions/:id` - Cancel subscription
- `GET /invoices` - List invoices
- `GET /invoices/:id` - Get invoice
- `POST /webhooks/stripe` - Stripe webhook

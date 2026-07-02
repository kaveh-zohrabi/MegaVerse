# Subscription Service

Subscription and plan management.

## Features

- Plan management
- Subscription lifecycle
- Upgrade/downgrade
- Cancellation
- Renewal handling
- Trial management
- Usage tracking
- Subscription analytics

## Endpoints

- `GET /plans` - List plans
- `POST /plans` - Create plan
- `GET /plans/:id` - Get plan
- `PUT /plans/:id` - Update plan
- `POST /subscriptions` - Create subscription
- `GET /subscriptions/:id` - Get subscription
- `PUT /subscriptions/:id` - Update subscription
- `DELETE /subscriptions/:id` - Cancel subscription
- `POST /subscriptions/:id/upgrade` - Upgrade plan

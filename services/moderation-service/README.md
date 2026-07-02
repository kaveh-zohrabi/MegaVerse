# Moderation Service

Content moderation and safety.

## Features

- Content moderation (text, image, video)
- User reporting
- Automated filtering
- Manual review queue
- Appeal process
- Policy management
- Audit logging
- Appeal handling

## Endpoints

- `POST /moderate` - Moderate content
- `POST /reports` - Submit report
- `GET /reports` - List reports
- `PUT /reports/:id` - Update report
- `GET /queue` - Get moderation queue
- `PUT /queue/:id` - Process item
- `GET /policies` - List policies
- `POST /policies` - Create policy

# Notification Service

Manages push notifications, in-app notifications, and alerts.

## Features

- Push notifications (FCM, APNs)
- In-app notifications
- Email notifications
- SMS notifications
- Notification preferences
- Notification history
- Batch notifications
- Scheduled notifications

## Endpoints

- `GET /notifications` - Get notifications
- `PUT /notifications/:id/read` - Mark as read
- `POST /notifications/read-all` - Mark all as read
- `DELETE /notifications/:id` - Delete notification
- `GET /notifications/preferences` - Get preferences
- `PUT /notifications/preferences` - Update preferences
- `POST /notifications/send` - Send notification (internal)

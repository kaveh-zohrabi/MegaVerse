# User Service

Manages user profiles, preferences, and account settings.

## Features

- User profile management
- User preferences
- Account settings
- Avatar upload
- Profile search
- User blocking/reporting
- Privacy settings

## Endpoints

- `GET /users/:id` - Get user profile
- `PUT /users/:id` - Update profile
- `DELETE /users/:id` - Delete account
- `GET /users/:id/avatar` - Get avatar
- `PUT /users/:id/avatar` - Upload avatar
- `GET /users/:id/settings` - Get settings
- `PUT /users/:id/settings` - Update settings
- `GET /users/search` - Search users
- `POST /users/:id/block` - Block user
- `POST /users/:id/report` - Report user

# Social Service

Handles social features like posts, comments, reactions, and followers.

## Features

- Posts (create, read, update, delete)
- Comments and nested replies
- Reactions (like, love, etc.)
- Following/followers system
- News feed generation
- Trending posts
- Content sharing
- Bookmarks

## Endpoints

- `GET /posts` - Get feed
- `POST /posts` - Create post
- `GET /posts/:id` - Get post
- `PUT /posts/:id` - Update post
- `DELETE /posts/:id` - Delete post
- `POST /posts/:id/comments` - Add comment
- `POST /posts/:id/reactions` - Add reaction
- `GET /users/:id/followers` - Get followers
- `GET /users/:id/following` - Get following
- `POST /users/:id/follow` - Follow user
- `DELETE /users/:id/follow` - Unfollow user

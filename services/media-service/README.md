# Media Service

Handles file uploads, processing, and media management.

## Features

- Image upload and processing
- Video upload and transcoding
- Audio processing
- File storage (S3-compatible)
- CDN integration
- Thumbnail generation
- Media metadata
- Content moderation

## Endpoints

- `POST /media/upload` - Upload file
- `GET /media/:id` - Get media info
- `GET /media/:id/download` - Download file
- `DELETE /media/:id` - Delete media
- `POST /media/:id/process` - Process media
- `GET /media/:id/thumbnail` - Get thumbnail
- `GET /media/:id/metadata` - Get metadata

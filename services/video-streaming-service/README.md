# Video Streaming Service

Video upload, transcoding, and streaming.

## Features

- Video upload
- Adaptive streaming (HLS/DASH)
- Transcoding
- Thumbnail generation
- Video analytics
- Live streaming
- DRM protection
- CDN distribution

## Endpoints

- `POST /videos/upload` - Upload video
- `GET /videos/:id` - Get video info
- `GET /videos/:id/stream` - Get stream URL
- `GET /videos/:id/thumbnail` - Get thumbnail
- `DELETE /videos/:id` - Delete video
- `POST /videos/:id/process` - Process video
- `GET /videos/:id/analytics` - Video analytics

# Storage Service

File storage and management.

## Features

- Object storage (S3-compatible)
- File upload/download
- File versioning
- Access control
- Storage quotas
- File sharing
- CDN integration
- Backup management

## Endpoints

- `POST /storage/upload` - Upload file
- `GET /storage/:id` - Get file
- `GET /storage/:id/download` - Download file
- `DELETE /storage/:id` - Delete file
- `GET /storage/:id/versions` - List versions
- `POST /storage/:id/share` - Share file
- `GET /storage/quota` - Get storage quota

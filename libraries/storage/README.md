# MegaVerse Storage Library

File storage abstraction for S3, GCS, Azure Blob.

## Features

- S3-compatible storage
- Google Cloud Storage
- Azure Blob Storage
- Local file system
- File streaming
- Multipart upload
- Presigned URLs

## Usage

```typescript
import { Storage } from '@megaverse/storage';

const storage = new Storage({
  provider: 's3',
  bucket: 'megaverse-storage',
  region: 'us-east-1',
});

// Upload
await storage.put('avatars/user-123.jpg', fileBuffer);

// Download
const file = await storage.get('avatars/user-123.jpg');

// Presigned URL
const url = await storage.getSignedUrl('avatars/user-123.jpg', { expiresIn: 3600 });
```

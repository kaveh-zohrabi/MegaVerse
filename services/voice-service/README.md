# Voice Service

Voice calls and voice features.

## Features

- Voice calls (WebRTC)
- Voice messages
- Speech-to-text
- Text-to-speech
- Voice recording
- Call recording
- Voice analytics
- IVR (Interactive Voice Response)

## Endpoints

- `POST /calls` - Start call
- `PUT /calls/:id` - Join/end call
- `GET /calls/:id` - Get call info
- `POST /voice-messages` - Send voice message
- `GET /voice-messages/:id` - Get voice message
- `POST /transcribe` - Transcribe audio
- `POST /synthesize` - Text to speech

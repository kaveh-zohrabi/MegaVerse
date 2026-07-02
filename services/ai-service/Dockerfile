FROM python:3.12-slim AS builder
WORKDIR /app
COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt
COPY . .
EXPOSE 8089
CMD ["uvicorn", "src.main:app", "--host", "0.0.0.0", "--port", "8089"]

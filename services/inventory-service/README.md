# Inventory Service

Product and inventory management.

## Features

- Product catalog
- Stock management
- SKU management
- Warehouse tracking
- Inventory alerts
- Batch operations
- Barcode/QR support
- Inventory reports

## Endpoints

- `GET /products` - List products
- `POST /products` - Create product
- `GET /products/:id` - Get product
- `PUT /products/:id` - Update product
- `GET /inventory` - Get inventory
- `PUT /inventory/:id` - Update stock
- `GET /inventory/alerts` - Low stock alerts
- `POST /inventory/transfer` - Transfer stock

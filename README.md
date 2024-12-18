# Cipher Tool Backend

This is the **backend** repository for the Cipher Tool app — a powerful cryptographic API that allows users to **hash, encrypt, decrypt**, and **encode** data. The backend is designed to be secure, efficient, and scalable.

## Features

- **Hashing**: Supports MD5, SHA-256, and other hashing algorithms.  
- **Encryption/Decryption**: Implements RSA, AES, and other encryption methods.  
- **Data Encoding**: Base64 encoding/decoding for text and files.  
- **Identicon Generation**: Generates unique visual representations for hashed data.  
- **File Support**: Handles text and file-based cryptographic operations.  

## Tech Stack

- **Golang**: The backend is built with Go for high performance and concurrency.  
- **Vercel**: API is hosted on Vercel's serverless infrastructure.  
- **Standard Libraries**: Utilizes Go's built-in `crypto`, `net/http`, and other libraries for secure cryptography and API handling.

## Endpoints

| **Method** | **Endpoint**       | **Description**                     |
|------------|--------------------|-------------------------------------|
| `POST`     | `/hash`            | Hashes input data (text or file).   |
| `POST`     | `/encrypt`         | Encrypts data using RSA/AES.        |
| `POST`     | `/decrypt`         | Decrypts data using RSA/AES.        |
| `POST`     | `/encode`          | Encodes data (Base64, etc.).        |
| `POST`     | `/identicon`       | Generates identicons for hashed data. |

## How to Run Locally

1. **Clone the repository**:  
   ```bash
   git clone https://github.com/yourusername/cipher-tool-backend.git
   cd cipher-tool-backend
   ```

2. **Install dependencies** (Go must be installed):  
   ```bash
   go mod tidy
   ```

3. **Run the server**:  
   ```bash
   go run main.go
   ```

4. **Access the API**:  
   By default, the server runs at `http://localhost:8080`. Test endpoints using tools like Postman or curl.

## Deployment

The API is deployed on **Vercel** using serverless functions. Due to Vercel's limitations:
- Generated files or images (like identicons) are sent **directly as a response** instead of being stored on the server.

## Example Usage

### Hash Endpoint (Request)
```json
POST /hash
{
    "data": "Hello World",
    "algorithm": "sha256"
}
```

### Response
```json
{
    "hashed": "a591a6d40bf420404a011733cfb7b190d62c65bf0bcda32b1eb2f8c3"
}
```

## Security

- RSA and AES encryption adhere to secure cryptographic practices.
- Sensitive data is not logged or stored.
- Server responses are handled securely over HTTPS.

## Future Enhancements

- Add user authentication (JWT-based) for restricted endpoints.
- Support for more encryption algorithms.
- Rate limiting and logging for better API management.


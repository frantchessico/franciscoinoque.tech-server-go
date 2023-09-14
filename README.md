
# Go Contact API & Send Email

This API is designed to handle contact form submissions, store contact information in a MongoDB database, and send a confirmation email. It provides endpoints to submit contact information and receive confirmation messages.

## Table of Contents

- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
- [Usage](#usage)
  - [API Endpoints](#api-endpoints)
- [Environment Variables](#environment-variables)
- [Contributing](#contributing)
- [License](#license)

## Getting Started

### Prerequisites

Before running the API, ensure you have the following prerequisites installed:

- [Go](https://golang.org/doc/install)
- [MongoDB](https://docs.mongodb.com/manual/installation/)
- [SMTP Server](https://en.wikipedia.org/wiki/Simple_Mail_Transfer_Protocol)

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/contact-api.git
   ```

2. Navigate to the project directory:

   ```bash
   cd contact-api
   ```

3. Create a `.env` file in the project root and configure the following environment variables:

   ```
   MONGODB_URI=your-mongodb-uri
   SMTP_HOST=your-smtp-host
   SMTP_PORT=your-smtp-port
   SMTP_USERNAME=your-smtp-username
   SMTP_PASSWORD=your-smtp-password
   PORT=your-api-port
   ```

4. Build and run the API:

   ```bash
   go build
   ./contact-api
   ```

The API should now be running on the specified port.

## Usage

### API Endpoints

#### Submit Contact Information

- **Endpoint**: `/api/contact`
- **Method**: `POST`
- **Request Body**:
  ```json
  {
    "firstName": "John",
    "lastName": "Doe",
    "email": "johndoe@example.com",
    "linkedin": "linkedin.com/in/johndoe",
    "tech": "Go, JavaScript",
    "message": "Hello, I'm interested in your services."
  }
  ```
- **Response**:
  ```json
  {
    "message": "Data saved successfully and email sent!"
  }
  ```

## Environment Variables

- `MONGODB_URI`: MongoDB connection URI.
- `SMTP_HOST`: SMTP server host.
- `SMTP_PORT`: SMTP server port.
- `SMTP_USERNAME`: SMTP server username.
- `SMTP_PASSWORD`: SMTP server password.
- `PORT`: Port on which the API should run.

## Contributing

Contributions are welcome! Please feel free to open issues or pull requests.

## Developed By

**Francisco Inoque**


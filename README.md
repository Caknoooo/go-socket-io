# Chat Application

This is a simple chat application built with Go and [Socket.IO](https://socket.io/). It allows real-time communication between clients via WebSocket.

> Todo: Integrated with database

## Prerequisites

Before you begin, ensure you have the following installed:

- [Go](https://golang.org/doc/install) (latest version recommended)
- [Socket.IO](https://socket.io/)

### Demo

**Demo with multiple client** 
https://sock.cakno.online/public/

https://github.com/user-attachments/assets/e8124140-95b0-4b70-a89a-f2e2ca4d8c76

**Demo with specific client (on branch -> chat-specific-user)**
https://specific-sock.cakno.online/public/ 

https://github.com/user-attachments/assets/3308a7b0-a0d6-4fe8-8577-a5b1930910a0


## Installation

1. **Clone the Repository:**

    ```bash
    git clone https://github.com/Caknoooo/go-socket-io
    cd go-socket-io
    ```

2. **Install Dependencies**
Make sure you are in the project directory and run:
    ```bash
    go mod tidy
    ```

## Running the Application
1. **Start the server**
To start the server, run:
    ```bash
    go run main.go
    ```

2. **Access the Application in a Browser:**
Open two or more browser tabs and navigate to:
    ```bash
    http://localhost:8222/public
    ```
    Each tab will act as a chat client. You can send messages from one tab and see them appear in the other tab in real-time.

## Troubleshooting
- **Connection Issues:** Ensure there are no network issues or firewall settings blocking the connection.
- **Dependency Errors:** If you encounter dependency errors, make sure to run go mod tidy to update the modules.
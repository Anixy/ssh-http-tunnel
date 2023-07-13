# SSH HTTP Tunnel

SSH HTTP Tunnel is a lightweight tool written in Go that allows you to create an encrypted tunnel between a client machine and a server using SSH and tunnel SSH connections to the browser's HTTP protocol.

# Features

- Establish an encrypted SSH tunnel between the client and server.
- Tunnel SSH through the HTTP/HTTPS traffic.


# Prerequisites

Before using SSH HTTP Tunnel, ensure that you have the following prerequisites installed:

- Go programming language: Download Go

# Installation

1. Clone the SSH HTTP Tunnel repository to your local machine using the following command:
   ```
   git clone https://github.com/your-username/ssh-http-tunnel.git
   ```

3. Change to the project directory:
   ```
   cd ssh-http-tunnel
   ```

4. Build the SSH HTTP Tunnel binary:
   ```
   go build
   ```

# Usage

1. Start the SSH HTTP Tunnel:
   ```
   ./ssh-http-tunnel
   ```

3. Start another terminal, and try to send the data.json file using ssh:
   
   ```
   ssh localhost -p 2222 < data.json
   ```

4. Back to the first terminal and copy the ID, then open your browser try to access `http://localhost:3000?id={id}`.
5. The content of the data.json will be shown on the browser.


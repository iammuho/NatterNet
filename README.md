# NatterNet

Your real-time chat application built with Golang, MongoDB, WebSockets, and Fasthttp.

## Table of Contents

1. [Introduction](#introduction)
2. [Features](#features)
3. [Local Development](#local-development)
4. [API Documentation](#api-documentation-with-swagger)
5. [License](#license)

## Introduction

NatterNet is a fast and lightweight chat WebSocket service built with Golang, MongoDB, WebSockets, and Fasthttp. It aims to provide a scalable and efficient platform for real-time chat applications.

## Features

Our plan for the next iterations of NatterNet is outlined below. We welcome contributions that align with this roadmap!

- [x] **Private Chats (1-1)**
  - Enable direct, private chat between two users
  
- [x] **Group Chats**
  - Allow users to create and participate in group chats
  - Admin/Moderation features in groups
  
- [ ] **Different Message Formats**
  - [x] **Text Messages**
    - Support for plain text messages
    - Emoji support
  - [ ] **Images**
    - Sending and receiving image messages
    - Image preview and download
  - [ ] **Audio**
    - Sending and receiving audio messages
    - Play audio within the app
  - [ ] **Location**
    - Share current location with other users
    - View shared locations on a map
  - [ ] **Video**
    - Sending and receiving video messages
    - Play video within the app
  - [ ] **Contact**
    - Sharing contact details through chat
    - Add shared contacts to the user's address book
- [ ] **Security and Authentication**
  - User authentication (e.g., JWT, OAuth)
  - End-to-end encryption of messages
  - Account recovery options
  
- [ ] **Notifications**
  - Push notifications for new messages
  - Email notifications for missed messages
  
- [ ] **User Presence and Status**
  - Show when users are online/offline
  - Custom user status messages (e.g., "Busy", "Away")
  
- [ ] **Chat History and Backup**
  - Option to export chat history
  - Cloud backup and sync across devices
  
- [ ] **Moderation and Reporting**
  - Block and report abusive users
  - Moderation tools for group chats
  
- [ ] **API and Extensibility**
  - Public API for third-party integrations
  - Plugins or extensions support
  
- [ ] **Internationalization**
  - Multi-language support
  
- [ ] **Performance and Optimization**
  - Efficient handling of large numbers of concurrent connections
  - Database optimizations for scalable message history retrieval


## Local Development

This section provides a guide to set up NatterNet for local development using Docker and Docker Compose. This will create an isolated environment with its own database and dependencies.

### Prerequisites

- [Docker](https://www.docker.com/products/docker-desktop)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [Go](https://golang.org/dl/)

### Steps

1. **Clone the Repository**
   ```sh
   git clone https://github.com/iammuho/natternet.git
   ```

2. **Navigate to the Project Directory**
   ```sh
   cd natternet
   ```

3. **Build and Start the Docker Services**

   This command will start the services defined in `docker-compose.yml` under the project name `natternet`.
   ```sh
   docker-compose -p natternet up -d
   ```

4. **Start the NatterNet Application**
    ```sh
    cd cmd/app/ && go run .
    ```

    At this point, your NatterNet application should be running, and you can access it in your browser at the specified address (e.g., http://localhost:8080).

5. **Access MongoDB via Mongo-Express**

   You can access the Mongo-Express web interface at `http://localhost:8081`.

6. **Stop the Docker Services**

   When you are done with development, you can stop the Docker services with the following command:
   ```sh
   docker-compose -p natternet down
   ```

### Tips

- The `-d` flag in the `docker-compose up -d` command makes the services run in the background. To view the logs for the services, you can run:
   ```sh
   docker-compose -p natternet logs
   ```

- To rebuild the Docker images (e.g., after changing a `Dockerfile`), you can run:
   ```sh
   docker-compose -p natternet build
   ```
## API Documentation with Swagger

You can access the API documentation through Swagger at the following endpoint after starting the application:

[http://localhost:8080/api/v1/swagger/](http://localhost:8080/api/v1/swagger/)

Tip: Before delving into the documentation, ensure to execute `make swagger` at the project root. This step regenerates essential Swagger components, keeping the docs updated.

## License

NatterNet is licensed under the MIT License. See the [LICENSE](LICENSE.md) file for more details.

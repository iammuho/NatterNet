# NatterNet

Your real-time chat application built with Golang, MongoDB, WebSockets, and Fasthttp.

## Table of Contents

1. [Introduction](#introduction)
2. [Features](#features)
3. [Resources & Community](#resources--community)
4. [API Documentation](#api-documentation-with-swagger)
5. [Blog Posts](#📚-blog-posts-for-deeper-understanding)
6. [License](#license)

## Introduction

NatterNet is a fast and lightweight chat WebSocket service built with Golang, MongoDB, WebSockets, and Fasthttp. It aims to provide a scalable and efficient platform for real-time chat applications.

## Features

### Completed Features

- **User Authentication**
  - JWT-based authentication.
  - RSA keys for JWT signing.

- **Real-time Communication**
  - WebSocket support for real-time updates.

- **Rooms**
  - Create, join, and leave rooms.

- **Message Types**
  - Support for text messages.

- **Distributed Systems**
  - NATS Streaming for event-driven architecture.
  
- **Logging**
  - Comprehensive logging with Uber's Zap logger.

- **API and Web Server**
  - API built on top of Fiber (FastHTTP).

- **Database**
  - MongoDB for persistent storage.

### Upcoming Features

- Support for more message types like images, audio, video, etc.
- Push notifications via Firebase/OneSignal.
- Real-time "typing..." status display.
- User presence system.
- Message editing and deletion.
- Multi-language support.
- Rate limiting for API requests.

## Resources & Community

We've compiled a list of resources that will help you understand, use, and contribute to NatterNet more effectively:

- **Postman Team**: Join our [Postman Team](https://app.getpostman.com/join-team?invite_code=cc442cb6e4c1bcd8c4bebe1246451a6e&target_code=ddee05fca163c199c6e7df2f31a3672b) to explore NatterNet's API.
  
- **Community Chat**: Have questions, feedback, or just want to chat? Join our [Discord Server](https://discord.gg/ewxXez4z).
  
- **Wiki**: For detailed guides, FAQs, and other information, check out the [NatterNet Wiki](https://github.com/iammuho/NatterNet/wiki).
  
- **Blog Posts**: Deep dives into the architecture, domain-driven design, and other facets of NatterNet are available in our [Substack Blog](https://muhammetarslan.substack.com/t/natternet).

Feel free to contribute by submitting PRs, opening issues, or spreading the word.

## API Documentation with Swagger

You can access the API documentation through Swagger at the following endpoint after starting the application:

[http://localhost:8080/api/v1/swagger/](http://localhost:8080/api/v1/swagger/)

Tip: Before delving into the documentation, ensure to execute `make swagger` at the project root. This step regenerates essential Swagger components, keeping the docs updated.

## 📚 Blog Posts for Deeper Understanding

Want to get more in-depth insights into NatterNet? Check out our blog posts:

1. **[Introducing NatterNet: A Domain-Driven Real-Time Chat Application](https://muhammetarslan.substack.com/p/introducing-natternet-a-domain-driven)**
    - Introduction to NatterNet and its architecture.

2. **[NatterNet: A Deep Dive into Clean Architecture](https://muhammetarslan.substack.com/p/natternet-a-deep-dive-into-clean)**
    - Learn the details of how Clean Architecture is implemented in NatterNet.

3. **[Deep Dive into Domain-Driven Design in NatterNet](https://muhammetarslan.substack.com/p/deep-dive-into-domain-driven-design)**
    - Understand the nuances of Domain-Driven Design as it's applied in NatterNet.

## License

NatterNet is licensed under the MIT License. See the [LICENSE](LICENSE.md) file for more details.

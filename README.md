# Go-Sid-Server

![Go](https://img.shields.io/badge/Go-1.16-blue.svg)
![License](https://img.shields.io/badge/License-MIT-green.svg)

Welcome to **Go-Sid-Server**! This is a simple HTTP server written in Go. It serves static files and handles form submissions.

## Table of Contents

- [Introduction](#introduction)
- [Features](#features)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Usage](#usage)
- [Folder Structure](#folder-structure)
- [Contributing](#contributing)
- [License](#license)
- [Contact](#contact)

## Introduction

**Go-Sid-Server** is a basic web server implemented in Go. It demonstrates handling static files, processing form data, and simple routing.

## Features

- Serve static HTML files
- Handle form submissions
- Simple routing for endpoints

## Prerequisites

- [Go](https://golang.org/doc/install) (version 1.16 or later)
- A terminal or command prompt
- Git (for cloning the repository)

## Installation

1. **Clone the repository**:
   ```sh
   git clone https://github.com/siddharttth/go-server.git
   cd go-server
2. **Initialize the Go module:**
   ```sh
   go mod init github.com/siddharttth/go-server
3. **Run the server:**
   ```sh
   go run main.go

## Usage
- Access the server:

- Open your web browser and go to http://localhost:3000/.
- Endpoints:

- GET / - Serves the index.html file.
- POST /form.html - Handles form submissions and fill the form.
- GET /hello - Simple hello endpoint.

## Folder Structure Diagram
go-server
├── static
│   ├── form.html
│   └── index.html
├── main.go
└── README.md


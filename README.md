# GoBankAPI
A JSON API for a Bank entirely built in Golang.

## Overview
The Bank JSON API is a lightweight and secure web service that provides banking functionalities via a JSON-based API. 
This project is written entirely in Golang and aims to offer a reliable and extensible platform for managing financial transactions, accounts, and customer data.

## Features
* Account Management: Create, view, update, and delete customer accounts.
* Transactions: Perform deposits, withdrawals, and fund transfers between accounts.
* Authentication: Secure user authentication and authorization mechanisms.
* Data Persistence: Store customer and transaction data in a robust database.
* JSON API: Communicate with the API using JSON for easy integration with various client applications.

## Getting Started
Follow these steps to get the Bank JSON API up and running on your local development environment:

### Prerequisites
* Golang (v1.21.1 or higher)
* PostgreSQL (v10 or higher)

### Installation
1. Clone repository
  ```bash
  git clone https://github.com/Clinton-Nmereole/GoBankAPI.git
  ```
2. Install go dependencies
  ```bash
  go install
  ```
3. Set up the PostgreSQL database
4. Start the server:
  ```bash
  make run
  ```

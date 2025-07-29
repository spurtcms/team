# FoodTraze

## Table of Contents

- [Introduction](#1.-introduction)
- [Prerequisites](#prerequisites)
- [Network Setup](#network-setup)


## 1. Introduction

**FoodTraze** is a blockchain-based food traceability application built on **Hyperledger Fabric**.

- Enables secure, transparent tracking of food products across the supply chain.
- Designed to improve food safety and ensure accountability.
- Each transaction is immutably recorded, enabling real-time traceability and auditability.
- Ideal for producers, distributors, retailers, and regulators.
- Built with a modular, scalable architecture for enterprise-grade performance.

---

## 2. Prerequisites

### Step 1: Install GIT
Download the latest version of Git (if not already installed):  
[https://enterprise.github.com/releases](https://enterprise.github.com/releases)

### Step 2: Install cURL
Download and install the latest version of cURL:  
[https://curl.se/download.html](https://curl.se/download.html)

### Step 3: Docker and Docker Compose
- **Docker Version Required:** 17.06.2-ce or greater
- **Platforms:** macOS, Linux, Windows 10 (Docker Toolbox for older Windows versions)

> Note: Installing Docker for Mac/Windows or Docker Toolbox will also install Docker Compose.  
> Ensure Docker Compose version is **1.14.0 or greater**.

### Step 4: Install Go
Download and install **Go v1.20.12** from:  
[https://go.dev/dl/](https://go.dev/dl/)

### Step 5: Install Node.js & npm
Run the following commands to install Node.js and dependencies:

```bash
sudo apt-get install nodejs
npm install
```

Once these steps are completed, you are ready to set up and configure the FoodTraze application.

---

## 3. Network Setup

### Step 1: Clone the FoodTraze Network Repository
Clone the official GitHub repository:

```bash
git clone -b predev https://github.com/hyperledger-foodtraze/foodtraze-network.git
cd foodtraze-network
```

> This clones the `predev` branch. Make sure to remain on this branch to avoid compatibility issues.

### Step 2: Download Hyperledger Fabric Binaries & Docker Images

Run the following script from inside the `foodtraze-network` directory:

```bash
curl -sSL https://bit.ly/2ysbOFE | bash -s -- -- 1.5.6
```

This script will:
- Download Hyperledger Fabric v1.5.6 binaries (`peer`, `orderer`, `configtxgen`, etc.)
- Pull required Docker images (e.g., peer, orderer, CA, CouchDB)

> Make sure your internet connection is stable, as this step downloads several large files.

### Step 3: Start the FoodTraze Blockchain Network

To launch the blockchain network, run:

```bash
./network.sh up createChannel -ca -s couchdb
```

This script will:
- Start all Docker containers (peers, orderers, CA, CouchDB, etc.)
- Create and join channels
- Install and instantiate chaincode

Once the script completes successfully, your **FoodTraze network is live** and ready for interaction via API or the Explorer UI.

---

Feel free to contribute or raise issues via the [FoodTraze GitHub repository](https://github.com/hyperledger-foodtraze/foodtraze-network).

# Go-App-K8s 🌐🚀

A hands‑on Kubernetes demo application built in Go — designed to demonstrate core Kubernetes concepts, multi‑node cluster behavior, and a full **CI/CD pipeline** using **GitHub Actions**.

---

## 🧠 Project Overview

This repository contains a simple Go web application that:

✨ Shows which Kubernetes **pod** handled each request  
✨ Uses **nginx** as a lightweight load balancer  
✨ Runs on a **local 3‑node Minikube cluster**  
✨ Is deployed and updated automatically via **GitHub Actions CI/CD**

This project was created for **learning and practicing real‑world Kubernetes workflows**, including containerization, scheduling across nodes, service balancing, and automated delivery.

---

## 🛠 Architecture & Concepts

### 🐳 1. Go Application

The Go app listens on a port and responds with its own ID or pod name, allowing you to visualize how requests are load‑balanced across replicas.

### 📦 2. Containerization

Built using a lightweight Docker image defined in `Dockerfile`, which packages the Go server ready for Kubernetes.

### ☸️ 3. Kubernetes Cluster (Minikube)

I've set up a **local multi‑node Minikube Kubernetes cluster** to simulate a distributed environment. Each node can host pods, and Kubernetes **services** route traffic between them.

### ⚖️ 4. Load Balancing with Nginx

An **nginx** service load balances incoming requests across the deployed pods, providing a simple way to observe pod scheduling and server response distribution.

### 🚦 5. CI/CD with GitHub Actions

I've implemented an automated pipeline:

1. Build & test the Go app
2. Build Docker image
3. Push to DockerHub
4. Apply updated manifests to Kubernetes

Using GitHub Actions ensures every push to `main` triggers an automated deployment.

---

## 📁 Repository Structure

```
.
├── .github/
│   └── workflows/          # GitHub Actions CI/CD pipeline configs
├── k8s/                    # Kubernetes manifests (Deployment, Service, LoadBalancer, etc.)
├── Dockerfile              # App container definition
├── main.go                 # Go application source
├── go.mod                  # Go modules
├── README.md               # Documentation
```

---

## 🚀 Getting Started

### 🧩 Prerequisites

Before running the app locally you'll need:
- Go (>= 1.17)
- Docker
- Minikube
- kubectl configured
- GitHub repository with Actions enabled

### 📦 Build & Deploy

#### 📊 Step 1 — Clone the Repo

```bash
git clone https://github.com/Gerges-Wageh/Go-app-k8s.git
cd Go-app-k8s
```

#### 🐋 Step 2 — Build the Image

```bash
docker build -t yourusername/go-app-k8s:latest .
```

#### ☸️ Step 3 — Deploy to Kubernetes

Ensure your Minikube cluster is running:

```bash
minikube start --nodes=3
kubectl apply -f k8s/
```

#### 🌐 Step 4 — Test Load Balancing

Run:

```bash
minikube service nginx-service --url
```

Then curl repeatedly to see traffic distributed across pods.

---

## 📈 CI/CD Pipeline

Your GitHub Actions workflow (in `.github/workflows/`) is configured to:

1. Build Docker image
2. Push to registry
3. Apply Kubernetes manifests to cluster

Every **push to `main`** automatically triggers this pipeline.

---

## 🧠 What I've learned From This Project

✔ Building & containerizing a real Go application  
✔ Creating and managing deployments on Kubernetes  
✔ Running multi‑node Minikube clusters  
✔ Configuring services & load balancing  
✔ Designing a CI/CD workflow using GitHub Actions

---

## 📌 Notes

- Workloads are scheduled across 3 Minikube nodes for learning purposes
- nginx demonstrates a simple load balancer
- This repo is ideal for beginners practicing Kubernetes in local environments

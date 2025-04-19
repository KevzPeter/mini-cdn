# MiniCDN

## Kubernetes-Based Static Site with Go Cache Controller

This project is a simplified, self-hosted Content Delivery Network (CDN) system designed to demonstrate core infrastructure concepts like edge caching, routing, and control planes using Kubernetes.

It serves static content through an NGINX reverse proxy, routes traffic via Kubernetes Ingress, and includes a custom-built Go-based controller to simulate cache invalidation.

---

## 🌐 Features

- Static site deployed via NGINX in Kubernetes
- Ingress-based routing with custom domain `cdn.local`
- Prometheus + Grafana monitoring stack for observability
- Go microservice for triggering cache purge/invalidation
- Fully containerized and Minikube-compatible

---

## 🛠 Architecture Overview

```
                 +----------------------------+
                 |        Browser (User)      |
                 +----------------------------+
                              |
                              v
                      [cdn.local domain]
                              |
                              v
               +-------------------------------+
               |    NGINX Ingress Controller   |
               +-------------------------------+
                        |            |
                        v            v
              +----------------+   +---------------------+
              |  Static Site   |   |  Go Cache Controller|
              |  (Nginx Pod)   |   |  (/purge endpoint)  |
              +----------------+   +---------------------+

                        ⭣
          [Optional: Prometheus scrapes metrics from both]
```

---

## 🛠 Stack

- **Kubernetes** (via Minikube)
- **Nginx** (reverse proxy + static server)
- **Go** (custom purge controller)
- **Docker** (image packaging)
- **Prometheus + Grafana** (monitoring & dashboards)

---

## 🚀 Getting Started

### 1. Clone the Repo

```bash
git clone https://github.com/yourusername/minicdn-go-controller.git
cd minicdn-go-controller
```

### 2. Start Minikube (if not running)

```bash
minikube start --driver=docker
minikube addons enable ingress
```

### 3. Build and Load Docker Images

```bash
# Static site
docker build -t minicdn-static-site ./static-site
minikube image load minicdn-static-site

# Go controller
docker build -t purge-controller:latest ./go-purge-controller
minikube image load purge-controller:latest
```

### 4. Deploy to Kubernetes

```bash
kubectl apply -f k8s/static-site-deploy.yaml
kubectl apply -f k8s/purge-controller-deploy.yaml
kubectl apply -f k8s/cdn-ingress.yaml
```

### 5. Tunnel + Access the Site

```bash
minikube tunnel
```

Then in your browser:  
**http://cdn.local**

> Be sure to add `127.0.0.1 cdn.local` to your `/etc/hosts` (Windows: `C:\Windows\System32\drivers\etc\hosts`)

---

## 🔀 Purge the Cache (Simulated)

```bash
kubectl port-forward svc/purge-controller-svc 8081:80
curl -X POST http://localhost:8081/purge
```

Expected response:

```
📅 Cache purged! Total purges: 1
```

---

## 📊 Monitoring (Prometheus + Grafana)

To install:

```bash
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm install kube-monitoring prometheus-community/kube-prometheus-stack --namespace monitoring --create-namespace
```

Access Grafana:

```bash
kubectl port-forward svc/kube-monitoring-grafana -n monitoring 3000:80
# Then open http://localhost:3000
```

Get admin password:

```bash
kubectl get secret --namespace monitoring kube-monitoring-grafana -o jsonpath="{.data.admin-password}" | base64 --decode
```

---

## 🧠 Future Enhancements

- Add Redis to simulate real caching layer
- Expose Prometheus metrics from Go controller (`/metrics`)
- Add PromQL-based dashboards for cache purge frequency, traffic stats
- Integrate with cert-manager for HTTPS

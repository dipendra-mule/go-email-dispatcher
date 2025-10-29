# go-email-dispatcher


<img width="507" height="516" alt="Flow to visualize_ - visual selection" src="https://github.com/user-attachments/assets/67ba323a-93ee-46ac-9c97-ca785c416276" />

A compact, opinionated **email dispatching microservice** written in Go. It demonstrates a producer/consumer pattern for reading email data (CSV or other sources), queueing messages, and sending them via SMTP with simple retry/logging behaviour.

---

## 🔍 Project overview

This repository implements a small, easy-to-read email dispatcher with three main parts:

* **Producer** — reads input (CSV / other) and enqueues email jobs.
* **Consumer** — dequeues jobs and sends emails over SMTP.
* **Main / Runner** — wiring, configuration and example of how to run the producer and consumer(s).

The code is minimal and intended as a learning example or a starting point for a production-ready dispatcher (add TLS, templates, queuing backend, metrics, etc.).

---

## Features

* Simple SMTP email sending.
* CSV-driven sample producer (`email.csv`) to demonstrate bulk jobs.
* Producer/consumer separation for easy extension (add Redis/RabbitMQ later).
* Small, readable codebase — easy to extend with templates, attachments, and retries.

---

## Files of interest

* `main.go` — example wiring / entrypoint.
* `producer.go` — reads `email.csv` and publishes email jobs.
* `consumer.go` — receives jobs and sends emails through SMTP.
* `email.csv` — sample CSV with columns used by the producer.
* `image.png`, `image00.png` — repository images used below.

---
![image](https://raw.githubusercontent.com/dipendra-mule/go-email-dispatcher/master/image.png)

## Architecture 
<img width="2772" height="2016" alt="Flow to visualize_ - visual selection (1)" src="https://github.com/user-attachments/assets/746a981d-3dc6-467e-9151-6cde118e91ad" />



> **Notes:** The repository uses an in-process queue by default (simple channel-based queue). For high-scale use, replace `In-memory Queue` with Redis / RabbitMQ / Kafka and add durable retries and backoff policies.

---

## Quick Start

**Requirements**

* Go 1.18+
* Access to an SMTP server (e.g., Gmail/SendGrid) or a local SMTP test server (mailhog, smtp4dev)

**Clone & install**

```bash
git clone https://github.com/dipendra-mule/go-email-dispatcher.git
cd go-email-dispatcher
go mod tidy
```

**Run (example)**

1. Edit the SMTP credentials and sender/recipient settings in `main.go` or provide them via environment variables if you wire that in.
2. Use `email.csv` as a sample source — the producer will read it and push jobs.

```bash
# build and run
go run .
```

> The repository contains a tiny example wiring in `main.go`. It demonstrates reading the CSV and running producer/consumer in the same process for demo purposes.

---

## Implementation highlights (what to look for in code)

* `producer.go`

  * Reads `email.csv` and maps each line to a simple `EmailJob` structure.
  * Publishes jobs to a channel/queue for consumers.
* `consumer.go`

  * Receives `EmailJob` objects and composes an SMTP message.
  * Uses Go's `net/smtp` (or a similar SMTP helper) to deliver messages.
  * Logs success and errors.
* `main.go`

  * Starts producer and a configurable number of consumers.
  * Demonstrates graceful shutdown (or can be extended to do so).

---

## Example `email.csv`

You can find the sample CSV in the repo — here's a preview of the file contents (keeps columns simple for demonstration):

```
from,to,subject,body
from@example.com,to@example.com,Hello,This is a test
```

Raw file: [https://raw.githubusercontent.com/dipendra-mule/go-email-dispatcher/master/email.csv](https://raw.githubusercontent.com/dipendra-mule/go-email-dispatcher/master/email.csv)

---

## Extending & Productionizing (recommendations)

If you'd like to move this toward production-ready use, consider:

1. **Secure credentials** — load SMTP credentials from a secrets manager or environment variables; avoid hard-coded secrets.
2. **TLS & Auth** — ensure SMTP connections use STARTTLS/TLS and support modern auth (OAuth2 for providers like Gmail/Google Workspace).
3. **Durable queue** — replace the in-memory queue with Redis / RabbitMQ / SQS for persistence and horizontal scale.
4. **Retry & backoff** — implement exponential backoff and dead-letter queues for failed deliveries.
5. **Templates & attachments** — add HTML templates and attachment handling.
6. **Observability** — add structured logs, metrics (Prometheus), and tracing (OpenTelemetry).
7. **Rate limiting & batching** — throttle sends to respect provider limits and implement batch sends if useful.

---

## Images from the repo

Repository includes two images — embedded here for convenience:

![Feature Illustration](https://raw.githubusercontent.com/dipendra-mule/go-email-dispatcher/master/image00.png)

---

## Contributing

1. Fork the repo
2. Create a feature branch (`git checkout -b feat/your-feature`)
3. Add tests where applicable
4. Open a PR with a clear description
---

---

## Checklist / TODOs for this README (optional)

* [ ] Add configuration examples (env var names) and a sample `config.yaml`.
* [ ] Include snippet of `main.go` wiring with exact function names once confirmed.
* [ ] Add a sequence diagram image if you want a visual PNG/SVG instead of Mermaid.

---

*Generated by ChatGPT — improved, detailed README and napkin flow included.*

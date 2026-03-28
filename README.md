# Crawly 🕷️

A simple web crawler written in Go — built to practice core Go concepts like HTTP requests, parsing, modular design, and package organization.


## ⚙️ Features

* Crawl a given URL
* Extract links from HTML pages
* Remove duplicate URLs
* Check HTTP status codes
* Simple logging system
* Modular structure (separation of concerns)

## 🧱 Project Structure

```
crawly/
├── cmd/            # Entry point (CLI)
├── crawly/         # Core crawling logic
├── checker/        # URL validation & utilities

## 🛠️ How It Works

1. Start with a root URL
2. Fetch HTML content
3. Parse links (`<a>` tags)
4. Normalize & deduplicate URLs
5. Optionally check status codes
6. Log results

## ▶️ Usage

```bash
go run /cmd/main.go <url>
```

Example:

```bash
go run /cmd/main.go https://example.com
```

## 📄 License

Apache 2.0

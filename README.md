# 📌 c45-decision-tree : A Fast & Scalable Decision Tree

 A **high-performance and scalable Decision Tree (C4.5) classifier** in Go. It allows users to train a C4.5 decision tree model and make predictions using the trained model.

## Table of Contents

- [📌 c45-decision-tree : A Fast & Scalable Decision Tree](#-c45-decision-tree--a-fast--scalable-decision-tree)
- [🚀 Features](#features)
- [📂 Project Structure](#project-structure)
- [📥 Installation](#installation)
  - [1️⃣ Install Go (if not installed)](#1️⃣install-go-if-not-installed)
  - [2️⃣ Clone the Repository](#2️⃣clone-the-repository)
- [🔧 Usage](#usage)
  - [CLI Usage](#cli-usage)
- [📜 License](#license)
- [🙌 Contributors](#contributors)

## 🚀 Features

✔ **Parallel Processing**: It employs goroutines to facilitate faster processing of files.

## 📂 Project Structure

```plaintext
├── cmd/                # CLI commands
│   ├── root.go         # Root command
│   ├── train.go        # Train command
│   ├── predict.go      # Predict command
│   ├── evaluate.go     # (Optional: for model evaluation)
├── internal/           # Core logic (separated for modularity)
│   ├── c45/            # C4.5 algorithm implementation
│   │   ├── train.go
│   │   ├── predict.go
│   │   ├── tree.go     # Decision tree struct & functions
│   │   ├── utils.go    # Data processing functions
├── models/             # Stored trained models
│   ├── model.json
├── pkg/                # Reusable utility packages
│   ├── config/         # Configurations
│   ├── logger/         # Logging utilities
├── go.mod
├── main.go             # Entry point 
```

## 📜 License

This project is licensed under [MIT]()

## 🙌 Contributors

This project is open to contributions. Follow the steps below:

- Open an issue.
- Make your contributions.
- Create a pull request.

Developed with ❤️ in Go.
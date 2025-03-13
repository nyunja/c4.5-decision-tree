# 📌 c45-decision-tree : A Fast & Scalable Decision Tree

 A **high-performance and scalable Decision Tree (C4.5) classifier** in Go. It allows users to train a C4.5 decision tree model and make predictions using the trained model.

## Table of Contents

- [📌 c45-decision-tree : A Fast & Scalable Decision Tree](#-c45-decision-tree--a-fast--scalable-decision-tree)
- [🚀 Features](#-features)
- [📂 Project Structure](#-project-structure)
- [📥 Installation](#-installation)
    - [1️⃣ Install Go (if not installed)](#1️⃣-install-go-if-not-installed)
    - [2️⃣ Clone the Repository](#2️⃣-clone-the-repository)
- [🔧 Usage](#-usage)
    - [Build](#build)
    - [CLI Usage](#cli-usage)
    - [Training a Decision Tree](#training-a-decision-tree)
- [📜 License](#-license)
- [🙌 Contributors](#-contributors)

## 🚀 Features

✔ **Read CSV data file**: It reads CSV files and extracts labels and probabilities.

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
│   │   ├── readFile.go # Read input CSV file
│   │   ├── inferTypes.go # Infer data types for each column data
│   │   ├── parseData.go # Parse metadata
├── models/             # Stored trained models
│   ├── model.json
├── pkg/                # Reusable utility packages
│   ├── config/         # Configurations
│   ├── logger/         # Logging utilities
├── go.mod
├── main.go             # Entry point 
```

## 📥 Installation

### **1️⃣ Install Go (if not installed)**

Ensure you have Go installed on your system.🔗 [Download Go](https://golang.org/dl/)

### **2️⃣ Clone the Repository**

```shellscript
git clone https://learn.zone01kisumu.ke/git/tesiaka/c4.5-decision-tree.git
cd text-indexer
```

Install dependencies

```shellscript
go mod tidy
```

## 🔧 Usage

### Build

- First build the project:

```go
go build
```

### CLI Usage

### Training a Decision Tree

| Flag | Description
|-----|-----
| -c | Training a decision tree
| -i | Input CSV file path containing training dataset
| -t | Name of column in the dataset containing the target labels
| -o | Output JSON serialised format file path

- To train a decision tree

```bash
dt -c train -i <input_data_file.csv> -t <target_column> -o <output_tree.dt>
```

### Prediction

| Flag | Description
|-----|-----
| -c | Specify the predict command
| -i | Input CSV file path containing training dataset
| -m | Path to the trained decision tree model file
| -o | Path to save predictions as a CSV file

- To predict using a trained model

```bash
dt -c predict -i <input_data_file.csv> -m <model.dt> -o <output_tree.csv>
```

## 📜 License

This project is licensed under [MIT](https://learn.zone01kisumu.ke/git/tesiaka/c4.5-decision-tree/src/branch/main/LICENSE)

## 🙌 Contributors

This project is open to contributions. Follow the steps below:

- Open an issue.
- Make your contributions.
- Create a pull request.

Developed with ❤️ in Go.
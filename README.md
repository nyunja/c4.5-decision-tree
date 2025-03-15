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
- [ Contribution](#-contributors)

## 🚀 Features

✔ **Read CSV data file**: It reads CSV files and extracts labels and probabilities.

✔ **Parallel Processing**: It employs goroutines to facilitate faster processing of files.

## 📂 Project Structure

```plaintext
── cmd/                # CLI commands and argument parsing
│   ├── root.go         # Defines and parses the root CLI command
│
├── internal/model/     # Core logic for decision tree training and predictions
│   ├── cache/         # Caches computed values to optimize performance
│   ├── counter/       # Computes class distributions (e.g., mode in a class)
│   ├── entropy/       # Calculates data uncertainty (entropy calculation)
│   ├── model/         # Trains the decision tree based on input data
│   ├── node/          # Defines tree node structure and utility functions
│   ├── parser/        # Parses input files (CSV, dt, etc.) into usable data
│   ├── predict/       # Uses the trained model to make predictions
│   ├── split/         # Finds the best feature split to maximize information gain
│   ├── types/         # Defines tree structure and related data types
│   ├── utils/         # Utility functions for data preprocessing
│
├── tree_models/       # Stores serialized trained decision tree models
│
├── util/              # Utility functions (error handling)
│
├── go.mod             # Go module dependencies
├── go.sum             # Go dependency checksums
├── LICENSE            # License information
├── main.go            # Entry point of the application
```

## 📥 Installation

### **1️⃣ Install Go (if not installed)**

Ensure you have Go installed on your system.🔗 [Download Go](https://golang.org/dl/)

### **2️⃣ Clone the Repository**

```shellscript
git clone https://learn.zone01kisumu.ke/git/tesiaka/c4.5-decision-tree.git
cd c4.5-decision-tree
```

Install dependencies

```shellscript
go mod tidy
```

## 🔧 Usage

### Build

- First build the project:

```go
go build -o dt
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
./dt -c train -i <input_data_file.csv> -t <target_column> -o <output_tree.dt>
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
./dt -c predict -i <input_data_file.csv> -m <model.dt> -o <output_tree.csv>
```

## 📜 License

This project is licensed under [MIT](https://learn.zone01kisumu.ke/git/tesiaka/c4.5-decision-tree/src/branch/main/LICENSE)

## 🙌 Contributors

1. [John Paul](https://github.com/nyunja)
2. [Antony Odour](https://github.com/oduortoni)
3. [Teddy Siaka](https://github.com/Siak385)
4. [David Jesse](https://github.com/DavJesse)
5. [Amos Joel](https://github.com/Murzuqisah)

## Contribution

This project is open to contributions. Follow the steps below:

- Open an issue.
- Make your contributions.
- Create a pull request.

Developed with ❤️ in Go.
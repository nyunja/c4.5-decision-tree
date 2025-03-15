# 📌 **C4.5 Decision Tree** – Fast & Scalable Decision Tree in Go  

A **high-performance and scalable C4.5 decision tree classifier** implemented in Go. This tool enables users to efficiently train a **decision tree model** on structured data and make **accurate predictions** using the trained model.  

🚀 **Optimized for speed, parallel execution, and large datasets.**  

## 📑 **Table of Contents**  

- [🚀 Features](#-features)  
- [⚙️ How It Works](#️-how-it-works)  
- [📂 Project Structure](#-project-structure)  
- [📥 Installation](#-installation)  
- [🔧 Usage](#-usage)  
  - [Building the Project](#building-the-project)  
  - [Training a Decision Tree](#training-a-decision-tree)  
  - [Making Predictions](#making-predictions)  
- [📜 License](#-license)  
- [🙌 Contributors](#-contributors)  
- [🤝 Contributing](#-contributing)  

---

## 🚀 **Features**  

✔ **CSV Data Processing** – Reads CSV files, extracts features, and identifies the target labels.  
✔ **Parallel Processing** – Uses Go **goroutines** to speed up data handling and decision tree building.  
✔ **C4.5 Algorithm** – Implements the **C4.5 decision tree** with entropy-based splitting and pruning.  
✔ **Feature Selection** – Selects the **best feature** at each node to maximize **information gain**.  
✔ **Handles Missing Values** – Uses smart imputation techniques to handle missing data.  
✔ **Fast Predictions** – Efficiently classifies new data points using the trained decision tree model.  
✔ **Serialization** – Saves trained models as JSON files for later use in predictions.  
✔ **Command-Line Interface** – Simple CLI for training and predicting with decision trees.  

---

## ⚙️ **How It Works**  

1️⃣ **Data Processing**: Parses CSV files and detects headers.  
2️⃣ **Feature Selection**: Uses **entropy and information gain** to find the best splits.  
3️⃣ **Tree Building**: Recursively builds the decision tree, using pruning for efficiency.  
4️⃣ **Model Storage**: Saves the trained decision tree in a **serializable JSON format**.  
5️⃣ **Predictions**: Uses the trained tree to **classify new input data**.  

---

## 📂 **Project Structure**  

```plaintext
|─ cmd/                # CLI commands and argument parsing  
│   ├── root.go        # CLI entry point for commands  
│  
├── internal/model/    # Core logic for decision tree training and predictions  
│   ├── cache/        # Caches computed values for performance optimization  
│   ├── counter/      # Computes class distributions (e.g., mode in a class)  
│   ├── entropy/      # Calculates data uncertainty (entropy calculation)  
│   ├── model/        # Trains the decision tree based on input data  
│   ├── node/         # Defines tree node structure and utility functions  
│   ├── parser/       # Parses CSV files and converts data into structured format  
│   ├── predict/      # Uses the trained model to make predictions  
│   ├── split/        # Finds the best feature split for information gain  
│   ├── types/        # Defines tree structure and related data types  
│   ├── utils/        # Utility functions for data preprocessing  
│  
├── decision_model/    # Stores serialized trained decision tree models  
├── go.mod             # Go module dependencies  
├── go.sum             # Go dependency checksums  
├── LICENSE            # License information  
├── main.go            # Application entry point  
```

---

## 📥 **Installation**  

### **1️⃣ Install Go (if not installed)**  

Ensure you have Go installed. 🔗 [Download Go](https://golang.org/dl/)  

### **2️⃣ Clone the Repository**  

```bash
git clone https://learn.zone01kisumu.ke/git/tesiaka/c4.5-decision-tree.git
cd c4.5-decision-tree
```

### **3️⃣ Install Dependencies**  

```bash
go mod tidy
```

---

## 🔧 **Usage**  

### **Building the Project**  

```bash
go build -o dt
```

This creates an executable **`dt`** for running commands.  

---

### **Training a Decision Tree**  

| Flag | Description |
|------|------------|
| `-c` | Train a decision tree (`train`) |
| `-i` | Input CSV file path containing the training dataset |
| `-t` | Name of the column in the dataset containing the target labels |
| `-o` | Output file to save the trained decision tree (JSON format) |

#### Example (training):  

```bash
./dt -c train -i dataset.csv -t target_column -o model.dt
```

---

### **Making Predictions**  

| Flag | Description |
|------|------------|
| `-c` | Predict command (`predict`) |
| `-i` | Input CSV file containing test data |
| `-m` | Path to the trained decision tree model file |
| `-o` | Path to save predictions as a CSV file |

#### Example (prediction):  

```bash
./dt -c predict -i test_data.csv -m model.dt -o predictions.csv
```

---

## 📜 **License**  

This project is licensed under the **MIT License**.  
🔗 [MIT License](https://learn.zone01kisumu.ke/git/tesiaka/c4.5-decision-tree/src/branch/main/LICENSE)  

---

## 🙌 **Contributors**  

1. [John Paul](https://github.com/nyunja)  
2. [Antony Odour](https://github.com/oduortoni)  
3. [Teddy Siaka](https://github.com/Siak385)  
4. [David Jesse](https://github.com/DavJesse)  
5. [Amos Joel](https://github.com/Murzuqisah)  

---

## 🤝 **Contributing**  

🚀 This project is open for contributions!  

🔹 **How to contribute:**  

1. Fork the repository.
2. Create a new branch for your feature.
3. Commit your changes.
4. Push the branch to your fork.
5. Open a pull request.
6. Submit your pull request.
7. Review and merge.
8. Update the documentation.
9. Update the changelog.
 

💡 **Let's build a faster and more efficient Decision Tree model together!**  

---

**Developed with ❤️ in Go.** 🚀
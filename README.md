# Currency Converter CLI in Go

This is a simple command-line tool that converts one currency to another using a free online exchange rate API. It allows you to input the amount of money and specify the source and target currencies, and then it displays the converted amount.

## Features

- Convert any amount from one currency to another.
- Fetch real-time exchange rates using a free API.
- Supports error handling for invalid input and network issues.
- Interactive command-line input using `fmt.Scanln`.

## Prerequisites

Before running the program, make sure you have the following installed:

- [Go Programming Language](https://golang.org/dl/) (version 1.15 or higher)
- An active internet connection to fetch exchange rates from the API.

## Installation

1. Clone this repository to your local machine:

    ```bash
    git clone https://github.com/your-username/currency-converter-cli.git
    cd currency-converter-cli
    ```

2. Install Go on your system (if you haven't already) by following the [official installation guide](https://golang.org/doc/install).

3. Ensure you have a Go workspace set up. You can verify this by running:

    ```bash
    go version
    ```

## Usage

1. **Clone the repository**:

   Clone the repository to your local machine:

   ```bash
   git clone https://github.com/your-username/currency-converter-cli.git
   cd currency-converter-cli

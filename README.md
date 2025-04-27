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

   To get started, you need to clone the repository to your local machine. Open your terminal and run the following command:

   ```bash
   git clone https://github.com/your-username/currency-converter-cli.git
   ```

   After cloning, navigate into the project directory:

   ```bash
   cd currency-converter-cli
   ```

2. **Install Go** (if you don't have it already):

   To check if Go is installed on your system, you can run:

   ```bash
   go version
   ```

   If Go is not installed, follow the official [Go installation guide](https://golang.org/doc/install) to set it up.

3. **Run the application**:

   Once you have the project set up and Go installed, you can run the application with the following command:

   ```bash
   go run main.go
   ```

   The program will ask you for the following inputs:

   - **Amount**: Enter the amount of money you want to convert.
   - **Source Currency**: Enter the source currency (e.g., USD, EUR).
   - **Target Currency**: Enter the target currency (e.g., GBP, JPY).

   Example input:

   ```
   Enter amount: 100
   Enter source currency (e.g., USD): USD
   Enter target currency (e.g., EUR): EUR
   ```

   After entering the required data, the program will output the converted amount:

   ```
   100.00 USD is equivalent to 85.30 EUR
   ```

4. **Handling errors**:

   If the program encounters any issues, such as an invalid amount, unrecognized currency code, or network issues, it will display an error message.

   Example of an invalid currency code error:

   ```
   Enter amount: 100
   Enter source currency (e.g., USD): USD
   Enter target currency (e.g., ABC): ABC
   Invalid currency code. Please check the currency codes.
   ```

   The program will guide you to provide valid input.

5. **Exit the program**:

   After receiving the converted amount or an error message, the program will terminate. You can simply rerun the program to make another conversion.

---

### Example Session:

When you run the program, it will prompt you for inputs. Here's an example of what a typical session looks like:

```
$ go run main.go
Enter amount: 100
Enter source currency (e.g., USD): USD
Enter target currency (e.g., EUR): EUR
100.00 USD is equivalent to 85.30 EUR
```

If an invalid currency is entered, you might see this:

```
$ go run main.go
Enter amount: 100
Enter source currency (e.g., USD): USD
Enter target currency (e.g., XYZ): XYZ
Invalid currency code. Please check the currency codes.
```

---

### Debugging

If you run into any issues, here are some common troubleshooting steps:

- **Network issues**: If the program can't fetch exchange rates, ensure you have an active internet connection. The API requires a network connection to retrieve real-time data.
  
- **Invalid currency codes**: Make sure you are entering the currency codes in uppercase letters (e.g., USD, EUR, GBP). Currency codes are case-sensitive.

- **Invalid amount**: The program only accepts numeric input for the amount. If you accidentally enter a non-numeric value, the program will prompt you with an error message.


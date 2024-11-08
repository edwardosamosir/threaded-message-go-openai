
# Simple Threaded Message CLI App Using OpenAI

This project is a simple CLI (Command Line Interface) application built with Go that interacts with OpenAI's API using the [go-openai](https://github.com/sashabaranov/go-openai) wrapper. The application supports threaded message interactions, making it easy to engage in multi-turn conversations.

## Features

- **Threaded Message Support**: Enables continuous conversations that remember context.
- **User-friendly CLI**: Easy-to-use interface for initiating and managing interactions with OpenAI.
- **Quick Setup**: Minimal configuration required to get started.

## Prerequisites

Before running this application, ensure you have the following installed:

- [Go](https://golang.org/dl/) (version 1.18 or higher)
- An [OpenAI API Key](https://platform.openai.com/account/api-keys)
- Access to the [go-openai](https://github.com/sashabaranov/go-openai) package

## Installation

1. **Clone the repository**:
   ```bash
   git clone https://github.com/edwardosamosir/threaded-message-go-openai
   cd threaded-message-go-openai
   ```

2. **Install dependencies**:
   ```bash
   go mod tidy
   ```

3. **Set up your OpenAI API Key**:
   Create an environment variable named `OPENAI_API_KEY` and assign your API key to it:
   ```bash
   export OPENAI_API_KEY="your_openai_api_key"
   ```

## Usage

Run the CLI application with the following command:

```bash
go run main.go
```

### How to Use

1. Start the app and follow the prompts to input your message.
2. The app will display OpenAI's responses, allowing you to continue the conversation in a threaded format.

## Example Interaction

```text
Enter a question: What is adapter pattern in design pattern?
Assistant Response:
 The Adapter pattern is a structural design pattern that allows incompatible interfaces to work together. It acts as a bridge between two incompatible interfaces by converting the interface of a class into another interface that a client expects. This pattern enables classes to collaborate that otherwise couldn't because of incompatible interfaces.

### Key Components:
1. **Target Interface**: This is the interface the client expects and uses.
2. **Adaptee**: This is the existing class with a different interface.
3. **Adapter**: This implements the target interface and is responsible for converting calls to methods into calls to methods on the adaptee.

### How it Works:
- The client interacts with the target interface.
- The adapter takes calls from the client and translates them into a format that the adaptee can understand.
- The adapter handles all the necessary conversions and method calls.
- The adaptee processes the requests as per its usual functionality.

### Advantages:
- It helps match interfaces of different classes.
- It improves code reusability.
- It provides flexibility in integrating new components without altering existing code.
- It follows Open/Closed Principle, allowing new functionalities via adapters without modifying existing code.

### Use Cases:
- When you want to use an existing class, and its interface is incompatible with the system.
- When you need to create a reusable class that cooperates with unrelated or unforeseen classes.

### Example:
Imagine you have a `ModernPaymentSystem` class that new applications are using, but there's an older `LegacyPaymentSystem` that your application still needs to interact with. You can create an adapter class that implements the interface expected by `ModernPaymentSystem` but internally uses `LegacyPaymentSystem` to fulfill requests.

In summary, the adapter pattern is crucial for ensuring that systems can interoperate seamlessly with existing systems or third-party libraries that have different interfaces.



Enter a follow-up question (or type 'exit' to quit): what make it different from other pattern?
Assistant Response:
 The Adapter pattern is distinct from other design patterns due to its specific purpose and structural characteristics. Here's how it stands out compared to other common design patterns:

### Specific Purpose:
- **Adapter vs. Facade**: While both are structural patterns, the Adapter pattern is specifically focused on allowing two incompatible interfaces to work together, often dealing with issues of legacy code integration. In contrast, the Facade pattern provides a simplified interface to a set of interfaces in a subsystem, effectively "flattening" complex interfaces into a more understandable one without necessarily dealing with incompatibility.

- **Adapter vs. Decorator**: Both Adapter and Decorator patterns wrap objects. However, the Adapter changes the interface to make incompatible systems work together, while the Decorator extends or alters the behavior of an object without changing its interface. The Decorator pattern is used to add responsibilities to objects dynamically.

- **Adapter vs. Proxy**: The Proxy pattern also wraps another object. However, it is primarily used to control access to the object, possibly by lazy loading, caching, logging, or permission checking, without changing the original interface of the object. The Adapter focuses on converting the interface without adding new behavior beyond what is required for compatibility.

### Structural Characteristics:
- **Different Focus on Interfaces**: The Adapter pattern is deeply concerned with interface compatibility. It specifically deals with the challenge of interfacing two systems that are not directly compatible without altering either system's code. This is less of a concern in patterns like Singleton, Observer, or Factory, which address different architectural or behavioral issues.

### Use Cases:
- **Specific Use Case**: It is predominantly used for compatibility and reusability, making it ideal for integrating third-party code or transitioning from legacy systems to newer systems.

### Implementation Details:
- **Class vs. Object Adapter**: The Adapter pattern can be implemented in two ways:
  - **Class Adapter** (often using multiple inheritance): In languages that support multiple inheritance, a class adapter can inherit interfaces from both the target and the adaptee.
  - **Object Adapter**: Uses composition to hold an instance of the adaptee and implements the target interface, forwarding requests appropriately.

In summary, the Adapter pattern is unique in its primary focus on interface compatibility between incompatible systems, leveraging these distinctions to ensure seamless integration without needing extensive changes to existing code bases. Its specialization in solving integration issues differentiates it from other design patterns, each with its own specific intent and application context.
```



## Acknowledgments

- [OpenAI](https://platform.openai.com) for providing the API.
- [Sasha Baranov](https://github.com/sashabaranov) for the `go-openai` wrapper.

# Forward Proxy Server in Go  

This project implements a **forward proxy server** using Go's low-level `net` package.  
The proxy acts as an intermediary between a client and a target server, forwarding requests and responses while demonstrating key concepts of **network programming**.  

---

## Features  

 **Forward Proxy Functionality**  
- Accepts requests from clients and forwards them to a target server.  
- Relays the response from the target server back to the client.  

 **Dynamic Address Exchange**  
- The client dynamically sends its listening address to the proxy server, enabling flexible communication without hardcoding values.  

**Concurrency**  
- Handles multiple client connections concurrently using **goroutines**.  

 **Custom Protocol**  
- Implements a **custom protocol** for communication between the client and proxy server, ensuring proper data flow.  

---

## Key Achievements  

 **Built a fully functional forward proxy server** using Go's low-level `net` package, which provides raw **TCP socket programming** capabilities.  
 **Demonstrated understanding of core networking concepts** such as:  
- Establishing and managing **TCP connections**.  
- Reading from and writing to connections using **buffers**.  
- Handling **concurrency with goroutines**.  
- Implementing **dynamic address exchange** for scalable communication.  

---

## How It Works  

### **Client**  
1. The client connects to the **proxy server** and sends the **target server’s address**.  
2. After a short delay, the client sends its own **listening address (IP + port)** to the proxy server.  
3. The client listens on a **specific port** for responses from the proxy server.  

### **Proxy Server**  
1. The proxy server receives the **target server’s address** and the **client’s listening address**.  
2. It establishes a **connection to the target server**, forwards the client’s request, and relays the **response** back to the client.  

### **Target Server**  
- The **target server** processes the request and sends a **response** back to the **proxy server**.  

---

## Why Use Go's `net` Package?  

This project was implemented using Go’s **low-level `net` package** to:  

 Gain hands-on experience with **raw TCP socket programming**.  
 Understand the **underlying mechanics** of network communication without relying on higher-level abstractions like `net/http`.  
 Build a **lightweight and efficient proxy server**.  

---

## Future Improvements  

**Replace artificial delays** (`time.Sleep`) with a more robust communication mechanism.  
Use **structured data formats** (e.g., JSON) for requests and responses.  
Simplify the architecture by **reusing a single connection** for both requests and responses.  
Add **error handling, timeouts, and logging** for better reliability.  

---

  

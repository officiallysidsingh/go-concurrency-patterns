# Go Concurrency Patterns
A Step By Step Guide With Building Blocks For Go Concurrency Learning

## Go Routines
- Used for **running concurrent processes** in golang.
- Head to the folder named **goroutines**.
- Read the **code & comments** to learn about goroutines.
  - Yes, creating goroutines are as simple as that.
  - You just need to add **"go"** keyword before any function.

## Channels
- Used for **communication between different goroutines**.
- **Joining part** of the **Fork & Join Model of Concurrency** in golang.
- Head to the folder named **channels**.
- Read the **code & comments** to learn about channels.
  - **For reading data** from channels we have **<- before channel name**.
  ```go
    msg := <- channel
  ```
  - **For sending data** to channel we have **<- after channel name**.
  ```go
    channel <- "Sending Data Into Channel"
  ```
  - **For Recieving channels in a function** as arguments we use the similar arrow format.
  ```go
    // This func is going to read data from channel
    func(channel <-chan string){
    }

    // This func is going to send data into channel
    func(channel chan<- string){
    }
  ```
# Go Concurrency Patterns
A Step By Step Guide With Building Blocks For Go Concurrency Learning

## Go Routines
- Used for **running concurrent processes** in golang.
- You just need to add **"go"** keyword before any function. Yes, creating goroutines are as simple as that.
- Head to the folder named **goroutines**.
- Read the **code & comments** to learn more about goroutines.

## Channels
- Used for **communication between different goroutines**.
- **Joining part** of the **Fork & Join Model of Concurrency** in golang.
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
- Head to the folder named **channels**.
- Read the **code & comments** to learn more about channels.

## Select
- Used for **waiting to get messages from different channels**.
- Executes code based on **first operation that becomes available**.
- If **both the channels become available simultaneously** then it **picks one channel at random**.
```go
  select {
	case msg := <-channel1:
		fmt.Println(msg)
	case msg := <-channel2:
		fmt.Println(msg)
	}
```
- Head to the folder named **select**.
- Read the **code & comments** to learn more about select.
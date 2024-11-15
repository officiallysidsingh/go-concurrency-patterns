# Go Concurrency Patterns
A Step By Step Guide With Building Blocks For Go Concurrency Learning

## Concurrency Primitives
### 1. Go Routines
  - Used for **running concurrent processes** in golang.
  - You just need to add **"go"** keyword before any function. Yes, creating goroutines are as simple as that.
  - Head to the folder **primitives/goroutines**.
  - Read the **code & comments** to learn more about goroutines.

### 2. Channels
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
  - Head to the folder **primitives/channels**.
  - Read the **code & comments** to learn more about channels.

### 3. Select
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
  - Head to the folder **primitives/select**.
  - Read the **code & comments** to learn more about select.

## Concurrency Patterns
### 1. For-Select-Done Pattern
  - In this pattern the parent has the power to stop the working of a goroutine by closing it's done channel.
  ```go
  for {
		select {
		case <-done:
			return
		default:
			fmt.Println("DOING WORK")
		}
	}
  ```
  - Head to the folder **patterns/for-select-done**.
  - Read the **code & comments** to learn more about select.

### 2. Pipeline Pattern
  - Seperating concerns to make code more modular and organized.

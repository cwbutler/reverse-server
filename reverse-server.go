package main

import (
  "fmt"
  "net"
  "bufio"
)

// Reverse returns its argument string reversed rune-wise left to right.
// Wrote this while using the GO tutorial.
func reverse(s string) string {
  r := []rune(s)

  for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
    r[i], r[j] = r[j], r[i]
  }

  return string(r)
}

// Reads in a newline-terminated string, reverses it,
// and returns to client.
func processRequest(connection net.Conn) {
  // Read input from connection until the first occurrence of newline,
  // returning a string.
  request, error := bufio.NewReader(connection).ReadString('\n')

  // Handle error
  if (error != nil) {
    fmt.Println("Error: " + error.Error())
    return
  }

  // Respond with reverse of string
  connection.Write([]byte(reverse(request)))

  // Close connection when done
  connection.Close()
}


// Create a TCP server that listens on port 8080.
func main() {
  server, error := net.Listen("tcp", ":8080")

  // Exit if error occurs
  if error != nil {
    fmt.Println("Error: " + error.Error())
  }

  fmt.Println("Listening @ 127.0.0.1:8080.")
  fmt.Println("Try running, 'echo \"Hello World\" | nc localhost 8080 from terminal")

  // Loop that runs FOREVER!! (OR until ctrl-c)
  for {
    // Block loop until a connection is made
    connection, error := server.Accept()

    // Skip if error
    if error != nil {
      fmt.Println("Error: " + error.Error())
      continue
    }

    // In order to accept multiple connections,
    // seperate the logic of handling connections in
    // a go routine. This allows the loop to continue
    // to accept connections and not block.
    go processRequest(connection)
  }
}

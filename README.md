# reverse-server
A simple server in Go that will accept connections on a TCP socket and read in a newline-terminated string. The server reverses the string, and send the result back over the socket to the client.

# Test
go run reverse-server

Then in another command prompt:
echo "Hello World\n" | nc localhost 8080

This uses netcat to send string to localhost on port 8080.

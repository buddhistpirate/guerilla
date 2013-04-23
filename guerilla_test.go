package main

import "testing"
import "net"
import "bufio"
import "encoding/json"
import "time"

func TestMainForReal(t *testing.T) {
	go main()

	network_interface := ":8080"
	conn, err := net.Dial("tcp", network_interface)
	for i := 0; err != nil && i < 10; i++ {
		conn, err = net.Dial("tcp", network_interface)
		t.Logf("Sleeping for 50 ms to allow server to start. Loop %v", i)
		time.Sleep(50 * time.Millisecond)
	}

	if err != nil {
		t.Fatalf("Error connecting to %v: %v\n", network_interface, err)
	}

	_, err = conn.Write([]byte("2\n"))
	if err != nil {
		t.Errorf("Error writing to %v: %v", network_interface, err)
	}

	reader := bufio.NewReader(conn)
	json_bytes, err := reader.ReadString('\n')

	if err != nil {
		t.Errorf("Error while reading from server. json_bytes= %v, err= %v", json_bytes, err)
	}

	lines := make([]string, 2)

	err = json.Unmarshal([]byte(json_bytes), &lines)
	if err != nil {
		t.Errorf("Error Decoding JSON: %v, %v", json_bytes, err)
	}
	if 2 != len(lines) {
		t.Errorf("Length of Returned Array was %v not 2", len(lines))
	}
}

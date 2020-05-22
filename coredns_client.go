package main
import (
"fmt"
"net"
"os"
"strconv"
)

func sendPacketToDNSCore(conn *net.UDPConn) {
data := []byte("client query for example.com")
    _, err := conn.Write(data)
    if err != nil {
        fmt.Printf("Couldn't send response %v", err)
    }
}

func recvPacketfromDNSCore(conn *net.UDPConn) {
    buffer := make([]byte, 2048)
                n, addr, err := conn.ReadFromUDP(buffer)
                if err != nil {
                        fmt.Println(err)
                        return
                }
				fmt.Println("UDP Server : ", addr)
                fmt.Printf("Response from DNS core is: %s\n", string(buffer[0:n]))
}

func main(){

arguments := os.Args
if len(arguments) == 1 {
                fmt.Println("Please provide remote IP & port!")
                return
        }
dnscore_IP := "192.168.0.103"


dnscore_IP = arguments[1]
dnscore_port1 := arguments[2]
dnscore_port, err := strconv.Atoi(dnscore_port1)
fmt.Println("DNS client is started ...", dnscore_IP, dnscore_port1);




dnscore_addr := net.UDPAddr{
        Port: dnscore_port,
        IP: net.ParseIP(dnscore_IP),
    }

    conn, err := net.DialUDP("udp", nil, &dnscore_addr)
    if err != nil {
        fmt.Printf("Error1 %v", err)
        return
    }
    sendPacketToDNSCore(conn)
	recvPacketfromDNSCore(conn)
	
    conn.Close()
}

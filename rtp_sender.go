package main

import (
    "log"
    "io/ioutil"
    "os"
    "strconv"
    "net"
)

func net_init(_ip string, port int) (*net.UDPConn, error) {
    ip := net.ParseIP(_ip)
	srcAddr := &net.UDPAddr{IP: net.IPv4zero, Port: 0}
	dstAddr := &net.UDPAddr{IP: ip, Port: port}
	conn, err := net.DialUDP("udp", srcAddr, dstAddr)
	if err != nil {
        log.Println(err)
	}
    return conn,err
}

func main() {
    log.SetFlags(log.LstdFlags | log.Lshortfile)
    rtp_pkts,err := ioutil.ReadFile(os.Args[1])
    if err != nil {
        log.Printf("open file: %s error", os.Args[1])
        return
    }
    port,_ := strconv.Atoi(os.Args[3])
    conn, err := net_init(os.Args[2], port)
    if err != nil {
        log.Println(err)
        return
    }
    defer conn.Close()
    idx := 0
    for idx+1400 <= len(rtp_pkts) {
        _, err := conn.Write(rtp_pkts[idx:idx+1400])
        if err != nil {
            log.Printf("send rtp pkt error, idx:%d err: %v", idx, err)
            return
        }
        idx += 1400
    }
}

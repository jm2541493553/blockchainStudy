package main

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/quic-go/quic-go"
	"log"
	"math"
	"math/big"
)

const Size = math.MaxUint16

func clientUDPListen() {
	//listen, err := net.Listen("tcp", clientAddr)
	listen, err := quic.ListenAddr(clientAddr, generateTLSConfig(), nil)
	if err != nil {
		log.Panic(err)
	}
	defer listen.Close()

	for {
		//conn, err := listen.Accept()
		conn, err := listen.Accept(context.Background())
		if err != nil {
			log.Panic(err)
		}
		stream, err := conn.AcceptStream(context.Background())
		if err != nil {
			panic(err)
		}
		buf := make([]byte, Size)
		length, _ := stream.Read(buf)
		fmt.Println(string(buf[0:length]))
	}

}

// 节点使用的udp监听
func (p *pbft) udpListen() {
	//listen, err := net.Listen("tcp", p.node.addr)
	listen, err := quic.ListenAddr(p.node.addr, generateTLSConfig(), nil)
	//fmt.Printf("节点开启监听，地址：%s\n", p.node.addr)
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("节点开启监听，地址：%s\n", p.node.addr)
	defer listen.Close()

	for {
		conn, err := listen.Accept(context.Background())
		if err != nil {
			log.Panic(err)
		}
		stream, err := conn.AcceptStream(context.Background())
		if err != nil {
			panic(err)
		}
		buf := make([]byte, Size)
		length, _ := stream.Read(buf)
		p.handleRequest(buf[0:length])
	}

}

// 使用udp发送消息
func udpDial(message []byte, addr string) {
	tlsConf := &tls.Config{
		InsecureSkipVerify: true,
		NextProtos:         []string{"quic-echo-example"},
	}
	conn, err := quic.DialAddr(context.Background(), addr, tlsConf, nil)
	if err != nil {
		panic(err)
	}

	stream, err := conn.OpenStreamSync(context.Background())
	if err != nil {
		panic(err)
	}
	_, err = stream.Write([]byte(message))
	if err != nil {
		panic(err)
	}
	stream.Close()
}
func generateTLSConfig() *tls.Config {
	key, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		panic(err)
	}
	template := x509.Certificate{SerialNumber: big.NewInt(1)}
	certDER, err := x509.CreateCertificate(rand.Reader, &template, &template, &key.PublicKey, key)
	if err != nil {
		panic(err)
	}
	keyPEM := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(key)})
	certPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certDER})

	tlsCert, err := tls.X509KeyPair(certPEM, keyPEM)
	if err != nil {
		panic(err)
	}
	return &tls.Config{
		Certificates: []tls.Certificate{tlsCert},
		NextProtos:   []string{"quic-echo-example"},
	}
}

package main

import (
	"net"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

func spoofedICMP(spoof, dstIP net.IP) ([]byte, error) {
	ip := &layers.IPv4{
		SrcIP:    spoof, // spoofed source IP
		DstIP:    dstIP,
		Protocol: layers.IPProtocolICMPv4,
		Version:  4,
		TTL:      20,
	}

	icmp := &layers.ICMPv4{
		TypeCode: layers.CreateICMPv4TypeCode(layers.ICMPv4TypeEchoRequest, 0),
	}

	buf := gopacket.NewSerializeBuffer()
	opts := gopacket.SerializeOptions{FixLengths: true, ComputeChecksums: true}

	if err := gopacket.SerializeLayers(buf, opts, ip, icmp); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

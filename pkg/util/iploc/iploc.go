package iploc

import (
	_ "embed"
	"encoding/binary"
	"net"
	"strings"

	"github.com/yinheli/mahonia"
)

// Note: This file is a modified version of https://github.com/yinheli/qqwry.

//go:embed qqwry.dat
var qqwry []byte

const (
	cIndexLen      = 7
	cRedirectMode1 = 0x01
	cRedirectMode2 = 0x02
)

// Find get country and city base ip. return empty country
// and city when ip that pass from argument is empty string
// or invalid ip address that net.ParseIP(...).To4 return nil
// eg: "::1"
func Find(ip string) (string, string) {
	ip = strings.Trim(ip, " ")
	if len(ip) == 0 {
		return "", ""
	}

	to4 := net.ParseIP(ip).To4()
	// If ip is "::1", To4 returns nil.
	if to4 == nil {
		return "", ""
	}

	offset := searchIndex(binary.BigEndian.Uint32(to4))
	if offset <= 0 {
		return "", ""
	}

	var country, area []byte
	mode := readMode(offset + 4)
	if mode == cRedirectMode1 {
		countryOffset := readUInt24(offset + 5)
		mode = readMode(countryOffset)
		if mode == cRedirectMode2 {
			c := readUInt24(countryOffset + 1)
			country = readString(c)
			countryOffset += 4
		} else {
			country = readString(countryOffset)
			countryOffset += uint32(len(country) + 1)
		}
		area = readArea(countryOffset)
	} else if mode == cRedirectMode2 {
		countryOffset := readUInt24(offset + 5)
		country = readString(countryOffset)
		area = readArea(offset + 8)
	} else {
		country = readString(offset + 4)
		area = readArea(offset + uint32(5+len(country)))
	}
	enc := mahonia.NewDecoder("gbk")
	Country := enc.ConvertString(string(country))
	City := enc.ConvertString(string(area))
	return Country, City
}

func readMode(offset uint32) byte {
	return qqwry[offset]
}

func readArea(offset uint32) []byte {
	mode := readMode(offset)
	if mode == cRedirectMode1 || mode == cRedirectMode2 {
		areaOffset := readUInt24(offset + 1)
		if areaOffset == 0 {
			return []byte("")
		} else {
			return readString(areaOffset)
		}
	} else {
		return readString(offset)
	}
}

func readString(offset uint32) []byte {
	data := make([]byte, 0, 30)
	for {
		if qqwry[offset] == 0 {
			break
		}
		data = append(data, qqwry[offset])
		offset++
	}
	return data
}

func searchIndex(ip uint32) uint32 {
	header := qqwry[:8]
	start := binary.LittleEndian.Uint32(header[:4])
	end := binary.LittleEndian.Uint32(header[4:])
	for {
		mid := getMiddleOffset(start, end)
		buf := qqwry[mid : mid+cIndexLen]
		ipaddr := binary.LittleEndian.Uint32(buf[:4])
		if end-start == cIndexLen {
			offset := byte3ToUInt32(buf[4:])
			// TODO: 这里可能有bug，需要优化
			if ip < binary.LittleEndian.Uint32(qqwry[end:end+4]) {
				return offset
			} else {
				return 0
			}
		}
		// 找到的比较大，向前移
		if ipaddr > ip {
			end = mid
		} else if ipaddr < ip { // 找到的比较小，向后移
			start = mid
		} else {
			return byte3ToUInt32(buf[4:])
		}
	}
}

func readUInt24(offset uint32) uint32 {
	return byte3ToUInt32(qqwry[offset : offset+3])
}

func getMiddleOffset(start uint32, end uint32) uint32 {
	records := ((end - start) / cIndexLen) >> 1
	return start + records*cIndexLen
}

func byte3ToUInt32(data []byte) uint32 {
	i := uint32(data[0]) & 0xff
	i |= (uint32(data[1]) << 8) & 0xff00
	i |= (uint32(data[2]) << 16) & 0xff0000
	return i
}

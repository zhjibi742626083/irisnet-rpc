// Autogenerated by Thrift Compiler (0.11.0)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package main

import (
	"context"
	"flag"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"math"
	"model"
	"net"
	"net/url"
	"os"
	"strconv"
	"strings"
)

func Usage() {
	fmt.Fprintln(os.Stderr, "Usage of ", os.Args[0], " [-h host:port] [-u url] [-f[ramed]] function [arg1 [arg2...]]:")
	flag.PrintDefaults()
	fmt.Fprintln(os.Stderr, "\nFunctions:")
	fmt.Fprintln(os.Stderr, "   GetCandidateList(CandidateListRequest req)")
	fmt.Fprintln(os.Stderr, "  Candidate GetCandidateDetail(CandidateDetailRequest req)")
	fmt.Fprintln(os.Stderr, "  ValidatorExRateResponse GetValidatorExRate(ValidatorExRateRequest req)")
	fmt.Fprintln(os.Stderr, "   GetDelegatorCandidateList(DelegatorCandidateListRequest req)")
	fmt.Fprintln(os.Stderr, "  TotalShareResponse GetDelegatorTotalShares(TotalShareRequest req)")
	fmt.Fprintln(os.Stderr)
	os.Exit(0)
}

func main() {
	flag.Usage = Usage
	var host string
	var port int
	var protocol string
	var urlString string
	var framed bool
	var useHttp bool
	var parsedUrl *url.URL
	var trans thrift.TTransport
	_ = strconv.Atoi
	_ = math.Abs
	flag.Usage = Usage
	flag.StringVar(&host, "h", "localhost", "Specify host and port")
	flag.IntVar(&port, "p", 9090, "Specify port")
	flag.StringVar(&protocol, "P", "binary", "Specify the protocol (binary, compact, simplejson, json)")
	flag.StringVar(&urlString, "u", "", "Specify the url")
	flag.BoolVar(&framed, "framed", false, "Use framed transport")
	flag.BoolVar(&useHttp, "http", false, "Use http")
	flag.Parse()

	if len(urlString) > 0 {
		var err error
		parsedUrl, err = url.Parse(urlString)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
			flag.Usage()
		}
		host = parsedUrl.Host
		useHttp = len(parsedUrl.Scheme) <= 0 || parsedUrl.Scheme == "http"
	} else if useHttp {
		_, err := url.Parse(fmt.Sprint("http://", host, ":", port))
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
			flag.Usage()
		}
	}

	cmd := flag.Arg(0)
	var err error
	if useHttp {
		trans, err = thrift.NewTHttpClient(parsedUrl.String())
	} else {
		portStr := fmt.Sprint(port)
		if strings.Contains(host, ":") {
			host, portStr, err = net.SplitHostPort(host)
			if err != nil {
				fmt.Fprintln(os.Stderr, "error with host:", err)
				os.Exit(1)
			}
		}
		trans, err = thrift.NewTSocket(net.JoinHostPort(host, portStr))
		if err != nil {
			fmt.Fprintln(os.Stderr, "error resolving address:", err)
			os.Exit(1)
		}
		if framed {
			trans = thrift.NewTFramedTransport(trans)
		}
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating transport", err)
		os.Exit(1)
	}
	defer trans.Close()
	var protocolFactory thrift.TProtocolFactory
	switch protocol {
	case "compact":
		protocolFactory = thrift.NewTCompactProtocolFactory()
		break
	case "simplejson":
		protocolFactory = thrift.NewTSimpleJSONProtocolFactory()
		break
	case "json":
		protocolFactory = thrift.NewTJSONProtocolFactory()
		break
	case "binary", "":
		protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
		break
	default:
		fmt.Fprintln(os.Stderr, "Invalid protocol specified: ", protocol)
		Usage()
		os.Exit(1)
	}
	iprot := protocolFactory.GetProtocol(trans)
	oprot := protocolFactory.GetProtocol(trans)
	client := model.NewIRISHubServiceClient(thrift.NewTStandardClient(iprot, oprot))
	if err := trans.Open(); err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to ", host, ":", port, " ", err)
		os.Exit(1)
	}

	switch cmd {
	case "GetCandidateList":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetCandidateList requires 1 args")
			flag.Usage()
		}
		arg14 := flag.Arg(1)
		mbTrans15 := thrift.NewTMemoryBufferLen(len(arg14))
		defer mbTrans15.Close()
		_, err16 := mbTrans15.WriteString(arg14)
		if err16 != nil {
			Usage()
			return
		}
		factory17 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt18 := factory17.GetProtocol(mbTrans15)
		argvalue0 := model.NewCandidateListRequest()
		err19 := argvalue0.Read(jsProt18)
		if err19 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.GetCandidateList(context.Background(), value0))
		fmt.Print("\n")
		break
	case "GetCandidateDetail":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetCandidateDetail requires 1 args")
			flag.Usage()
		}
		arg20 := flag.Arg(1)
		mbTrans21 := thrift.NewTMemoryBufferLen(len(arg20))
		defer mbTrans21.Close()
		_, err22 := mbTrans21.WriteString(arg20)
		if err22 != nil {
			Usage()
			return
		}
		factory23 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt24 := factory23.GetProtocol(mbTrans21)
		argvalue0 := model.NewCandidateDetailRequest()
		err25 := argvalue0.Read(jsProt24)
		if err25 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.GetCandidateDetail(context.Background(), value0))
		fmt.Print("\n")
		break
	case "GetValidatorExRate":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetValidatorExRate requires 1 args")
			flag.Usage()
		}
		arg26 := flag.Arg(1)
		mbTrans27 := thrift.NewTMemoryBufferLen(len(arg26))
		defer mbTrans27.Close()
		_, err28 := mbTrans27.WriteString(arg26)
		if err28 != nil {
			Usage()
			return
		}
		factory29 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt30 := factory29.GetProtocol(mbTrans27)
		argvalue0 := model.NewValidatorExRateRequest()
		err31 := argvalue0.Read(jsProt30)
		if err31 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.GetValidatorExRate(context.Background(), value0))
		fmt.Print("\n")
		break
	case "GetDelegatorCandidateList":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetDelegatorCandidateList requires 1 args")
			flag.Usage()
		}
		arg32 := flag.Arg(1)
		mbTrans33 := thrift.NewTMemoryBufferLen(len(arg32))
		defer mbTrans33.Close()
		_, err34 := mbTrans33.WriteString(arg32)
		if err34 != nil {
			Usage()
			return
		}
		factory35 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt36 := factory35.GetProtocol(mbTrans33)
		argvalue0 := model.NewDelegatorCandidateListRequest()
		err37 := argvalue0.Read(jsProt36)
		if err37 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.GetDelegatorCandidateList(context.Background(), value0))
		fmt.Print("\n")
		break
	case "GetDelegatorTotalShares":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetDelegatorTotalShares requires 1 args")
			flag.Usage()
		}
		arg38 := flag.Arg(1)
		mbTrans39 := thrift.NewTMemoryBufferLen(len(arg38))
		defer mbTrans39.Close()
		_, err40 := mbTrans39.WriteString(arg38)
		if err40 != nil {
			Usage()
			return
		}
		factory41 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt42 := factory41.GetProtocol(mbTrans39)
		argvalue0 := model.NewTotalShareRequest()
		err43 := argvalue0.Read(jsProt42)
		if err43 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.GetDelegatorTotalShares(context.Background(), value0))
		fmt.Print("\n")
		break
	case "":
		Usage()
		break
	default:
		fmt.Fprintln(os.Stderr, "Invalid function ", cmd)
	}
}
package main

import (
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

func main() {
	//fmt.Println(getAddress(46.7803, 23.6044))
	//getDarkSkyData(46.7803, 23.6044)
	//getAccurateWeatherData(46.7803, 23.6044)
	http.HandleFunc("/mightWork/observation", handleObservationsRequest)
	http.HandleFunc("/mightWork/data", handleDataRequest)
	http.HandleFunc("/mightWork/user", handleUserLogIn)
	http.HandleFunc("/mightWork/register", handleUserRegister)
	http.HandleFunc("/mightWork/answersToObservations", handleAnswersToObservationRequest)
	http.HandleFunc("/mightWork/question", handleQuestionRequest)
	http.HandleFunc("/mightWork/answers", handleAnswerRequest)
	http.HandleFunc("/mightWork/alerts", handleAlertRequest)
	http.ListenAndServe(":8080", nil)
}
//func main() {
//	dev := flag.String("d", "COM6", "path to modem device")
//	baud := flag.Int("b", 115200, "baud rate")
//	num := flag.String("n", "+40773986102", "number to send to, in international format")
//	msg := flag.String("m", "GoLang is a Genius!", "the message to send")
//	timeout := flag.Duration("t", 5000*time.Millisecond, "command timeout period")
//	verbose := flag.Bool("v", false, "log modem interactions")
//	pdumode := flag.Bool("p", false, "send in PDU mode")
//	hex := flag.Bool("x", false, "hex dump modem responses")
//	flag.Parse()
//
//	m, err := serial.New(*dev, *baud)
//	if err != nil {
//		log.Fatal(err)
//	}
//	var mio io.ReadWriter = m
//	if *hex {
//		mio = trace.New(m, log.New(os.Stdout, "", log.LstdFlags), trace.ReadFormat("r: %v"))
//	} else if *verbose {
//		mio = trace.New(m, log.New(os.Stdout, "", log.LstdFlags))
//	}
//	g := gsm.New(mio)
//	ctx, cancel := context.WithTimeout(context.Background(), *timeout)
//	defer cancel()
//	if *pdumode {
//		g.SetPDUMode()
//	}
//	if err = g.Init(ctx); err != nil {
//		log.Fatal(err)
//	}
//	if !*pdumode {
//		mr, err := g.SendSMS(ctx, *num, *msg)
//		// !!! check CPIN?? on failure to determine root cause??  If ERROR 302
//		log.Printf("%v %v\n", mr, err)
//		return
//	}
//	sendPDU(ctx, g, *num, *msg)
//}
//
//func sendPDU(ctx context.Context, g *gsm.GSM, number string, msg string) {
//	ude, err := tpdu.NewUDEncoder()
//	if err != nil {
//		log.Fatal(err)
//	}
//	ude.AddAllCharsets()
//	me := message.NewEncoder(ude, sar.NewSegmenter())
//	pdus, err := me.Encode(number, msg)
//	if err != nil {
//		log.Fatal(err)
//	}
//	for i, p := range pdus {
//		tp, err := p.MarshalBinary()
//		if err != nil {
//			log.Fatal(err)
//		}
//		mr, err := g.SendSMSPDU(ctx, tp)
//		if err != nil {
//			// !!! check CPIN?? on failure to determine root cause??  If ERROR 302
//			log.Fatal(err)
//		}
//		log.Printf("PDU %d: %v\n", i+1, mr)
//	}
//}
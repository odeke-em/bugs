package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"
)

func main() {
	urls := []string{
		"https://www.legifrance.gouv.fr/affichCodeArticle.do;jsessionid=91E19D18FF040CFB3D743B58C5D21E12.tplgfr34s_2?idArticle=LEGIARTI000028743898&cidTexte=LEGITEXT000005634379&dateTexte=20190127&fastPos=1&fastReqId=1742940652&oldAction=rechCodeArticl",
		"https://www.legifrance.gouv.fr/affichCodeArticle.do;jsessionid=91E19D18FF040CFB3D743B58C5D21E12.tplgfr34s_2?idArticle=LEGIARTI000032955730&cidTexte=LEGITEXT000005634379&dateTexte=20190127&fastPos=2&fastReqId=1742940652&oldAction=rechCodeArticl",
		"https://www.legifrance.gouv.fr/affichCodeArticle.do;jsessionid=91E19D18FF040CFB3D743B58C5D21E12.tplgfr34s_2?idArticle=LEGIARTI000032132092&cidTexte=LEGITEXT000005634379&dateTexte=20190127&fastPos=3&fastReqId=1742940652&oldAction=rechCodeArticl",
		"https://www.legifrance.gouv.fr/affichCodeArticle.do;jsessionid=91E19D18FF040CFB3D743B58C5D21E12.tplgfr34s_2?idArticle=LEGIARTI000032944499&cidTexte=LEGITEXT000005634379&dateTexte=20190127&fastPos=4&fastReqId=1742940652&oldAction=rechCodeArticl",
		"https://www.legifrance.gouv.fr/affichCodeArticle.do;jsessionid=91E19D18FF040CFB3D743B58C5D21E12.tplgfr34s_2?idArticle=LEGIARTI000022496429&cidTexte=LEGITEXT000005634379&dateTexte=20190127&fastPos=5&fastReqId=1742940652&oldAction=rechCodeArticl",
		"https://www.legifrance.gouv.fr/affichCodeArticle.do;jsessionid=91E19D18FF040CFB3D743B58C5D21E12.tplgfr34s_2?idArticle=LEGIARTI000020196876&cidTexte=LEGITEXT000005634379&dateTexte=20190127&fastPos=6&fastReqId=1742940652&oldAction=rechCodeArticl",
		"https://www.legifrance.gouv.fr/affichCodeArticle.do;jsessionid=91E19D18FF040CFB3D743B58C5D21E12.tplgfr34s_2?idArticle=LEGIARTI000020196883&cidTexte=LEGITEXT000005634379&dateTexte=20190127&fastPos=7&fastReqId=1742940652&oldAction=rechCodeArticl",
		"https://www.legifrance.gouv.fr/affichCodeArticle.do;jsessionid=91E19D18FF040CFB3D743B58C5D21E12.tplgfr34s_2?idArticle=LEGIARTI000026454267&cidTexte=LEGITEXT000005634379&dateTexte=20190127&fastPos=8&fastReqId=1742940652&oldAction=rechCodeArticl",
		"https://www.legifrance.gouv.fr/affichCodeArticle.do;jsessionid=91E19D18FF040CFB3D743B58C5D21E12.tplgfr34s_2?idArticle=LEGIARTI000030985098&cidTexte=LEGITEXT000005634379&dateTexte=20190127&fastPos=9&fastReqId=1742940652&oldAction=rechCodeArticl",
		"https://www.legifrance.gouv.fr/affichCodeArticle.do;jsessionid=91E19D18FF040CFB3D743B58C5D21E12.tplgfr34s_2?idArticle=LEGIARTI000030985094&cidTexte=LEGITEXT000005634379&dateTexte=20190127&fastPos=10&fastReqId=1742940652&oldAction=rechCodeArticl",
		"https://www.legifrance.gouv.fr/affichCodeArticle.do;jsessionid=91E19D18FF040CFB3D743B58C5D21E12.tplgfr34s_2?idArticle=LEGIARTI000030044870&cidTexte=LEGITEXT000005634379&dateTexte=20190127&fastPos=11&fastReqId=1742940652&oldAction=rechCodeArticl",
		"https://www.legifrance.gouv.fr/affichCodeArticle.do;jsessionid=91E19D18FF040CFB3D743B58C5D21E12.tplgfr34s_2?idArticle=LEGIARTI000030985126&cidTexte=LEGITEXT000005634379&dateTexte=20190127&fastPos=12&fastReqId=1742940652&oldAction=rechCodeArticl",
		"https://www.legifrance.gouv.fr/affichCodeArticle.do;jsessionid=91E19D18FF040CFB3D743B58C5D21E12.tplgfr34s_2?idArticle=LEGIARTI000034182699&cidTexte=LEGITEXT000005634379&dateTexte=20190127&fastPos=13&fastReqId=1742940652&oldAction=rechCodeArticl",
		"https://www.legifrance.gouv.fr/affichCodeArticle.do;jsessionid=91E19D18FF040CFB3D743B58C5D21E12.tplgfr34s_2?idArticle=LEGIARTI000006241532&cidTexte=LEGITEXT000005634379&dateTexte=20190127&fastPos=14&fastReqId=1742940652&oldAction=rechCodeArticl",
		"https://www.legifrance.gouv.fr/affichCodeArticle.do;jsessionid=91E19D18FF040CFB3D743B58C5D21E12.tplgfr34s_2?idArticle=LEGIARTI000034161980&cidTexte=LEGITEXT000005634379&dateTexte=20190127&fastPos=15&fastReqId=1742940652&oldAction=rechCodeArticl",
		"https://www.legifrance.gouv.fr/affichCodeArticle.do;jsessionid=91E19D18FF040CFB3D743B58C5D21E12.tplgfr34s_2?idArticle=LEGIARTI000033461935&cidTexte=LEGITEXT000005634379&dateTexte=20190127&fastPos=16&fastReqId=1742940652&oldAction=rechCodeArticl",
		"https://www.legifrance.gouv.fr/affichCodeArticle.do;jsessionid=91E19D18FF040CFB3D743B58C5D21E12.tplgfr34s_2?idArticle=LEGIARTI000033739520&cidTexte=LEGITEXT000005634379&dateTexte=20190127&fastPos=17&fastReqId=1742940652&oldAction=rechCodeArticl",
		"https://www.legifrance.gouv.fr/affichCodeArticle.do;jsessionid=91E19D18FF040CFB3D743B58C5D21E12.tplgfr34s_2?idArticle=LEGIARTI000037266553&cidTexte=LEGITEXT000005634379&dateTexte=20190127&fastPos=18&fastReqId=1742940652&oldAction=rechCodeArticl",
		"https://www.legifrance.gouv.fr/affichCodeArticle.do;jsessionid=91E19D18FF040CFB3D743B58C5D21E12.tplgfr34s_2?idArticle=LEGIARTI000006219296&cidTexte=LEGITEXT000005634379&dateTexte=20190127&fastPos=19&fastReqId=1742940652&oldAction=rechCodeArticl",
		"https://www.legifrance.gouv.fr/affichCodeArticle.do;jsessionid=91E19D18FF040CFB3D743B58C5D21E12.tplgfr34s_2?idArticle=LEGIARTI000006219127&cidTexte=LEGITEXT000005634379&dateTexte=20190127&fastPos=20&fastReqId=1742940652&oldAction=rechCodeArticl",
		"https://www.legifrance.gouv.fr/affichCodeArticle.do;jsessionid=3B2A2AFFC85D4227A434B7DACCAD80F0.tplgfr28s_3?idArticle=LEGIARTI000006224130&cidTexte=LEGITEXT000005634379&dateTexte=20190127&fastPos=21&fastReqId=1338593856&oldAction=rechCodeArticl",
		"https://www.legifrance.gouv.fr/affichCodeArticle.do;jsessionid=3B2A2AFFC85D4227A434B7DACCAD80F0.tplgfr28s_3?idArticle=LEGIARTI000006226098&cidTexte=LEGITEXT000005634379&dateTexte=20190127&fastPos=22&fastReqId=1338593856&oldAction=rechCodeArticl",
		"https://www.legifrance.gouv.fr/affichCodeArticle.do;jsessionid=3B2A2AFFC85D4227A434B7DACCAD80F0.tplgfr28s_3?idArticle=LEGIARTI000006228726&cidTexte=LEGITEXT000005634379&dateTexte=20190127&fastPos=23&fastReqId=1338593856&oldAction=rechCodeArticl",
		"https://www.legifrance.gouv.fr/affichCodeArticle.do;jsessionid=3B2A2AFFC85D4227A434B7DACCAD80F0.tplgfr28s_3?idArticle=LEGIARTI000006241532&cidTexte=LEGITEXT000005634379&dateTexte=20190127&fastPos=24&fastReqId=1338593856&oldAction=rechCodeArticl",
		"https://www.legifrance.gouv.fr/affichCodeArticle.do;jsessionid=3B2A2AFFC85D4227A434B7DACCAD80F0.tplgfr28s_3?idArticle=LEGIARTI000006224359&cidTexte=LEGITEXT000005634379&dateTexte=20190127&fastPos=25&fastReqId=1338593856&oldAction=rechCodeArticl",
		"https://www.legifrance.gouv.fr/affichCodeArticle.do;jsessionid=3B2A2AFFC85D4227A434B7DACCAD80F0.tplgfr28s_3?idArticle=LEGIARTI000006222351&cidTexte=LEGITEXT000005634379&dateTexte=20190127&fastPos=26&fastReqId=1338593856&oldAction=rechCodeArticl",
		"https://www.legifrance.gouv.fr/affichCodeArticle.do;jsessionid=3B2A2AFFC85D4227A434B7DACCAD80F0.tplgfr28s_3?idArticle=LEGIARTI000006223239&cidTexte=LEGITEXT000005634379&dateTexte=20190127&fastPos=27&fastReqId=1338593856&oldAction=rechCodeArticl",
		"https://www.legifrance.gouv.fr/affichCodeArticle.do;jsessionid=3B2A2AFFC85D4227A434B7DACCAD80F0.tplgfr28s_3?idArticle=LEGIARTI000006262799&cidTexte=LEGITEXT000005634379&dateTexte=20190127&fastPos=28&fastReqId=1338593856&oldAction=rechCodeArticl",
		"https://www.legifrance.gouv.fr/affichCodeArticle.do;jsessionid=3B2A2AFFC85D4227A434B7DACCAD80F0.tplgfr28s_3?idArticle=LEGIARTI000006239978&cidTexte=LEGITEXT000005634379&dateTexte=20190127&fastPos=29&fastReqId=1338593856&oldAction=rechCodeArticl",
		"https://www.legifrance.gouv.fr/affichCodeArticle.do;jsessionid=3B2A2AFFC85D4227A434B7DACCAD80F0.tplgfr28s_3?idArticle=LEGIARTI000006234274&cidTexte=LEGITEXT000005634379&dateTexte=20190127&fastPos=30&fastReqId=1338593856&oldAction=rechCodeArticl",
		"https://www.legifrance.gouv.fr/affichCodeArticle.do;jsessionid=3B2A2AFFC85D4227A434B7DACCAD80F0.tplgfr28s_3?idArticle=LEGIARTI000006219167&cidTexte=LEGITEXT000005634379&dateTexte=20190127&fastPos=31&fastReqId=1338593856&oldAction=rechCodeArticl",
		"https://www.legifrance.gouv.fr/affichCodeArticle.do;jsessionid=3B2A2AFFC85D4227A434B7DACCAD80F0.tplgfr28s_3?idArticle=LEGIARTI000006219304&cidTexte=LEGITEXT000005634379&dateTexte=20190127&fastPos=32&fastReqId=1338593856&oldAction=rechCodeArticl",
		"https://www.legifrance.gouv.fr/affichCodeArticle.do;jsessionid=3B2A2AFFC85D4227A434B7DACCAD80F0.tplgfr28s_3?idArticle=LEGIARTI000006219821&cidTexte=LEGITEXT000005634379&dateTexte=20190127&fastPos=33&fastReqId=1338593856&oldAction=rechCodeArticl",
		"https://www.legifrance.gouv.fr/affichCodeArticle.do;jsessionid=3B2A2AFFC85D4227A434B7DACCAD80F0.tplgfr28s_3?idArticle=LEGIARTI000006219127&cidTexte=LEGITEXT000005634379&dateTexte=20190127&fastPos=34&fastReqId=1338593856&oldAction=rechCodeArticl",
		"https://www.legifrance.gouv.fr/affichCodeArticle.do;jsessionid=3B2A2AFFC85D4227A434B7DACCAD80F0.tplgfr28s_3?idArticle=LEGIARTI000006219296&cidTexte=LEGITEXT000005634379&dateTexte=20190127&fastPos=35&fastReqId=1338593856&oldAction=rechCodeArticl",
		"https://www.legifrance.gouv.fr/affichCodeArticle.do;jsessionid=3B2A2AFFC85D4227A434B7DACCAD80F0.tplgfr28s_3?idArticle=LEGIARTI000006219812&cidTexte=LEGITEXT000005634379&dateTexte=20190127&fastPos=36&fastReqId=1338593856&oldAction=rechCodeArticl",
		"https://www.legifrance.gouv.fr/affichCodeArticle.do;jsessionid=3B2A2AFFC85D4227A434B7DACCAD80F0.tplgfr28s_3?idArticle=LEGIARTI000006259055&cidTexte=LEGITEXT000005634379&dateTexte=20190127&fastPos=37&fastReqId=1338593856&oldAction=rechCodeArticl",
		"https://www.legifrance.gouv.fr/affichCodeArticle.do;jsessionid=3B2A2AFFC85D4227A434B7DACCAD80F0.tplgfr28s_3?idArticle=LEGIARTI000006219609&cidTexte=LEGITEXT000005634379&dateTexte=20190127&fastPos=38&fastReqId=1338593856&oldAction=rechCodeArticl",
		"https://www.legifrance.gouv.fr/affichCodeArticle.do;jsessionid=3B2A2AFFC85D4227A434B7DACCAD80F0.tplgfr28s_3?idArticle=LEGIARTI000006222461&cidTexte=LEGITEXT000005634379&dateTexte=20190127&fastPos=39&fastReqId=1338593856&oldAction=rechCodeArticl",
		"https://www.legifrance.gouv.fr/affichCodeArticle.do;jsessionid=3B2A2AFFC85D4227A434B7DACCAD80F0.tplgfr28s_3?idArticle=LEGIARTI000006222473&cidTexte=LEGITEXT000005634379&dateTexte=20190127&fastPos=40&fastReqId=1338593856&oldAction=rechCodeArticl",
	}

	checkDuplicates(processPayloads(urls))
}

func processPayloads(urls []string) (ch <-chan *payload) {
	payloadsCh := make(chan *payload)

	go func() {
		var wg sync.WaitGroup
		defer func() {
			wg.Wait()
			close(payloadsCh)
		}()

		for _, link := range urls {
			wg.Add(1)

			go func(link string) {
				pIn := &payload{
					url: link,
				}
				defer func() {
					payloadsCh <- pIn
					wg.Done()
				}()

				DefaultClient := &http.Client{
					Timeout: time.Duration(5) * time.Second,
					Transport: &http.Transport{
						DisableKeepAlives: true,
						DialContext: (&net.Dialer{
							Timeout:   3 * time.Second,
							KeepAlive: 0,
							DualStack: true,
						}).DialContext,
					},
				}

				res, err := DefaultClient.Get(link)
				if err != nil {
					pIn.err = err
					return
				}
				defer res.Body.Close()

				cksum := md5.New()
				tr := io.TeeReader(res.Body, cksum)
				pIn.body, pIn.err = ioutil.ReadAll(tr)
				pIn.checksum = fmt.Sprintf("%x", cksum.Sum(nil))
			}(link)
		}
	}()

	return payloadsCh
}

type payload struct {
	url      string
	body     []byte
	checksum string
	err      error
}

func checkDuplicates(payloadsIn <-chan *payload) {
	groupingsByChecksum := make(map[string][]string)
	for pIn := range payloadsIn {
		key := pIn.checksum
		value := pIn.url
		groupingsByChecksum[key] = append(groupingsByChecksum[key], value)
	}

	// Finally provide the summary.
	var clashesDetected int
	for checksum, urls := range groupingsByChecksum {
		if len(urls) > 1 {
			fmt.Printf("Clashes detected at checksum %s: with %d URLs as below\n%s\n", checksum,
				strings.Join(append([]string{""}, urls...), "\t\n"))
			clashesDetected++
		}
	}

	if clashesDetected == 0 {
		fmt.Println("No clashes detected!\n")
	}
}

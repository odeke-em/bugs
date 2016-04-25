package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		log.Printf("Received post request whose Headers are %v", r.Header)
		if false {
			io.Copy(ioutil.Discard, r.Body)
		}
		err := r.Body.Close()
		log.Printf("req.Body.Close() here, err=%v\n", err)
		if false {
			time.Sleep(700 * time.Millisecond)
		}
		//http.Error(w, error, code)
		w.WriteHeader(413)

		b, err := io.Copy(w, strings.NewReader(strings.Repeat("XXXXXX", 4192)))
		log.Printf("Wrote %d bytes. Error: %v", b, err)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(`
    <body>
    <script type="text/javascript" src="https://cdnjs.cloudflare.com/ajax/libs/jquery/3.0.0-beta1/jquery.min.js"></script>
    <script>
      var reqs = 1;
      $('body').text('Sending ' + reqs + ' requests...');
      function send(i) {
  var el = $('<div class="req">Waiting...</div>').attr('id', i).appendTo('body');
  var data = new FormData();
  data.append('file', new File([new Blob([new Array(5 * 1024 * 1024).join('0')], {type: 'image/jpeg'})], 'test.jpg'));
  var done = false;
  var req = new XMLHttpRequest();
  req.upload.onprogress = function(e) {
    if (!e.lengthComputable || done) return;
    el.text((e.loaded / e.total * 100) + '%');
  }
  req.onreadystatechange = function(e) {
    if (done || req.readyState !== 3) return;
    done = true;
    el.text('Done. Got ' + req.status);
  }
  req.open('POST', '/', true);
  req.send(data);
      }
      for (var i = 0; i < reqs; i++)
  send(i);
    </script>
    </body>
  `))
}

func main() {
	http.HandleFunc("/", handler)
	if err := http.ListenAndServeTLS(":1333", "key.crt", "key.key", nil); err != nil {
		log.Fatal(err)
	}
}

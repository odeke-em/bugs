const http = require('http');
const util = require('util');

var options = {
  hostname: 'https://api.nexthot.net',
  port: 7001,
  path: '/?num=100',
  method: 'GET',
  headers: {
    // set connectiont to close otherwise
    // node.js trips out with `HPE_INVALID_CONSTANT`
    // if by default `connection: keep-alive`
    connection: 'close',
  },
};

var req = http.request(options, (res, err) => {
  console.log(`STATUS: ${res.statusCode}`);
  console.log(`HEADERS: ${JSON.stringify(res.headers)}`);
  res.setEncoding('utf8');
  res.on('data', (chunk) => {
    console.log(`BODY: ${chunk}`);
  });
  res.on('end', () => {
    console.log('No more data in response.');
  });
});

req.on('error', (e) => {
  console.log(`problem with request:`, util.inspect(e, {depth:10}));
});

req.end();

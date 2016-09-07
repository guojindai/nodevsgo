"use strict";

const http = require('http');

const hostname = 'localhost';
const port = 35033;

const server = http.createServer((req, res) => {
  let sum = 0;
  res.statusCode = 200;
  for (let i = 0; i < 10000; i ++) {
    sum += i;
  }
  res.end(sum + '');
});

server.listen(port, hostname, () => {
  console.log(`Server running at http://${hostname}:${port}/`);
});
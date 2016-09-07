"use strict";

const http = require('http'),
      fs = require('fs');

const hostname = 'localhost';
const port = 35033;

const server = http.createServer((req, res) => {
  res.statusCode = 200;
  fs.readFile('../data/resnd', {encoding: 'utf8'}, (err, text) => {
    if (err) {
      res.end(err.toString());
    } else {
      res.end(text);
    }
  });  
});

server.listen(port, hostname, () => {
  console.log(`Server running at http://${hostname}:${port}/`);
});
const whois = require('whois');
const dig = require('node-dig-dns');
const express = require('express');
const app = express();
const port = process.env.PORT || 8000
const usagewhois =  'whois: GET /whois/[URL or IP]';
const usagedig = 'dig: GET /dig/[URL]/[ANY, A, MX, NS, TXT...]';

app.get('/', (req, res) => {
  res.send(usagewhois + '<br>' + usagedig);
});

app.get('/whois', (req, res) => {
  res.send(usagewhois);
});

app.get('/whois/:url', (req, res) => {
  console.log('whois: ' + req.params.url);
  try {
    whois.lookup(req.params.url, function(err, data) {
      res.setHeader('Content-Type', 'text/plain');
      res.send(data);
    });
  } catch(e) {
	  res.status(500).send('Error');
  }
});

app.get('/dig', (req, res) => {
  res.send(usagedig);
});

app.get('/dig/:url/:type', (req, res) => {
  console.log('dig: ' + req.params.url + ' ' + req.params.type);
  dig(['@1.1.1.1', req.params.url, req.params.type])
    .then((result) => {
      result2 = JSON.stringify(result, null, 1);
      res.setHeader('Content-Type', 'text/plain');
      res.send(result2);
    })
    .catch((err) => {
      res.status(500).send('Error');
    });
  });

app.listen(port, () => {
  console.log('Server listening on port ' + port)
});

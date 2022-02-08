const whois = require('whois');
const dig = require('node-dig-dns');
const express = require('express');
const app = express();
const port = process.env.PORT || 8000
const usagewhois =  'whois: GET /whois/[URL or IP]';
const usagedig = 'dig: GET /dig/[URL]/[ANY, A, MX, NS, TXT...]';

app.get('/', (req, res) => {
  res.send('whois: GET /whois?domain=[DOMAIN or IP]<br>dig: GET /dig?domain=[DOMAIN]&type=[TYPE]<br>client: GET /client');
});

app.get('/client', (req, res) => {
  res.sendFile(__dirname + '/public/client.html');
});

app.get('/whois', (req, res) => {
  if (!req.query.domain) {
    res.send('Usage: GET /whois?domain=[DOMAIN or IP]');
  } else {
    console.log('whois: ' + req.query.domain);
    try {
      whois.lookup(req.query.domain, function(err, data) {
        res.setHeader('Content-Type', 'text/plain');
        res.send(data);
      });
    } catch(e) {
	    res.status(500).send('Error');
    };
  };
});

app.get('/dig', (req, res) => {
  if (!req.query.domain) {
    res.send('Usage: GET /dig?domain=[DOMAIN]&type=[ANY, A, MX, NS, TXT...]');
  } else {
    let rectype;
    if(!req.query.type) {
      type = 'any';
    } else {
      type = req.query.type;
    };
    console.log('dig: ' + req.query.domain + ' ' + type);
    dig(['@1.1.1.1', req.query.domain, type])
    .then((result) => {
      const ans = JSON.stringify(result, null, 1);
      res.setHeader('Content-Type', 'text/plain');
      res.send(ans);
    })
    .catch((err) => {
      res.status(500).send('Error');
    });
  };
});

app.listen(port, () => {
  console.log('Server listening on port ' + port);
});

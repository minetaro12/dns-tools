const whois = require('whois');
const dig = require('node-dig-dns');
const express = require('express');
const app = express();
const port = process.env.PORT || 8000
const usagewhois =  'whois: GET /whois?domain=[DOMAIN or IP]';
const usagedig = 'dig: GET /dig?domain=[DOMAIN]&type=[TYPE]';

app.use(express.static('public'));

app.get('/', (req, res) => {
  res.sendFile(__dirname + '/public/client.html');
});

app.get('/whois', (req, res) => {
  if (!req.query.domain) {
    res.status(400).send(usagewhois);
  } else {
    const domain = req.query.domain;
    console.log(`whois: ${domain}`);
    try {
      whois.lookup(domain, function(err, data) {
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
    res.status(400).send(usagedig);
  } else {
    const domain = req.query.domain;
    let type;
    if(!req.query.type) {
      type = 'any';
    } else {
      type = req.query.type;
    };
    console.log(`dig: ${domain} ${type}`);
    dig(['@1.1.1.1', domain, type])
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

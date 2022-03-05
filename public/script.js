const whoisDomain = document.getElementById('whoisDomain');
const digDomain = document.getElementById('digDomain');
const digType = document.getElementById('digType');

const whoisResult = document.getElementById('whoisResult');
const digResult =document.getElementById('digResult');

const whoisSubmit = function() {
  if (!whoisDomain.value) {
    whoisResult.innerHTML = 'ドメインを入力してください';
    return;
  };
  const domain = whoisDomain.value;

  whoisResult.innerHTML = '<p>Please Wait...<img src="https://www.benricho.org/loading_images/img-transparent/712-24.gif"></p>'

  fetch(`/whois?domain=${domain}`)
    .then((res) => {
      if (!res.ok) {
        whoisResult.innerHTML = '<p>Error</p>';
        return;
      };
      return res.text();
    })
    .then((text) => {
      whoisResult.innerHTML = `<pre>${text}</pre>`;
      return
    })
    .catch((error) => {
      whoisResult.innerHTML = '<p>Error</p>';
      return;
    });
};

const digSubmit = function() {
  if (!digDomain.value) {
    digResult.innerHTML = 'ドメインを入力してください';
    return;
  };
  const domain = digDomain.value;
  let typeQuery;

  if (!digType.value) {
    typeQuery = '&type=any';
  } else {
    typeQuery = `&type=${digType.value}`;
  };

  fetch(`/dig?domain=${domain}${typeQuery}`)
    .then((res) => {
      if (!res.ok) {
        digResult.innerHTML = '<p>Error</p>';
        return;
      };
      return res.text();
    })
    .then((text) => {
      digResult.innerHTML = `<pre>${text}</pre>`;
      return;
    })
    .catch((error) => {
      digResult.innerHTML = '<p>Error</p>';
      return;
    });
};
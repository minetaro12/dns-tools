const whoisDomain = document.getElementById('whoisDomain');
const digDomain = document.getElementById('digDomain');
const digType = document.getElementById('digType');

const whoisResult = document.getElementById('whoisResult');
const digResult =document.getElementById('digResult');

const waitHtml = '<p>Please Wait...<span class="spinner-border text-primary" role="status"></span></p>';

const whoisSubmit = function() {
  if (!whoisDomain.value) {
    whoisResult.innerHTML = 'ドメインを入力してください';
    return;
  };
  
  const domain = whoisDomain.value;

  whoisResult.innerHTML = waitHtml;

  fetch(`/whois?domain=${domain}`)
    .then((res) => {
      if (!res.ok) {
        throw new Error;
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

  digResult.innerHTML = waitHtml;

  fetch(`/dig?domain=${domain}${typeQuery}`)
    .then((res) => {
      if (!res.ok) {
        throw new Error;
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

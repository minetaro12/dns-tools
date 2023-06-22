const wDomain = document.querySelector("#wDomain")
const qFqdn = document.querySelector("#qFqdn")
const qDns = document.querySelector("#qDns")
const qType = document.querySelector("#qType")

const resultBox = document.querySelector("#resultBox")
const result = document.querySelector("#result")
const errDiag = document.querySelector("#errDiag")

function submit(m) {
  errDiag.classList.add("hide")
  resultBox.classList.remove("hide")
  result.textContent = "Please wait..."
  const fd = new FormData()
  if(m == "whois") {
    fd.append("domain", wDomain.value)
  } else {
    fd.append("fqdn", qFqdn.value)
    fd.append("dns", qDns.value)
    fd.append("type", qType.value)
  }

  fetch("/" + m, {
    method: "POST",
    body: fd
  })
  .then((res) => {
    if(!res.ok) {
      throw new Error
    }
    return res.text()
  })
  .then((text) => {
    result.textContent = text
    return
  })
  .catch((error) => {
    resultBox.classList.add("hide")
    errDiag.classList.remove("hide")
    return
  })
}
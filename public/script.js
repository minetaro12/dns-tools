const { createApp } = Vue

createApp({
  data() {
    return {
      result: "",
      whoisDomain: "",
      qDns: "",
      qDomain: "",
      qType: "A"
    }
  },
  methods: {
    submit(m) {
      this.result = "Please wait..."
      const fd = new FormData()
      if(m == "whois") {
        fd.append("domain", this.whoisDomain)
      } else {
        fd.append("domain", this.qDomain)
        fd.append("dns", this.qDns)
        fd.append("type", this.qType)
      }
      fetch("/" + m + "/", {
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
        this.result = text
        return
      })
      .catch((error) =>{
        this.result = "Error"
        return
      })
    }
  }
}).mount('#app')
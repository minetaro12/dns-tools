const { createApp } = Vue

createApp({
  data() {
    return {
      result: "",
      whoisDomain: "",
      digDns: "",
      digDomain: "",
      digType: "A"
    }
  },
  methods: {
    whoisSubmit() {
      this.result = "Please wait..."
      const fd = new FormData()
      fd.append("domain", this.whoisDomain)
      fetch("/whois/", {
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
    },
    digSubmit() {
      this.result = "Please wait..."
      const fd = new FormData()
      fd.append("dns", this.digDns)
      fd.append("domain", this.digDomain)
      fd.append("type", this.digType)
      fetch("/dig/", {
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
      .catch((error) => {
        this.result = "Error"
        return
      })
    },
    nslookupSubmit() {
      this.result = "Please wait..."
      const fd = new FormData()
      fd.append("dns", this.digDns)
      fd.append("domain", this.digDomain)
      fd.append("type", this.digType)
      fetch("/nslookup/", {
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
      .catch((error) => {
        this.result = "Error"
        return
      })
    }
  }
}).mount('#app')
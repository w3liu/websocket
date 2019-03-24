### golang websocket demo

#### vue client

```
<template>
<div>
    {{msg}}
</div>
</template>
<script>
export default {
  data () {
    return {
      websock: null,
      msg: ''
    }
  },
  methods: {
    init: function () {
      const wsurl = 'ws://127.0.0.1:88/ws'
      this.websock = new WebSocket(wsurl)
      this.websock.onmessage = this.onmessage
      this.websock.onopen = this.onopen
      this.websock.onerror = this.onerror
      this.websock.onclose = this.onclose
    },
    onopen: function () {
      this.send('{"userid":1, "name":"zhang san", "age":"30"}')
    },
    send: function (data) {
      for (var i = 0; i < 10; i++) {
        this.websock.send(data)
      }
    },
    onclose: function (e) {
      console.log('ws close', e)
    },
    onmessage: function (e) {
      let _this = this
      console.log(e.data)
      _this.msg = e.data
    },
    onerror: function () {
      console.log('ws error')
      this.init()
    }
  },
  mounted: function () {
    this.init()
  },
  watch: {
  }
}
</script>
```
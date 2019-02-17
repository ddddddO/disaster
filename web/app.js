const vm = new Vue({
    el: '#app',
    data: {
      radigo_health: []
    },
    mounted() {
      axios.get("http://localhost:8888/health")
      .then(res => {this.radigo_health = res.data})
    }
  });
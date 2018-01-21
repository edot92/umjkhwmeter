<template>
  <div>
    <!-- <v-layout wrap row> -->
    <div id="container2" style="margin-top:0"></div>
    <!-- </v-layout> -->
    <v-layout wrap row style="background: white ;">
      <v-flex xs6 style=" border:1px solid gray">
        <H6>
          Arus
        </H6>
        <typegauge :dataChannel="gaugeArus" :updateVal="gaugeArusVal" :isReadyDetailChannel="isReady">

        </typegauge>
      </v-flex>
      <v-flex xs6 style=" border:1px solid gray">
        <H6>
          Cospi
        </H6>
        <typegauge :dataChannel="gaugeCospi" :updateVal="gaugeCospiVal" :isReadyDetailChannel="isReady">

        </typegauge>
      </v-flex>
    </v-layout>
    <v-layout wrap row>

      <v-flex xs6 style=" border:1px solid gray">
        <H6>
          Tegangan
        </H6>
        <typegauge :dataChannel="gaugeTegangan" :updateVal="gaugeTeganganVal" :isReadyDetailChannel="isReady">

        </typegauge>
      </v-flex>
      <v-flex xs6 style=" border:1px solid gray">
        <H6>
          Frekuensi
        </H6>
        <typegauge :dataChannel="gaugeFrekuensi" :updateVal="gaugeFrekuensiVal" :isReadyDetailChannel="isReady">
        </typegauge>
      </v-flex>
    </v-layout>

  </div>
</template>
<script>
  import typegauge from './typegauge.vue';
  export default {
    components: {
      typegauge
    },
    data() {
      return {
        gaugeArus: {},
        gaugeCospi: {},
        gaugeTegangan: {},
        gaugeFrekuensi: {},
        gaugeArusVal: '0.0',
        gaugeCospiVal: '0.0',
        gaugeTeganganVal: '0.0',
        gaugeFrekuensiVal: '0.0',
        isReady: false,
        timer: null
      };
    },
    destroyed() {
      this.timer = null
    },
    mounted() {
      // this.ChartInit2();
      this.gaugeArus = {
        containerID: 'gauge1',
        data: [
          {
            value: 0.0,
            waktu: Date.now()
          }
        ],
        rangeMin: 0.0,
        rangeMax: 5.0,
        unit: 'Ampere'
      };
      this.gaugeCospi = {
        containerID: 'gauge2',
        data: [
          {
            value: 0.0,
            waktu: Date.now()
          }
        ],
        rangeMin: 0.0,
        rangeMax: 1.2,
        unit: 'CosPi'
      };
      this.gaugeTegangan = {
        containerID: 'gauge3',
        data: [
          {
            value: 0.7,
            waktu: Date.now()
          }
        ],
        rangeMin: 0.0,
        rangeMax: 300.0,
        unit: 'Volt'
      };
      this.gaugeFrekuensi = {
        containerID: 'gauge4',
        data: [
          {
            value: 0.95,
            waktu: Date.now()
          }
        ],
        rangeMin: 0.0,
        rangeMax: 70.0,
        unit: 'Hz'
      };
      this.isReady = true;
      const thisV = this
      // setTimeout(() => {
      // window.socketIO.on('chat datapm', function (message) {
      //   try {
      //     var datake = JSON.parse(message);
      //     thisV.gaugeArusVal = datake.arus
      //     // console.log(thisV.gaugeArusVal);
      //   } catch (error) {
      //     // console.log(message);
      //   }
      // });

      // }, 2000);

      this.timer = setInterval(() => {
        thisV.tesInterval()
      }, 3000)
    },
    methods: {
      tesInterval() {
        const thisV = this
        this.$http(
          {
            url: '/alat/getnilaipm',
            method: 'get'
          }
        ).then(res => {
          try {
            const datak = res.data.payload.data
            console.log(datak
            )
            thisV.gaugeArusVal = datak.pmArus
            thisV.gaugeCospiVal = datak.pmCospi
            thisV.gaugeTeganganVal = datak.pmTegangan
            thisV.gaugeFrekuensiVal = datak.pmFrekuensi
          } catch (error) {

          }
        }).catch(err => {
          console.error(err)
        })
      },
      socketIO() {
        window.socketIO = window.io('http://localhost:8080');
        window.socketIO.on('connect', function () {
          console.log('connect');
          console.log(window.socketIO.id);
        });

        window.socketIO.on('disconnect', function () {
          console.log('diskonek');
        });
      }
    }
  };
</script>


<style>
  h6 {
    align-content: 'center';
  }
</style>

<template>
  <div>
    {{waktu}}
    <hr>
    <!-- <v-layout wrap row> -->
    <div id="conatinerMikro" style="margin-top:0"></div>
    <!-- </v-layout> -->
    <v-layout wrap row>
      <v-flex xs4>
        <H6>
          Beban 1
        </H6>
        <typegauge :dataChannel="gaugeDataMikro1" :updateVal="beban1" :isReadyDetailChannel="isReady">

        </typegauge>
      </v-flex>
      <v-flex xs4>
        <H6>
          Beban 2
        </H6>
        <typegauge :dataChannel="gaugeDataMikro2" :updateVal="beban2" :isReadyDetailChannel="isReady">

        </typegauge>
      </v-flex>
      <v-flex xs4>
        <H6>
          Beban 3
        </H6>
        <typegauge :dataChannel="gaugeDataMikro3" :updateVal="beban3" :isReadyDetailChannel="isReady">

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
        gaugeDataMikro1: {},
        gaugeDataMikro2: {},
        gaugeDataMikro3: {},
        isReady: false,
        beban1: 0,
        beban2: 0,
        beban3: 0,
        mikroTegangan: 0,
        mikroArus1: 0,
        mikroArus2: 0,
        mikroArus3: 0,
        waktu: ''
      };
    },
    mounted() {
      this.tesInterval()
      this.ChartInit2();
      this.gaugeDataMikro1 = {
        containerID: 'gauge1',
        data: [
          {
            value: 0.0,
            waktu: Date.now()
          }
        ],
        rangeMin: 0.0,
        rangeMax: 1100.0,
        unit: 'Watt'
      };
      this.gaugeDataMikro2 = {
        containerID: 'gauge2',
        data: [
          {
            value: 0.0,
            waktu: Date.now()
          }
        ],
        rangeMin: 0.0,
        rangeMax: 1100.0,
        unit: 'Watt'
      };
      this.gaugeDataMikro3 = {
        containerID: 'gauge3',
        data: [
          {
            value: 0.7,
            waktu: Date.now()
          }
        ],
        rangeMin: 0.0,
        rangeMax: 1100.0
      };
      this.isReady = true;
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
            thisV.waktu = datak.Waktu
            thisV.mikroTegangan = datak.mikroTegangan
            thisV.mikroArus1 = datak.mikroArus1
            thisV.mikroArus2 = datak.mikroArus2
            thisV.mikroArus3 = datak.mikroArus3
            thisV.beban1 = (parseFloat(datak.mikroArus1) * parseFloat(datak.mikroTegangan) * parseFloat(datak.PMCospi))
            thisV.beban2 = (parseFloat(datak.mikroArus2) * parseFloat(datak.mikroTegangan) * parseFloat(datak.PMCospi))
            thisV.beban3 = (parseFloat(datak.mikroArus3) * parseFloat(datak.mikroTegangan) * parseFloat(datak.PMCospi))
          } catch (error) {

          }
        }).catch(err => {
          console.error(err)
        })
      },
      ChartInit2() {
        $(document).ready(function () {
          window.Highcharts.setOptions({
            global: {
              useUTC: false
            }
          });
          window.Highcharts.chart('conatinerMikro', {
            chart: {
              type: 'spline',
              animation: window.Highcharts.svg, // don't animate in old IE
              marginRight: 10,
              events: {
                load: function () {
                  // set up the updating of the chart each second
                  // var series = this.series[0];
                  // var series2 = this.series[1];
                  // var series3 = this.series[2];
                  // setInterval(function () {
                  //   var x = new Date().getTime(); // current time
                  //   var y = Math.random();
                  //   series.addPoint([x, y], true, true);
                  //   x = new Date().getTime(); // current time
                  //   y = Math.random();
                  //   series2.addPoint([x, y], true, true);
                  //   x = new Date().getTime(); // current time
                  //   y = Math.random();
                  //   series3.addPoint([x, y], true, true);
                  // }, 1000);
                }
              }
            },
            title: {
              text: 'Monitoring Daya Aktif'
            },
            xAxis: {
              type: 'datetime',
              tickPixelInterval: 150
            },
            yAxis: {
              title: {
                text: 'Kw'
              },
              plotLines: [
                {
                  value: 0,
                  width: 1,
                  color: '#808080'
                }
              ]
            },
            tooltip: {
              formatter: function () {
                return (
                  '<b>' +
                  this.series.name +
                  '</b><br/>' +
                  window.Highcharts.dateFormat('%Y-%m-%d %H:%M:%S', this.x) +
                  '<br/>' +
                  window.Highcharts.numberFormat(this.y, 2)
                );
              }
            },
            legend: {
              enabled: true
            },
            exporting: {
              enabled: true
            },
            series: [
              {
                name: 'Beban 1',
                data: (function () {
                  // generate an array of random data
                  var data = [];
                  var time = new Date().getTime();
                  var i;

                  for (i = -19; i <= 0; i += 1) {
                    data.push({
                      x: time + i * 1000,
                      y: Math.random()
                    });
                  }
                  return data;
                })()
              },
              {
                name: 'Beban 2',
                data: (function () {
                  // generate an array of random data
                  var data = [];
                  var time = new Date().getTime();
                  var i;

                  for (i = -19; i <= 0; i += 1) {
                    data.push({
                      x: time + i * 1000,
                      y: Math.random()
                    });
                  }
                  return data;
                })()
              },
              {
                name: 'Beban 3',
                data: (function () {
                  // generate an array of random data
                  var data = [];
                  var time = new Date().getTime();
                  var i;

                  for (i = -19; i <= 0; i += 1) {
                    data.push({
                      x: time + i * 1000,
                      y: Math.random()
                    });
                  }
                  return data;
                })()
              }
            ]
          });
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

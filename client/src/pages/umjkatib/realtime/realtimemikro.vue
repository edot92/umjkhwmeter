<template>
  <div>
    <br>
    <v-layout row wrap style="background: #b3d4fc;">
      <v-flex xs12 style=" border:1px solid gray">
        <h5 class="headline" style="  text-align: center;">
          Waktu: {{waktu}}
        </h5>
      </v-flex>
    </v-layout>

    <v-layout wrap row>
      <v-flex xs12 sm4 style=" border:1px solid gray">
        <h5 class="headline" style="  text-align: center;">
          Beban 1
        </H5>
        <typegauge :dataChannel="gaugeDataMikro1" :updateVal="beban1" :isReadyDetailChannel="isReady">

        </typegauge>
      </v-flex>
      <v-flex xs12 sm4 style=" border:1px solid gray">
        <h5 class="headline" style="  text-align: center;">
          Beban 2
        </H5>
        <typegauge :dataChannel="gaugeDataMikro2" :updateVal="beban2" :isReadyDetailChannel="isReady">

        </typegauge>
      </v-flex>
      <v-flex xs12 sm4 style=" border:1px solid gray">
        <h5 class="headline" style="  text-align: center;">
          Beban 3
        </H5>
        <typegauge :dataChannel="gaugeDataMikro3" :updateVal="beban3" :isReadyDetailChannel="isReady">

        </typegauge>
      </v-flex>
      <!-- <v-flex xs4>
            <h5 class="headline" style="  text-align: center;">
          Tegangan
        </H5>
        <typegauge :dataChannel="gaugeDataMikroTegangan" :updateVal="tegangan" :isReadyDetailChannel="isReady">

        </typegauge>
      </v-flex> -->
    </v-layout>
    <v-layout wrap row>
      <v-flex xs12 style=" border:1px solid gray">
        <div id="conatinerMikro" style="margin-top:0"></div>
      </v-flex>
    </v-layout>
  </div>
</template>
<script>
import typegauge from './typegauge.vue';
var chartku;
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
      waktu: '',
      gaugeDataMikroTegangan: 0,
      tegangan: 0
    };
  },
  mounted() {
    // this.ChartInit2();
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
          value: 0.0,
          waktu: Date.now()
        }
      ],
      rangeMin: 0.0,
      rangeMax: 1100.0,
      unit: 'Watt'
    };

    this.isReady = true;
    const thisV = this;
    this.timer = setInterval(() => {
      thisV.tesInterval();
    }, 4000);
    this.ChartInit2();
  },
  methods: {
    tesInterval() {
      const thisV = this;
      this.$http({
        url: '/alat/getnilaipm',
        method: 'get'
      })
        .then(res => {
          try {
            const datak = res.data.payload.data;
            thisV.waktu = window
              .moment(datak.waktu)
              .format('HH:mm:ss DD-MM-YYYY');
            thisV.beban1 = (
              parseFloat(datak.mikroArus1) *
              parseFloat(datak.mikroTegangan) *
              parseFloat(datak.pmCospi)
            ).toPrecision(4);
            thisV.beban2 = (
              parseFloat(datak.mikroArus2) *
              parseFloat(datak.mikroTegangan) *
              parseFloat(datak.pmCospi)
            ).toPrecision(4);
            thisV.beban3 = (
              parseFloat(datak.mikroArus3) *
              parseFloat(datak.mikroTegangan) *
              parseFloat(datak.pmCospi)
            ).toPrecision(4);
            var dayaAktif = (
              parseFloat(datak.pmArus) *
              parseFloat(datak.pmCospi) *
              parseFloat(datak.pmTegangan)
            ).toPrecision(4);
            var resultMax = Math.max(
              parseFloat(thisV.beban1),
              parseFloat(thisV.beban2),
              parseFloat(thisV.beban3)
            );
            var valArr = [
              parseFloat(thisV.beban1),
              parseFloat(thisV.beban2),
              parseFloat(thisV.beban3)
            ];
            var indexArr = valArr.findIndex(function(age) {
              return age === resultMax;
            });
            console.log(indexArr);
            if (indexArr === 0) {
              console.log('beban 1 aktif');
              thisV.beban1 = dayaAktif;
            }
            if (indexArr === 1) {
              console.log('beban 2 aktif');
              thisV.beban2 = dayaAktif;
            }
            if (indexArr === 2) {
              console.log('beban 3 aktif');
              thisV.beban3 = dayaAktif;
            }
            thisV.updatedChart(thisV.beban1, thisV.beban2, thisV.beban3);
          } catch (err) {
            console.error(err);
          }
        })
        .catch(err => {
          console.error(err);
        });
    },
    updatedChart(y1, y2, y3) {
      var series = chartku.series[0];
      var series2 = chartku.series[1];
      var series3 = chartku.series[2];
      var x = new Date().getTime(); // current time
      // var y = Math.random();
      // console.log(y1);
      series.addPoint([x, parseFloat(y1)], true, true);
      x = new Date().getTime(); // current time
      // y = Math.random();
      series2.addPoint([x, parseFloat(y2)], true, true);
      x = new Date().getTime(); // current time
      // y = Math.random();
      series3.addPoint([x, parseFloat(y3)], true, true);
      chartku.redraw();
    },
    updateChart() {},
    ChartInit2() {
      $(document).ready(function() {
        window.Highcharts.setOptions({
          global: {
            useUTC: false
          }
        });
        chartku = window.Highcharts.chart('conatinerMikro', {
          chart: {
            type: 'spline',
            animation: window.Highcharts.svg, // don't animate in old IE
            marginRight: 10,
            events: {
              load: function() {
                // set up the updating of the chart each second
                // var series = this.series[0];
                // var series2 = this.series[1];
                // var series3 = this.series[2];
                // $(document).on('updateChartMikro', function(e, arg1, arg2) {
                //   var x = new Date().getTime(); // current time
                //   var y = Math.random();
                //   series.addPoint([x, y], true, true);
                //   x = new Date().getTime(); // current time
                //   y = Math.random();
                //   series2.addPoint([x, y], true, true);
                //   x = new Date().getTime(); // current time
                //   y = Math.random();
                //   series3.addPoint([x, y], true, true);
                // });
                // setInterval(function () {
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
              text: 'Watt'
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
            formatter: function() {
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
              data: (function() {
                // generate an array of random data
                var data = [];
                var time = new Date().getTime();
                var i;

                for (i = -19; i <= 0; i += 1) {
                  data.push({
                    x: time + i * 1000,
                    y: 0
                  });
                }
                return data;
              })()
            },
            {
              name: 'Beban 2',
              data: (function() {
                // generate an array of random data
                var data = [];
                var time = new Date().getTime();
                var i;

                for (i = -19; i <= 0; i += 1) {
                  data.push({
                    x: time + i * 1000,
                    y: 0
                  });
                }
                return data;
              })()
            },
            {
              name: 'Beban 3',
              data: (function() {
                // generate an array of random data
                var data = [];
                var time = new Date().getTime();
                var i;

                for (i = -19; i <= 0; i += 1) {
                  data.push({
                    x: time + i * 1000,
                    y: 0
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

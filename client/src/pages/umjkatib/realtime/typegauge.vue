<template>
  <div>
    <!-- <div style="background:red"> -->
    <div class="typegauge1" :id="containerID" style=" width:80%;height: 200px; ">
      <p :id="waktuID"></p>
    </div>
  </div>
</template>
<script>
  /* eslint one-var: ["error", "never"] */
  var Highcharts = window.Highcharts;
  export default {
    name: 'channelgauge1',
    props: ['dataChannel', 'isReadyDetailChannel',
      'updateVal'
    ],
    data() {
      return {
        waktuID: '',
        containerID: ''
      };
    },
    watch: {
      isReadyDetailChannel(param) {
        if (param === true) {
          this.firstInit();
        }
      },
      updateVal(param) {
        console.log(param)
        this.updatedValue(parseFloat(param), window.moment());
      }
    },
    mounted() {
      const thisVue = this;

      if (thisVue.isReadyDetailChannel === true) {
        thisVue.firstInit();
      }
    },
    methods: {
      firstInit() {
        const thisVue = this;
        if (this.dataChannel.containerID === undefined) {
          return;
        }
        var containerTemp = 'container' + this.dataChannel.containerID;
        if (document.querySelectorAll('#' + containerTemp).length) {
          return;
        }
        this.containerID = 'container' + this.dataChannel.containerID;
        setTimeout(() => {
          thisVue.initGauge1(thisVue.dataChannel);
        }, 1000);
      },
      updatedValue(val, time) {
        const thisVue = this;
        $(document).trigger('updatedValueGauge1_' + thisVue.containerID, [
          val,
          time
        ]);
      },
      initGauge1(datak) {
        // var pointGauge1;
        const thisVue = this;
        var widgetVar;
        try {
          const qty = datak.data.length;
          var valUnit = datak.data[qty - 1].value;
          if (valUnit === undefined) {
            return;
          }
          if (valUnit.toString().indexOf('.') === -1) {
            valUnit = parseInt(valUnit);
          } else {
            valUnit = parseFloat(valUnit);
          }
          thisVue.waktuID = datak.data[qty - 1].waktu;
          widgetVar = Highcharts.chart(
            thisVue.containerID,
            {
              chart: {
                type: 'gauge',
                plotBackgroundColor: null,
                plotBackgroundImage: null,
                plotBorderWidth: 0,
                plotShadow: false,
                events: {
                  load: function () {
                    // var thisHighchart = this;
                  }
                }
              },

              title: {
                text: datak.title
              },

              pane: {
                startAngle: -150,
                endAngle: 150,
                background: [
                  {
                    backgroundColor: {
                      linearGradient: { x1: 0, y1: 0, x2: 0, y2: 1 },
                      stops: [[0, '#FFF'], [1, '#333']]
                    },
                    borderWidth: 0,
                    outerRadius: '109%'
                  },
                  {
                    backgroundColor: {
                      linearGradient: { x1: 0, y1: 0, x2: 0, y2: 1 },
                      stops: [[0, '#333'], [1, '#FFF']]
                    },
                    borderWidth: 1,
                    outerRadius: '107%'
                  },
                  {
                    // default background
                  },
                  {
                    backgroundColor: '#DDD',
                    borderWidth: 0,
                    outerRadius: '105%',
                    innerRadius: '103%'
                  }
                ]
              },
              // the value axis
              yAxis: {
                min: datak.rangeMin,
                max: datak.rangeMax,

                minorTickInterval: 'auto',
                minorTickWidth: 1,
                minorTickLength: 10,
                minorTickPosition: 'inside',
                minorTickColor: '#666',

                tickPixelInterval: 30,
                tickWidth: 2,
                tickPosition: 'inside',
                tickLength: 10,
                tickColor: '#666',
                labels: {
                  step: 2,
                  rotation: 'auto'
                },
                title: {
                  text: datak.unit
                }
              },

              series: [
                {
                  name: datak.unit,
                  data: [valUnit],
                  tooltip: {
                    valueSuffix: datak.unit
                  }
                }
              ]
            },
            function (chart) {
              if (!chart.renderer.forExport) {
                $(document).on(
                  'updatedValueGauge1_' + thisVue.containerID,
                  function (e, arg1, arg2) {
                    var point = widgetVar.series[0].points[0];
                    var valTemp;
                    if (arg1.toString().indexOf('.') === -1) {
                      valTemp = parseInt(arg1);
                    } else {
                      valTemp = parseFloat(arg1);
                    }
                    point.update(valTemp);
                  }
                );
              }
            }
          );
          // const roomID = 'chat ' + datak.channelID + '_' + datak.namaTags;
          // window.socketIO.on(roomID, function (message) {
          //   var datake = JSON.parse(message);
          //   try {
          //     thisVue.updatedValue(datake.Value, window.moment());
          //   } catch (error) {
          //     console.error(error);
          //   }
          // });
        } catch (error) {
          console.error(error);
        }
      }
    }
  };
</script>

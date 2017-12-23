<template>
  <div>
    <div id="container"></div>
  </div>
</template>
<script>
export default {
  mounted () {
    this.ChartInit();
  },
  methods: {
    socketIO () {
      var socket = window.io('http://localhost');
      socket.on('connect', function () {
        console.log('connect');
      });
      socket.on('event', function (data) {
        console.log(data);
      });
      socket.on('disconnect', function () {
        console.log('diskonek');
      });
    },
    ChartInit () {
      $(document).ready(function () {
        window.Highcharts.setOptions({
          global: {
            useUTC: false
          }
        });

        window.Highcharts.chart('container', {
          chart: {
            type: 'spline',
            animation: window.Highcharts.svg, // don't animate in old IE
            marginRight: 10,
            events: {
              load: function () {
                // set up the updating of the chart each second
                var series = this.series[0];
                var series2 = this.series[1];
                setInterval(function () {
                  var x = new Date().getTime(); // current time
                  var y = Math.random();
                  series.addPoint([x, y], true, true);
                  x = new Date().getTime(); // current time
                  y = Math.random();
                  series2.addPoint([x, y], true, true);
                }, 1000);
              }
            }
          },
          title: {
            text: 'Monitoring Daya Aktif & Reaktif'
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
              name: 'Daya AKtif',
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
              name: 'Daya Reaktif',
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

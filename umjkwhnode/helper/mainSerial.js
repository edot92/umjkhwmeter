
var SerialPort = require('serialport');
var chalk = require('chalk');

var addrList={
    PF1:3293,
    Frekuensi:3915,
    Vl1:3927,
    A1:3929
};
// console.log(addrList);
var port; 
port = {
    // arduino:new SerialPort(process.env.SERIAL_ARDUINO, {
    //     baudRate: 9600
    // }),
    pm:null,
};
// // Open errors will be emitted as an error event
// // Switches the port into "flowing mode"
// port.arduino.on('data', function (data) {
//     // console.log('Data:', data);
// });
      
// // Read data that is available but keep the stream from entering "flowing mode"
// port.arduino.on('readable', function () {
//     // console.log('Data:', port.read());
// });
    

var ModbusRTU = require('modbus-serial');
port.pm= new ModbusRTU();

// open connection to a serial port
port.pm.connectRTU(process.env.SERIAL_PM, { baudRate: 9600 })
    .then(setClient);
// .then(function() {
//     console.log('Connected'); })
// .catch(function(e) {
//     if(e.errno) {
//         if(networkErrors.includes(e.errno)) {
//             console.log('we have to reconnect');
//         }
//     }
//     console.log(e.message); });
function setClient() {
    // set the client's unit id
    // set a timout for requests default is null (no timeout)
    port.pm.setID(1);
    port.pm.setTimeout(1000);
        
    // run program
    readRegisters();
}
function readRegisters() {
    // read the 2 registers starting at address 5
    // on device number 1.
    port.pm.readHoldingRegisters(3905, 1)
        .then(res=>{
            console.log(res);
            // var buf = new ArrayBuffer(4);
            // // Create a data view of it
            // var view = new DataView(buf);
            
            // // set bytes
            // res.forEach(function (b, i) {
            //     view.setUint8(i, b);
            // });
            
            // // Read the bits as a float; note that by doing this, we're implicitly
            // // converting it from a 32-bit float into JavaScript's native 64-bit double
            // var num = view.getFloat32(0);
            // // Done
            // console.log('f=',chalk.green(num));
        }
        ).catch(err=>{
            console.log('errornya=',chalk.red(err));
        });
}

        
module.exports = port;